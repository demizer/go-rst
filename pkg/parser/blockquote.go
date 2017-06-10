package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) emptyblockquote(i *tok.Item) {
	//
	//  FIXME: Blockquote parsing is NOT fully implemented.
	//
	sec := doc.NewEmptyBlockQuote(i)
	p.nodeTarget.Append(sec)
	p.nodeTarget.SetParent(sec)
	p.bqLevel = sec
}

func (p *Parser) blockquote(i *tok.Item) {
	//
	//  FIXME: Blockquote parsing is NOT fully implemented.
	//
	if p.bqLevel != nil {
		// Parser Test 03.02.07.00
		p.Msg("Adding blockquote text as NodeText to existing blockquote")
		p.bqLevel.NodeList.Append(doc.NewParagraphWithNodeText(i))
		return
	}
	p.Msg("Creating blockquote")
	sec := doc.NewBlockQuote(i)
	p.nodeTarget.Append(sec)
	p.nodeTarget.SetParent(sec)
	p.bqLevel = sec
}
