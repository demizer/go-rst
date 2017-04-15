package parser

import (
	"fmt"

	. "github.com/demizer/go-rst"
	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/token"
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

func (i *indentQueue) lastStartPosition() StartPosition {
	var n doc.Node
	var nl *doc.NodeList
	var sp StartPosition

	// FIXME: Type system pain here

	// Get the NodeList of the indent item
	switch nt := (*i)[len(*i)-1].node.(type) {
	case *doc.BulletListItemNode:
		nl = &nt.NodeList
		if len(*nl) == 0 {
			return sp
		}
	default:
		log.Err(fmt.Errorf("Unhandled type = %T", nt))
		return sp
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
			return sp
		}
		// Get the first start position in the paragraph node list
		switch nt2 := nt.NodeList[0].(type) {
		case *doc.TextNode:
			sp = nt2.StartPosition
		default:
			log.Err(fmt.Errorf("Unhandled sub Node type = %T", nt))
			return sp
		}
	default:
		log.Err(fmt.Errorf("Unhandled child NodeList type = %T", n))
		return sp
	}

	return sp
}

func (i *indentQueue) topNodeList() *doc.NodeList {
	l := (*i)[len(*i)-1]
	switch n := l.node.(type) {
	case *doc.BulletListItemNode:
		return &n.NodeList
	// case *ParagraphNode:
	// return &n.NodeList
	default:
		log.Err(fmt.Errorf("Unhandled type = %T", n))
	}
	return nil
}

func (i *indentQueue) topNode() doc.Node { return (*i)[len(*i)-1].node }
