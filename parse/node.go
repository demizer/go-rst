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

func NodeTypeFromString(str string) NodeType {
	for num, val := range nodeTypes {
		if val == str {
			return NodeType(num)
		}
	}
	return -1
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

func newSection(item item, level int, overAdorn item, underAdorn item) *SectionNode {
	n := &SectionNode{Text: item.Value.(string),
		Type:          NodeSection,
		Level:         level,
		StartPosition: item.StartPosition,
		Length:        item.Length,
		Line:          1,
	}

	if overAdorn.Value != nil {
		Rune := rune(overAdorn.Value.(string)[0])
		n.OverLine = &AdornmentNode{
			Rune:          Rune,
			Type:          NodeAdornment,
			StartPosition: overAdorn.StartPosition,
			Line:          overAdorn.Line,
			Length:        overAdorn.Length,
		}
	}

	Rune := rune(underAdorn.Value.(string)[0])
	n.UnderLine = &AdornmentNode{
		Rune:          Rune,
		Type:          NodeAdornment,
		StartPosition: underAdorn.StartPosition,
		Line:          underAdorn.Line,
		Length:        underAdorn.Length,
	}

	return n
}

type AdornmentNode struct {
	Type          NodeType `json:"type"`
	Rune          rune     `json:"rune"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func (a AdornmentNode) NodeType() NodeType {
	return a.Type
}

func newBlankLine(i item) *BlankLineNode {
	return &BlankLineNode{
		Type:          NodeBlankLine,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

type BlankLineNode struct {
	Type          NodeType `json:"nodetype"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func (b BlankLineNode) NodeType() NodeType {
	return b.Type
}

type ParagraphNode struct {
	Type          NodeType `json:"type"`
	Text          string   `json:"text"`
	Length        int      `json:"length"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
}

func newParagraph(i item) *ParagraphNode {
	return &ParagraphNode{
		Type:          NodeParagraph,
		Text:          i.Value.(string),
		Length:        i.Length,
		Line:          i.Line,
		StartPosition: i.StartPosition,
	}
}

func (p ParagraphNode) NodeType() NodeType {
	return p.Type
}
