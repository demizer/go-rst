package tokenizer

import "unicode"

func isComment(l *lexer) bool {
	if l.lastItem != nil && l.lastItem.Type == itemTitle {
		return false
	}
	nMark := l.peek(1)
	nMark2 := l.peek(2)
	if l.mark == '.' && nMark == '.' && (unicode.IsSpace(nMark2) || nMark2 == EOL) {
		if isHyperlinkTarget(l) {
			logl.Msg("Found hyperlink target!")
			return false
		}
		logl.Msg("Found comment!")
		return true
	}
	logl.Msg("Comment not found!")
	return false
}

func lexComment(l *lexer) stateFn {
	for l.mark == '.' {
		l.next()
	}
	l.emit(itemCommentMark)
	if l.mark != EOL {
		l.next()
		lexSpace(l)
		lexText(l)
	}
	return lexStart
}
