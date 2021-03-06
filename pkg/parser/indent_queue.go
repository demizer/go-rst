package parser

import (
	"fmt"

	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

type indent struct {
	token *tok.Item
	node  doc.Node
}

// indentQueue is a LIFO stack
type indentQueue []*indent

func (i *indentQueue) len() int { return len(*i) }

func (i *indentQueue) add(t *tok.Item, n doc.Node) doc.Node {
	ni := &indent{t, n}
	*i = append(*i, ni)
	return n
}

func (i *indentQueue) pop() { *i = (*i)[:len(*i)-1] }

func (i *indentQueue) lastStartPosition() (int, error) {
	var sp int
	var n doc.Node
	var nl *doc.NodeList

	// Get the NodeList of the indent item
	switch nt := (*i)[len(*i)-1].node.(type) {
	case *doc.BulletListItemNode:
		nl = &nt.NodeList
		if len(*nl) == 0 {
			return sp, nil
		}
	default:
		return sp, fmt.Errorf("Unhandled type = %T", nt)
	}

	// Get node that isn't a comment
	for i := len(*nl) - 1; i >= 0; i-- {
		n = (*nl)[i]
		if n.NodeType() != doc.NodeComment {
			break
		}
	}

	// Return the StartPosition of the last NodeList item
	switch nt := n.(type) {
	case *doc.ParagraphNode:
		if len(nt.NodeList) == 0 {
			return sp, nil
		}
		// Get the first start position in the paragraph node list
		switch nt2 := nt.NodeList[0].(type) {
		case *doc.TextNode:
			sp = nt2.StartPosition
		default:
			return sp, fmt.Errorf("Unhandled sub Node type = %T", nt)
		}
	default:
		return sp, fmt.Errorf("Unhandled child NodeList type = %T", n)
	}

	return sp, nil
}

func (i *indentQueue) topNodeList() (*doc.NodeList, error) {
	l := (*i)[len(*i)-1]
	switch n := l.node.(type) {
	case *doc.BulletListItemNode:
		return &n.NodeList, nil
	// case *ParagraphNode:
	// return &n.NodeList
	default:
		return nil, fmt.Errorf("Unhandled type = %T", n)
	}
	return nil, nil
}

func (i *indentQueue) topNode() doc.Node { return (*i)[len(*i)-1].node }
