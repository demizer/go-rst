package token

import (
	"errors"
	"fmt"
	"strings"
	"unicode"
	"unicode/utf8"

	. "github.com/demizer/go-rst"
)

// Function prototype for scanner functions
type stateFn func(*Lexer) stateFn

// The Lexer struct tracks the state of the Lexer
type Lexer struct {
	Name    string   // The name of the current lexer
	input   string   // The input text
	line    int      // Line number of the parser, from 0
	numLine int      // Total number of input lines
	state   stateFn  // The current state of the lexer
	lines   []string // The input split into lines

	id    int  // Unique ID for each item emitted
	mark  rune // The current lexed rune
	start int  // Start position of the token in the line
	index int  // Position in input
	width int  // The width of the current position

	items            chan Item // The channel items are emitted to
	lastItem         *Item     // The last item emitted to the channel
	lastItemPosition *StartPosition

	indentLevel int    // For tracking indentation with indentable items
	indentWidth string // For tracking indent width
}

func newLexer(name string, input []byte) (l *Lexer, err error) {
	if len(input) == 0 {
		err = errors.New("no input given")
		return
	}

	l = &Lexer{Name: name}

	ni, err := normalize(input)
	if err != nil {
		err = fmt.Errorf("could not normalize input: %s", err)
		return
	}

	lines := strings.Split(string(ni), "\n")

	mark, width := utf8.DecodeRuneInString(lines[0][0:])
	log.Log("mark", mark, "index", 0, "line", 1)

	l.input = string(ni) // stored string is never altered
	l.lines = lines
	l.items = make(chan Item)
	l.index = 0
	l.mark = mark
	l.width = width

	return
}

// lex is the entry point of the lexer. Name should be any name that signifies the purporse of the lexer. It is mostly used
// to identify the lexing process in debugging.
func Lex(name string, input []byte) (l *Lexer, err error) {
	l, err = newLexer(name, input)
	if err != nil {
		return
	}
	go l.run()
	return
}

// run is the engine of the lexing process.
func (l *Lexer) run() {
	for l.state = lexStart; l.state != nil; {
		l.state = l.state(l)
	}
}

// emit passes an item back to the client.
func (l *Lexer) emit(t Type) {
	var tok string

	if t == BlankLine {
		tok = "\n"
	} else if t == Space && l.start == l.index {
		// For linebreaks and vertical tabs at the end of the line in a paragraph
		tok = " "
	} else if t == EOF {
		tok = ""
	} else {
		tok = l.lines[l.line][l.start:l.index]
	}

	l.id++
	length := utf8.RuneCountInString(tok)

	nItem := Item{
		ID:   ID(l.id),
		Type: t,
		Text: tok,
		Line: Line(l.lineNumber()),
		// +1 because positions begin at 1, not 0
		StartPosition: StartPosition(l.start + 1),
		Length:        length,
	}

	log.Log("ID", ID(l.id)+1, t.String(), fmt.Sprintf("%q", tok), "l.start+1", l.start+1, "l.index",
		l.index, "line", l.lineNumber())

	l.items <- nItem
	l.lastItem = &nItem
	l.start = l.index
	log.Log("msg", "Position after EMIT", "l.mark", fmt.Sprintf("%q", l.mark), "l.start", l.start,
		"l.index", l.index, "line", l.lineNumber())
}

// backup backs up the lexer position by a number of rune positions (pos).  backup cannot backup off the input, in that case
// the index of the lexer is set to the starting position on the input. The run
func (l *Lexer) backup(pos int) {
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
		if r == EOL && lLen != 0 && lLen != l.index {
			l.backup(1)
		}
	}
	// log.Log("mark", l.mark)
}

// peek looks ahead in the input by a number of locations (locs) and returns the rune at that location in the input. Peek
// works across lines.
func (l *Lexer) peek(locs int) rune {
	pos := saveLexerPosition(l)
	var r rune
	x := 0
	// FIXME: NO NEED FOR FOR LOOP HERE
	for x < locs {
		l.next()
		if x == locs-1 {
			r = l.mark
		}
		x++
	}
	pos.restore(l)
	// log.Log("mark", fmt.Sprintf("%#v", r), "index", l.index)
	return r
}

func (l *Lexer) peekBack(locs int) rune {
	if l.start == l.index {
		return EOL
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
	// log.Log("mark", string(r), "index", l.index)
	return r
}

func (l *Lexer) peekNextLine() string {
	if l.isLastLine() {
		return ""
	}
	return l.lines[l.line+1]
}

// next advances the position of the lexer by one rune and returns that rune.
func (l *Lexer) next() (rune, int) {
	if l.isEndOfLine() && !l.isLastLine() {
		l.nextLine()
	}
	l.index += l.width
	r, width := utf8.DecodeRuneInString(l.currentLine()[l.index:])
	l.width = width
	l.mark = r
	// log.Log("mark", fmt.Sprintf("%#U", r), "start", l.start, "index", l.index, "line", l.lineNumber())
	return r, width
}

func (l *Lexer) nextLine() string {
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
func (l *Lexer) NextItem() *Item {
	item, ok := <-l.items
	if ok == false {
		return nil
	}
	l.lastItemPosition = &item.StartPosition
	return &item

}

func (l *Lexer) skip(locs int) {
	for x := 1; x <= locs; x++ {
		l.next()
	}
	l.start = l.index
}

// gotoLine advances the lexer to a line and index within that line. Line numbers start at 1.
func (l *Lexer) gotoLocation(start, line int) {
	l.line = line - 1
	l.index = start
	r, width := utf8.DecodeRuneInString(l.currentLine()[l.index:])
	l.width = width
	l.mark = r
	return
}

func (l *Lexer) currentLine() string {
	return l.lines[l.line]
}

func (l *Lexer) lineNumber() int {
	return l.line + 1
}

func (l *Lexer) isLastLine() bool {
	return len(l.lines) == l.lineNumber()
}

func (l *Lexer) lastLineIsBlankLine() bool {
	if l.line == 0 {
		return false
	}
	m, _ := utf8.DecodeRuneInString(l.lines[l.line-1])
	if m == EOL {
		return true
	}
	return false
}

func (l *Lexer) isEndOfLine() bool {
	return len(l.lines[l.line]) == l.index
}

// lexStart is the first stateFn called by run(). From here other stateFn's are called depending on the input. When this
// function returns nil, the lexing is finished and run() will exit.
func lexStart(l *Lexer) stateFn {
	for {
		if l.index == 0 && l.start == 0 {
			log.Log("msg", "lexing line", "text", l.currentLine(), "line", l.lineNumber())
		}
		if l.index-l.start <= l.width && l.width > 0 && !l.isEndOfLine() {
			if l.index == 0 && l.mark != ' ' {
				l.indentLevel = 0
				l.indentWidth = ""
			}
			log.Log("mark", fmt.Sprintf("%#U", l.mark), "start", l.start, "index", l.index,
				"width", l.width, "line", l.lineNumber())
			if isComment(l) {
				return lexComment
			} else if isHyperlinkTarget(l) {
				return lexHyperlinkTarget
			} else if isInlineReference(l) {
				return lexInlineReference
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
				return lexText
			}
		} else if l.isEndOfLine() {
			log.Msg("isEndOfLine == true")
			if l.start == l.index {
				if l.start == 0 && len(l.currentLine()) == 0 {
					log.Msg("Found blank line")
					l.emit(BlankLine)
					if l.isLastLine() {
						break
					}
				} else if l.isLastLine() {
					log.Msg("Found end of last line")
					break
				}
			}
		}
		l.next()
	}

	l.emit(EOF)
	close(l.items)
	return nil
}

// lexSpace consumes space characters (space and tab) in the input and emits a Space token.
func lexSpace(l *Lexer) stateFn {
	log.Log("l.mark", l.mark)
	for unicode.IsSpace(l.mark) {
		log.Log("msg", "found space rune", "isSpace", unicode.IsSpace(l.mark))
		if r := l.peek(1); unicode.IsSpace(r) {
			l.next()
		} else {
			log.Msg("Next mark is not space!")
			l.next()
			break
		}
	}
	log.Log("start", l.start, "index", l.index)
	if l.start < l.index {
		l.emit(Space)
	}
	return lexStart
}

func lexEscape(l *Lexer) stateFn {
	l.next()
	l.emit(Escape)
	if unicode.IsSpace(l.mark) {
		lexSpace(l)
	}
	return lexStart
}
