package tokenizer

func isTransition(l *lexer) bool {
	if r := l.peek(4); !isSectionAdornment(l.mark) || !isSectionAdornment(r) {
		logl.Msg("Transition not found")
		return false
	}
	pBlankLine := l.lastItem != nil && l.lastItem.Type == itemBlankLine
	nBlankLine := l.peekNextLine() == ""
	if l.line == 0 && nBlankLine {
		logl.Msg("Found transition (followed by newline)")
		return true
	} else if pBlankLine && nBlankLine {
		logl.Msg("Found transition (surrounded by newlines)")
		return true
	}
	logl.Msg("Transition not found")
	return false
}

func lexTransition(l *lexer) stateFn {
	for {
		if len(l.lines[l.line]) == l.index {
			break
		}
		l.next()
	}
	l.emit(itemTransition)
	l.nextLine()
	return lexStart
}
