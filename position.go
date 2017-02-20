package rst

import "strconv"

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
