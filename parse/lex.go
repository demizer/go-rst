// go-rst - A reStructuredText parser for Go
// 2014,2015 (c) The go-rst Authors
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

type lexPosition struct {
	index int
	start int
	line  int
	mark  rune
	width int
}

func saveLexerPosition(lexState *lexer) *lexPosition {
	return &lexPosition{
		index: lexState.index,
		start: lexState.start,
		line:  lexState.line,
		mark:  lexState.mark,
		width: lexState.width,
	}
}

func (l *lexPosition) restore(lexState *lexer) {
	lexState.index = l.index
	lexState.start = l.start
	lexState.line = l.line
	lexState.mark = l.mark
	lexState.width = l.width
}

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
	itemInlineStrong
	itemInlineEmphasis
	itemInlineLiteral
	itemInlineInterpretedText
	itemInlineInterpretedTextRole
	itemDefinitionTerm
	itemBullet
	itemEscape
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
	"itemInlineStrong",
	"itemInlineEmphasis",
	"itemInlineLiteral",
	"itemInlineInterpretedText",
	"itemInlineInterpretedTextRole",
	"itemDefinitionTerm",
	"itemBullet",
	"itemEscape",
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

// Runes that must precede inline markup. Includes whitespace and unicode
// categories Pd, Po, Pi, Pf, and Ps. These are checked for in isInlineMarkup()
var inlineMarkupStartStringOpeners = []rune{'-', ':', '/', '\'', '"', '<', '(',
	'[', '{'}

// Runes that must immediately follow inline markup end strings (if not at the
// end of a text block). Includes whitespace and unicode categories Pd, Po, Pi,
// Pf, and Pe. These categories are checked for in isInlineMarkupClosed()
var inlineMarkupEndStringClosers = []rune{'-', '.', ',', ':', ';', '!', '?',
	'\\', '/', '\'', '"', ')', ']', '}', '>'}

var bullets = []rune{'*', '+', '-', '•', '‣', '⁃'}

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
	id               int    // Unique ID for each item emitted
	mark             rune   // The current lexed rune
	indentLevel      int    // For tracking indentation with indentable items
	indentWidth      string // For tracking indent width
}

// getu4 decodes a unicode literal from s length q
func getu4(s []byte, q int) rune {
	r, _ := strconv.ParseUint(string(s[2:q]), 16, 64)
	return rune(r)
}

func newLexer(name string, input []byte) *lexer {
	if len(input) == 0 {
		return nil
	}

	// Convert unicode literals to runes and strip escaped whitespace
	var tInput []byte
	r := 0
	for r < len(input) {
		if input[r] == '\\' && input[r+1] == 'u' {
			tInput = append(tInput, []byte(string(getu4(input[r:], 6)))...)
			r += 6
		} else if input[r] == '\\' && input[r+1] == 'x' {
			tInput = append(tInput, []byte(string(getu4(input[r:], 4)))...)
			r += 4
		} else if input[r] == '\\' && (input[r+1] == '\\') {
			tInput = append(tInput, '\\')
			r += 2
		} else {
			tInput = append(tInput, input[r])
			r++
		}
	}

	var nInput []byte
	if !norm.NFC.IsNormal(tInput) {
		log.Infoln("Normalizing input")
		nInput = norm.NFC.Bytes(tInput)
		tInput = nInput
	}

	lines := strings.Split(string(tInput), "\n")

	mark, width := utf8.DecodeRuneInString(lines[0][0:])
	log.Debugf("mark: %#U, index: %d, line: %d\n", mark, 0, 1)

	return &lexer{
		name:  name,
		input: string(tInput), // stored string is never altered
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
func lex(name string, input []byte) *lexer {
	l := newLexer(name, input)
	if l == nil {
		return nil
	}
	go l.run()
	return l
}

// run is the engine of the lexing process.
func (l *lexer) run() {
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
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

	log.Infof("[ID: %d]: %s: %q l.start: %d (%d) l.index: %d (%d) line: %d\n",
		ID(l.id)+1, t, tok, l.start, l.start+1, l.index, l.index+1, l.lineNumber())

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
	log.Infof("Position after EMIT: l.mark: %q, l.start: %d (%d) l.index: %d (%d) line: %d\n", l.mark, l.start, l.start+1, l.index, l.index+1, l.lineNumber())
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

		l.index -= 1

		if l.index < 0 {
			l.index = 0
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
	log.Debugf("l.mark backed up to: %q\n", l.mark)
}

// peek looks ahead in the input by a number of locations (locs) and returns
// the rune at that location in the input. Peek works across lines.
func (l *lexer) peek(locs int) rune {
	pos := saveLexerPosition(l)
	defer func() {
		pos.restore(l)
	}()
	var r rune
	x := 0
	for x < locs {
		l.next()
		if x == locs-1 {
			r = l.mark
		}
		x++
	}
	log.Debugf("peek() found %q at index %d\n", r, l.index)
	return r
}

func (l *lexer) peekBack(locs int) rune {
	if l.start == l.index {
		return utf8.RuneError
	}
	pos := saveLexerPosition(l)
	defer func() {
		pos.restore(l)
	}()
	var r rune
	x := locs
	for x != 0 {
		l.backup(1)
		r = l.mark
		x--
	}
	log.Debugf("peekBack found %q at index %d\n", string(r), l.index)
	return r
}

func (l *lexer) peekNextLine() string {
	if l.isLastLine() {
		return ""
	}
	return l.lines[l.line+1]
}

// next advances the position of the lexer by one rune and returns that rune.
func (l *lexer) next() (rune, int) {
	if l.isEndOfLine() && !l.isLastLine() {
		log.Debugln("Getting next line")
		l.nextLine()
	}
	l.index += l.width
	r, width := utf8.DecodeRuneInString(l.currentLine()[l.index:])
	l.width = width
	l.mark = r
	log.Debugf("mark: %#U, start: %d, index: %d, line: %d\n",
		r, l.start, l.index, l.lineNumber())

	return r, width
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

func (l *lexer) skip(locs int) {
	for x := 1; x <= locs; x++ {
		l.next()
	}
	l.start = l.index
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
	return unicode.In(r, unicode.Zs, unicode.Zl)
}

// isArabic returns true if rune r is an Arabic numeral.
func isArabic(r rune) bool {
	return r > '0' && r < '9'
}

// isSection compares a number of positions (skipping whitespace) to determine
// if the runes are sectionAdornments and returns a true if the positions match
// each other. Rune comparison begins at the current lexer position. isSection
// returns false if there is a blank line between the positions or if there is
// a rune mismatch between positions.
func isSection(l *lexer) bool {
	log.Debugln("START Checking for section...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()

	// Check two positions to see if the line contains a section adornment
	checkLine := func(input string) bool {
		var first, last rune
		for j := 0; j < len(input); j++ {
			r, _ := utf8.DecodeRuneInString(input[j:])
			if isSpace(r) {
				log.Debugln("Skipping space rune")
				continue
			}
			if first == '\x00' {
				first = r
				last = r
			}
			// log.Debugf("first: %q, last: %q, r: %q, j: %d\n", first, last, r, j)
			if !isSectionAdornment(r) || (r != first && last != first) {
				log.Debugln("Section not found")
				return false
			}
			last = r
		}
		return true
	}

	if isTransition(l) {
		log.Debugln("Returning (found transition)")
		return false
	}

	log.Debugln("Checking current line")
	if checkLine(l.currentLine()) {
		log.Debugln("Found section adornment")
		return true
	}

	log.Debugln("Checking next line")

	nLine := l.peekNextLine()
	if nLine != "" {
		if checkLine(nLine) {
			log.Debugln("Found section adornment (nextline)")
			return true
		}
	}
	log.Debugln("Section not found")
	return false
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
	log.Debugln("START Checking for transition...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	if r := l.peek(4); !isSectionAdornment(l.mark) || !isSectionAdornment(r) {
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
	}
	log.Debugln("Transition not found")
	return false
}

func isComment(l *lexer) bool {
	log.Debugln("START Checking for comment...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	if l.lastItem != nil && l.lastItem.Type == itemTitle {
		return false
	}
	if nMark := l.peek(1); l.mark == '.' && nMark == '.' {
		nMark2 := l.peek(2)
		if isSpace(nMark2) || nMark2 == utf8.RuneError {
			log.Debugln("Found comment!")
			return true
		}
	}
	log.Debugln("Comment not found!")
	return false
}

func isEnumList(l *lexer) (ret bool) {
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
	return
}

func isBulletList(l *lexer) bool {
	log.Debugln("START Checking for bullet...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	for _, x := range bullets {
		if l.mark == x && l.peek(1) == ' ' {
			log.Debugln("A bullet was found")
			return true
		}
	}
	log.Debugln("A bullet was not found")
	return false
}

func isDefinitionTerm(l *lexer) bool {
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
	return false
}

func isBlockquote(l *lexer) bool {
	if !l.lastLineIsBlankLine() || l.lastItem.Type != itemSpace {
		return false
	}
	if l.index != len(l.indentWidth) {
		return true
	}
	return false
}

func isInlineMarkup(l *lexer) bool {
	isOpenerRune := func(r rune) bool {
		for _, x := range inlineMarkupStartStringOpeners {
			if x == r {
				return true
			}
		}
		if unicode.In(r, unicode.Pd, unicode.Po, unicode.Pi, unicode.Pf,
			unicode.Ps, unicode.Zs, unicode.Zl) {
			return true
		}
		return false
	}
	isSurrounded := func(back, front rune) bool {
		if back == '\'' && front == '\'' {
			return true
		} else if back == '"' && front == '"' {
			return true
		} else if back == '<' && front == '>' {
			return true
		} else if back == '(' && front == ')' {
			return true
		} else if back == '[' && front == ']' {
			return true
		} else if back == '{' && front == '}' {
			return true
		} else if unicode.In(back, unicode.Ps) &&
			unicode.In(front, unicode.Pe, unicode.Pf, unicode.Pi) {
			return true
		} else if unicode.In(back, unicode.Pi) &&
			unicode.In(front, unicode.Pf, unicode.Ps) {
			return true
		} else if unicode.In(back, unicode.Pf) &&
			unicode.In(front, unicode.Pf) {
			return true
		} else if unicode.In(back, unicode.Pf) &&
			unicode.In(front, unicode.Pi) {
			return true
		}
		return false
	}
	if l.mark == '*' || l.mark == '`' {
		log.Debugln("START Checking for inline markup")
		log.SetIndent(log.Indent() + 1)
		defer func() {
			log.SetIndent(log.Indent() - 1)
			log.Debugln("END")
		}()
		b := l.peekBack(1)
		f := l.peek(1)
		if !isSurrounded(b, f) && (isOpenerRune(b) || l.start == l.index) && !isSpace(f) {
			log.Debugln("Found inline markup!")
			return true
		}
	}
	return false
}

func isInlineMarkupClosed(l *lexer, markup string) bool {
	log.Debugln("START Checking for closed inline markup")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	isEndAscii := func(r rune) bool {
		for _, x := range inlineMarkupEndStringClosers {
			if x == r {
				return true
			}
		}
		return false
	}
	var a, b rune
	b = l.peekBack(1)
	a = l.peek(1)
	if len(markup) > 1 {
		a = l.peek(2)
	}
	if (b == '\\' || b == '*') && !isSpace(a) {
		log.Debugln("Inline markup close not found (b == '\\' || b == '*')")
		return false
	}
	if !isSpace(b) && (isSpace(a) || isEndAscii(a) || unicode.In(a, unicode.Pd, unicode.Po, unicode.Pi, unicode.Pf, unicode.Pe, unicode.Ps)) {
		log.Debugln("Found inline markup close")
		return true
	}
	log.Debugln("Inline markup close not found")
	return false
}

func isEscaped(l *lexer) bool {
	// log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d\n",
	// l.mark, l.index, l.width, l.lineNumber())
	return (l.mark == '\\' &&
		(unicode.In(l.peek(1), unicode.Zs, unicode.Cc, unicode.Lu, unicode.Ll) ||
			l.peek(1) == utf8.RuneError))
}

// lexStart is the first stateFn called by run(). From here other stateFn's are
// called depending on the input. When this function returns nil, the lexing is
// finished and run() will exit.
func lexStart(l *lexer) stateFn {
	log.Debugln("START")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	for {
		// log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width:
		// %d, l.line: %d\n", l.mark, l.index, l.start, l.width,
		// l.lineNumber())
		if l.index-l.start <= l.width && l.width > 0 &&
			!l.isEndOfLine() {
			if l.index == 0 && l.mark != ' ' {
				l.indentLevel = 0
				l.indentWidth = ""
			}
			log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d\n",
				l.mark, l.index, l.width, l.lineNumber())
			if isComment(l) {
				return lexComment
			} else if isBulletList(l) {
				return lexBullet
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
			} else if isInlineMarkup(l) {
				return lexInlineMarkup
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
	return nil
}

// lexSpace consumes space characters (space and tab) in the input and emits a
// itemSpace token.
func lexSpace(l *lexer) stateFn {
	log.Debugln("START lexing space")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	log.Debugln("l.mark ==", l.mark)
	for isSpace(l.mark) {
		log.Debugln("isSpace ==", isSpace(l.mark))
		if r := l.peek(1); isSpace(r) {
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
	return lexStart
}

// lexSection is used after isSection() has determined that the next runes of
// input are section.  From here, the lexTitle() and lexSectionAdornment() are
// called based on the input.
func lexSection(l *lexer) stateFn {
	log.Debugln("START lexing section")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
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
	return lexStart
}

// lexTitle consumes input until newline and emits an itemTitle token. If
// spaces are detected at the start of the line, an itemSpace is emitted.
// Spaces after the title (and before newline) are ignored. On completion
// control is returned to lexSection.
func lexTitle(l *lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() {
			l.emit(itemTitle)
			break
		}
	}
	return lexSection
}

// lexSectionAdornment advances the lexer until a newline is encountered and
// emits a itemSectionAdornment token. Control is returned to lexSection() on
// completion.
func lexSectionAdornment(l *lexer) stateFn {
	for {
		if l.isEndOfLine() {
			l.emit(itemSectionAdornment)
			if l.mark == utf8.RuneError {
				break
			}
		}
		l.next()
	}
	return lexSection
}

func lexTransition(l *lexer) stateFn {
	for {
		if len(l.lines[l.line]) == l.index {
			break
		}
		l.next()
	}
	l.emit(itemTransition)
	l.nextLine()
	return lexStart
}

func lexEnumList(l *lexer) stateFn {
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
	return lexStart
}

func lexParagraph(l *lexer) stateFn {
	log.Debugln("START lexing paragraph")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	for {
		// log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d\n",
		// l.mark, l.index, l.width, l.lineNumber())
		if isEscaped(l) {
			l.emit(itemParagraph)
			lexEscape(l)
		}
		if isInlineMarkup(l) {
			if l.index > l.start {
				l.emit(itemParagraph)
			}
			lexInlineMarkup(l)
			if isEscaped(l) {
				lexEscape(l)
			}
			continue
		}
		if l.isEndOfLine() && l.mark == utf8.RuneError {
			if l.start == l.index {
				return lexStart
			} else {
				l.emit(itemParagraph)
				break
			}
		}
		l.next()
	}
	l.nextLine()
	return lexStart
}

func lexComment(l *lexer) stateFn {
	log.Debugln("START lexing comment")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	for l.mark == '.' {
		l.next()
	}
	l.emit(itemCommentMark)
	if l.mark != utf8.RuneError {
		l.next()
		lexSpace(l)
		lexParagraph(l)
	}
	return lexStart
}

func lexBlockquote(l *lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() && l.mark == utf8.RuneError {
			l.emit(itemBlockQuote)
			break
		}
	}
	l.nextLine()
	return lexStart
}

func lexDefinitionTerm(l *lexer) stateFn {
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
	return lexStart
}

func lexBullet(l *lexer) stateFn {
	log.Debugln("START lexing bullet")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	l.next()
	l.emit(itemBullet)
	lexSpace(l)
	l.indentWidth += l.lastItem.Text + " "
	lexParagraph(l)
	l.indentLevel++
	return lexStart
}

func lexInlineMarkup(l *lexer) stateFn {
	for {
		log.Debugf("l.mark: %q l.start: %d l.index: %d l.width: %d l.line: %d\n",
			l.mark, l.start, l.index, l.width, l.lineNumber())
		if l.mark == '*' && l.peek(1) == '*' {
			lexInlineStrong(l)
			break
		} else if l.mark == '*' {
			lexInlineEmphasis(l)
			break
		} else if l.mark == '`' && l.peek(1) == '`' {
			lexInlineLiteral(l)
			break
		} else if l.mark == '`' {
			lexInlineInterpretedText(l)
			break
		}
	}
	return lexStart
}

func lexInlineStrong(l *lexer) stateFn {
	log.Debugln("START lexing inline strong...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	// skip the '*'
	l.skip(2)
	for {
		l.next()
		if l.peekBack(1) != '\\' && l.mark == '*' && isInlineMarkupClosed(l, "**") {
			log.Debugln("Found strong close")
			l.emit(itemInlineStrong)
			break
		}
	}
	// skip the '*'
	l.skip(2)
	return lexStart
}

func lexInlineEmphasis(l *lexer) stateFn {
	log.Debugln("START lexing inline emphasis...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	// skip the '*'
	l.skip(1)
	for {
		l.next()
		if l.peekBack(1) != '\\' && l.mark == '*' && isInlineMarkupClosed(l, "*") {
			log.Debugln("Found emphasis close")
			l.emit(itemInlineEmphasis)
			break
		} else if l.mark == '*' && l.peek(1) == utf8.RuneError {
			log.Debugln("Found emphasis close at end-of-line")
			l.emit(itemInlineEmphasis)
			break
		} else if l.isEndOfLine() && l.mark == utf8.RuneError {
			log.Debugln("Found end-of-line")
			l.emit(itemInlineEmphasis)
			l.emit(itemBlankLine)
			l.nextLine()
		}
	}
	// skip the '*'
	l.skip(1)
	return lexStart
}

func lexEscape(l *lexer) stateFn {
	log.Debugln("START lexing escape code")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	if l.peek(1) == utf8.RuneError {
		// Ignore escape sequences at the end of a line
		log.Debugln("Found escape at end of line, ignoring")
		l.skip(2)
		return lexStart
	}
	l.next()
	l.next()
	l.emit(itemEscape)
	return lexStart
}

func lexInlineLiteral(l *lexer) stateFn {
	log.Debugln("START lexing inline literal...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	// skip the '`'
	l.skip(2)
	for {
		l.next()
		if l.mark == '`' && isInlineMarkupClosed(l, "``") {
			log.Debugln("Found literal close")
			l.emit(itemInlineLiteral)
			break
		}
	}
	// skip the '`'
	l.skip(2)
	return lexStart
}

func lexInlineInterpretedText(l *lexer) stateFn {
	log.Debugln("START lexing inline interpreted text...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	// skip the '`'
	l.skip(1)
	for {
		l.next()
		if l.mark == '`' && isInlineMarkupClosed(l, "`") {
			log.Debugln("Found literal close")
			l.emit(itemInlineInterpretedText)
			break
		}
	}
	// skip the '`'
	l.skip(1)
	if l.mark == ':' {
		lexInlineInterpretedTextRole(l)
	}
	return lexStart
}

func lexInlineInterpretedTextRole(l *lexer) stateFn {
	log.Debugln("START lexing inline interpreted text role...")
	log.SetIndent(log.Indent() + 1)
	defer func() {
		log.SetIndent(log.Indent() - 1)
		log.Debugln("END")
	}()
	// skip the :
	l.skip(1)
	for {
		l.next()
		if l.mark == ':' {
			l.emit(itemInlineInterpretedTextRole)
			break
		}
	}
	// skip the :
	l.skip(1)
	return lexStart
}
