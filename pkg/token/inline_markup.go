package token

import (
	"fmt"
	"unicode"
)

func isInlineMarkup(l *Lexer) bool {
	isOpenerRune := func(r rune) bool {
		for _, x := range inlineMarkupStartStringOpeners {
			if x == r {
				l.Msg("Found inline markup!")
				return true
			}
		}
		if unicode.In(r, unicode.Pd, unicode.Po, unicode.Pi, unicode.Pf, unicode.Ps, unicode.Zs, unicode.Zl) {
			l.Msg("Found inline markup!")
			return true
		}
		l.Msg("Not inline markup!")
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
		} else if unicode.In(back, unicode.Ps) && unicode.In(front, unicode.Pe, unicode.Pf, unicode.Pi) {
			return true
		} else if unicode.In(back, unicode.Pi) && unicode.In(front, unicode.Pf, unicode.Ps) {
			return true
		} else if unicode.In(back, unicode.Pf) && unicode.In(front, unicode.Pf) {
			return true
		} else if unicode.In(back, unicode.Pf) && unicode.In(front, unicode.Pi) {
			return true
		}
		return false
	}
	if l.mark == '*' || l.mark == '`' {
		var b rune
		// Look back up to three runes looking for * or `
		for x := 1; x != 3; x++ {
			b = l.peekBack(x)
			if l.mark != b {
				break
			}
		}
		// Get the next rune, if the rune is * or ` then get the rune after that
		f := l.peek(1)
		if l.mark == f {
			f = l.peek(2)
		}
		if b != '\\' && (isOpenerRune(b) || l.start == l.index) && !isSurrounded(b, f) &&
			!unicode.IsSpace(f) && f != EOL {
			l.Msg("Found inline markup!")
			return true
		}
	}
	l.Msg("Not inline markup!")
	return false
}

func isInlineMarkupClosed(l *Lexer, markup string) bool {
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
		unicode.Pi, unicode.Pf, unicode.Pe, unicode.Ps) || a == EOL))

	// If the closing markup is two runes, such as '**', make sure the next rune is not '*' and the rune after that is
	// not '*'. The spec is completely silent on this, (and somewhat confusing), but it is clearly how the ref compiler
	// works.
	validNext := (len(markup) == 1 && l.mark != l.peek(1) || len(markup) > 1 && l.mark == l.peek(1) && l.mark !=
		l.peek(2))

	// If the closing markup is one rune, then do nothing.
	if validEnd && validNext {
		l.Msg("Found inline markup close")
		return true
	}

	l.Msg("Inline markup close not found")
	return false
}

func isInlineReference(l *Lexer) bool {
	isNotSurroundedByUnderscores := l.peekBack(1) != '_' && l.peek(1) != '_'
	lastItemIsNotSpace := l.lastItem == nil || l.lastItem.Type != Space
	isAnon := l.lastItem != nil && l.lastItem.Type == BlankLine && l.mark == '_' && l.peek(1) == '_'

	isQuotedAnon := func() bool {
		x := l.index
		if l.mark != '`' {
			return false
		}
		// Check for end quote
		for {
			lp := l.peek(x)
			lp2 := l.peek(x + 1)
			if lp == '`' {
				if l.peek(x+1) == '_' {
					l.Msg("FOUND quoted inline anonymous hyperlink reference!")
					return true
				}
			} else if lp == EOL && lp2 == EOL {
				l.Msg("FOUND blank line")
				break
			}
			x++
		}
		l.Msg("NOT FOUND quoted inline anonymous hyperlink reference")
		return false
	}

	if l.mark == '_' && isNotSurroundedByUnderscores && lastItemIsNotSpace {
		l.Msg("FOUND inlineReference!")
		return true
	} else if isAnon || isQuotedAnon() {
		l.Msg("FOUND anonymous inlineReference!")
		return true
	}

	l.Msg("NOT FOUND isInlineReference")
	return false
}

func lexInlineMarkup(l *Lexer) stateFn {
	for {
		l.Log("mark", fmt.Sprintf("%#U", l.mark), "start", l.start, "index", l.index,
			"width", l.width, "line", l.lineNumber())
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
			if isInlineReference(l) {
				lexInlineReference(l)
				break
			}
			lexInlineInterpretedText(l)
			break
		}
	}
	return lexStart
}

func lexInlineStrong(l *Lexer) stateFn {
	// Log.funcName = "lexInlineStrong"
	l.next()
	l.next()
	l.emit(InlineStrongOpen)
	for {
		l.next()
		if l.peekBack(1) != '\\' && l.mark == '*' && isInlineMarkupClosed(l, "**") {
			l.Msg("Found strong close")
			l.emit(InlineStrong)
			break
		} else if l.isEndOfLine() && l.mark == EOL {
			if l.peekNextLine() == "" {
				l.Msg("Found EOF (unclosed strong)")
				l.emit(InlineStrong)
				return lexStart
			}
			l.Msg("Found end-of-line")
			l.emit(InlineStrong)
			l.emit(BlankLine)
			l.nextLine()
		}
	}
	l.next()
	l.next()
	l.emit(InlineStrongClose)
	return lexStart
}

func lexInlineEmphasis(l *Lexer) stateFn {
	l.next()
	l.emit(InlineEmphasisOpen)
	for {
		l.next()
		if l.peekBack(1) != '\\' && l.mark == '*' && isInlineMarkupClosed(l, "*") {
			l.Msg("Found emphasis close")
			l.emit(InlineEmphasis)
			break
		} else if l.isEndOfLine() && l.mark == EOL {
			if l.peekNextLine() == "" {
				l.Msg("Found EOF (unclosed emphasis)")
				l.emit(InlineEmphasis)
				return lexStart
			}
			l.Msg("Found end-of-line")
			l.emit(InlineEmphasis)
			l.emit(BlankLine)
			l.nextLine()
		}
	}
	l.next()
	l.emit(InlineEmphasisClose)
	return lexStart
}

func lexInlineLiteral(l *Lexer) stateFn {
	l.next()
	l.next()
	l.emit(InlineLiteralOpen)
	for {
		l.next()
		if l.mark == '`' && isInlineMarkupClosed(l, "``") {
			l.Msg("Found literal close")
			l.emit(InlineLiteral)
			break
		} else if l.isEndOfLine() && l.mark == EOL {
			if l.peekNextLine() == "" {
				l.Msg("Found EOF (unclosed inline literal)")
				l.emit(InlineLiteral)
				return lexStart
			}
			l.Msg("Found end-of-line")
			l.emit(InlineLiteral)
			l.emit(BlankLine)
			l.nextLine()
		}
	}
	l.next()
	l.next()
	l.emit(InlineLiteralClose)
	return lexStart
}

func lexInlineInterpretedText(l *Lexer) stateFn {
	l.next()
	l.emit(InlineInterpretedTextOpen)
	for {
		l.next()
		if l.mark == '`' && isInlineMarkupClosed(l, "`") {
			l.Msg("Found literal close")
			l.emit(InlineInterpretedText)
			break
		}
	}
	l.next()
	l.emit(InlineInterpretedTextClose)
	if l.mark == ':' {
		lexInlineInterpretedTextRole(l)
	}
	return lexStart
}

func lexInlineInterpretedTextRole(l *Lexer) stateFn {
	l.next()
	l.emit(InlineInterpretedTextRoleOpen)
	for {
		l.next()
		if l.mark == ':' {
			l.emit(InlineInterpretedTextRole)
			break
		}
	}
	l.next()
	l.emit(InlineInterpretedTextRoleClose)
	return lexStart
}

func lexInlineReference(l *Lexer) stateFn {
	if l.mark == '`' {
		l.next()
		l.emit(InlineReferenceOpen)
		for {
			l.next()
			if l.mark == '`' {
				l.emit(InlineReferenceText)
				l.next()
				break
			} else if l.start == l.index && l.mark == ' ' && l.peek(1) != ' ' {
				lexSpace(l)
				continue
			} else if l.mark == EOL && l.peek(1) != EOL {
				l.emit(InlineReferenceText)
				l.next()
			} else if l.mark == EOL && l.peek(1) == EOL {
				break
			}
		}
		if lp := l.peek(1); l.mark == '_' && lp == '_' {
			l.next()
			l.next()
		}
		l.emit(InlineReferenceClose)
		return lexStart
	}
	l.emit(InlineReferenceText)
	l.next()
	if l.mark == '_' {
		// Anonymous hyperlink reference
		l.next()
	}
	l.emit(InlineReferenceClose)
	return lexStart
}
