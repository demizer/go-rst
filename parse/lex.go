// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"fmt"
	"github.com/demizer/go-elog"
	"unicode"
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
	itemSpace
	itemBlankLine
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
	"itemSpace",
	"itemBlankLine",
}

func (t itemElement) String() string { return elements[t] }

var sectionAdornments = []rune{'!', '"', '#', '$', '\'', '%', '&', '(', ')', '*',
	'+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\',
	']', '^', '_', '`', '{', '|', '}', '~'}

const EOF rune = -1

type stateFn func(*lexer) stateFn

type item struct {
	ElementName string      `json:"element-name"`
	ElementType itemElement `json: "-"`
	Position    Pos         `json: "position"`
	Line        int         `json: "line"`
	Value       interface{} `json: "value"`
}

type lexer struct {
	name     string
	input    string
	state    stateFn
	pos      Pos
	start    Pos
	width    Pos
	lastPos  Pos
	items    chan item
	lastItem *item
	line     int
}

func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		line:  1,
		items: make(chan item),
	}
	go l.run()
	return l
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemElement) {
	if l.start == l.pos && int(l.pos) < len(l.input) {
		l.pos += 1
	}
	log.Debugf("#### %s: %q start: %d pos: %d\n", t, l.input[l.start:l.pos], l.start, l.pos)
	nItem := item{ElementType: t, ElementName: fmt.Sprint(t), Position: l.start+1, Line: l.line,
		Value: l.input[l.start:l.pos]}
	l.items <- nItem
	l.lastItem = &nItem
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

func (l *lexer) advance(to rune) {
	for {
		if l.next() == EOF || to == l.current() {
			break
		}
	}
}

func (l *lexer) next() rune {
	if int(l.pos) >= len(l.input) {
		log.Debugln("Reached EOF!")
		l.width = 0
		return EOF
	}
	r, w := utf8.DecodeRuneInString(l.input[l.pos:])
	l.width = Pos(w)
	l.pos += l.width
	log.Debugf("cur: %q start: %d pos: %d\n", r, l.start, l.pos)
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
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

// isSection compares a number of positions (skipping whitespace) to
// determine if the runes are sectionAdornments and returns a true if the
// positions match each other. Rune comparison begins at the current lexer
// position. isSection returns false if there is a blank line between the
// positions or if there is a rune mismatch between positions.
func isSection(l *lexer) bool {
	log.Debugln("Start")
	var lookPositions = 2
	var lastAdornment rune
	var newLineNum int
	var runePositions []rune
	var matchCount int
	cPos := l.pos

	exit := func(value bool) bool {
		l.pos = cPos
		log.Debugln("Returning", value)
		return value
	}

	log.Debugln("Looking ahead", lookPositions, "position(s) for sectionAdornments...")

	// Check two runes to see if they are section adornments, if not, we will check the next
	// line (after whitespace).
	if isSectionAdornment(l.current()) && isSectionAdornment(l.peek()) &&
		l.current() == l.peek() {
		return true
	}

	// Advance to the end of the line
	l.advance('\n')
	if l.current() == EOF {
		return false
	}

	for j := 0; j < lookPositions; j++ {
		for isSpace(l.current()) {
			if isEndOfLine(l.current()) {
				if newLineNum == 1 {
					log.Debugln("Too many newlines!")
					return exit(false)
				}
				newLineNum += 1
				log.Debugln("newLineNum:", newLineNum)
			}
			l.next()
		}

		if isSectionAdornment(l.current()) {
			log.Debugf("Found adornment: \"%s\" pos: %d\n", string(l.current()), l.pos)
			if lastAdornment != 0 && l.current() != lastAdornment {
				log.Debugf("Adornment mismatch, last: %s current: %s\n",
					string(lastAdornment), string(l.current()))
				return exit(false)
			}
			runePositions = append(runePositions, l.current())
			lastAdornment = l.current()
			matchCount += 1
		}
		l.next()
	}

	if len(runePositions) == 0 || matchCount != lookPositions {
		return exit(false)
	}

	return exit(true)
}

func isSectionAdornment(r rune) bool {
	for _, a := range sectionAdornments {
		if a == r {
			return true
		}
	}
	return false
}

func lexStart(l *lexer) stateFn {
	log.Debugln("Start")
	for {
		var tokenLength = l.pos - l.start
		switch r := l.current(); {
		case tokenLength == 1:
			log.Debugln("tokenLength == 1; Start of new token")
			if isEndOfLine(r) {
				l.emit(itemBlankLine)
				l.start += 1
				l.line += 1
			} else if isSpace(r) {
				return lexSpace
			} else if isSection(l) {
				return lexSection
			}
		case isEndOfLine(r):
			log.Debugln("isEndOfLine == true")
			if l.pos > l.start {
				l.emit(itemParagraph)
				l.start += 1 // Skip the new line
			} else if l.start == l.pos {
				l.emit(itemBlankLine)
			}
			l.line += 1

		}
		if l.next() == EOF {
			break
		}
	}

	// Correctly reached EOF.
	if l.pos > l.start {
		l.emit(itemParagraph)
	}

	l.emit(itemEOF)
	log.Debugln("End")
	return nil
}

func lexSpace(l *lexer) stateFn {
	log.Debugln("Start")
	for isSpace(l.current()) {
		l.next()
	}
	if l.start < l.pos {
		l.emit(itemSpace)
		l.next()
	}
	log.Debugln("End")
	return lexStart
}

func lexSection(l *lexer) stateFn {
	log.Debugln("Start")
	// The order of the case statement matter here
	switch r := l.next(); {
	case isSectionAdornment(r):
		if l.lastItem.ElementType != itemTitle {
			return lexSectionAdornment
		}
		lexSectionAdornment(l)
	case isSpace(r):
		return lexSpace
	case r <= unicode.MaxASCII && unicode.IsPrint(r):
		return lexTitle
	case isEndOfLine(r):
		l.start += 1
		l.line += 1
	}
	log.Debugln("Exit")
	return lexStart
}

func lexTitle(l *lexer) stateFn {
	log.Debugln("Start")
	for {
		l.next()
		if l.peek() == '\n' {
			l.emit(itemTitle)
			l.start += 1
			l.pos += 1
			l.line += 1
			break
		}
	}
	log.Debugln("End")
	return lexSection
}

func lexSectionAdornment(l *lexer) stateFn {
	log.Debugln("Start")
	for {
		//TODO: Add adornment rune check
		l.next()
		if l.peek() == '\n' {
			l.emit(itemSectionAdornment)
			l.start += 1
			l.pos += 1
			l.line += 1
			break
		}
	}
	log.Debugln("End")
	return lexSection
}
