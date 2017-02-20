package parser

import (
	"unicode/utf8"

	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/tokenizer"
)

func (p *Parser) inlineEmphasis(i *tok.Item) {
	ni := p.next(1)
	if len(*p.Nodes) == 0 {
		np := doc.NewParagraph()
		p.nodeTarget.Append(np)
		p.nodeTarget.SetParent(np)
	}
main:
	for {
		ci := p.next(1)
		if ci == nil {
			break
		}
		switch ci.Type {
		case tok.ItemInlineEmphasis:
			ni.Text += "\n" + ci.Text
		case tok.ItemBlankLine:
			continue
		default:
			p.backup()
			break main
		}
	}
	ni.Length = utf8.RuneCountInString(ni.Text)
	p.nodeTarget.Append(doc.NewInlineEmphasis(ni))
	p.next(1)
}

func (p *Parser) inlineStrong(i *tok.Item) {
	ni := p.next(1)
	if len(*p.Nodes) == 0 {
		np := doc.NewParagraph()
		p.nodeTarget.Append(np)
		p.nodeTarget.SetParent(np)
	}
main:
	for {
		ci := p.next(1)
		if ci == nil {
			break
		}
		switch ci.Type {
		case tok.ItemInlineStrong:
			ni.Text += "\n" + ci.Text
		case tok.ItemBlankLine:
			continue
		default:
			p.backup()
			break main
		}
	}
	ni.Length = utf8.RuneCountInString(ni.Text)
	p.nodeTarget.Append(doc.NewInlineStrong(ni))
	p.next(1)
}

func (p *Parser) inlineLiteral(i *tok.Item) {
	ni := p.next(1)
	if len(*p.Nodes) == 0 {
		np := doc.NewParagraph()
		p.nodeTarget.Append(np)
		p.nodeTarget.SetParent(np)
	}
main:
	for {
		ci := p.next(1)
		if ci == nil {
			break
		}
		switch ci.Type {
		case tok.ItemInlineLiteral:
			ni.Text += "\n" + ci.Text
		case tok.ItemBlankLine:
			continue
		default:
			p.backup()
			break main
		}
	}
	ni.Length = utf8.RuneCountInString(ni.Text)
	p.nodeTarget.Append(doc.NewInlineLiteral(ni))
	p.next(1)
}

func (p *Parser) inlineInterpretedText(i *tok.Item) {
	p.next(1)
	n := doc.NewInlineInterpretedText(p.token[zed])
	p.nodeTarget.Append(n)
	p.next(1)
	if p.peek(1).Type == tok.ItemInlineInterpretedTextRoleOpen {
		p.next(2)
		n.NodeList.Append(doc.NewInlineInterpretedTextRole(p.token[zed]))
		p.next(1)
	}
}

func (p *Parser) inlineInterpretedTextRole(i *tok.Item) {
	p.next(1)
	p.nodeTarget.Append(doc.NewInlineInterpretedTextRole(p.token[zed]))
	p.next(1)
}
