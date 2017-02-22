package token

import (
	"unicode"
	// . "github.com/demizer/go-rst"
)

func isComment(l *Lexer) bool {
	if l.lastItem != nil && l.lastItem.Type == ItemTitle {
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
	l.emit(ItemCommentMark)
	if l.mark != EOL {
		l.next()
		lexSpace(l)
		lexText(l)
	}
	return lexStart
}
