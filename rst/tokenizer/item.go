package tokenizer

import (
	"encoding/json"
	"fmt"
)

// itemElement are the types that are emitted by the lexer.
type itemElement int

const (
	itemEOF itemElement = iota
	itemError
	itemTitle
	itemSectionAdornment
	itemText
	itemBlockQuote
	itemLiteralBlock
	itemSystemMessage
	itemSpace
	itemBlankLine
	itemTransition
	itemCommentMark
	itemEnumListAffix
	itemEnumListArabic
	itemHyperlinkTargetStart
	itemHyperlinkTargetPrefix
	itemHyperlinkTargetQuote
	itemHyperlinkTargetName
	itemHyperlinkTargetSuffix
	itemHyperlinkTargetURI
	itemInlineStrongOpen
	itemInlineStrong
	itemInlineStrongClose
	itemInlineEmphasisOpen
	itemInlineEmphasis
	itemInlineEmphasisClose
	itemInlineLiteralOpen
	itemInlineLiteral
	itemInlineLiteralClose
	itemInlineInterpretedTextOpen
	itemInlineInterpretedText
	itemInlineInterpretedTextClose
	itemInlineInterpretedTextRoleOpen
	itemInlineInterpretedTextRole
	itemInlineInterpretedTextRoleClose
	itemInlineReferenceOpen
	itemInlineReferenceText
	itemInlineReferenceClose
	itemDefinitionTerm
	itemDefinitionText
	itemBullet
	itemEscape
)

var elements = [...]string{
	"itemEOF",
	"itemError",
	"itemTitle",
	"itemSectionAdornment",
	"itemText",
	"itemBlockQuote",
	"itemLiteralBlock",
	"itemSystemMessage",
	"itemSpace",
	"itemBlankLine",
	"itemTransition",
	"itemCommentMark",
	"itemEnumListAffix",
	"itemEnumListArabic",
	"itemHyperlinkTargetStart",
	"itemHyperlinkTargetPrefix",
	"itemHyperlinkTargetQuote",
	"itemHyperlinkTargetName",
	"itemHyperlinkTargetSuffix",
	"itemHyperlinkTargetURI",
	"itemInlineStrongOpen",
	"itemInlineStrong",
	"itemInlineStrongClose",
	"itemInlineEmphasisOpen",
	"itemInlineEmphasis",
	"itemInlineEmphasisClose",
	"itemInlineLiteralOpen",
	"itemInlineLiteral",
	"itemInlineLiteralClose",
	"itemInlineInterpretedTextOpen",
	"itemInlineInterpretedText",
	"itemInlineInterpretedTextClose",
	"itemInlineInterpretedTextRoleOpen",
	"itemInlineInterpretedTextRole",
	"itemInlineInterpretedTextRoleClose",
	"itemInlineReferenceOpen",
	"itemInlineReferenceText",
	"itemInlineReferenceClose",
	"itemDefinitionTerm",
	"itemDefinitionText",
	"itemBullet",
	"itemEscape",
}

// String implements the Stringer interface for printing itemElement types.
func (t itemElement) String() string { return elements[t] }

func (t *itemElement) UnmarshalJSON(data []byte) error {
	for num, elm := range elements {
		if elm == string(data[1:len(data)-1]) {
			*t = itemElement(num)
		}
	}
	return nil
}

// Struct for tokens emitted by the scanning process
type item struct {
	ID            `json:"id"`
	Type          itemElement `json:"type"`
	Text          string      `json:"text"`
	Line          `json:"line"`
	StartPosition `json:"startPosition"`
	Length        int `json:"length"`
}

// MarshalJSON satisfies the Marshaler interface.
func (i item) MarshalJSON() ([]byte, error) {
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
func (i *item) String() string {
	return fmt.Sprintf("ID=%d Type=%s text=%q Line=%d StartPosition=%d Length=%d",
		i.ID, i.Type, i.Text, i.Line, i.StartPosition, i.Length)
}
