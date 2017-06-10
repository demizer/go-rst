package token

import "unicode"

func isHyperlinkTarget(l *Lexer) bool {
	nMark := l.peek(1)
	nMark2 := l.peek(2)
	if l.mark == '.' && nMark == '.' && nMark2 != EOL {
		nMark3 := l.peek(3)
		if unicode.IsSpace(nMark2) && nMark3 == '_' {
			if isReferenceNameSimple(l, 4) {
				l.Msg("FOUND hyperlink simple target")
				return true
			} else if isReferenceNamePhrase(l, 4) {
				l.Msg("FOUND hyperlink phrase target")
				return true
			} else if len(l.currentLine()) > 4 {
				l.Msg("FOUND malformed hyperlink target")
				return true
			}
		}
	} else if l.mark == '_' && nMark == '_' && unicode.IsSpace(nMark2) {
		l.Msg("FOUND anonymous hyperlink target")
		return true
	}
	l.Msg("NOT FOUND Hyperlink target")
	return false
}

func lexHyperlinkTarget(l *Lexer) stateFn {
	var anonstart bool
	for l.mark == '.' || l.mark == '_' {
		if l.mark == '_' {
			anonstart = true
		}
		l.next()
	}
	lp := l.peek(1)
	lp2 := l.peek(2)
	if (unicode.IsSpace(l.mark) && lp == '_' && lp2 == '_') || (anonstart && unicode.IsSpace(l.mark)) {
		lexAnonymousHyperlinkTarget(l)
		return lexStart
	}
	lexHyperlinkTargetStart(l)
	lexHyperlinkTargetPrefix(l)
	lexHyperlinkTargetName(l)
	if l.mark == ':' {
		l.next()
		l.emit(HyperlinkTargetSuffix)
		if unicode.IsSpace(l.mark) && l.index < len(l.currentLine()) {
			lexSpace(l)
			lexHyperlinkTargetBlock(l)
		}
	}
	if lp := l.peek(1); lp != EOL && lp != '\n' && unicode.IsSpace(lp) {
		l.next()
		lexSpace(l)
		lexHyperlinkTargetBlock(l)
	}
	l.next()
	return lexStart
}

func lexAnonymousHyperlinkTarget(l *Lexer) stateFn {
	// l.emit(HyperlinkTargetStart)
	lexHyperlinkTargetStart(l)
	if l.mark == '_' {
		lexHyperlinkTargetPrefix(l)
	}
	// lexSpace(l)
	// lp := l.peek(2)
	if l.mark == ':' {
		// l.next()
		// l.next()
		// l.emit(HyperlinkTargetPrefix)
		// lexHyperlinkTargetPrefix(l)
		// l.next()
		// l.emit(HyperlinkTargetSuffix)
		// lexSpace(l)
		lexHyperlinkTargetSuffix(l)
		// } else if l.mark != ':' {
		// lexHyperlinkTargetName(l)
		// lexHyperlinkTargetSuffix(l)
	}
	lexSpace(l)
	lexAnonymousHyperlinkTargetBlock(l)
	return lexStart
}

func lexHyperlinkTargetName(l *Lexer) stateFn {
	var inquote bool
	for {
		if l.mark == '`' {
			if !inquote {
				inquote = true
				l.next()
				l.emit(HyperlinkTargetQuote)
				l.next()
			} else {
				l.emit(HyperlinkTargetName)
				l.next()
				l.emit(HyperlinkTargetQuote)
				break
			}
			continue
		}
		lb := l.peekBack(1)
		lp := l.peek(1)
		// make sure the : mark is not escaped, i.e., \\:
		if l.mark == ':' && !inquote && lb != '\\' {
			if l.index != l.start {
				// There are runes in the "buffer" that need to be emitted. This is a malformed link
				l.emit(HyperlinkTargetName)
			}
			break
		} else if unicode.IsSpace(l.mark) && (lp != EOL && unicode.IsSpace(lp)) {
			lexSpace(l)
		} else if l.mark == EOL && !unicode.IsSpace(lp) {
			l.emit(HyperlinkTargetName)
			break
		} else if l.mark == EOL {
			// hyperlink target name is multi-line
			l.emit(HyperlinkTargetName)
		}
		l.next()
	}
	return lexStart
}

func lexHyperlinkTargetBlock(l *Lexer) stateFn {
	var inquote bool
	for {
		if l.mark == '`' {
			if !inquote {
				inquote = true
				l.next()
				l.emit(InlineReferenceOpen)
				l.next()
			} else {
				l.emit(InlineReferenceText)
				l.next()
				l.next()
				l.emit(InlineReferenceClose)
				break
			}
			continue
		}
		lb := l.peekBack(1)
		lp := l.peek(1)
		// First check for indirect reference
		if lb != '\\' && l.mark == '_' && lp == EOL {
			l.emit(InlineReferenceText)
			l.next()
			l.emit(InlineReferenceClose)
			break
		} else if !inquote && l.mark == EOL {
			// end of current line
			l.emit(HyperlinkTargetURI)
			if lp == EOL {
				break
			}
			// uri continues on next line
			l.next()
			lexSpace(l)
		} else if inquote && l.lastItem.Type == InlineReferenceOpen && l.mark == EOL {
			// end of current line, reference continues on next line
			l.emit(InlineReferenceText)
			l.next()
			lexSpace(l)
		}
		l.next()
	}
	return lexStart
}

func lexAnonymousHyperlinkTargetBlock(l *Lexer) stateFn {
	var inquote bool
	var containsSpaces bool
	for {
		if unicode.IsSpace(l.mark) {
			containsSpaces = true
		}
		if l.mark == '`' {
			if !inquote {
				inquote = true
				l.next()
				l.emit(InlineReferenceOpen)
				l.next()
			} else {
				l.emit(InlineReferenceText)
				l.next()
				l.next()
				l.emit(InlineReferenceClose)
				break
			}
			continue
		}
		lb := l.peekBack(1)
		lp := l.peek(1)
		if !containsSpaces && lb != '\\' && l.mark == '_' && lp == EOL {
			l.emit(InlineReferenceText)
			l.next()
			l.emit(InlineReferenceClose)
			break
		} else if l.mark == EOL {
			// end of current line
			l.emit(HyperlinkTargetURI)
			if lp == EOL {
				break
			}
			// uri continues on next line
			l.next()
			lexSpace(l)
		}
		l.next()
	}
	return lexStart
}

func lexHyperlinkTargetStart(l *Lexer) stateFn {
	for {
		if l.mark != '.' {
			break
		}
		l.next()
	}
	l.emit(HyperlinkTargetStart)
	lexSpace(l)
	return lexStart
}

func lexHyperlinkTargetPrefix(l *Lexer) stateFn {
	for {
		if l.mark != '_' {
			break
		}
		l.next()
	}
	l.emit(HyperlinkTargetPrefix)
	return lexStart
}

func lexHyperlinkTargetSuffix(l *Lexer) stateFn {
	for {
		if l.mark != ':' {
			break
		}
		l.next()
	}
	l.emit(HyperlinkTargetSuffix)
	return lexStart
}
