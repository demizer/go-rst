package parse

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"

	"github.com/apex/log"
	"golang.org/x/text/unicode/norm"
)

// ID is a consecutive number for identication of a lexed item and parsed item. Primarily for the purpose of debugging lexer
// and parser output when compared to the JSON encoded tests.
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
	itemInlineStrongOpen
	itemInlineStrong
	itemInlineStrongClose
	itemInlineEmphasisOpen
	itemInlineEmphasis
	itemInlineEmphasisClose
	itemInlineLiteralOpen
	itemInlineLiteral
	itemInlineLiteralClose
	itemInlineInterpretedTextOpen
	itemInlineInterpretedText
	itemInlineInterpretedTextClose
	itemInlineInterpretedTextRoleOpen
	itemInlineInterpretedTextRole
	itemInlineInterpretedTextRoleClose
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
	"itemInlineStrongOpen",
	"itemInlineStrong",
	"itemInlineStrongClose",
	"itemInlineEmphasisOpen",
	"itemInlineEmphasis",
	"itemInlineEmphasisClose",
	"itemInlineLiteralOpen",
	"itemInlineLiteral",
	"itemInlineLiteralClose",
	"itemInlineInterpretedTextOpen",
	"itemInlineInterpretedText",
	"itemInlineInterpretedTextClose",
	"itemInlineInterpretedTextRoleOpen",
	"itemInlineInterpretedTextRole",
	"itemInlineInterpretedTextRoleClose",
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
var sectionAdornments = []rune{'!', '"', '#', '$', '\'', '%', '&', '(', ')', '*', '+', ',', '-', '.', '/', ':', ';', '<',
	'=', '>', '?', '@', '[', '\\', ']', '^', '_', '`', '{', '|', '}', '~'}

// Runes that must precede inline markup. Includes whitespace and unicode categories Pd, Po, Pi, Pf, and Ps. These are
// checked for in isInlineMarkup()
var inlineMarkupStartStringOpeners = []rune{'-', ':', '/', '\'', '"', '<', '(', '[', '{'}

// Runes that must immediately follow inline markup end strings (if not at the end of a text block). Includes whitespace and
// unicode categories Pd, Po, Pi, Pf, and Pe. These categories are checked for in isInlineMarkupClosed()
var inlineMarkupEndStringClosers = []rune{'-', '.', ',', ':', ';', '!', '?', '\\', '/', '\'', '"', ')', ']', '}', '>'}

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
	log              *log.Entry
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
	l := &lexer{name: name, log: log.NewEntry(Log).WithField("unit", "lexer")}
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
		l.log.Info("nexLexer: Normalizing input")
		nInput = norm.NFC.Bytes(tInput)
		tInput = nInput
	}

	lines := strings.Split(string(tInput), "\n")

	mark, width := utf8.DecodeRuneInString(lines[0][0:])
	l.log.WithFields(log.Fields{"mark": mark, "index": 0, "line": 1}).Debug("nexLexer: mark")

	l.input = string(tInput) // stored string is never altered
	l.lines = lines
	l.items = make(chan item)
	l.index = 0
	l.mark = mark
	l.width = width
	return l
}

// lex is the entry point of the lexer. Name should be any name that signifies the purporse of the lexer. It is mostly used
// to identify the lexing process in debugging.
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

	Log.Infof("[ID: %d]: %s: %q l.start: %d (%d) l.index: %d (%d) line: %d",
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
	Log.Infof("Position after EMIT: l.mark: %q, l.start: %d (%d) l.index: %d (%d) line: %d",
		l.mark, l.start, l.start+1, l.index, l.index+1, l.lineNumber())
}

// backup backs up the lexer position by a number of rune positions (pos).  backup cannot backup off the input, in that case
// the index of the lexer is set to the starting position on the input. The run
func (l *lexer) backup(pos int) {
	for i := 0; i < pos; i++ {
		if l.index == 0 && l.line != 0 && i < pos {
			l.line--
			l.index = len(l.lines[l.line]) + 1
		}
		l.index--
		if l.index < 0 {
			l.index = 0
		}
		r, w := utf8.DecodeRuneInString(l.currentLine()[l.index:])
		l.mark = r
		l.width = w
		// Backup again if iteration has landed on part of a multi-byte rune
		lLen := len(l.currentLine())
		if r == utf8.RuneError && lLen != 0 && lLen != l.index {
			l.backup(1)
		}
	}
	Log.Debugf("l.mark backed up to: %q", l.mark)
}

// peek looks ahead in the input by a number of locations (locs) and returns the rune at that location in the input. Peek
// works across lines.
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
	l.log.WithFields(log.Fields{"mark": string(r), "index": l.index}).Debugf("peek: mark")
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
	l.log.WithFields(log.Fields{"mark": string(r), "index": l.index}).Debugf("peekBack: mark")
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
		l.log.Debug("next: Getting next line")
		l.nextLine()
	}
	l.index += l.width
	r, width := utf8.DecodeRuneInString(l.currentLine()[l.index:])
	l.width = width
	l.mark = r
	// Log.Debugf("mark: %#U, start: %d, index: %d, line: %d", r, l.start, l.index, l.lineNumber())
	l.log.WithFields(log.Fields{"mark": fmt.Sprintf("%#U", r),
		"start": l.start, "index": l.index, "line": l.lineNumber()}).Debug("next: mark")
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

// gotoLine advances the lexer to a line and index within that line. Line numbers start at 1.
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

// isArabic returns true if rune r is an Arabic numeral.
func isArabic(r rune) bool {
	return r > '0' && r < '9'
}

// isSection compares a number of positions (skipping whitespace) to determine if the runes are sectionAdornments and returns
// a true if the positions match each other. Rune comparison begins at the current lexer position. isSection returns false if
// there is a blank line between the positions or if there is a rune mismatch between positions.
func isSection(l *lexer) bool {
	// Check two positions to see if the line contains a section adornment
	checkLine := func(input string) bool {
		var first, last rune
		for j := 0; j < len(input); j++ {
			r, _ := utf8.DecodeRuneInString(input[j:])
			if unicode.IsSpace(r) {
				l.log.Debug("isSeciton: Skipping space rune")
				continue
			}
			if first == '\x00' {
				first = r
				last = r
			}
			// Log.Debugf("first: %q, last: %q, r: %q, j: %d", first, last, r, j)
			if !isSectionAdornment(r) || (r != first && last != first) {
				l.log.Debug("isSeciton: Section not found")
				return false
			}
			last = r
		}
		return true
	}

	if isTransition(l) {
		l.log.Debug("isSeciton: Returning (found transition)")
		return false
	}

	l.log.Debug("isSeciton: Checking current line")
	if checkLine(l.currentLine()) {
		l.log.Debug("isSeciton: Found section adornment")
		return true
	}

	l.log.Debug("isSeciton: Checking next line")

	nLine := l.peekNextLine()
	if nLine != "" {
		if checkLine(nLine) {
			l.log.Debug("isSeciton: Found section adornment (nextline)")
			return true
		}
	}
	l.log.Debug("isSeciton: Section not found")
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
	if r := l.peek(4); !isSectionAdornment(l.mark) || !isSectionAdornment(r) {
		l.log.Debug("isTransition: Transition not found")
		return false
	}
	pBlankLine := l.lastItem != nil && l.lastItem.Type == itemBlankLine
	nBlankLine := l.peekNextLine() == ""
	if l.line == 0 && nBlankLine {
		l.log.Debug("isTransition: Found transition (followed by newline)")
		return true
	} else if pBlankLine && nBlankLine {
		l.log.Debug("isTransition: Found transition (surrounded by newlines)")
		return true
	}
	l.log.Debug("isTransition: Transition not found")
	return false
}

func isComment(l *lexer) bool {
	if l.lastItem != nil && l.lastItem.Type == itemTitle {
		return false
	}
	if nMark := l.peek(1); l.mark == '.' && nMark == '.' {
		nMark2 := l.peek(2)
		if unicode.IsSpace(nMark2) || nMark2 == utf8.RuneError {
			l.log.Debug("isComment: Found comment!")
			return true
		}
	}
	l.log.Debug("isComment: Comment not found!")
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
					l.log.Debug("isComment: Found arabic enum list!")
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
	for _, x := range bullets {
		if l.mark == x && l.peek(1) == ' ' {
			l.log.Debug("isBulletList: A bullet was found")
			return true
		}
	}
	l.log.Debug("isBulletList: A bullet was not found")
	return false
}

func isDefinitionTerm(l *lexer) bool {
	// Definition terms are preceded by a blankline
	if l.line != 0 && !l.lastLineIsBlankLine() {
		l.log.Debug("isDefinitionTerm: Not definition, lastLineIsBlankLine == false")
		return false
	}
	nL := l.peekNextLine()
	sCount := 0
	for {
		if sCount < len(nL) && unicode.IsSpace(rune(nL[sCount])) {
			sCount++
		} else {
			break
		}
	}
	l.log.WithField("sCount", sCount).Debug("isDefinitionTerm: section count")
	if sCount >= 2 {
		l.log.Debug("isDefinitionTerm: Found definition term!")
		return true
	}
	l.log.Debug("isDefinitionTerm: Did not find definition term.")
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
		if unicode.In(r, unicode.Pd, unicode.Po, unicode.Pi, unicode.Pf, unicode.Ps, unicode.Zs, unicode.Zl) {
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
		var b rune
		for x := 1; x != 3; x++ {
			b = l.peekBack(x)
			if l.mark != b {
				break
			}
		}
		f := l.peek(1)
		if l.mark == f {
			f = l.peek(2)
		}
		// l.log.Debugf("back: %q forward: %q", b, f)
		if b != '\\' && (isOpenerRune(b) || l.start == l.index) && !isSurrounded(b, f) &&
			!unicode.IsSpace(f) && f != utf8.RuneError {
			l.log.Debug("isInlineMarkup: Found inline markup!")
			return true
		}
	}
	return false
}

func isInlineMarkupClosed(l *lexer, markup string) bool {
	isEndASCII := func(r rune) bool {
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

	// A valid end string is made up of one of the following items, notice unicode.Po is troublesome with '*' (emphasis
	// and strong) runes. Special logic is needed in these cases (below).
	validEnd := (!unicode.IsSpace(b) && (unicode.IsSpace(a) || isEndASCII(a) || unicode.In(a, unicode.Pd, unicode.Po,
		unicode.Pi, unicode.Pf, unicode.Pe, unicode.Ps) || a == utf8.RuneError))

	// If the closing markup is two runes, such as '**', make sure the next rune is not '*' and the rune after that is
	// not '*'. The spec is completely silent on this, (and somewhat confusing), but it is clearly how the ref compiler
	// works.
	validNext := (len(markup) == 1 && l.mark != l.peek(1) || len(markup) > 1 && l.mark == l.peek(1) && l.mark !=
		l.peek(2))

	// If the closing markup is one rune, then do nothing.
	if validEnd && validNext {
		l.log.Debug("isInlineMarkupClosed: Found inline markup close")
		return true
	}

	l.log.Debug("isInlineMarkupClosed: Inline markup close not found")
	return false
}

func isEscaped(l *lexer) bool {
	// l.log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d", l.mark, l.index, l.width, l.lineNumber())
	return (l.mark == '\\' && (unicode.In(l.peek(1), unicode.Zs, unicode.Cc, unicode.Lu, unicode.Ll) || l.peek(1) ==
		utf8.RuneError))
}

// lexStart is the first stateFn called by run(). From here other stateFn's are called depending on the input. When this
// function returns nil, the lexing is finished and run() will exit.
func lexStart(l *lexer) stateFn {
	for {
		// l.log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width: %d, l.line: %d", l.mark, l.index, l.start,
		// l.width, l.lineNumber())
		if l.index-l.start <= l.width && l.width > 0 &&
			!l.isEndOfLine() {
			if l.index == 0 && l.mark != ' ' {
				l.indentLevel = 0
				l.indentWidth = ""
			}
			l.log.WithFields(log.Fields{"mark": fmt.Sprintf("%#U", l.mark), "start": l.start, "index": l.index,
				"width": l.width, "line": l.lineNumber()}).Debug("lexStart: mark")
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
			} else if unicode.IsSpace(l.mark) {
				return lexSpace
			} else if isBlockquote(l) {
				return lexBlockquote
			} else if isDefinitionTerm(l) {
				return lexDefinitionTerm
			} else {
				return lexParagraph
			}
		} else if l.isEndOfLine() {
			l.log.Debug("lexStart: isEndOfLine == true")
			if l.start == l.index {
				if l.start == 0 && len(l.currentLine()) == 0 {
					l.log.Debug("lexStart: Found blank line")
					l.emit(itemBlankLine)
					if l.isLastLine() {
						break
					}
				} else if l.isLastLine() {
					l.log.Debug("lexStart: Found end of last line")
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

// lexSpace consumes space characters (space and tab) in the input and emits a itemSpace token.
func lexSpace(l *lexer) stateFn {
	l.log.WithField("l.mark", l.mark).Debug("lexSpace: mark")
	for unicode.IsSpace(l.mark) {
		l.log.WithField("isSpace", unicode.IsSpace(l.mark)).Debug("lexSpace: found space rune")
		if r := l.peek(1); unicode.IsSpace(r) {
			l.next()
		} else {
			l.log.Debug("lexSpace: Next mark is not space!")
			l.next()
			break
		}
	}
	l.log.WithFields(log.Fields{"start": l.start, "index": l.index}).Debugf("lexSpace")
	if l.start < l.index {
		l.emit(itemSpace)
	}
	return lexStart
}

// lexSection is used after isSection() has determined that the next runes of input are section.  From here, the lexTitle()
// and lexSectionAdornment() are called based on the input.
func lexSection(l *lexer) stateFn {
	// l.log.Debugf("l.mark: %#U, l.index: %d, l.start: %d, l.width: %d, " + "l.line: %d", l.mark, l.index, l.start,
	// l.width, l.lineNumber())
	if isSectionAdornment(l.mark) {
		if l.lastItem != nil && l.lastItem.Type != itemTitle {
			return lexSectionAdornment
		}
		lexSectionAdornment(l)
	} else if unicode.IsSpace(l.mark) {
		return lexSpace
	} else if l.mark == utf8.RuneError {
		l.next()
	} else if unicode.IsPrint(l.mark) {
		return lexTitle
	}
	return lexStart
}

// lexTitle consumes input until newline and emits an itemTitle token. If spaces are detected at the start of the line, an
// itemSpace is emitted. Spaces after the title (and before newline) are ignored. On completion control is returned to
// lexSection.
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

// lexSectionAdornment advances the lexer until a newline is encountered and emits a itemSectionAdornment token. Control is
// returned to lexSection() on completion.
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
	for {
		// l.log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d", l.mark, l.index, l.width, l.lineNumber())
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
			}
			l.emit(itemParagraph)
			break
		}
		l.next()
	}
	l.nextLine()
	return lexStart
}

func lexComment(l *lexer) stateFn {
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
	l.log.WithField("line", l.currentLine()).Debugf("lexDefinitionTerm: current line")
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
		l.log.WithFields(log.Fields{"mark": fmt.Sprintf("%#U", l.mark), "start": l.start, "index": l.index,
			"width": l.width, "line": l.lineNumber()}).Debug("lexInlineMarkup: mark")
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
	l.next()
	l.next()
	l.emit(itemInlineStrongOpen)
	for {
		l.next()
		if l.peekBack(1) != '\\' && l.mark == '*' && isInlineMarkupClosed(l, "**") {
			l.log.Debug("lexInlineStrong: Found strong close")
			l.emit(itemInlineStrong)
			break
		} else if l.isEndOfLine() && l.mark == utf8.RuneError {
			if l.peekNextLine() == "" {
				l.log.Debug("lexInlineStrong: Found EOF (unclosed strong)")
				l.emit(itemInlineStrong)
				return lexStart
			}
			l.log.Debug("lexInlineStrong: Found end-of-line")
			l.emit(itemInlineStrong)
			l.emit(itemBlankLine)
			l.nextLine()
		}
	}
	l.next()
	l.next()
	l.emit(itemInlineStrongClose)
	return lexStart
}

func lexInlineEmphasis(l *lexer) stateFn {
	l.next()
	l.emit(itemInlineEmphasisOpen)
	for {
		l.next()
		if l.peekBack(1) != '\\' && l.mark == '*' && isInlineMarkupClosed(l, "*") {
			l.log.Debug("lexInlineEmphasis: Found emphasis close")
			l.emit(itemInlineEmphasis)
			break
		} else if l.isEndOfLine() && l.mark == utf8.RuneError {
			if l.peekNextLine() == "" {
				l.log.Debug("lexInlineEmphasis: Found EOF (unclosed emphasis)")
				l.emit(itemInlineEmphasis)
				return lexStart
			}
			l.log.Debug("lexInlineEmphasis: Found end-of-line")
			l.emit(itemInlineEmphasis)
			l.emit(itemBlankLine)
			l.nextLine()
		}
	}
	l.next()
	l.emit(itemInlineEmphasisClose)
	return lexStart
}

func lexEscape(l *lexer) stateFn {
	if l.peek(1) == utf8.RuneError {
		// Ignore escape sequences at the end of a line
		l.log.Debug("lexEscape: Found escape at end of line, ignoring")
		l.skip(2)
		return lexStart
	}
	l.next()
	l.next()
	l.emit(itemEscape)
	return lexStart
}

func lexInlineLiteral(l *lexer) stateFn {
	l.next()
	l.next()
	l.emit(itemInlineLiteralOpen)
	for {
		l.next()
		if l.mark == '`' && isInlineMarkupClosed(l, "``") {
			l.log.Debug("lexInlineLiteral: Found literal close")
			l.emit(itemInlineLiteral)
			break
		} else if l.isEndOfLine() && l.mark == utf8.RuneError {
			if l.peekNextLine() == "" {
				l.log.Debug("lexInlineLiteral: Found EOF (unclosed inline literal)")
				l.emit(itemInlineLiteral)
				return lexStart
			}
			l.log.Debug("lexInlineLiteral: Found end-of-line")
			l.emit(itemInlineLiteral)
			l.emit(itemBlankLine)
			l.nextLine()
		}
	}
	l.next()
	l.next()
	l.emit(itemInlineLiteralClose)
	return lexStart
}

func lexInlineInterpretedText(l *lexer) stateFn {
	l.next()
	l.emit(itemInlineInterpretedTextOpen)
	for {
		l.next()
		if l.mark == '`' && isInlineMarkupClosed(l, "`") {
			l.log.Debug("lexInlineInterpretedText: Found literal close")
			l.emit(itemInlineInterpretedText)
			break
		}
	}
	l.next()
	l.emit(itemInlineInterpretedTextClose)
	if l.mark == ':' {
		lexInlineInterpretedTextRole(l)
	}
	return lexStart
}

func lexInlineInterpretedTextRole(l *lexer) stateFn {
	l.next()
	l.emit(itemInlineInterpretedTextRoleOpen)
	for {
		l.next()
		if l.mark == ':' {
			l.emit(itemInlineInterpretedTextRole)
			break
		}
	}
	l.next()
	l.emit(itemInlineInterpretedTextRoleClose)
	return lexStart
}
