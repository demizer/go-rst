package parser

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"golang.org/x/text/unicode/norm"
	// doc "github.com/demizer/go-rst/rst/document"
)

// Used for debugging only
var spd = spew.ConfigState{ContinueOnMethod: true, Indent: "\t", MaxDepth: 0} //, DisableMethods: true}

const (
	// The middle of the ParserState.token buffer so that there are three possible "backup" token positions and three
	// forward "peek" positions.
	zed = 4

	// Default indent width
	indentWidth = 4
)

// ParserState contains the parser ParserState. The Nodes field contains the parsed nodes of the input input data.
type ParserState struct {
	Name string // The name of the current parser input
	// Nodes    *NodeList // The root node list
	// Messages NodeList  // Messages generated by the parser

	// nodeTarget *nodeTarget  // Used to append nodes to a target NodeList
	text    string       // The input text
	lex     *lexer       // The place where tokens come from
	token   [9]*item     // Token buffer, number 4 is the middle. AKA the "zed" token
	indents *indentQueue // Indent level tracking

	bqLevel *BlockQuoteNode // FIXME: will be replaced with blockquoteLevels

	sectionLevels *sectionLevels // Encountered section levels
	sections      []*SectionNode // Pointers to encountered sections

	openList Node // Open Bullet List, Enum List, or Definition List
}

// New returns a fresh parser ParserState.
func New(name, text string) *ParserState {
	nl := make(NodeList, 0)
	logp.Log("msg", "ParserState.Nodes pointer", "nodeListPointer", fmt.Sprintf("%p", nl))
	t := &ParserState{
		Name:          name,
		Nodes:         &nl,
		text:          text,
		sectionLevels: new(sectionLevels),
		indents:       new(indentQueue),
		nodeTarget:    newNodeTarget(&nl),
	}
	return t
}

// Parse is the entry point for the reStructuredText parser. Errors generated by the parser are returned as a NodeList.
func Parse(name, text string) (t *ParserState, errors NodeList) {
	t = New(name, text)
	if !norm.NFC.IsNormalString(text) {
		text = norm.NFC.String(text)
	}
	t.Parse(text, t)
	errors = t.Messages
	return
}

// startParse initializes the parser, using the lexer.
func (t *ParserState) startParse(lex *lexer) {
	t.lex = lex
}

// Parse activates the parser using text as input data. A parse ParserState is returned on success or failure. Users of the
// Parse package should use the Top level Parse function.
func (t *ParserState) Parse(text string, p *ParserState) *ParserState {
	t.startParse(lex(t.Name, []byte(text)))
	t.text = text
	t.parse(p)
	return t
}

// parse is where items are retrieved from the parser and dispatched according to the itemElement type.
func (t *ParserState) parse(p *ParserState) {
	for {
		var n Node

		token := t.next(1)
		if token == nil || token.Type == itemEOF {
			break
		}

		logp.Log("msg", "Parser got token", "token", token)

		switch token.Type {
		case itemText:
			t.paragraph(token)
		case itemInlineEmphasisOpen:
			t.inlineEmphasis(token)
		case itemInlineStrongOpen:
			t.inlineStrong(token)
		case itemInlineLiteralOpen:
			t.inlineLiteral(token)
		case itemInlineInterpretedTextOpen:
			t.inlineInterpretedText(token)
		case itemInlineInterpretedTextRoleOpen:
			t.inlineInterpretedTextRole(token)
		case itemTransition:
			// FIXME: Workaround until transitions are supported
			t.nodeTarget.append(newTransition(token))
		case itemCommentMark:
			t.comment(token)
		case itemSectionAdornment:
			t.section(token)
		case itemEnumListArabic:
			n = t.enumList(token)
			// FIXME: This is only until enumerated list are properly implemented.
			if n == nil {
				continue
			}
			t.nodeTarget.append(n)
		case itemSpace:
			//
			//  FIXME: Blockquote parsing is NOT fully implemented.
			//
			if t.peekBack(1).Type == itemBlankLine && t.bqLevel == nil {
				// Ignore if next item is a blockquote from the lexer
				if pn := t.peek(1); pn != nil && pn.Type == itemBlockQuote {
					logp.Msg("Next item is blockquote; not creating empty blockquote")
					continue
				}
				logp.Msg("Creating empty blockquote!")
				t.emptyblockquote(token)
			} else if t.peekBack(1).Type == itemBlankLine {
				t.nodeTarget.setParent(t.bqLevel)
			}
		case itemBlankLine, itemTitle, itemEscape:
			// itemTitle is consumed when evaluating itemSectionAdornment
			continue
		case itemBlockQuote:
			t.blockquote(token)
		case itemDefinitionTerm:
			t.definitionTerm(token)
		case itemBullet:
			t.bulletList(token)
		default:
			logp.Msg(fmt.Sprintf("Token type: %q is not yet supported in the parser", token.Type.String()))
		}

	}
}

func (t *ParserState) subParseBodyElements(token *item) Node {
	logp.Log("msg", "Have token", "tokenType", token.Type, "tokenText", fmt.Sprintf("%q", token.Text))
	var n Node
	switch token.Type {
	case itemText:
		n = t.paragraph(token)
	case itemInlineEmphasisOpen:
		t.inlineEmphasis(token)
	case itemInlineStrongOpen:
		t.inlineStrong(token)
	case itemInlineLiteralOpen:
		t.inlineLiteral(token)
	case itemInlineInterpretedTextOpen:
		t.inlineInterpretedText(token)
	case itemInlineInterpretedTextRoleOpen:
		t.inlineInterpretedTextRole(token)
	case itemCommentMark:
		t.comment(token)
	case itemEnumListArabic:
		t.enumList(token)
	case itemSpace:
	case itemBlankLine, itemEscape:
	case itemBlockQuote:
		t.blockquote(token)
	default:
		logp.Msg(fmt.Sprintf("Token type: %q is not yet supported in the parser", token.Type.String()))
	}
	return n
}

// backup shifts the token buffer right one position.
func (t *ParserState) backup() {
	t.token[0] = nil
	for x := len(t.token) - 1; x > 0; x-- {
		t.token[x] = t.token[x-1]
		t.token[x-1] = nil
	}
	if t.token[zed] == nil {
		logp.Msg("Current token is: <nil>")
	} else {
		logp.Msg(fmt.Sprintf("Current token is: %T", t.token[zed].Type))
	}
}

// peekBack uses the token buffer to "look back" a number of positions (pos). Looking back more positions than the
// ParserState.token
// buffer allows (3) will generate a panic.
func (t *ParserState) peekBack(pos int) *item {
	return t.token[zed-pos]
}

func (t *ParserState) peekBackTo(item itemElement) (tok *item) {
	for i := zed - 1; i >= 0; i-- {
		if t.token[i] != nil && t.token[i].Type == item {
			tok = t.token[i]
			break
		}
	}
	return
}

// peek looks ahead in the token stream a number of positions (pos) and gets the next token from the lexer. A pointer to the
// token is kept in the ParserState.token buffer. If a token pointer already exists in the buffer, that token is used instead
// and no tokens are received the the lexer stream (channel).
func (t *ParserState) peek(pos int) *item {
	nItem := t.token[zed]
	for i := 1; i <= pos; i++ {
		if t.token[zed+i] != nil {
			nItem = t.token[zed+i]
			logp.Log("msg", "Have token", "token", nItem)
			continue
		} else {
			if t.lex == nil {
				continue
			}
			logp.Msg("Getting next item")
			t.token[zed+i] = t.lex.nextItem()
			nItem = t.token[zed+i]
		}
	}
	return nItem
}

// peekSkip looks ahead one position skipiing a specified itemElement. If that element is found, a pointer is returned,
// otherwise nil is returned.
func (t *ParserState) peekSkip(iSkip itemElement) *item {
	var nItem *item
	count := 1
	for {
		nItem = t.peek(count)
		if nItem.Type != iSkip {
			break
		}
		count++
	}
	return nItem
}

// next token already exists in the token buffer, than the token buffer is shifted left and the pointer to the "zed" token is
// returned. pos specifies the number of times to call next.
func (t *ParserState) next(pos int) *item {
	if pos == 0 {
		return t.token[zed]
	}
	for x := 0; x < len(t.token)-1; x++ {
		t.token[x] = t.token[x+1]
		t.token[x+1] = nil
	}
	if t.token[zed] == nil && t.lex != nil {
		t.token[zed] = t.lex.nextItem()
	}
	pos--
	if pos > 0 {
		t.next(pos)
	}
	return t.token[zed]
}

// clearTokens sets tokens from begin to end to nil.
func (t *ParserState) clearTokens(begin, end int) {
	for i := begin; i <= end; i++ {
		t.token[i] = nil
	}
}
