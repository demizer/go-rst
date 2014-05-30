// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"code.google.com/p/go.text/unicode/norm"
	"github.com/demizer/go-elog"
	"strings"
	"unicode"
	"unicode/utf8"
	"strconv"
)

type Id int

func (i Id) IdNumber() Id { return i }

func (i Id) String() string { return strconv.Itoa(int(i)) }

// The line number of an item in the input string
type Line int

func (l Line) LineNumber() Line { return l }

func (l Line) String() string { return strconv.Itoa(int(l)) }

// The begining location of an item in the input.
type StartPosition int

func (s StartPosition) Position() StartPosition { return s }

func (s StartPosition) String() string { return strconv.Itoa(int(s)) }

// itemElement are the types that are emitted by the lexer.
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

// String implements the Stringer interface for printing itemElement types.
func (t itemElement) String() string { return elements[t] }

func (t *itemElement) UnmarshalJSON(data []byte) error {
	for num, elm := range elements {
		if elm == string(data[1:len(data)-1]) {
			*t = itemElement(num)
		}
	}
	return nil
}

// Valid section adornment runes
var sectionAdornments = []rune{'!', '"', '#', '$', '\'', '%', '&', '(', ')', '*',
	'+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[', '\\',
	']', '^', '_', '`', '{', '|', '}', '~'}

// Emitted by the lexer on End of File
const eof rune = -1

// Function prototype for scanner functions
type stateFn func(*lexer) stateFn

// Struct for tokens emitted by the scanning process
type item struct {
	Id            `json:"id"`
	Type          itemElement `json:"type"`
	Text          interface{} `json:"text"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	Length        int `json:"length"`
}

// The lexer struct tracks the state of the lexer
type lexer struct {
	name             string    // The name of the current lexer
	input            string    // The input text
	state            stateFn   // The current state of the lexer
	index            int       // Position in input
	start            int       // The start of the current token
	width            int       // The width of the current position
	items            chan item // The channel items are emitted to
	lastItem         *item     // The last item emitted to the channel
	lastItemPosition StartPosition
	id               int // Unique id for each item emitted
}

// lex is the entry point of the lexer
func lex(name, input string) *lexer {
	log.Debugln("Start")
	if !norm.NFC.IsNormalString(input) {
		input = norm.NFC.String(input)
	}
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run()
	log.Debugln("End")
	return l
}

// run is the engine of the lexing process.
func (l *lexer) run() {
	log.Debugln("Start")
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
	log.Debugln("End")
}

// emit passes an item back to the client.
func (l *lexer) emit(t itemElement) {
	if l.start == l.index && int(l.index) < len(l.input) {
		l.index += 1
	}
	log.Infof("#### %s: %q start: %d pos: %d line: %d\n", t,
		l.input[l.start:l.index], l.start, l.index, l.lineNumber())
	l.id++
	length := utf8.RuneCountInString(l.input[l.start:l.index])
	nItem := item{
		Id:            Id(l.id),
		Type:          t,
		Text:          l.input[l.start:l.index],
		Line:          Line(l.lineNumber()),
		StartPosition: StartPosition(l.start + 1), // Positions start at 1 not 0
		Length:        length,
	}
	l.items <- nItem
	l.lastItem = &nItem
	l.start = l.index
}

// backup backs up the lexer by one position using the width of the last rune retrieved from the
// input.
func (l *lexer) backup(pos int) {
	for i := 0; i < pos; i++ {
		l.index -= l.width
	}
}

// current returns the rune at the current position in the input.
func (l *lexer) current() rune {
	r, _ := utf8.DecodeRuneInString(l.input[l.index:])
	return r
}

// peek looks ahead in the input by one position and returns the rune.
func (l *lexer) peek() rune {
	r := l.next()
	l.backup(1)
	return r
}

// peekNextLine returns rune at pos from the next line of input.
func (l *lexer) peekNextLine(pos int) (r rune) {
	var bCount int
	next := func() {
		bCount++
		r = l.next()
	}
	// Advance to the next line
	for {
		next()
		if r == '\n' {
			next()
			break
		}
	}
	for i := 1; i < pos; i++ {
		next()
	}
	l.backup(bCount)
	return r
}

// advanceToRune fast forwards the lexer to the next rune in the input specified by "to".
func (l *lexer) advanceToRune(to rune) {
	for {
		if l.next() == eof || to == l.current() {
			break
		}
	}
}

// next advances the position of the lexer by one rune and returns that rune.
func (l *lexer) next() rune {
	if int(l.index) >= len(l.input) {
		log.Debugln("Reached eof!")
		l.width = 0
		return eof
	}
	r, w := utf8.DecodeRuneInString(l.input[l.index:])
	l.width = w
	l.index += l.width
	if unicode.In(r, unicode.Diacritic) {
		log.Debugf("DIACRITIC char found: %#U width: %d\n", r, w)
	}
	log.Debugf("cur: %q start: %d pos: %d\n", r, l.start, l.index)
	return r
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() *item {
	item := <-l.items
	l.lastItemPosition = item.StartPosition
	return &item

}

func (l *lexer) lineNumber() int {
	return strings.Count(l.input[:l.index-1], "\n") + 1
}

// isStartOfLine calculates if the current position of the lexer in the input is the beginning of a
// new line.
func (l *lexer) isStartOfLine() bool {
	return (l.index - l.start) == 0
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
	var newLineNum int
	var runePositions []rune
	var matchCount int
	cPos := l.index

	exit := func(value bool) bool {
		l.index = cPos
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
	l.advanceToRune('\n')

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
			log.Debugf("Found adornment: \"%s\" pos: %d\n", string(l.current()), l.index)
			runePositions = append(runePositions, l.current())
			matchCount += 1
		}
		if l.next() == eof {
			return exit(false)
		}
	}

	if len(runePositions) == 0 || matchCount != lookPositions {
		return exit(false)
	}

	return exit(true)
}

// isSectionAdornment returns true if r matches a section adornment.
func isSectionAdornment(r rune) bool {
	for _, a := range sectionAdornments {
		if a == r {
			return true
		}
	}
	return false
}

// lexStart is the first stateFn called by run(). From here other stateFn's are called depending on
// the input. When this function returns nil, the lexing is finished and run() will exit.
func lexStart(l *lexer) stateFn {
	log.Debugln("Start")
	for {
		var tokenLength = l.index - l.start
		r := l.current()
		// log.Debugln("tokenLength:", tokenLength, l.width)
		if tokenLength <= l.width && l.width > 0 && !isEndOfLine(r) {
			log.Debugln("Start of new token")
			if isSection(l) {
				return lexSection
			} else if isSpace(r) {
				return lexSpace
			}
		} else if isEndOfLine(r) {
			log.Debugln("isEndOfLine == true")
			if l.index > l.start {
				l.emit(itemParagraph)
				l.start += 1 // Skip the new line
			} else if l.start == l.index {
				l.emit(itemBlankLine)
			}
		}
		if l.next() == eof {
			break
		}
	}

	// Correctly reached eof.
	l.emit(itemEOF)
	log.Debugln("End")
	return nil
}

// lexSpace consumes space characters (space and tab) in the input and emits a itemSpace token.
func lexSpace(l *lexer) stateFn {
	log.Debugln("Start")
	for isSpace(l.current()) {
		l.next()
	}
	if l.start < l.index {
		l.emit(itemSpace)
		l.next()
	}
	log.Debugln("End")
	return lexStart
}

// lexSection is used after isSection() has determined that the next runes of input are section.
// From here, the lexTitle() and lexSectionAdornment() are called based on the input.
func lexSection(l *lexer) stateFn {
	log.Debugln("Start")
	// The order of the case statements matter here
	switch r := l.next(); {
	case isSectionAdornment(r):
		if l.lastItem != nil && l.lastItem.Type != itemTitle {
			return lexSectionAdornment
		}
		lexSectionAdornment(l)
	case unicode.IsPrint(r):
		return lexTitle
	}
	log.Debugln("Exit")
	return lexStart
}

// lexTitle consumes input until newline and emits an itemTitle token. If spaces are detected at the
// start of the line, an itemSpace is emitted. Spaces after the title (and before newline) are
// ignored. On completion control is returned to lexSection.
func lexTitle(l *lexer) stateFn {
	log.Debugln("Start")
	l.backup(2) // Start from the newline of the previous line
	for {
		if isSpace(l.current()) && l.isStartOfLine() {
			return lexSpace
		}
		l.next()
		if l.peek() == '\n' {
			l.emit(itemTitle)
			l.start += 1
			l.index += 1
			break
		}
	}
	log.Debugln("End")
	return lexSection
}

// lexSectionAdornment advances the lexer until a newline is encountered and emits a
// itemSectionAdornment token. Control is returned to lexSection() on completion.
func lexSectionAdornment(l *lexer) stateFn {
	log.Debugln("Start")
	for {
		//TODO: Add adornment rune check
		l.next()
		if l.peek() == '\n' {
			l.emit(itemSectionAdornment)
			l.start += 1
			l.index += 1
			break
		}
	}
	log.Debugln("End")
	return lexSection
}
