// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

// NodeType identifies the type of a parse tree node.
type NodeType int

const (
	// NodeSection is a section element.
	NodeSection NodeType = iota

	// NodeParagraph is a paragraph element.
	NodeParagraph

	// NodeAdornment is the overline or underline of a section.
	NodeAdornment

	// NodeBlockQuote is a blockquote element.
	NodeBlockQuote

	// NodeSystemMessage contains an error encountered by the parser.
	NodeSystemMessage

	// NodeLiteralBlock is a literal block element.
	NodeLiteralBlock

	// NodeIndent is indention encountered by the lexer. It can be any number
	// of spaces found before any other element type.
	NodeIndent

	// NodeTransition is a transition element. Transitions are very similar to
	// NodeSection except that they have newlines before and after.
	NodeTransition

	// NodeTitle is a section title element to be used inside SectionNodes.
	NodeTitle
)

var nodeTypes = [...]string{
	"NodeSection",
	"NodeParagraph",
	"NodeAdornment",
	"NodeBlockQuote",
	"NodeSystemMessage",
	"NodeLiteralBlock",
	"NodeIndent",
	"NodeTransition",
	"NodeTitle",
}

// Type returns the type of a node element.
func (n NodeType) Type() NodeType {
	return n
}

func (n NodeType) String() string {
	return nodeTypes[n]
}

// Node is the interface used to implement parser nodes.
type Node interface {
	IDNumber() ID
	NodeType() NodeType
}

// NodeList is a list of parser nodes that implement Node.
type NodeList []Node

func newList() *NodeList {
	return new(NodeList)
}

func (l *NodeList) append(n Node) {
	*l = append(*l, n)

}

// SectionNode is a a single section node. It contains overline, underline, and
// indentation nodes. NodeList contains nodes that are children of the section.
type SectionNode struct {
	ID   `json:"id"`
	Type NodeType `json:"type"`

	// Level is the hierarchical level of the section. The first level is level
	// 1, any further sections encountered after the first level are given
	// consecutive level numbers.
	Level int `json:"level"`

	// OverLine and UnderLine are the parsed Nodes that make up the section.
	Title     *TitleNode     `json:"title"`
	OverLine  *AdornmentNode `json:"overLine"`
	UnderLine *AdornmentNode `json:"underLine"`

	// Indent is indentation encountered by the parser before the SectionNode.
	// Sections cannot be indented, so this is primarily for error detection.
	Indent *IndentNode `json:"underLine"`

	// NodeList contains
	NodeList NodeList `json:"nodeList"`
}

// NodeType returns the Node type of the SectionNode.
func (s *SectionNode) NodeType() NodeType {
	return s.Type
}

func newSection(title *item, overSec *item, underSec *item, indent *item, id *int) *SectionNode {
	*id++
	n := &SectionNode{
		ID:   ID(*id),
		Type: NodeSection,
	}

	*id++
	n.Title = &TitleNode{
		ID:            ID(*id),
		Type:          NodeTitle,
		Text:          title.Text.(string),
		StartPosition: title.StartPosition,
		Length:        title.Length,
		Line:          title.Line,
	}

	if indent != nil && indent.Text != nil {
		*id++
		n.Indent = &IndentNode{
			ID:            ID(*id),
			Type:          NodeIndent,
			Text:          indent.Text.(string),
			StartPosition: indent.StartPosition,
			Line:          indent.Line,
			Length:        indent.Length,
		}
	}

	if overSec != nil && overSec.Text != nil {
		*id++
		Rune := rune(overSec.Text.(string)[0])
		n.OverLine = &AdornmentNode{
			ID:            ID(*id),
			Type:          NodeAdornment,
			Rune:          Rune,
			StartPosition: overSec.StartPosition,
			Line:          overSec.Line,
			Length:        overSec.Length,
		}
	}

	*id++
	Rune := rune(underSec.Text.(string)[0])
	n.UnderLine = &AdornmentNode{
		ID:            ID(*id),
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
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

// NodeType returns the Node type of the TitleNode.
func (t TitleNode) NodeType() NodeType {
	return t.Type
}

// AdornmentNode contains the parsed data for a section overline or underline.
type AdornmentNode struct {
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Rune          rune     `json:"rune"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

// NodeType returns the Node type of the AdornmentNode.
func (a AdornmentNode) NodeType() NodeType {
	return a.Type
}

// ParagraphNode is a parsed paragraph.
type ParagraphNode struct {
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newParagraph(i *item, id *int) *ParagraphNode {
	*id++
	return &ParagraphNode{
		ID:            ID(*id),
		Type:          NodeParagraph,
		Text:          i.Text.(string),
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the ParagraphNode.
func (p ParagraphNode) NodeType() NodeType {
	return p.Type
}

// BlockQuoteNode contains a parsed blockquote Node. Any nodes that are
// children of the blockquote are contained in NodeList.
type BlockQuoteNode struct {
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Level         int      `json:"level"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	// NodeList contains Nodes parsed as children of the BlockQuoteNode.
	NodeList NodeList `json:"nodeList"`
}

func newBlockQuote(i *item, indentLevel int, id *int) *BlockQuoteNode {
	*id++
	return &BlockQuoteNode{
		ID:            ID(*id),
		Type:          NodeBlockQuote,
		Level:         indentLevel,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

// NodeType returns the Node type of the BlockQuoteNode.
func (b BlockQuoteNode) NodeType() NodeType {
	return b.Type
}

// SystemMessageNode are messages generated by the parser. System messages are
// leveled by severity and can be one of either Warning, Error, Info, and
// Severe.
type SystemMessageNode struct {
	ID   `json:"id"`
	Type NodeType `json:"type"`
	Line `json:"line"`

	// Severity is the level of importance of the message. It can be one of
	// either info, warning, error, and severe.
	Severity systemMessageLevel `json:"severity"`

	// NodeList contains children Nodes of the systemMessage. Typically
	// containing the first list item as a NodeParagraph which contains the
	// message, and a NodeLiteralBlock which contains the input data causing
	// the systemMessage to be generated.
	NodeList NodeList `json:"nodeList"`
}

func newSystemMessage(i *item, severity systemMessageLevel, id *int) *SystemMessageNode {
	*id++
	return &SystemMessageNode{
		ID:       ID(*id),
		Type:     NodeSystemMessage,
		Severity: severity,
		Line:     i.Line,
	}
}

// NodeType returns the Node type of the SystemMessageNode.
func (s SystemMessageNode) NodeType() NodeType {
	return s.Type
}

// LiteralBlockNode is a parsed literal block element.
type LiteralBlockNode struct {
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newLiteralBlock(i *item, id *int) *LiteralBlockNode {
	*id++
	return &LiteralBlockNode{
		ID:            ID(*id),
		Type:          NodeLiteralBlock,
		Text:          i.Text.(string),
		Length:        i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

// NodeType returns the Node type of LiteralBlockNode.
func (l LiteralBlockNode) NodeType() NodeType {
	return l.Type
}

// IndentNode is any indentation encountered by the parser before block level
// elements.
type IndentNode struct {
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

// NodeType returns the Node type of IndentNode.
func (i IndentNode) NodeType() NodeType {
	return i.Type
}

// TransitionNode is a parsed transition element. Transition elements are very
// similar to AdornmentNodes.
type TransitionNode struct {
	ID            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newTransition(i *item, id *int) *TransitionNode {
	*id++
	return &TransitionNode{
		ID:            ID(*id),
		Type:          NodeTransition,
		Text:          i.Text.(string),
		Length:        i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

// NodeType returns the Node type of the TransitionNode.
func (t TransitionNode) NodeType() NodeType {
	return t.Type
}
