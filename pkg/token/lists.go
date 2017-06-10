package token

import "unicode"

// isArabic returns true if rune r is an Arabic numeral.
func isArabic(r rune) bool {
	return r > '0' && r < '9'
}

func isEnumList(l *Lexer) (ret bool) {
	bCount := 0
	if isSection(l) {
		goto exit
	}
	if isArabic(l.mark) {
		for {
			bCount++
			if nMark, _ := l.next(); !isArabic(nMark) {
				if nMark == '.' || nMark == ' ' {
					l.Msg("Found arabic enum list!")
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

func isBulletList(l *Lexer) bool {
	for _, x := range bullets {
		if l.mark == x && l.peek(1) == ' ' {
			l.Msg("A bullet was found")
			return true
		}
	}
	l.Msg("A bullet was not found")
	return false
}

func isDefinitionTerm(l *Lexer) bool {
	// Definition terms are preceded by a blankline
	if l.line != 0 && !l.lastLineIsBlankLine() {
		l.Msg("Not definition, lastLineIsBlankLine == false")
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
	l.Msgr("Section count", "sCount", sCount)
	if sCount >= 2 {
		l.Msg("FOUND definition term!")
		return true
	}
	l.Msg("NOT FOUND definition term.")
	return false
}

func lexEnumList(l *Lexer) stateFn {
	if isArabic(l.mark) {
		for {
			if nMark, _ := l.next(); !isArabic(nMark) {
				l.emit(EnumListArabic)
				l.next()
				if nMark == '.' {
					l.emit(EnumListAffix)
					l.next()
				}
				l.emit(Space)
				break
			}
		}
	}
	return lexStart
}

func lexDefinitionTerm(l *Lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() && l.mark == EOL {
			l.emit(DefinitionTerm)
			break
		}
	}
	l.nextLine()
	l.next()
	l.Msgr("Current line", "line", l.currentLine())
	lexSpace(l)
	for {
		l.next()
		if l.isEndOfLine() && l.mark == EOL {
			l.emit(DefinitionText)
			break
		}
	}
	return lexStart
}

func lexBullet(l *Lexer) stateFn {
	l.next()
	l.emit(Bullet)
	lexSpace(l)
	l.indentWidth += l.lastItem.Text + " "
	lexText(l)
	l.indentLevel++
	return lexStart
}
