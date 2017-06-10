package token

func lexText(l *Lexer) stateFn {
	l.Msg("lexText start")
	for {
		// l.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d", l.mark, l.index, l.width, l.lineNumber())
		if isEscaped(l) {
			l.emit(Text)
			lexEscape(l)
		}
		if isInlineMarkup(l) {
			l.Msg("FOUND inline reference!")
			if l.index > l.start {
				l.emit(Text)
			}
			lexInlineMarkup(l)
			if isEscaped(l) {
				lexEscape(l)
			}
			continue
		} else if isInlineReference(l) {
			l.Msg("FOUND inline reference!")
			lexInlineReference(l)
			continue
		}
		if l.isEndOfLine() && l.mark == EOL {
			if l.start == l.index {
				return lexStart
			}
			l.emit(Text)
			// if !l.isLastLine() {
			// l.emit(Space) // We hit a "newline", which is converted to a space when in a paragraph
			// }
			break
		}
		l.next()
	}
	l.nextLine()
	return lexStart
}
