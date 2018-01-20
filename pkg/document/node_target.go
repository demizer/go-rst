package document

import (
	"fmt"

	"github.com/demizer/go-rst/pkg/log"
)

// NodeTarget contains the NodeList where subsequent nodes will be added during parsing. It also contains a pointer to the
// parent Node of the NodeTarget NodeList.
type NodeTarget struct {
	MainList *NodeList // The default NodeList for reset()
	SubList  *NodeList // The nodelist contained in target
	Parent   Node      // If set, a parent Node containing a NodeList. Can be nil.

	logConf log.Config
	log.Logger
}

// NewNodeTarget creates a NodeTarget with a context logger.
func NewNodeTarget(pNodes *NodeList, logConf log.Config) *NodeTarget {
	conf := logConf
	conf.Name = "document"
	return &NodeTarget{
		MainList: pNodes,
		SubList:  pNodes,
		logConf:  conf,
		Logger:   log.NewLogger(conf),
	}
}

// Reset sets the NodeTarget to the root document.
func (nt *NodeTarget) Reset() {
	nt.Msgr("Resetting Tree.Nodes", "nodePointer", fmt.Sprintf("%p", nt.MainList))
	nt.SubList = nt.MainList
	nt.Parent = nil
}

// Append add a node to the NodeTarget list.
func (nt *NodeTarget) Append(n ...Node) {
	for _, node := range n {
		// panic("SHOW ME THE STACKS!")
		// nt.Msgr("Adding node", "nodePointer", fmt.Sprintf("%p", &node),
		// "nodeListPointer", fmt.Sprintf("%p", nt.SubList), "node", node.String())
		// switch t := node.(type) {
		// case *InlineEmphasisNode:
		// if t.Text == "*" {
		// // nt.DumpExit(t)
		// panic("foo")
		// }
		// }
		nt.printNode("Adding node", node)
		nt.SubList.Append(node)
	}
}

// SetParent sets the NodeTarget to the NodeList of a Node
func (nt *NodeTarget) SetParent(n Node) {
	// nt.Msgr("setParent have node", "node", n.(Node).String())
	// nt.Msgr("NodeTarget before", "nodeParentPointer", fmt.Sprintf("%p", nt.parent),
	// "nodeListPointer", fmt.Sprintf("%p", nt.subList))
	switch t := n.(type) {
	case *ParagraphNode:
		nt.SubList = &n.(*ParagraphNode).NodeList
		nt.Parent = n
	case *InlineInterpretedText:
		nt.SubList = &n.(*InlineInterpretedText).NodeList
		nt.Parent = n
	case *BlockQuoteNode:
		nt.SubList = &n.(*BlockQuoteNode).NodeList
		nt.Parent = n
	case *SystemMessageNode:
		nt.SubList = &n.(*SystemMessageNode).NodeList
		nt.Parent = n
	case *BulletListNode:
		nt.SubList = &n.(*BulletListNode).NodeList
		nt.Parent = n
	case *BulletListItemNode:
		nt.SubList = &n.(*BulletListItemNode).NodeList
		nt.Parent = n
	case *EnumListNode:
		nt.SubList = &n.(*EnumListNode).NodeList
		nt.Parent = n
	case *DefinitionListNode:
		nt.SubList = &n.(*DefinitionListNode).NodeList
		nt.Parent = n
	case *DefinitionNode:
		nt.SubList = &n.(*DefinitionNode).NodeList
		nt.Parent = n
	case *SectionNode:
		nt.SubList = &n.(*SectionNode).NodeList
		nt.Parent = n
	case *TitleNode:
		nt.SubList = &n.(*TitleNode).NodeList
		nt.Parent = n
	default:
		nt.Msgr("WARNING: type not supported or doesn't have a NodeList!", "type", fmt.Sprintf("%T", t))
	}
	// nt.Msgr("NodeTarget after", "nodeMainListPointer", fmt.Sprintf("%p", nt.mainList),
	// "nodeSubListPointer", fmt.Sprintf("%p", nt.subList), "nodeParentPointer", fmt.Sprintf("%p", nt.parent))
}

// IsParentParagraph will return true if the parent Node of the NodeTarget is a paragraph.
func (nt *NodeTarget) IsParagraphNode() bool {
	switch nt.Parent.(type) {
	case *ParagraphNode:
		nt.Msg("nt.parent is type *ParagraphNode!")
		return true
	default:
		nt.Msg(fmt.Sprintf("nt.parent is type '%T' not type *ParagraphNode!", nt.Parent))
	}
	return false
}

func (nt *NodeTarget) printNode(msg string, n Node) {
	var typ, text, nlp, np string
	var length, line, spos int

	// SIGH ...
	switch t := n.(type) {
	case *TextNode:
		typ = "TextNode"
		text, line, spos, length = t.Text, t.Line, t.StartPosition, t.Length
	case *InlineEmphasisNode:
		typ = "InlineEmphasisNode"
		text, line, spos, length = t.Text, t.Line, t.StartPosition, t.Length
	case *InlineStrongNode:
		typ = "InlineStrongNode"
		text, line, spos, length = t.Text, t.Line, t.StartPosition, t.Length
	case *InlineLiteralNode:
		typ = "InlineLiteralNode"
		text, line, spos, length = t.Text, t.Line, t.StartPosition, t.Length
	case *InlineInterpretedText:
		typ = "InlineInterpretedText"
		text, line, spos, length = t.Text, t.Line, t.StartPosition, t.Length
	case *ParagraphNode: // do nothing
	default:
		nt.Msgr("WARNING: type not supported yet in NodeTarget.printNode()!", "type", fmt.Sprintf("%T", t))
	}

	np = fmt.Sprintf("%p", &n)
	nlp = fmt.Sprintf("%p", nt.SubList)

	if len(text) > 0 {
		log.WithCallDepth(nt.Logger, nt.Logger.CallDepth+1).Msgr("Adding node",
			"nodePointer", np,
			"nodeListPointer", nlp,
			"type", typ,
			"line", line,
			"startPosition", spos,
			"length", length,
			"text", text,
		)
	} else if line > 0 {
		log.WithCallDepth(nt.Logger, nt.Logger.CallDepth+1).Msgr("Adding node",
			"nodePointer", np,
			"nodeListPointer", nlp,
			"type", typ,
			"line", line,
			"startPosition", spos,
			"length", length,
		)
	} else {
		log.WithCallDepth(nt.Logger, nt.Logger.CallDepth+1).Msgr("Adding node",
			"nodePointer", np,
			"nodeListPointer", nlp,
			"type", typ,
		)
	}

}
