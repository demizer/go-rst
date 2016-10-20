package parse

import "fmt"

// NodeType identifies the type of a parse tree node.
type NodeType int

const (
	// NodeSection is a section element.
	NodeSection NodeType = iota

	// NodeText is ordinary text.
	NodeText

	// NodeParagraph is a paragraph container that contains text and inline markup.
	NodeParagraph

	// NodeAdornment is the overline or underline of a section.
	NodeAdornment

	// NodeBlockQuote is a blockquote element.
	NodeBlockQuote

	// NodeSystemMessage contains an error encountered by the parser.
	NodeSystemMessage

	// NodeLiteralBlock is a literal block element.
	NodeLiteralBlock

	// NodeTransition is a transition element. Transitions are very similar
	// to NodeSection except that they have newlines before and after.
	NodeTransition

	// NodeTitle is a section title element to be used inside SectionNodes.
	NodeTitle

	// NodeComment is a comment element
	NodeComment

	// NodeBulletList is the beginning of a bullet list
	NodeBulletList

	// NodeBulletListItem is a bullet list item
	NodeBulletListItem

	// NodeEnumList is an enumerated list
	NodeEnumList

	// NodeDefinitionList is the beginning of a definition list element
	NodeDefinitionList

	// NodeDefinitionListItem is a definition list item
	NodeDefinitionListItem

	// NodeDefinitionTerm is a definition list term element
	NodeDefinitionTerm

	// NodeDefinition is a definition element
	NodeDefinition

	// NodeInlineEmphasis is the italicized text element
	NodeInlineEmphasis

	// NodeInlineStrong is the bold text element
	NodeInlineStrong

	// NodeInlineLiteral defines inline literal markup
	NodeInlineLiteral

	// NodeInlineInterpretedText is part of an interpreted text role
	NodeInlineInterpretedText

	// NodeInlineInterpretedTextRole is the role of the interpreted text
	NodeInlineInterpretedTextRole
)

var nodeTypes = [...]string{
	"NodeSection",
	"NodeText",
	"NodeParagraph",
	"NodeAdornment",
	"NodeBlockQuote",
	"NodeSystemMessage",
	"NodeLiteralBlock",
	"NodeTransition",
	"NodeTitle",
	"NodeComment",
	"NodeBulletList",
	"NodeBulletListItem",
	"NodeEnumList",
	"NodeDefinitionList",
	"NodeDefinitionListItem",
	"NodeDefinitionTerm",
	"NodeDefinition",
	"NodeInlineEmphasis",
	"NodeInlineStrong",
	"NodeInlineLiteral",
	"NodeInlineInterpretedText",
	"NodeInlineInterpretedTextRole",
}

// Type returns the type of a node element.
func (n NodeType) Type() NodeType { return n }

func (n NodeType) String() string { return nodeTypes[n] }

// Node is the interface used to implement parser nodes.
type Node interface {
	NodeType() NodeType
	String() string
}

// NodeList is a list of parser nodes that implement Node.
type NodeList []Node

func (l *NodeList) append(n ...Node) {
	for _, node := range n {
		*l = append(*l, node)
	}
}

// last returns the last item added to the slice
func (l *NodeList) lastNode(n ...Node) Node { return (*l)[len(*l)-1] }

// NodeTarget contains the NodeList where subsequent nodes will be added during parsing. It also contains a pointer to the
// parent Node of the NodeTarget NodeList.
type nodeTarget struct {
	mainList *NodeList // The default NodeList for reset()
	subList  *NodeList // The nodelist contained in target
	parent   Node      // If set, a parent Node containing a NodeList. Can be nil.
}

func newNodeTarget(pNodes *NodeList) *nodeTarget {
	return &nodeTarget{mainList: pNodes, subList: pNodes}
}

func (nt *nodeTarget) reset() {
	logp.Log("msg", "Resetting Tree.Nodes", "nodePointer", fmt.Sprintf("%p", nt.mainList))
	nt.subList = nt.mainList
	nt.parent = nil
}

func (nt *nodeTarget) append(n ...Node) {
	for _, node := range n {
		logp.Log("msg", "Adding node", "nodePointer", fmt.Sprintf("%p", node),
			"nodeListPointer", fmt.Sprintf("%p", nt.subList), "node", node.String())
		nt.subList.append(node)
	}
}

// setParent sets the nodeTarget to the NodeList of a Node
func (nt *nodeTarget) setParent(n Node) {
	// logp.Log("msg", "setParent have node", "node", n.(Node).String())
	// logp.Log("msg", "nodeTarget before", "nodeParentPointer", fmt.Sprintf("%p", nt.parent),
	// "nodeListPointer", fmt.Sprintf("%p", nt.subList))
	switch t := n.(type) {
	case *ParagraphNode:
		nt.subList = &n.(*ParagraphNode).NodeList
		nt.parent = n
	case *InlineInterpretedText:
		nt.subList = &n.(*InlineInterpretedText).NodeList
		nt.parent = n
	case *BlockQuoteNode:
		nt.subList = &n.(*BlockQuoteNode).NodeList
		nt.parent = n
	case *SystemMessageNode:
		nt.subList = &n.(*SystemMessageNode).NodeList
		nt.parent = n
	case *BulletListNode:
		nt.subList = &n.(*BulletListNode).NodeList
		nt.parent = n
	case *BulletListItemNode:
		nt.subList = &n.(*BulletListItemNode).NodeList
		nt.parent = n
	case *EnumListNode:
		nt.subList = &n.(*EnumListNode).NodeList
		nt.parent = n
	case *DefinitionListNode:
		nt.subList = &n.(*DefinitionListNode).NodeList
		nt.parent = n
	case *DefinitionNode:
		nt.subList = &n.(*DefinitionNode).NodeList
		nt.parent = n
	case *SectionNode:
		nt.subList = &n.(*SectionNode).NodeList
		nt.parent = n
	default:
		logp.Log("msg", "WARNING: type not supported or dosen't have a NodeList!", "type", fmt.Sprintf("%T", t))
	}
	// logp.Log("msg", "nodeTarget after", "nodeMainListPointer", fmt.Sprintf("%p", nt.mainList),
	// "nodeSubListPointer", fmt.Sprintf("%p", nt.subList), "nodeParentPointer", fmt.Sprintf("%p", nt.parent))
}

// isParentParagraph will return true if the parent Node of the NodeTarget is a paragraph.
func (nt *nodeTarget) isParagraphNode() bool {
	switch nt.parent.(type) {
	case *ParagraphNode:
		logp.Msg("nt.parent is type *ParagraphNode!")
		return true
	default:
		logp.Msg(fmt.Sprintf("nt.parent is type '%T' not type *ParagraphNode!", nt.parent))
	}
	return false
}

// EnumListType identifies the type of the enumeration list element
type EnumListType int

const (
	enumListArabic EnumListType = iota
	enumListUpperAlpha
	enumListLowerAlpha
	enumListUpperRoman
	enumListLowerRoman
	enumListAuto
)

var enumListTypes = [...]string{
	"enumListArabic",
	"enumListUpperAlpha",
	"enumListLowerAlpha",
	"enumListUpperRoman",
	"enumListLowerRoman",
	"enumListAuto",
}

func (e EnumListType) String() string { return enumListTypes[e] }

// EnumAffixType identifies the type of affix for the Enumerated list element
type EnumAffixType int

const (
	enumAffixPeriod EnumAffixType = iota
	enumAffixParenthesisSurround
	enumAffixParenthesisRight
)

var enumAffixesTypes = [...]string{
	"enumAffixPeriod",
	"enumAffixParenthesisSurround",
	"enumAffixParenthesisRight",
}

// String satisfies the Stringer interface
func (a EnumAffixType) String() string { return enumAffixesTypes[a] }

// SectionNode is a a single section node. It contains overline, title, and underline nodes. NodeList contains nodes that are
// children of the section.
type SectionNode struct {
	Type NodeType `json:"type"`

	// Level is the hierarchical level of the section. The first level is level 1, any further sections encountered after
	// the first level are given consecutive level numbers.
	Level int `json:"level"`

	// OverLine and UnderLine are the parsed Nodes that make up the section.
	Title     *TitleNode     `json:"title"`
	OverLine  *AdornmentNode `json:"overLine"`
	UnderLine *AdornmentNode `json:"underLine"`

	// NodeList contains
	NodeList `json:"nodeList"`
}

// NodeType returns the Node type of the SectionNode.
func (s *SectionNode) NodeType() NodeType { return s.Type }

// String satisfies the Stringer interface
func (s *SectionNode) String() string { return fmt.Sprintf("%#v", s) }

func newSection(title *item, overSec *item, underSec *item, indent *item) *SectionNode {
	var indentLen int
	n := &SectionNode{Type: NodeSection}
	if indent != nil {
		indentLen = indent.Length
	}
	n.Title = &TitleNode{
		Type:          NodeTitle,
		Text:          title.Text,
		StartPosition: title.StartPosition,
		IndentLength:  indentLen,
		Length:        title.Length,
		Line:          title.Line,
	}
	if overSec != nil && overSec.Text != "" {
		Rune := rune(overSec.Text[0])
		n.OverLine = &AdornmentNode{
			Type:          NodeAdornment,
			Rune:          Rune,
			StartPosition: overSec.StartPosition,
			Line:          overSec.Line,
			Length:        overSec.Length,
		}
	}
	Rune := rune(underSec.Text[0])
	n.UnderLine = &AdornmentNode{
		Rune:          Rune,
		Type:          NodeAdornment,
		StartPosition: underSec.StartPosition,
		Line:          underSec.Line,
		Length:        underSec.Length,
	}

	return n
}

// TitleNode contains the parsed data for a section titles.
type TitleNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	IndentLength  int      `json:"indentLength"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

// NodeType returns the Node type of the TitleNode.
func (t TitleNode) NodeType() NodeType { return t.Type }

// String satisfies the Stringer interface
func (t TitleNode) String() string { return fmt.Sprintf("%#v", t) }

// AdornmentNode contains the parsed data for a section overline or underline.
type AdornmentNode struct {
	Type          NodeType `json:"type"`
	Rune          rune     `json:"rune"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

// NodeType returns the Node type of the AdornmentNode.
func (a AdornmentNode) NodeType() NodeType { return a.Type }

// String satisfies the Stringer interface
func (a AdornmentNode) String() string { return fmt.Sprintf("%#v", a) }

// TextNode is ordinary text. Typically added to the nodelist of parapgraphs.
type TextNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newText(i *item) *TextNode {
	return &TextNode{
		Type:          NodeText,
		Text:          i.Text,
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the TextNode.
func (p TextNode) NodeType() NodeType { return p.Type }

// String satisfies the Stringer interface
func (p TextNode) String() string { return fmt.Sprintf("%#v", p) }

// ParagraphNode is a parsed paragraph.
type ParagraphNode struct {
	Type     NodeType          `json:"type"`
	NodeList `json:"nodeList"` // NodeList contains children of the ParagraphNode, even other ParagraphNodes!
}

func newParagraph() *ParagraphNode { return &ParagraphNode{Type: NodeParagraph} }

func newParagraphWithNodeText(i *item) *ParagraphNode {
	pn := &ParagraphNode{Type: NodeParagraph}
	pn.append(newText(i))
	return pn
}

// NodeType returns the Node type of the ParagraphNode.
func (p ParagraphNode) NodeType() NodeType { return p.Type }

// String satisfies the Stringer interface
func (p ParagraphNode) String() string { return fmt.Sprintf("%#v", p) }

// InlineEmphasisNode is parsed inline italicized text.
type InlineEmphasisNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newInlineEmphasis(i *item) *InlineEmphasisNode {
	return &InlineEmphasisNode{
		Type:          NodeInlineEmphasis,
		Text:          i.Text,
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the InlineEmphasisNode.
func (e InlineEmphasisNode) NodeType() NodeType { return e.Type }

// String satisfies the Stringer interface
func (e InlineEmphasisNode) String() string { return fmt.Sprintf("%#v", e) }

// InlineStrongNode is a parsed inline bold text.
type InlineStrongNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newInlineStrong(i *item) *InlineStrongNode {
	return &InlineStrongNode{
		Type:          NodeInlineStrong,
		Text:          i.Text,
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the InlineStrongNode.
func (s InlineStrongNode) NodeType() NodeType { return s.Type }

// String satisfies the Stringer interface
func (s InlineStrongNode) String() string { return fmt.Sprintf("%#v", s) }

// InlineLiteralNode is a parsed inline literal node.
type InlineLiteralNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newInlineLiteral(i *item) *InlineLiteralNode {
	return &InlineLiteralNode{
		Type:          NodeInlineLiteral,
		Text:          i.Text,
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the InlineStrongNode.
func (l InlineLiteralNode) NodeType() NodeType { return l.Type }

// String satisfies the Stringer interface
func (l InlineLiteralNode) String() string { return fmt.Sprintf("%#v", l) }

// InlineInterpretedText is a parsed interpreted text role.
type InlineInterpretedText struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	// NodeList contains Nodes parsed as children of the BlockQuoteNode.
	NodeList `json:"nodeList"`
}

func newInlineInterpretedText(i *item) *InlineInterpretedText {
	return &InlineInterpretedText{
		Type:          NodeInlineInterpretedText,
		Text:          i.Text,
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the InlineInterpretedText.
func (i InlineInterpretedText) NodeType() NodeType { return i.Type }

// String satisfies the Stringer interface
func (i InlineInterpretedText) String() string { return fmt.Sprintf("%#v", i) }

// InlineInterpretedTextRole is a parsed interpreted text role.
type InlineInterpretedTextRole struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newInlineInterpretedTextRole(i *item) *InlineInterpretedTextRole {
	return &InlineInterpretedTextRole{
		Type:          NodeInlineInterpretedTextRole,
		Text:          i.Text,
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the InlineInterpretedTextRole
func (i InlineInterpretedTextRole) NodeType() NodeType { return i.Type }

// String satisfies the Stringer interface
func (i InlineInterpretedTextRole) String() string { return fmt.Sprintf("%#v", i) }

// BlockQuoteNode contains a parsed blockquote Node. Any nodes that are children of the blockquote are contained in NodeList.
type BlockQuoteNode struct {
	Type          NodeType `json:"type"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	// NodeList contains Nodes parsed as children of the BlockQuoteNode.
	NodeList `json:"nodeList"`
}

func newEmptyBlockQuote(i *item) *BlockQuoteNode {
	bq := &BlockQuoteNode{
		Type:          NodeBlockQuote,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
	return bq
}

func newBlockQuote(i *item) *BlockQuoteNode {
	bq := &BlockQuoteNode{
		Type:          NodeBlockQuote,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
	bq.NodeList.append(newParagraphWithNodeText(i))
	return bq
}

// NodeType returns the Node type of the BlockQuoteNode.
func (b BlockQuoteNode) NodeType() NodeType { return b.Type }

// String satisfies the Stringer interface
func (b BlockQuoteNode) String() string { return fmt.Sprintf("%#v", b) }

// SystemMessageNode are messages generated by the parser. System messages are leveled by severity and can be one of either
// Warning, Error, Info, and Severe.
type SystemMessageNode struct {
	Type NodeType `json:"type"`
	Line `json:"line"`

	// The type of parser message that generated the systemMessage.
	MessageType parserMessage `json:"messageType"`

	// Severity is the level of importance of the message. It can be one of either info, warning, error, and severe.
	Severity systemMessageLevel `json:"severity"`

	// NodeList contains children Nodes of the systemMessage. Typically containing the first list item as a NodeParagraph
	// which contains the message, and a NodeLiteralBlock which contains the input data causing the systemMessage to be
	// generated.
	NodeList `json:"nodeList"`
}

func newSystemMessage(i *item, m parserMessage) *SystemMessageNode {
	return &SystemMessageNode{
		Type:        NodeSystemMessage,
		MessageType: m,
		Severity:    m.Level(),
		Line:        i.Line,
	}
}

// NodeType returns the Node type of the SystemMessageNode.
func (s SystemMessageNode) NodeType() NodeType { return s.Type }

// String satisfies the Stringer interface
func (s SystemMessageNode) String() string { return fmt.Sprintf("%#v", s) }

// LiteralBlockNode is a parsed literal block element.
type LiteralBlockNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newLiteralBlock(i *item) *LiteralBlockNode {
	return &LiteralBlockNode{
		Type:          NodeLiteralBlock,
		Text:          i.Text,
		Length:        i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

// NodeType returns the Node type of LiteralBlockNode.
func (l LiteralBlockNode) NodeType() NodeType { return l.Type }

// String satisfies the Stringer interface
func (l LiteralBlockNode) String() string { return fmt.Sprintf("%#v", l) }

// TransitionNode is a parsed transition element. Transition elements are very similar to AdornmentNodes.
type TransitionNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newTransition(i *item) *TransitionNode {
	return &TransitionNode{
		Type:          NodeTransition,
		Text:          i.Text,
		Length:        i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

// NodeType returns the Node type of the TransitionNode.
func (t TransitionNode) NodeType() NodeType { return t.Type }

// String satisfies the Stringer interface
func (t TransitionNode) String() string { return fmt.Sprintf("%#v", t) }

// CommentNode is a parsed comment element. Comment elements do not appear as visible elements in document transformations.
type CommentNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newComment(i *item) *CommentNode {
	return &CommentNode{
		Type:          NodeComment,
		Text:          i.Text,
		Length:        i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

// NodeType returns the Node type of the CommentNode.
func (t CommentNode) NodeType() NodeType { return t.Type }

// String satisfies the Stringer interface
func (t CommentNode) String() string { return fmt.Sprintf("%#v", t) }

// BulletListNode defines a bullet list element.
type BulletListNode struct {
	Type     NodeType `json:"type"`
	Bullet   string   `json:"bullet"`
	NodeList `json:"nodeList"`
}

// newEnumListNode initializes a new BulletListNode.
func newBulletListNode(i *item) *BulletListNode {
	return &BulletListNode{
		Type:   NodeBulletList,
		Bullet: i.Text,
	}
}

// NodeType returns the type of Node for the bullet list.
func (b BulletListNode) NodeType() NodeType { return b.Type }

// String satisfies the Stringer interface
func (b BulletListNode) String() string { return fmt.Sprintf("%#v", b) }

// BulletListItemNode defines a Bullet List Item element.
type BulletListItemNode struct {
	Type     NodeType `json:"type"`
	NodeList `json:"nodeList"`
}

// newBulletListNode initializes a new EnumListNode.
func newBulletListItemNode(i *item) *BulletListItemNode {
	return &BulletListItemNode{Type: NodeBulletListItem}
}

// NodeType returns the type of Node for the bullet list item.
func (b BulletListItemNode) NodeType() NodeType { return b.Type }

// String satisfies the Stringer interface
func (b BulletListItemNode) String() string { return fmt.Sprintf("%#v", b) }

// EnumListNode defines an enumerated list element.
type EnumListNode struct {
	Type     NodeType      `json:"type"`
	EnumType EnumListType  `json:"enumType"`
	Affix    EnumAffixType `json:"affix"`
	NodeList `json:"nodeList"`
}

// newEnumListNode initializes a new EnumListNode.
func newEnumListNode(enumList *item, affix *item) *EnumListNode {
	var enType EnumListType
	switch enumList.Type {
	case itemEnumListArabic:
		enType = enumListArabic
	}

	var afType EnumAffixType
	switch affix.Text {
	case ".":
		afType = enumAffixPeriod
		// case "(":
		// afType = enumAffixParenthesisSurround
		// case ")":
		// afType = enumAffixParenthesisRight
	}

	return &EnumListNode{
		Type:     NodeEnumList,
		EnumType: enType,
		Affix:    afType,
	}
}

// NodeType returns the Node type of the EnumListNode.
func (e EnumListNode) NodeType() NodeType { return e.Type }

// String satisfies the Stringer interface
func (e EnumListNode) String() string { return fmt.Sprintf("%#v", e) }

// DefinitionListNode defines a definition list element.
type DefinitionListNode struct {
	Type     NodeType `json:"type"`
	NodeList `json:"nodeList"`
}

func newDefinitionList(i *item) *DefinitionListNode {
	return &DefinitionListNode{Type: NodeDefinitionList}
}

// NodeType returns the Node type of DefinitionListNode.
func (d DefinitionListNode) NodeType() NodeType { return d.Type }

// String satisfies the Stringer interface
func (d DefinitionListNode) String() string { return fmt.Sprintf("%#v", d) }

// DefinitionListItemNode defines a definition list item element.
type DefinitionListItemNode struct {
	Type       NodeType            `json:"type"`
	Term       *DefinitionTermNode `json:"term"`
	Definition *DefinitionNode     `json:"definition"`
}

func newDefinitionListItem(defTerm *item, def *item) *DefinitionListItemNode {
	n := &DefinitionListItemNode{Type: NodeDefinitionListItem}
	ndt := &DefinitionTermNode{
		Type:          NodeDefinitionTerm,
		Text:          defTerm.Text,
		Length:        defTerm.Length,
		StartPosition: defTerm.StartPosition,
		Line:          defTerm.Line,
	}
	nd := &DefinitionNode{Type: NodeDefinition}
	n.Term = ndt
	n.Definition = nd
	return n
}

// NodeType returns the Node type of DefinitionListItemNode.
func (d DefinitionListItemNode) NodeType() NodeType { return d.Type }

// String satisfies the Stringer interface
func (d DefinitionListItemNode) String() string { return fmt.Sprintf("%#v", d) }

// DefinitionTermNode defines a definition list term element.
type DefinitionTermNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

// NodeType returns the Node type of DefinitionTermNode.
func (d DefinitionTermNode) NodeType() NodeType { return d.Type }

// String satisfies the Stringer interface
func (d DefinitionTermNode) String() string { return fmt.Sprintf("%#v", d) }

// DefinitionNode defines a difinition element.
type DefinitionNode struct {
	Type     NodeType `json:"type"`
	Line     `json:"line"`
	NodeList `json:"nodeList"`
}

// NodeType returns the Node type of DefinitionNode.
func (d DefinitionNode) NodeType() NodeType { return d.Type }

// String satisfies the Stringer interface
func (d DefinitionNode) String() string { return fmt.Sprintf("%#v", d) }
