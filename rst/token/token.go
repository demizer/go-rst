package token

import (
	"strconv"
	"unicode"
	"unicode/utf8"

	. "github.com/demizer/go-rst"
)

var log *LogContext

func init() {
	log = RegisterNewLogContext("lexer", StdLogger())
}

// EOL is denoted by a utf8.RuneError
var EOL rune = utf8.RuneError

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

// ID is a consecutive number for identication of a lexed item and parsed item. Primarily for the purpose of debugging lexer
// and parser output when compared to the JSON encoded tests.
type ID int

// IDNumber returns the ID from an item.
func (i ID) IDNumber() ID { return i }

// String implements Stringer and returns ID as a string.
func (i ID) String() string { return strconv.Itoa(int(i)) }

type lexPosition struct {
	index int
	start int
	line  int
	mark  rune
	width int
}

func saveLexerPosition(lexState *Lexer) *lexPosition {
	return &lexPosition{
		index: lexState.index,
		start: lexState.start,
		line:  lexState.line,
		mark:  lexState.mark,
		width: lexState.width,
	}
}

func (l *lexPosition) restore(lexState *Lexer) {
	lexState.index = l.index
	lexState.start = l.start
	lexState.line = l.line
	lexState.mark = l.mark
	lexState.width = l.width
}

func isEscaped(l *Lexer) bool {
	return (l.mark == '\\' && (unicode.In(l.peek(1), unicode.Zs, unicode.Cc, unicode.Lu, unicode.Ll) || l.peek(1) ==
		EOL))
}
