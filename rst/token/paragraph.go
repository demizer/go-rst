package token

func lexText(l *Lexer) stateFn {
	log.Msg("lexText start")
	for {
		// log.Log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d", l.mark, l.index, l.width, l.lineNumber())
		if isEscaped(l) {
			l.emit(ItemText)
			lexEscape(l)
		}
		if isInlineMarkup(l) {
			log.Msg("FOUND inline reference!")
			if l.index > l.start {
				l.emit(ItemText)
			}
			lexInlineMarkup(l)
			if isEscaped(l) {
				lexEscape(l)
			}
			continue
		} else if isInlineReference(l) {
			log.Msg("FOUND inline reference!")
			lexInlineReference(l)
			continue
		}
		if l.isEndOfLine() && l.mark == EOL {
			if l.start == l.index {
				return lexStart
			}
			l.emit(ItemText)
			// if !l.isLastLine() {
			// l.emit(ItemSpace) // We hit a "newline", which is converted to a space when in a paragraph
			// }
			break
		}
		l.next()
	}
	l.nextLine()
	return lexStart
}
