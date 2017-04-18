package token

import (
	"unicode"
)

func isComment(l *Lexer) bool {
	if l.lastItem != nil && l.lastItem.Type == Title {
		return false
	}
	nMark := l.peek(1)
	nMark2 := l.peek(2)
	if l.mark == '.' && nMark == '.' && (unicode.IsSpace(nMark2) || nMark2 == EOL) {
		if isHyperlinkTarget(l) {
			log.Msg("Found hyperlink target!")
			return false
		}
		log.Msg("Found comment!")
		return true
	}
	log.Msg("Comment not found!")
	return false
}

func lexComment(l *Lexer) stateFn {
	for l.mark == '.' {
		l.next()
	}
	l.emit(CommentMark)
	if l.mark != EOL {
		l.next()
		lexSpace(l)
		lexText(l)
	}
	return lexStart
}
