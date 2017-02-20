package parser

import (
	"fmt"

	// doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/tokenizer"
)

type indent struct {
	token *tok.Item
	node  doc.Node
}

// indentQueue is a LIFO stack
type indentQueue []*indent

func (i *indentQueue) len() int { return len(*i) }

func (i *indentQueue) add(t *tok.Item, n Node) Node {
	ni := &indent{t, n}
	*i = append(*i, ni)
	return n
}

func (i *indentQueue) pop() { *i = (*i)[:len(*i)-1] }

func (i *indentQueue) lastStartPosition() StartPosition {
	var n Node
	var nl *NodeList
	var sp StartPosition

	// FIXME: Type system pain here

	// Get the NodeList of the indent item
	switch nt := (*i)[len(*i)-1].node.(type) {
	case *BulletListItemNode:
		nl = &nt.NodeList
		if len(*nl) == 0 {
			return sp
		}
	default:
		logp.Err(fmt.Errorf("Unhandled type = %T", nt))
		return sp
	}

	// Get node that isn't a comment
	for i := len(*nl) - 1; i >= 0; i-- {
		n = (*nl)[i]
		if n.NodeType() != NodeComment {
			break
		}
	}

	// Return the StartPosition of the last NodeList item
	switch nt := n.(type) {
	case *ParagraphNode:
		if len(nt.NodeList) == 0 {
			return sp
		}
		// Get the first start position in the paragraph node list
		switch nt2 := nt.NodeList[0].(type) {
		case *TextNode:
			sp = nt2.StartPosition
		default:
			logp.Err(fmt.Errorf("Unhandled sub Node type = %T", nt))
			return sp
		}
	default:
		logp.Err(fmt.Errorf("Unhandled child NodeList type = %T", n))
		return sp
	}

	return sp
}

func (i *indentQueue) topNodeList() *NodeList {
	l := (*i)[len(*i)-1]
	switch n := l.node.(type) {
	case *BulletListItemNode:
		return &n.NodeList
	// case *ParagraphNode:
	// return &n.NodeList
	default:
		logp.Err(fmt.Errorf("Unhandled type = %T", n))
	}
	return nil
}

func (i *indentQueue) topNode() Node { return (*i)[len(*i)-1].node }
