// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

import (
	// "strings"
	// "unicode"
	"fmt"
	"github.com/demizer/go-elog"
	"unicode/utf8"
)

type itemElement int

const (
	itemEOF itemElement = iota
	itemError
	itemTitle
	itemSectionAdornment
	itemParagraph
	itemBlockquote
	itemLiteralBlock
	itemSystemMessage
)

var elements = [...]string{
	"itemEOF",
	"itemError",
	"itemTitle",
	"itemSectionAdornment",
	"itemParagraph",
	"itemBlockquote",
	"itemLiteralBlock",
	"itemSystemMessage",
}

func (t itemElement) String() string { return elements[t] }

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
	ElementType itemElement
	Position    Pos
	Value       interface{}
}

func (i item) String() string {
	switch {
	case i.ElementType == itemEOF:
		return "EOF"
	case i.ElementType == itemError:
		return i.Value.(string)
	}
	return fmt.Sprintf("%q", i.Value)
}

type lexer struct {
	name    string
	input   string
	state   stateFn
	pos     Pos
	start   Pos
	width   Pos
	lastPos Pos
	items   chan item
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
	log.Debugf("\tEmit %s!\n", t)
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
	l.lastPos = item.Position
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
	log.Debugln("\nTransition lexStart...")
	for {
		if len(l.input) == 0 {
			log.Debugln("\tEmit EOF!")
			l.emit(itemEOF)
			return nil
		}

		log.Debugf("\tlexStart: %q, Start: %d, Pos: %d\n",
			l.input[l.start:l.pos], l.start, l.pos)

		switch r := l.current(); {
		case isSectionAdornment(r) && isSectionAdornment(l.peek()) && l.pos == 1:
			log.Debugln("Transition lexSection...")
			return lexSection
		case isEndOfLine(r):
			log.Debugln("\tFound newline!")
			if isSectionAdornment(rune(l.input[l.pos+2])) {
				log.Debugln("Transition lexSection...")
				return lexSection
			}
			if l.pos > l.start {
				l.emit(itemParagraph)
			}
			l.ignore()
		}

		if l.next() == eof {
			break
		}
	}

	// Correctly reached EOF.
	if l.pos > l.start {
		l.emit(itemParagraph)
	}

	l.emit(itemEOF)
	return nil
}

func lexSection(l *lexer) stateFn {
	if len(l.input) > 0 {
		log.Debugf("\tlexSection: %q, Pos: %d\n",
			l.input[l.start:l.pos], l.pos)
	}

	if isEndOfLine(l.peek()) {
		l.emit(itemTitle)
		l.ignore()
	}

Loop:
	for {
		switch r := l.next(); {
		case isSectionAdornment(r):
			if len(l.input) > 0 {
				log.Debugf("\tlexSection: %q, Pos: %d\n",
					l.input[l.start:l.pos], l.pos)
			}
		case isEndOfLine(r):
			l.backup()
			l.emit(itemSectionAdornment)
			l.ignore()
			break Loop
		}
	}
	return lexStart
}
