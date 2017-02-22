package token

func isTransition(l *Lexer) bool {
	if r := l.peek(4); !isSectionAdornment(l.mark) || !isSectionAdornment(r) {
		log.Msg("Transition not found")
		return false
	}
	pBlankLine := l.lastItem != nil && l.lastItem.Type == BlankLine
	nBlankLine := l.peekNextLine() == ""
	if l.line == 0 && nBlankLine {
		log.Msg("Found transition (followed by newline)")
		return true
	} else if pBlankLine && nBlankLine {
		log.Msg("Found transition (surrounded by newlines)")
		return true
	}
	log.Msg("Transition not found")
	return false
}

func lexTransition(l *Lexer) stateFn {
	for {
		if len(l.lines[l.line]) == l.index {
			break
		}
		l.next()
	}
	l.emit(Transition)
	l.nextLine()
	return lexStart
}
