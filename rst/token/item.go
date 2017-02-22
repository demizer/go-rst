package token

import (
	"encoding/json"
	"fmt"
	. "github.com/demizer/go-rst"
)

// Type are the types that are emitted by the lexer.
type Type int

const (
	ItemEOF Type = iota
	ItemError
	ItemTitle
	ItemSectionAdornment
	ItemText
	ItemBlockQuote
	ItemLiteralBlock
	ItemSystemMessage
	ItemSpace
	ItemBlankLine
	ItemTransition
	ItemCommentMark
	ItemEnumListAffix
	ItemEnumListArabic
	ItemHyperlinkTargetStart
	ItemHyperlinkTargetPrefix
	ItemHyperlinkTargetQuote
	ItemHyperlinkTargetName
	ItemHyperlinkTargetSuffix
	ItemHyperlinkTargetURI
	ItemInlineStrongOpen
	ItemInlineStrong
	ItemInlineStrongClose
	ItemInlineEmphasisOpen
	ItemInlineEmphasis
	ItemInlineEmphasisClose
	ItemInlineLiteralOpen
	ItemInlineLiteral
	ItemInlineLiteralClose
	ItemInlineInterpretedTextOpen
	ItemInlineInterpretedText
	ItemInlineInterpretedTextClose
	ItemInlineInterpretedTextRoleOpen
	ItemInlineInterpretedTextRole
	ItemInlineInterpretedTextRoleClose
	ItemInlineReferenceOpen
	ItemInlineReferenceText
	ItemInlineReferenceClose
	ItemDefinitionTerm
	ItemDefinitionText
	ItemBullet
	ItemEscape
)

var elements = [...]string{
	"ItemEOF",
	"ItemError",
	"ItemTitle",
	"ItemSectionAdornment",
	"ItemText",
	"ItemBlockQuote",
	"ItemLiteralBlock",
	"ItemSystemMessage",
	"ItemSpace",
	"ItemBlankLine",
	"ItemTransition",
	"ItemCommentMark",
	"ItemEnumListAffix",
	"ItemEnumListArabic",
	"ItemHyperlinkTargetStart",
	"ItemHyperlinkTargetPrefix",
	"ItemHyperlinkTargetQuote",
	"ItemHyperlinkTargetName",
	"ItemHyperlinkTargetSuffix",
	"ItemHyperlinkTargetURI",
	"ItemInlineStrongOpen",
	"ItemInlineStrong",
	"ItemInlineStrongClose",
	"ItemInlineEmphasisOpen",
	"ItemInlineEmphasis",
	"ItemInlineEmphasisClose",
	"ItemInlineLiteralOpen",
	"ItemInlineLiteral",
	"ItemInlineLiteralClose",
	"ItemInlineInterpretedTextOpen",
	"ItemInlineInterpretedText",
	"ItemInlineInterpretedTextClose",
	"ItemInlineInterpretedTextRoleOpen",
	"ItemInlineInterpretedTextRole",
	"ItemInlineInterpretedTextRoleClose",
	"ItemInlineReferenceOpen",
	"ItemInlineReferenceText",
	"ItemInlineReferenceClose",
	"ItemDefinitionTerm",
	"ItemDefinitionText",
	"ItemBullet",
	"ItemEscape",
}

// String implements the Stringer interface for printing Type types.
func (t Type) String() string { return elements[t] }

func (t *Type) UnmarshalJSON(data []byte) error {
	for num, elm := range elements {
		if elm == string(data[1:len(data)-1]) {
			*t = Type(num)
		}
	}
	return nil
}

// Struct for tokens emitted by the scanning process
type Item struct {
	ID            `json:"id"`
	Type          Type   `json:"type"`
	Text          string `json:"text"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	Length        int `json:"length"`
}

// MarshalJSON satisfies the Marshaler interface.
func (i Item) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		ID            int    `json:"id"`
		Type          string `json:"type"`
		Text          string `json:"text"`
		Line          int    `json:"line"`
		StartPosition int    `json:"startPosition"`
		Length        int    `json:"length"`
	}{
		ID:            int(i.IDNumber()),
		Type:          i.Type.String(),
		Text:          i.Text,
		Line:          int(i.Line),
		StartPosition: i.StartPosition.Int(),
		Length:        i.Length,
	})
}

// String satisfies the Stringer interface.
func (i *Item) String() string {
	return fmt.Sprintf("ID=%d Type=%s text=%q Line=%d StartPosition=%d Length=%d",
		i.ID, i.Type, i.Text, i.Line, i.StartPosition, i.Length)
}
