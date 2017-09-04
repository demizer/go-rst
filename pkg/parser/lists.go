package parser

import (
	"errors"
	"fmt"

	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) definitionTerm(i *tok.Item) doc.Node {
	//
	//  FIXME: Definition list parsing is NOT fully implemented.
	//
	dl := doc.NewDefinitionList(&tok.Item{Line: i.Line})
	p.nodeTarget.Append(dl)
	p.nodeTarget.SetParent(dl)
	p.next(1)

	// Container for definition items
	dli := doc.NewDefinitionListItem(i, p.peek(1))
	p.nodeTarget.Append(dli)
	p.nodeTarget.SetParent(dli.Definition)

	// Gather definitions and body elements
	for {
		ni := p.next(1)
		if ni == nil {
			break
		}
		p.Msgr("Have token", "token", ni)
		pb := p.peekBack(1)
		if ni.Type == tok.Space {
			p.Msg("continue; ni.Type == Space")
			continue
		} else if ni.Type == tok.EOF {
			p.Msg("break; ni.Type == EOF")
			break
		} else if ni.Type == tok.BlankLine {
			p.Msg("Setting nodeTarget to dli")
			p.nodeTarget.SetParent(dli.Definition)
		} else if ni.Type == tok.CommentMark && (pb != nil && pb.Type != tok.Space) {
			// Comment at start of the line breaks current definition list
			p.Msg("Have tok.CommentMark at start of the line!")
			p.nodeTarget.Reset()
			p.backup()
			break
		} else if ni.Type == tok.DefinitionText {
			// FIXME: This function is COMPLETELY not final. It is only setup for passing section test TitleNumberedGood0100.
			np := doc.NewParagraphWithNodeText(ni)
			p.nodeTarget.Append(np)
			p.nodeTarget.SetParent(np)
			p.Msg("continue; ni.Type == tok.DefinitionText")
			continue
		} else if ni.Type == tok.DefinitionTerm {
			dli2 := doc.NewDefinitionListItem(ni, p.peek(2))
			p.nodeTarget.SetParent(dl)
			p.nodeTarget.Append(dli2)
			p.nodeTarget.SetParent(dli2.Definition)
			p.Msg("continue; ni.Type == tok.DefinitionTerm")
			continue
		}
		p.subParseBodyElements(ni)
	}
	return dl
}

func (p *Parser) bulletList(i *tok.Item) error {
	//
	// FIXME: Bullet List Parsing is NOT fully implemented
	//
	bl := doc.NewBulletListNode(i)
	p.openList = bl
	p.nodeTarget.Append(bl)
	p.nodeTarget.SetParent(bl)

	// Get the bullet list paragraph
	p.next(1)
	bli := doc.NewBulletListItemNode(i)

	// Set the node target to the bullet list paragraph
	p.nodeTarget.Append(bli)
	p.nodeTarget.SetParent(bli)
	p.indents.add(i, bli)

	// Capture all bullet items until un-indent
	for {
		ni := p.next(1)
		topNodeList, err := p.indents.topNodeList()
		if err != nil {
			p.Err(err)
		}
		p.Msgr("Have token", "token", fmt.Sprintf("%+#v", ni))
		if ni == nil {
			p.Log("break next item == nil")
			break
		} else if ni.Type == tok.EOF {
			p.Log("break itemEOF")
			break
		} else if p.indents.len() > 0 && len(*topNodeList) > 0 && p.peekBack(1).Type == tok.Space &&
			p.peekBack(2).Type != tok.CommentMark {
			lastStartPos, err := p.indents.lastStartPosition()
			if err != nil {
				p.Err(err)
			}
			p.Msgr("Have indents", "lastStartPosition", lastStartPos, "ni.StartPosition", ni.StartPosition)
			if lastStartPos != ni.StartPosition {
				return errors.New("Unexpected un-indent!")
			}
		}

		p.subParseBodyElements(ni)
	}
	p.indents.pop()
	return nil
}

var lastEnum *doc.EnumListNode

func (p *Parser) enumList(i *tok.Item) (n doc.Node) {
	var eNode *doc.EnumListNode
	var affix *tok.Item
	if lastEnum == nil {
		p.next(1)
		affix = p.token
		p.next(1)
		eNode = doc.NewEnumListNode(i, affix)
		p.next(1)
		eNode.NodeList.Append(doc.NewParagraphWithNodeText(p.token))
	} else {
		p.next(3)
		lastEnum.NodeList.Append(doc.NewParagraphWithNodeText(p.token))
		return nil
	}
	lastEnum = eNode
	return eNode
}
