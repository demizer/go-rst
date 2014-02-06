// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

type itemElements int

const (
	itemEOF itemElements = iota
	itemBodyParagraph
	itemBodyText
	itemBodyInlineEmphasis
	itemBodyInlineStrong
	itemBodyInlineInterpreted
	itemBodyInlineInternalTarget
	itemBodyInlineFootnote
	itemBodyInlineHyperLink
	itemBodyBulletList
	itemBodyEnumeratedList
	itemBodyDefinitionList
	itemBodyFieldList
	itemBodyBibliographicField
	itemBodyOptionList
	itemBodyLiteralBlock
	itemBodyLiteralBlockIndented
	itemBodyQuotedBlock
)

type stateFn func(*lexer) stateFn

type item struct {
	typ itemElements
	pos Pos
	val string
}

type lexer struct {
	name    string
	input   string
	state   stateFn
	pos     Pos
	start   Pos
	width   Pos
	lastPos Pos
	items   chan item
}

func lex(name, input string) *lexer {
	l := &lexer{
		name:  name,
		input: input,
		items: make(chan item),
	}
	go l.run()
	return l
}

func (l *lexer) run() {
	for l.state = lexText; l.state != nil; {
		l.state = l.state(l)
	}
}

func lexText(l *lexer) stateFn {
	// for {
	// if strings.HasPrefix(l.input[l.pos:], l.leftDelim) {
	// if l.pos > l.start {
	// l.emit(itemText)
	// }
	// return lexLeftDelim
	// }
	// if l.next() == eof {
	// break
	// }
	// }
	// // Correctly reached EOF.
	// if l.pos > l.start {
	// l.emit(itemText)
	// }
	// l.emit(itemEOF)
	return nil
}
