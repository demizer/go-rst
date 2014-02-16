// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

import (
	// "strings"
	// "unicode"
	"fmt"
	"unicode/utf8"
	"github.com/demizer/go-elog"
)

type itemElement int

const (
	itemEOF itemElement = iota
	itemError
	itemTitle
	itemSectionAdornment
	itemParagraph
)

var sectionAdornments = []rune{'!', '"', '#', '$', '\'', '%', '&', '(', ')', '*',
	'+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\',
	']', '^', '_', '`', '{', '|', '}', '~'}

func isSectionAdornment(r rune) bool {
	for _, a := range sectionAdornments {
		if a == r {
			return true
		}
	}
	return false
}

const eof = -1

type stateFn func(*lexer) stateFn

type item struct {
	typ itemElement
	pos Pos
	val string
}

func (i item) String() string {
	switch {
	case i.typ == itemEOF:
		return "EOF"
	case i.typ == itemError:
		return i.val
	}
	return fmt.Sprintf("%q", i.val)
}

type lexer struct {
	name       string
	input      string
	state      stateFn
	pos        Pos
	start      Pos
	width      Pos
	lastPos    Pos
	items      chan item
}

func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemElement) {
	l.items <- item{t, l.start, l.input[l.start:l.pos]}
	l.start = l.pos

}

func (l *lexer) backup() {
	l.pos -= l.width
}

func (l *lexer) current() rune {
	r, _ := utf8.DecodeRuneInString(l.input[l.pos:])
	return r
}

func (l *lexer) peek() rune {
	r := l.next()
	l.backup()
	return r
}

func (l *lexer) ignore() {
	l.pos += 1
	l.start = l.pos
}

// next returns the next rune in the input.
func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	return r
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() item {
	item := <-l.items
	l.lastPos = item.pos
	return item

}

func (l *lexer) run() {
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

func lexStart(l *lexer) stateFn {
	for {
		if len(l.input) == 0 {
			log.Debugln("\tEmit EOF!")
			l.emit(itemEOF)
			return nil
		}

		log.Debugf("\tlexParagraph: %q, Start: %d, Pos: %d\n",
			l.input[l.start:l.pos], l.start, l.pos)
		if isEndOfLine(l.current()) || isEndOfLine(l.peek()) {
			log.Debugln("\tFound newline!")
			if isSectionAdornment(rune(l.input[l.pos+2])) {
				log.Debugln("Transition lexSection...")
				return lexSection
			}
			l.ignore()
		}

		if l.next() == eof {
			break
		}
	}

	// Correctly reached EOF.
	if l.pos > l.start {
		log.Debugln("\tEmit Paragraph!")
		l.emit(itemParagraph)
	}

	l.emit(itemEOF)
	return nil
}

func lexSection(l *lexer) stateFn {
	var titleWidth, sectionWidth Pos
	if len(l.input) > 0 {
		log.Debugf("\tlexSection: %q, Pos: %d, titleWidth: %d, sectionWidth: %d\n",
			l.input[l.start:l.pos], l.pos, titleWidth, sectionWidth)
	}

	if isEndOfLine(l.peek()) {
		titleWidth = l.pos - l.start
		log.Debugf("\tEmit itemTitle: %d-%d\n", l.start, l.pos)
		l.emit(itemTitle)
		l.ignore()
	}

Loop:
	for {
		switch r := l.next(); {
		case isSectionAdornment(r):
			if len(l.input) > 0 {
				log.Debugf("\tlexSection: %q, Pos: %d, " +
				           "titleWidth: %d, sectionWidth: %d\n",
					   l.input[l.start:l.pos], l.pos,
					   titleWidth, sectionWidth)
			}
			sectionWidth += 1
		case isEndOfLine(r):
			if sectionWidth == titleWidth {
				l.backup()
				log.Debugf("\tEmit itemSectionAdornment: " +
					   "%d-%d\n", l.start, l.pos)
				l.emit(itemSectionAdornment)
				l.ignore()
				break Loop
			}
		}
	}
	log.Debugln("Transition lexStart...")
	return lexStart
}
