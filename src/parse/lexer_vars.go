package parse

import "unicode/utf8"

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
