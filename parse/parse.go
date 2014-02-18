// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

type Tree struct {
	Name      string
	text      string
	Root      *ListNode
	lex       *lexer
	peekCount int
	token     [3]item // three-token lookahead for parser.
}

func Parse(name, text string) (t *Tree, err error) {
	t = New(name)
	t.text = text
	_, err = t.Parse(text, t)
	return
}

func New(name string) *Tree {
	return &Tree{
		Name: name,
	}
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(lex *lexer) {
	t.Root = nil
	t.lex = lex
}

// stopParse terminates parsing.
func (t *Tree) stopParse() {
	t.lex = nil
}

func (t *Tree) Parse(text string, treeSet *Tree) (tree *Tree, err error) {
	t.startParse(lex(t.Name, text))
	t.text = text
	t.parse(treeSet)
	return t, nil
}

func (t *Tree) parse(tree *Tree) (next Node) {
	t.Root = newList(t.peek().Position)
	for t.peek().ElementType != itemEOF {
		//
	}
	return nil

}

// peek returns but does not consume the next token.
func (t *Tree) peek() item {
	if t.peekCount > 0 {
		return t.token[t.peekCount-1]

	}
	t.peekCount = 1
	t.token[0] = t.lex.nextItem()
	return t.token[0]
}
