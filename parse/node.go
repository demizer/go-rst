// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

// NodeType identifies the type of a parse tree node.
type NodeType int

const (
	NodeSection NodeType = iota
	NodeParagraph
	NodeAdornment
	NodeBlockQuote
	NodeSystemMessage
	NodeLiteralBlock
)

var nodeTypes = [...]string{
	"NodeSection",
	"NodeParagraph",
	"NodeAdornment",
	"NodeBlockQuote",
	"NodeSystemMessage",
	"NodeLiteralBlock",
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
	IdNumber() Id
	LineNumber() Line
	NodeType() NodeType
}

type NodeList []Node

func newList() *NodeList {
	return new(NodeList)
}

func (l *NodeList) append(n Node) {
	*l = append(*l, n)

}

type SectionNode struct {
	Id            `json:"id"`
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

func newSection(title *item, overAdorn *item, underAdorn *item, id *int) *SectionNode {
	*id++
	n := &SectionNode{
		Id:            Id(*id),
		Type:          NodeSection,
		Text:          title.Text.(string),
		StartPosition: title.StartPosition,
		Length:        title.Length,
		Line:          title.Line,
	}

	if overAdorn != nil && overAdorn.Text != nil {
		*id++
		Rune := rune(overAdorn.Text.(string)[0])
		n.OverLine = &AdornmentNode{
			Id:            Id(*id),
			Type:          NodeAdornment,
			Rune:          Rune,
			StartPosition: overAdorn.StartPosition,
			Line:          overAdorn.Line,
			Length:        overAdorn.Length,
		}
	}

	*id++
	Rune := rune(underAdorn.Text.(string)[0])
	n.UnderLine = &AdornmentNode{
		Id:            Id(*id),
		Rune:          Rune,
		Type:          NodeAdornment,
		StartPosition: underAdorn.StartPosition,
		Line:          underAdorn.Line,
		Length:        underAdorn.Length,
	}

	return n
}

type AdornmentNode struct {
	Id            `json:"id"`
	Type          NodeType `json:"type"`
	Rune          rune     `json:"rune"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func (a AdornmentNode) NodeType() NodeType {
	return a.Type
}

type ParagraphNode struct {
	Id            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newParagraph(i *item, id *int) *ParagraphNode {
	*id++
	return &ParagraphNode{
		Id:            Id(*id),
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

type BlockQuoteNode struct {
	Id            `json:"id"`
	Type          NodeType `json:"type"`
	Level         int      `json:"level"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	NodeList      NodeList `json:"nodeList"`
}

func newBlockQuote(i *item, indentLevel int, id *int) *BlockQuoteNode {
	*id++
	return &BlockQuoteNode{
		Id:            Id(*id),
		Type:          NodeBlockQuote,
		Level:         indentLevel,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

func (b BlockQuoteNode) NodeType() NodeType {
	return b.Type
}

type SystemMessageNode struct {
	Id       `json:"id"`
	Type     NodeType           `json:"type"`
	Severity systemMessageLevel `json:"severity"`
	Line     `json:"line"`
	NodeList NodeList `json:"nodeList"`
}

func newSystemMessage(i *item, severity systemMessageLevel, id *int) *SystemMessageNode {
	*id++
	return &SystemMessageNode{
		Id:       Id(*id),
		Type:     NodeSystemMessage,
		Severity: severity,
		Line:     i.Line,
	}
}

func (s SystemMessageNode) NodeType() NodeType {
	return s.Type
}

type LiteralBlockNode struct {
	Id            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func newLiteralBlock(i *item, id *int) *LiteralBlockNode {
	*id++
	return &LiteralBlockNode{
		Id:            Id(*id),
		Type:          NodeLiteralBlock,
		Text:          i.Text.(string),
		Length:        i.Length,
		StartPosition: i.StartPosition,
		Line:          i.Line,
	}
}

func (l LiteralBlockNode) NodeType() NodeType {
	return l.Type
}
