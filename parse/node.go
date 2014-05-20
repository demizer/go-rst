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

func newSection(i item, id *int, level int, overAdorn item, underAdorn item) *SectionNode {
	*id++
	n := &SectionNode{
		Id:            *id,
		Type:          NodeSection,
		Text:          i.Text.(string),
		Level:         level,
		StartPosition: i.StartPosition,
		Length:        i.Length,
		Line:          i.Line,
	}

	if overAdorn.Text != nil {
		*id++
		Rune := rune(overAdorn.Text.(string)[0])
		n.OverLine = &AdornmentNode{
			Id:            *id,
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
		Id:            *id,
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
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func (b BlankLineNode) NodeType() NodeType {
	return b.Type
}

func newBlankLine(i item, id *int) *BlankLineNode {
	*id++
	return &BlankLineNode{
		Id:            *id,
		Type:          NodeBlankLine,
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

func newParagraph(i item, id *int) *ParagraphNode {
	*id++
	return &ParagraphNode{
		Id:            *id,
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
