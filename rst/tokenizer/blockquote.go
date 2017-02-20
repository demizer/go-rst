package tokenizer

func isBlockquote(l *Lexer) bool {
	if !l.lastLineIsBlankLine() || l.lastItem.Type != ItemSpace {
		return false
	}
	if l.index != len(l.indentWidth) {
		return true
	}
	return false
}

func lexBlockquote(l *Lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() && l.mark == EOL {
			l.emit(ItemBlockQuote)
			break
		}
	}
	l.nextLine()
	return lexStart
}
