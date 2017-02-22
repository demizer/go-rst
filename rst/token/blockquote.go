package token

func isBlockquote(l *Lexer) bool {
	if !l.lastLineIsBlankLine() || l.lastItem.Type != Space {
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
			l.emit(BlockQuote)
			break
		}
	}
	l.nextLine()
	return lexStart
}
