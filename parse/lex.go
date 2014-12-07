// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/demizer/go-elog"

	"code.google.com/p/go.text/unicode/norm"
)

// ID is a consecutive number for identication of a lexed item and parsed item.
// Primarily for the purpose of debugging lexer and parser output when compared
// to the JSON encoded tests.
type ID int

// IDNumber returns the ID from an item.
func (i ID) IDNumber() ID { return i }

// String implements Stringer and returns ID as a string.
func (i ID) String() string { return strconv.Itoa(int(i)) }

// Line contains the number of a lexed item, or parsed item, from the input
// data.
type Line int

// LineNumber returns the Line of an item.
func (l Line) LineNumber() Line { return l }

// String implements Stringer and returns Line converted to a string.
func (l Line) String() string { return strconv.Itoa(int(l)) }

// StartPosition is the starting location of an item in the line of input.
type StartPosition int

// Position returns the StartPosition of an item.
func (s StartPosition) Position() StartPosition { return s }

// String implements Stringer and returns StartPosition converted to a string.
func (s StartPosition) String() string { return strconv.Itoa(int(s)) }

// itemElement are the types that are emitted by the lexer.
type itemElement int

const (
	itemEOF itemElement = iota
	itemError
	itemTitle
	itemSectionAdornment
	itemParagraph
	itemBlockQuote
	itemLiteralBlock
	itemSystemMessage
	itemSpace
	itemBlankLine
	itemTransition
	itemCommentMark
	itemEnumListAffix
	itemEnumListArabic
	itemInlineEmphasis
	itemInlineLiteral
	itemDefinitionTerm
)

var elements = [...]string{
	"itemEOF",
	"itemError",
	"itemTitle",
	"itemSectionAdornment",
	"itemParagraph",
	"itemBlockQuote",
	"itemLiteralBlock",
	"itemSystemMessage",
	"itemSpace",
	"itemBlankLine",
	"itemTransition",
	"itemCommentMark",
	"itemEnumListAffix",
	"itemEnumListArabic",
	"itemInlineEmphasis",
	"itemInlineLiteral",
	"itemDefinitionTerm",
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
var sectionAdornments = []rune{'!', '"', '#', '$', '\'', '%', '&', '(', ')',
	'*', '+', ',', '-', '.', '/', ':', ';', '<', '=', '>', '?', '@', '[',
	'\\', ']', '^', '_', '`', '{', '|', '}', '~'}

// Emitted by the lexer on End of File
const eof rune = -1

// Function prototype for scanner functions
type stateFn func(*lexer) stateFn

// Struct for tokens emitted by the scanning process
type item struct {
	ID            `json:"id"`
	Type          itemElement `json:"type"`
	Text          string      `json:"text"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	Length        int `json:"length"`
}

// The lexer struct tracks the state of the lexer
type lexer struct {
	name             string    // The name of the current lexer
	input            string    // The input text
	line             int       // Line number of the parser, from 0
	lines            []string  // The input split into lines
	state            stateFn   // The current state of the lexer
	start            int       // Start position of the token in the line
	index            int       // Position in input
	width            int       // The width of the current position
	items            chan item // The channel items are emitted to
	lastItem         *item     // The last item emitted to the channel
	lastItemPosition StartPosition
	id               int  // Unique ID for each item emitted
	mark             rune // The current lexed rune
}

func newLexer(name, input string) *lexer {
	if len(input) == 0 {
		return nil
	}

	if !norm.NFC.IsNormalString(input) {
		input = norm.NFC.String(input)
	}

	lines := strings.Split(input, "\n")

	mark, width := utf8.DecodeRuneInString(lines[0][0:])

	log.Debugf("mark: %#U, index: %d, line: %d\n", mark, 0, 1)

	return &lexer{
		name:  name,
		input: input,
		lines: lines,
		items: make(chan item),
		index: 0,
		mark:  mark,
		width: width,
	}
}

// lex is the entry point of the lexer. Name should be any name that signifies
// the purporse of the lexer. It is mostly used to identify the lexing process
// in debugging.
func lex(name, input string) *lexer {
	l := newLexer(name, input)
	if l == nil {
		return nil
	}
	go l.run()
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
	var tok string

	if t == itemBlankLine {
		tok = "\n"
	} else if t == itemEOF {
		tok = ""
	} else {
		tok = l.lines[l.line][l.start:l.index]
	}

	log.Infof("%s: %q l.start: %d (%d) l.index: %d (%d) line: %d\n", t,
		tok, l.start, l.start+1, l.index, l.index+1, l.lineNumber())

	l.id++
	length := utf8.RuneCountInString(tok)

	nItem := item{
		ID:   ID(l.id),
		Type: t,
		Text: tok,
		Line: Line(l.lineNumber()),
		// +1 because positions begin at 1, not 0
		StartPosition: StartPosition(l.start + 1),
		Length:        length,
	}

	l.items <- nItem
	l.lastItem = &nItem
	l.start = l.index
}

// backup backs up the lexer position by a number of rune positions (pos).
// backup cannot backup off the input, in that case the index of the lexer is
// set to the starting position on the input. The run
func (l *lexer) backup(pos int) {
	for i := 0; i < pos; i++ {
		if l.index == 0 && l.line != 0 && i < pos {
			l.line--
			l.index = len(l.lines[l.line]) + 1
		}

		l.index -= l.width
		if l.index < 0 {
			l.index = 0
		} else if l.index > len(l.lines[l.line]) {
			l.index--
		}

		r, w := utf8.DecodeRuneInString(l.currentLine()[l.index:])
		l.mark = r
		l.width = w

		// Backup again if iteration has landed on part of a multi-byte
		// rune
		lLen := len(l.currentLine())
		if r == utf8.RuneError && lLen != 0 && lLen != l.index {
			l.backup(1)
		}
	}
	log.Debugln("l.mark backed up to:", string(l.mark))
}

// peek looks ahead in the input by one position and returns the rune.
func (l *lexer) peek() (r rune) {
	r, _ = l.next()
	l.backup(1)
	return
}

func (l *lexer) peekNextLine() string {
	if l.isLastLine() {
		return ""
	}
	return l.lines[l.line+1]
}

// next advances the position of the lexer by one rune and returns that rune.
func (l *lexer) next() (r rune, width int) {
	if l.isEndOfLine() && !l.isLastLine() {
		log.Debugln("Getting next line")
		l.nextLine()
	}

	l.index += l.width
	r, width = utf8.DecodeRuneInString(l.currentLine()[l.index:])
	l.width = width
	l.mark = r

	log.Debugf("mark: %#U, start: %d, index: %d, line: %d\n",
		r, l.start, l.index, l.lineNumber())

	return
}

func (l *lexer) nextLine() string {
	if len(l.lines) == l.line+1 {
		return ""
	}
	l.line++
	l.start = 0
	l.index = 0
	l.width = 0
	return l.lines[l.line]
}

// nextItem returns the next item from the input.
func (l *lexer) nextItem() *item {
	item, ok := <-l.items
	if ok == false {
		return nil
	}
	l.lastItemPosition = item.StartPosition
	return &item

}

// gotoLine advances the lexer to a line and index within that line. Line
// numbers start at 1.
func (l *lexer) gotoLocation(start, line int) {
	l.line = line - 1
	l.index = start
	r, width := utf8.DecodeRuneInString(l.currentLine()[l.index:])
	l.width = width
	l.mark = r
	return
}

func (l *lexer) currentLine() string {
	return l.lines[l.line]
}

func (l *lexer) lineNumber() int {
	return l.line + 1
}

func (l *lexer) isLastLine() bool {
	return len(l.lines) == l.lineNumber()
}

func (l *lexer) lastLineIsBlankLine() bool {
	if l.line == 0 {
		return false
	}
	m, _ := utf8.DecodeRuneInString(l.lines[l.line-1])
	if m == utf8.RuneError {
		return true
	}
	return false
}

func (l *lexer) isEndOfLine() bool {
	return len(l.lines[l.line]) == l.index
}

// isSpace reports whether r is a space character.
func isSpace(r rune) bool {
	return r == ' ' || r == '\t' || r == '\n' || r == '\r'
}

// isArabic returns true if rune r is an Arabic numeral.
func isArabic(r rune) bool {
	return r > '0' && r < '9'
}

// func isInlineMarkup(r rune) bool {
// // TODO: The check should include a peek for the matching rune markup and
// // proper spacing.
// return r == '*' || r == '`' || r == '_' || r == '|'
// }

// isSection compares a number of positions (skipping whitespace) to determine
// if the runes are sectionAdornments and returns a true if the positions match
// each other. Rune comparison begins at the current lexer position. isSection
// returns false if there is a blank line between the positions or if there is
// a rune mismatch between positions.
func isSection(l *lexer) (found bool) {
	log.Debugln("START")
	var nLine string

	checkLine := func(input string, skipSpace bool) (a bool) {
		end := 2
		for j := 0; j < end; j++ {
			r, _ := utf8.DecodeRuneInString(input[l.start+j:])
			if skipSpace && isSpace(r) {
				log.Debugln("Skipping space rune")
				end++
				continue
			}
			a = isSectionAdornment(r)
			if !a {
				return
			}
		}
		return
	}

	log.Debugln("Checking for transition...")
	if isTransition(l) {
		log.Debugln("Returning (found transition)")
		found = false
		goto exit
	}

	if checkLine(l.currentLine(), false) {
		log.Debugln("Found section adornment")
		found = true
		goto exit
	}

	nLine = l.peekNextLine()
	if nLine != "" {
		if checkLine(nLine, true) {
			log.Debugln("Found section adornment")
			found = true
		}
	} else {
		log.Debugln(`l.peekNextLine() == ""`)
	}

exit:
	if !found {
		log.Debugln("Section adornment not found")
	}
	log.Debugln("END")
	return
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

func isTransition(l *lexer) bool {
	log.Debugln("START")
	if r := l.peek(); !isSectionAdornment(l.mark) || !isSectionAdornment(r) {
		log.Debugln("Transition not found")
		return false
	}
	pBlankLine := l.lastItem != nil && l.lastItem.Type == itemBlankLine
	nBlankLine := l.peekNextLine() == ""
	if l.line == 0 && nBlankLine {
		log.Debugln("Found transition (followed by newline)")
		return true
	} else if pBlankLine && nBlankLine {
		log.Debugln("Found transition (surrounded by newlines)")
		return true
		// } else if pBlankLine && l.isLastLine() {
		// log.Debugln("Found transition (opened by newline)")
		// return true
	}
	log.Debugln("Transition not found")
	log.Debugln("END")
	return false
}

func isComment(l *lexer) bool {
	log.Debugln("START")
	if l.lastItem != nil && l.lastItem.Type == itemTitle {
		return false
	}
	if nMark := l.peek(); l.mark == '.' && nMark == '.' {
		l.next()
		nMark2 := l.peek()
		if isSpace(nMark2) || nMark2 == utf8.RuneError {
			log.Debugln("Found comment!")
			return true
		}
		l.backup(1)
	}
	log.Debugln("Comment not found!")
	log.Debugln("END")
	return false
}

func isEnumList(l *lexer) (ret bool) {
	log.Debugln("START")
	bCount := 0
	if isSection(l) {
		goto exit
	}
	if isArabic(l.mark) {
		for {
			bCount++
			if nMark, _ := l.next(); !isArabic(nMark) {
				if nMark == '.' || nMark == ' ' {
					log.Debugln("Found arabic enum list!")
					ret = true
					goto exit
				}
			}
		}
	}
exit:
	l.backup(bCount)
	log.Debugln("END")
	return
}

func isDefinitionTerm(l *lexer) bool {
	log.Debugln("START")
	// Definition terms are preceded by a blankline
	if l.line != 0 && !l.lastLineIsBlankLine() {
		log.Debugln("Not definition, lastLineIsBlankLine == false")
		return false
	}
	nL := l.peekNextLine()
	sCount := 0
	for {
		if sCount < len(nL) && isSpace(rune(nL[sCount])) {
			sCount++
		} else {
			break
		}
	}
	log.Debugln("sCount =", sCount)
	if sCount >= 2 {
		log.Debugln("Found definition term!")
		return true
	}
	log.Debugln("Did not find definition term.")
	log.Debugln("END")
	return false
}

func isBlockquote(l *lexer) bool {
	log.Debugln("START")
	if l.lastLineIsBlankLine() && l.lastItem.Type == itemSpace {
		return true
	}
	log.Debugln("END")
	return false
}

// lexStart is the first stateFn called by run(). From here other stateFn's are
// called depending on the input. When this function returns nil, the lexing is
// finished and run() will exit.
func lexStart(l *lexer) stateFn {
	log.Debugln("START")
	for {
		// log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width:
		// %d, l.line: %d\n", l.mark, l.index, l.start, l.width,
		// l.lineNumber())
		if l.index-l.start <= l.width && l.width > 0 &&
			!l.isEndOfLine() {

			log.Debugf("l.index: %d, l.width: %d, l.line: %d\n",
				l.index, l.width, l.lineNumber())
			if isComment(l) {
				if !l.isEndOfLine() {
					l.next()
				}
				l.emit(itemCommentMark)
				return lexStart
			} else if isEnumList(l) {
				return lexEnumList
			} else if isSection(l) {
				return lexSection
			} else if isTransition(l) {
				return lexTransition
			} else if isSpace(l.mark) {
				return lexSpace
			} else if isBlockquote(l) {
				return lexBlockquote
			} else if isDefinitionTerm(l) {
				return lexDefinitionTerm
			} else {
				return lexParagraph
			}

		} else if l.isEndOfLine() {
			log.Debugln("isEndOfLine == true")
			if l.start == l.index {
				if l.start == 0 && len(l.currentLine()) == 0 {
					log.Debugln("Found blank line")
					l.emit(itemBlankLine)
					if l.isLastLine() {
						break
					}
				} else if l.isLastLine() {
					log.Debugln("Found end of last line")
					break
				}
			}
		}
		l.next()
	}

	l.emit(itemEOF)
	close(l.items)
	log.Debugln("END")
	return nil
}

// lexSpace consumes space characters (space and tab) in the input and emits a
// itemSpace token.
func lexSpace(l *lexer) stateFn {
	log.Debugln("START")
	log.Debugln("l.mark ==", l.mark)
	for isSpace(l.mark) {
		log.Debugln("isSpace ==", isSpace(l.mark))
		if r := l.peek(); isSpace(r) {
			l.next()
		} else {
			log.Debugln("Next mark is not space!")
			l.next()
			break
		}
	}
	log.Debugf("l.start: %d, l.index: %d\n", l.index, l.start)
	if l.start < l.index {
		l.emit(itemSpace)
	}
	log.Debugln("END")
	return lexStart
}

// lexSection is used after isSection() has determined that the next runes of
// input are section.  From here, the lexTitle() and lexSectionAdornment() are
// called based on the input.
func lexSection(l *lexer) stateFn {
	log.Debugln("START")
	// log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width: %d, " +
	// "l.line: %d\n", l.mark, l.index, l.start, l.width, l.lineNumber())
	if isSectionAdornment(l.mark) {
		if l.lastItem != nil && l.lastItem.Type != itemTitle {
			return lexSectionAdornment
		}
		lexSectionAdornment(l)
	} else if isSpace(l.mark) {
		return lexSpace
	} else if l.mark == utf8.RuneError {
		l.next()
	} else if unicode.IsPrint(l.mark) {
		return lexTitle
	}
	log.Debugln("END")
	return lexStart
}

// lexTitle consumes input until newline and emits an itemTitle token. If
// spaces are detected at the start of the line, an itemSpace is emitted.
// Spaces after the title (and before newline) are ignored. On completion
// control is returned to lexSection.
func lexTitle(l *lexer) stateFn {
	log.Debugln("START")
	for {
		l.next()
		if l.isEndOfLine() {
			l.emit(itemTitle)
			break
		}
	}
	log.Debugln("END")
	return lexSection
}

// lexSectionAdornment advances the lexer until a newline is encountered and
// emits a itemSectionAdornment token. Control is returned to lexSection() on
// completion.
func lexSectionAdornment(l *lexer) stateFn {
	log.Debugln("START")
	for {
		if l.isEndOfLine() {
			l.emit(itemSectionAdornment)
			if l.mark == utf8.RuneError {
				break
			}
		}
		l.next()
	}
	log.Debugln("END")
	return lexSection
}

func lexTransition(l *lexer) stateFn {
	log.Debugln("START")
	for {
		if len(l.lines[l.line]) == l.index {
			break
		}
		l.next()
	}
	l.emit(itemTransition)
	l.nextLine()
	log.Debugln("END")
	return lexStart
}

func lexEnumList(l *lexer) stateFn {
	log.Debugln("START")
	if isArabic(l.mark) {
		for {
			if nMark, _ := l.next(); !isArabic(nMark) {
				l.emit(itemEnumListArabic)
				l.next()
				if nMark == '.' {
					l.emit(itemEnumListAffix)
					l.next()
				}
				l.emit(itemSpace)
				break
			}
		}
	}
	log.Debugln("END")
	return lexStart
}

func lexParagraph(l *lexer) stateFn {
	log.Debugln("START")
	for {
		l.next()
		if l.isEndOfLine() && l.mark == utf8.RuneError {
			l.emit(itemParagraph)
			break
		}
	}
	l.nextLine()
	log.Debugln("END")
	return lexStart
}

func lexBlockquote(l *lexer) stateFn {
	log.Debugln("START")
	for {
		l.next()
		if l.isEndOfLine() && l.mark == utf8.RuneError {
			l.emit(itemBlockQuote)
			break
		}
	}
	l.nextLine()
	log.Debugln("END")
	return lexStart
}

func lexDefinitionTerm(l *lexer) stateFn {
	log.Debugln("START")
	for {
		l.next()
		if l.isEndOfLine() && l.mark == utf8.RuneError {
			l.emit(itemDefinitionTerm)
			break
		}
	}
	l.nextLine()
	l.next()
	log.Debugf("Current line: %q\n", l.currentLine())
	lexSpace(l)
	for {
		l.next()
		if l.isEndOfLine() && l.mark == utf8.RuneError {
			l.emit(itemParagraph)
			break
		}
	}
	log.Debugln("END")
	return lexStart
}
