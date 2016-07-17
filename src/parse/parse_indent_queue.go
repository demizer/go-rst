package parse

import "fmt"

type indent struct {
	token *item
	node  Node
}

// indentQueue is a LIFO stack
type indentQueue []*indent

func (i *indentQueue) len() int { return len(*i) }

func (i *indentQueue) add(t *item, n Node) Node {
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
		Log.Error(fmt.Sprintf("topNodeList: Unhandled type = %q", spd.Sdump(nt)))
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
			Log.Error(fmt.Sprintf("topNodeList: Unhandled sub Node type = %q", spd.Sdump(nt)))
			return sp
		}
	default:
		Log.Error(fmt.Sprintf("topNodeList: Unhandled child NodeList type = %q", spd.Sdump(n)))
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
		Log.Error(fmt.Sprintf("topNodeList: Unhandled type = %q", spd.Sdump(n)))
	}
	return nil
}
