// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

type Tree struct {
	Name string
	text string
	// Document *ListNode
	lex *lexer
}

func New(name string) *Tree {
	return &Tree{
		Name: name,
	}
}

func (t *Tree) Parse(text string) (tree *Tree, err error) {
	return nil, nil
}
