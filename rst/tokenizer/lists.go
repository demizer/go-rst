package tokenizer

import "unicode"

// isArabic returns true if rune r is an Arabic numeral.
func isArabic(r rune) bool {
	return r > '0' && r < '9'
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
					logl.Msg("Found arabic enum list!")
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
			logl.Msg("A bullet was found")
			return true
		}
	}
	logl.Msg("A bullet was not found")
	return false
}

func isDefinitionTerm(l *lexer) bool {
	// Definition terms are preceded by a blankline
	if l.line != 0 && !l.lastLineIsBlankLine() {
		logl.Msg("Not definition, lastLineIsBlankLine == false")
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
	logl.Log("msg", "Section count", "sCount", sCount)
	if sCount >= 2 {
		logl.Msg("FOUND definition term!")
		return true
	}
	logl.Msg("NOT FOUND definition term.")
	return false
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

func lexDefinitionTerm(l *lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() && l.mark == EOL {
			l.emit(itemDefinitionTerm)
			break
		}
	}
	l.nextLine()
	l.next()
	logl.Log("msg", "Current line", "line", l.currentLine())
	lexSpace(l)
	for {
		l.next()
		if l.isEndOfLine() && l.mark == EOL {
			l.emit(itemDefinitionText)
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
	lexText(l)
	l.indentLevel++
	return lexStart
}
