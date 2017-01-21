package parse

func isBlockquote(l *lexer) bool {
	if !l.lastLineIsBlankLine() || l.lastItem.Type != itemSpace {
		return false
	}
	if l.index != len(l.indentWidth) {
		return true
	}
	return false
}

func lexBlockquote(l *lexer) stateFn {
	for {
		l.next()
		if l.isEndOfLine() && l.mark == EOL {
			l.emit(itemBlockQuote)
			break
		}
	}
	l.nextLine()
	return lexStart
}
