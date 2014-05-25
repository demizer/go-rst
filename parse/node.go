// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

// NodeType identifies the type of a parse tree node.
type NodeType int

const (
	NodeSection NodeType = iota
	NodeParagraph
	NodeBlankLine
	NodeAdornment
)

var nodeTypes = [...]string{
	"NodeSection",
	"NodeParagraph",
	"NodeBlankLine",
	"NodeAdornment",
}

func (n NodeType) Type() NodeType {
	return n
}

func (n NodeType) String() string {
	return nodeTypes[n]
}

func (n NodeType) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

type Node interface {
	LineNumber() Line
	NodeType() NodeType
	Position() StartPosition
}

type NodeList []Node

func newList() *NodeList {
	return new(NodeList)
}

func (l *NodeList) append(n Node) {
	*l = append(*l, n)

}

type SectionNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Level         int      `json:"level"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
	OverLine      *AdornmentNode `json:"overLine"`
	UnderLine     *AdornmentNode `json:"underLine"`
	NodeList      NodeList       `json:"nodeList"`
}

func (s *SectionNode) NodeType() NodeType {
	return s.Type
}

func newSection(title *item, overAdorn *item, underAdorn *item) *SectionNode {
	n := &SectionNode{
		Id:            title.Id,
		Type:          NodeSection,
		Text:          title.Text.(string),
		StartPosition: title.StartPosition,
		Length:        title.Length,
		Line:          title.Line,
	}

	if overAdorn != nil && overAdorn.Text != nil {
		Rune := rune(overAdorn.Text.(string)[0])
		n.OverLine = &AdornmentNode{
			Id:            title.Id,
			Type:          NodeAdornment,
			Rune:          Rune,
			StartPosition: overAdorn.StartPosition,
			Line:          overAdorn.Line,
			Length:        overAdorn.Length,
		}
	}

	Rune := rune(underAdorn.Text.(string)[0])
	n.UnderLine = &AdornmentNode{
		Id:            title.Id,
		Rune:          Rune,
		Type:          NodeAdornment,
		StartPosition: underAdorn.StartPosition,
		Line:          underAdorn.Line,
		Length:        underAdorn.Length,
	}

	return n
}

type AdornmentNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Rune          rune     `json:"rune"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func (a AdornmentNode) NodeType() NodeType {
	return a.Type
}

type BlankLineNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"nodetype"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func (b BlankLineNode) NodeType() NodeType {
	return b.Type
}

func newBlankLine(i *item) *BlankLineNode {
	return &BlankLineNode{
		Id:            i.Id,
		Type:          NodeBlankLine,
		Text:          i.Text.(string),
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

type ParagraphNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newParagraph(i *item) *ParagraphNode {
	return &ParagraphNode{
		Id:            i.Id,
		Type:          NodeParagraph,
		Text:          i.Text.(string),
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

func (p ParagraphNode) NodeType() NodeType {
	return p.Type
}

type SpaceNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newSpace(i *item) *SpaceNode {
	return &SpaceNode{
		Id:            i.Id,
		Type:          NodeParagraph,
		Text:          i.Text.(string),
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

func (s SpaceNode) NodeType() NodeType {
	return s.Type
}

type BlockQuoteNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	NodeList      NodeList `json:"nodeList"`
}

func newBlockQuote(i *item) *BlockQuoteNode {
	return &BlockQuoteNode{
		Id:            i.Id,
		Type:          NodeParagraph,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

func (b BlockQuoteNode) NodeType() NodeType {
	return b.Type
}

type SystemMessageNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Level         systemMessageLevel `json:"level"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	NodeList      NodeList `json:"nodeList"`
}

func newSystemMessage(i *item, level systemMessageLevel) *SystemMessageNode {
	return &SystemMessageNode{
		Id:            i.Id,
		Type:          NodeParagraph,
		Level:	       level,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

func (s SystemMessageNode) NodeType() NodeType {
	return s.Type
}

type LiteralBlockNode struct {
	Id            int      `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newLiteralBlock(i *item) *LiteralBlockNode {
	return &LiteralBlockNode{
		Id:            i.Id,
		Type:          NodeParagraph,
		Text:	       i.Text.(string),
		Length:	       i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

func (l LiteralBlockNode) NodeType() NodeType {
	return l.Type
}
