package parser

import (
	"errors"
	"fmt"
	"os"

	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/tokenizer"
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
		log.Log("msg", "Have token", "token", ni)
		pb := p.peekBack(1)
		if ni.Type == tok.ItemSpace {
			log.Msg("continue; ni.Type == ItemSpace")
			continue
		} else if ni.Type == tok.ItemEOF {
			log.Msg("break; ni.Type == ItemEOF")
			break
		} else if ni.Type == tok.ItemBlankLine {
			log.Msg("Setting nodeTarget to dli")
			p.nodeTarget.SetParent(dli.Definition)
		} else if ni.Type == tok.ItemCommentMark && (pb != nil && pb.Type != tok.ItemSpace) {
			// Comment at start of the line breaks current definition list
			log.Msg("Have tok.ItemCommentMark at start of the line!")
			p.nodeTarget.Reset()
			p.backup()
			break
		} else if ni.Type == tok.ItemDefinitionText {
			// FIXME: This function is COMPLETELY not final. It is only setup for passing section test TitleNumberedGood0100.
			np := doc.NewParagraphWithNodeText(ni)
			p.nodeTarget.Append(np)
			p.nodeTarget.SetParent(np)
			log.Msg("continue; ni.Type == tok.ItemDefinitionText")
			continue
		} else if ni.Type == tok.ItemDefinitionTerm {
			dli2 := doc.NewDefinitionListItem(ni, p.peek(2))
			p.nodeTarget.SetParent(dl)
			p.nodeTarget.Append(dli2)
			p.nodeTarget.SetParent(dli2.Definition)
			log.Msg("continue; ni.Type == tok.ItemDefinitionTerm")
			continue
		}
		p.subParseBodyElements(ni)
	}
	return dl
}

func (p *Parser) bulletList(i *tok.Item) {
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
		log.Log("msg", "Have token", "token", fmt.Sprintf("%+#v", ni))
		if ni == nil {
			log.Log("break next item == nil")
			break
		} else if ni.Type == tok.ItemEOF {
			log.Log("break itemEOF")
			break
		} else if p.indents.len() > 0 && len(*p.indents.topNodeList()) > 0 && p.peekBack(1).Type == tok.ItemSpace &&
			p.peekBack(2).Type != tok.ItemCommentMark {
			log.Log("msg", "Have indents",
				"lastStartPosition", p.indents.lastStartPosition(),
				"ni.StartPosition", ni.StartPosition)
			if p.indents.lastStartPosition() != ni.StartPosition {
				// FIXME: WE SHOULD NEVER EXIT IN LIBRARY !! This is just debug code, but we need to add
				// proper handling for this ...
				log.Log(errors.New("Unexpected un-indent!"))
				spd.Dump(p.indents)
				os.Exit(1)
			}
		}

		p.subParseBodyElements(ni)
	}
	p.indents.pop()
}

var lastEnum *doc.EnumListNode

func (p *Parser) enumList(i *tok.Item) (n doc.Node) {
	var eNode *doc.EnumListNode
	var affix *tok.Item
	if lastEnum == nil {
		p.next(1)
		affix = p.token[zed]
		p.next(1)
		eNode = doc.NewEnumListNode(i, affix)
		p.next(1)
		eNode.NodeList.Append(doc.NewParagraphWithNodeText(p.token[zed]))
	} else {
		p.next(3)
		lastEnum.NodeList.Append(doc.NewParagraphWithNodeText(p.token[zed]))
		return nil
	}
	lastEnum = eNode
	return eNode
}
