package tokenizer

import "unicode"

func isHyperlinkTarget(l *lexer) bool {
	nMark := l.peek(1)
	nMark2 := l.peek(2)
	if l.mark == '.' && nMark == '.' && nMark2 != EOL {
		nMark3 := l.peek(3)
		if unicode.IsSpace(nMark2) && nMark3 == '_' {
			if isReferenceNameSimple(l, 4) {
				logl.Msg("FOUND hyperlink simple target")
				return true
			} else if isReferenceNamePhrase(l, 4) {
				logl.Msg("FOUND hyperlink phrase target")
				return true
			}
			logl.Msg("FOUND malformed hyperlink target")
			return true
		}
	} else if l.mark == '_' && nMark == '_' && unicode.IsSpace(nMark2) {
		logl.Msg("FOUND anonymous hyperlink target")
		return true
	}
	logl.Msg("NOT FOUND Hyperlink target")
	return false
}

func lexHyperlinkTarget(l *lexer) stateFn {
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
		l.emit(ItemHyperlinkTargetSuffix)
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

func lexAnonymousHyperlinkTarget(l *lexer) stateFn {
	// l.emit(ItemHyperlinkTargetStart)
	lexHyperlinkTargetStart(l)
	if l.mark == '_' {
		lexHyperlinkTargetPrefix(l)
	}
	// lexSpace(l)
	// lp := l.peek(2)
	if l.mark == ':' {
		// l.next()
		// l.next()
		// l.emit(ItemHyperlinkTargetPrefix)
		// lexHyperlinkTargetPrefix(l)
		// l.next()
		// l.emit(ItemHyperlinkTargetSuffix)
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

func lexHyperlinkTargetName(l *lexer) stateFn {
	var inquote bool
	for {
		if l.mark == '`' {
			if !inquote {
				inquote = true
				l.next()
				l.emit(ItemHyperlinkTargetQuote)
				l.next()
			} else {
				l.emit(ItemHyperlinkTargetName)
				l.next()
				l.emit(ItemHyperlinkTargetQuote)
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
				l.emit(ItemHyperlinkTargetName)
			}
			break
		} else if unicode.IsSpace(l.mark) && (lp != EOL && unicode.IsSpace(lp)) {
			lexSpace(l)
		} else if l.mark == EOL && !unicode.IsSpace(lp) {
			l.emit(ItemHyperlinkTargetName)
			break
		} else if l.mark == EOL {
			// hyperlink target name is multi-line
			l.emit(ItemHyperlinkTargetName)
		}
		l.next()
	}
	return lexStart
}

func lexHyperlinkTargetBlock(l *lexer) stateFn {
	var inquote bool
	for {
		if l.mark == '`' {
			if !inquote {
				inquote = true
				l.next()
				l.emit(ItemInlineReferenceOpen)
				l.next()
			} else {
				l.emit(ItemInlineReferenceText)
				l.next()
				l.next()
				l.emit(ItemInlineReferenceClose)
				break
			}
			continue
		}
		lb := l.peekBack(1)
		lp := l.peek(1)
		// First check for indirect reference
		if lb != '\\' && l.mark == '_' && lp == EOL {
			l.emit(ItemInlineReferenceText)
			l.next()
			l.emit(ItemInlineReferenceClose)
			break
		} else if !inquote && l.mark == EOL {
			// end of current line
			l.emit(ItemHyperlinkTargetURI)
			if lp == EOL {
				break
			}
			// uri continues on next line
			l.next()
			lexSpace(l)
		} else if inquote && l.lastItem.Type == ItemInlineReferenceOpen && l.mark == EOL {
			// end of current line, reference continues on next line
			l.emit(ItemInlineReferenceText)
			l.next()
			lexSpace(l)
		}
		l.next()
	}
	return lexStart
}

func lexAnonymousHyperlinkTargetBlock(l *lexer) stateFn {
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
				l.emit(ItemInlineReferenceOpen)
				l.next()
			} else {
				l.emit(ItemInlineReferenceText)
				l.next()
				l.next()
				l.emit(ItemInlineReferenceClose)
				break
			}
			continue
		}
		lb := l.peekBack(1)
		lp := l.peek(1)
		if !containsSpaces && lb != '\\' && l.mark == '_' && lp == EOL {
			l.emit(ItemInlineReferenceText)
			l.next()
			l.emit(ItemInlineReferenceClose)
			break
		} else if l.mark == EOL {
			// end of current line
			l.emit(ItemHyperlinkTargetURI)
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

func lexHyperlinkTargetStart(l *lexer) stateFn {
	for {
		if l.mark != '.' {
			break
		}
		l.next()
	}
	l.emit(ItemHyperlinkTargetStart)
	lexSpace(l)
	return lexStart
}

func lexHyperlinkTargetPrefix(l *lexer) stateFn {
	for {
		if l.mark != '_' {
			break
		}
		l.next()
	}
	l.emit(ItemHyperlinkTargetPrefix)
	return lexStart
}

func lexHyperlinkTargetSuffix(l *lexer) stateFn {
	for {
		if l.mark != ':' {
			break
		}
		l.next()
	}
	l.emit(ItemHyperlinkTargetSuffix)
	return lexStart
}
