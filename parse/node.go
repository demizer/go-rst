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
	NodeIndent
)

var nodeTypes = [...]string{
	"NodeSection",
	"NodeParagraph",
	"NodeAdornment",
	"NodeBlockQuote",
	"NodeSystemMessage",
	"NodeLiteralBlock",
	"NodeIndent",
}

func (n NodeType) Type() NodeType {
	return n
}

func (n NodeType) String() string {
	return nodeTypes[n]
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
	Indent        *IndentNode    `json:"underLine"`
	NodeList      NodeList       `json:"nodeList"`
}

func (s *SectionNode) NodeType() NodeType {
	return s.Type
}

func newSection(title *item, overSec *item, underSec *item, indent *item, id *int) *SectionNode {
	*id++
	n := &SectionNode{
		Id:            Id(*id),
		Type:          NodeSection,
		Text:          title.Text.(string),
		StartPosition: title.StartPosition,
		Length:        title.Length,
		Line:          title.Line,
	}

	if indent != nil && indent.Text != nil {
		*id++
		n.Indent = &IndentNode{
			Id:            Id(*id),
			Type:          NodeIndent,
			Text:	       indent.Text.(string),
			StartPosition: indent.StartPosition,
			Line:          indent.Line,
			Length:        indent.Length,
		}
	}

	if overSec != nil && overSec.Text != nil {
		*id++
		Rune := rune(overSec.Text.(string)[0])
		n.OverLine = &AdornmentNode{
			Id:            Id(*id),
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
		Id:            Id(*id),
		Rune:          Rune,
		Type:          NodeAdornment,
		StartPosition: underSec.StartPosition,
		Line:          underSec.Line,
		Length:        underSec.Length,
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

type IndentNode struct {
	Id            `json:"id"`
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	StartPosition `json:"startPosition"`
	Line          `json:"line"`
}

func (i IndentNode) NodeType() NodeType {
	return i.Type
}
