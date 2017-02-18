package tokenizer

func lexText(l *lexer) stateFn {
	logl.Msg("lexText start")
	for {
		// logl.Log.Debugf("l.mark: %q, l.index: %d, l.width: %d, l.line: %d", l.mark, l.index, l.width, l.lineNumber())
		if isEscaped(l) {
			l.emit(itemText)
			lexEscape(l)
		}
		if isInlineMarkup(l) {
			logl.Msg("FOUND inline reference!")
			if l.index > l.start {
				l.emit(itemText)
			}
			lexInlineMarkup(l)
			if isEscaped(l) {
				lexEscape(l)
			}
			continue
		} else if isInlineReference(l) {
			logl.Msg("FOUND inline reference!")
			lexInlineReference(l)
			continue
		}
		if l.isEndOfLine() && l.mark == EOL {
			if l.start == l.index {
				return lexStart
			}
			l.emit(itemText)
			// if !l.isLastLine() {
			// l.emit(itemSpace) // We hit a "newline", which is converted to a space when in a paragraph
			// }
			break
		}
		l.next()
	}
	l.nextLine()
	return lexStart
}
