// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

import (
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

type systemMessageLevel int

const (
	levelInfo systemMessageLevel = iota
	levelWarning
	levelError
	levelSevere
)

var systemMessageLevels = [...]string{
	"INFO",
	"WARNING",
	"ERROR",
	"SEVERE",
}

func (s systemMessageLevel) String() string { return systemMessageLevels[s] }

type systemMessage struct {
	level  systemMessageLevel
	line   int
	source string
	items  []item
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
	l.items <- item{ElementType: t, ElementName: fmt.Sprint(t),
		Position: l.start, Line: l.line, Value: l.input[l.start:l.pos]}
	log.Debugf("%s: %q\n", t, l.input[l.start:l.pos])
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

// isWhiteSpace reports whether r is a space character.
func isWhiteSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isEndOfLine reports whether r is an end-of-line character.
func isEndOfLine(r rune) bool {
	return r == '\r' || r == '\n'
}

func lexStart(l *lexer) stateFn {
	log.Debugln("\nTransition...")
	for {
		if len(l.input) == 0 {
			l.emit(itemEOF)
			return nil
		}

		if l.pos > l.start {
			log.Debugf("%q, Current: %q, Start: %d, Pos: %d, Line: %d\n",
				l.input[l.start:l.pos], l.current(), l.start, l.pos, l.line)
		}

		isStartOfToken := l.start == l.pos-l.width

		switch r := l.current(); {
		case isStartOfToken:
			if isWhiteSpace(r) {
				lexWhiteSpace(l)
			}
			if isSection(l) {
				return lexSection
			}
			l.next()
		case isEndOfLine(r):
			if l.pos > l.start {
				l.emit(itemParagraph)
			}
			l.line += 1
			l.skip() // Skip the newline
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
	return nil
}

func lexSection(l *lexer) stateFn {
	var lexTitle bool
	log.Debugln("\nTransition...")

	if !isSectionAdornment(l.current()) {
		lexTitle = true
	}

	for {
		if len(l.input) > 0 {
			log.Debugf("%q, Start: %d, Pos: %d\n", l.input[l.start:l.pos], l.start, l.pos)
		}
		switch r := l.next(); {
		case isEndOfLine(r):
			l.backup()
			if lexTitle {
				l.emit(itemTitle)
				lexTitle = false
				l.line += 1
				l.skip()
				l.next()
				continue
			} else {
				l.emit(itemSectionAdornment)
				return lexStart
			}
		case isWhiteSpace(r):
			lexWhiteSpace(l)
		}
	}
}

func lexWhiteSpace(l *lexer) stateFn {
	log.Debugln("\nTransition...")
	if isEndOfLine(l.previous()) {
		l.emit(itemBlankLine)
		l.line += 1
		l.next()
	}
	for isWhiteSpace(l.peek()) {
		log.Debugf("%q, Start: %d, Pos: %d, Line: %d\n",
			l.input[l.start:l.pos], l.start, l.pos, l.line)
		l.next()
	}
	log.Debugf("%q, Start: %d, Pos: %d, Line: %d\n",
		l.input[l.start:l.pos], l.start, l.pos, l.line)
	l.emit(itemSpace)
	l.next()
	return lexStart
}

// isSection compares a number of positions (skipping whitespace) to
// determine if the runes are sectionAdornments and returns a true if the
// positions match each other. Rune comparison begins at the current lexer
// position. isSection returns false if there is a blank line between the
// positions or if there is a rune mismatch between positions.
func isSection(l *lexer) bool {
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

	log.Debugln("\nLooking ahead", lookPositions, "position(s) for sectionAdornments...")

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
		for isWhiteSpace(l.current()) {
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
