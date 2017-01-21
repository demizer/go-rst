package parse

import "strconv"

// Function prototype for scanner functions
type stateFn func(*lexer) stateFn

// ID is a consecutive number for identication of a lexed item and parsed item. Primarily for the purpose of debugging lexer
// and parser output when compared to the JSON encoded tests.
type ID int

// IDNumber returns the ID from an item.
func (i ID) IDNumber() ID { return i }

// String implements Stringer and returns ID as a string.
func (i ID) String() string { return strconv.Itoa(int(i)) }

// Line contains the number of a lexed item, or parsed item, from the input data.
type Line int

// LineNumber returns the Line of an item.
func (l Line) LineNumber() Line { return l }

// String implements Stringer and returns Line converted to a string.
func (l Line) String() string { return strconv.Itoa(int(l)) }

// StartPosition is the starting location of an item in the line of input.
type StartPosition int

// String implements Stringer and returns StartPosition converted to a string.
func (s StartPosition) String() string { return strconv.Itoa(int(s)) }

// Int return the StartPosition as an integer value.
func (s StartPosition) Int() int { return int(s) }

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
