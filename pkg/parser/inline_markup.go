package parser

import (
	"unicode/utf8"

	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) inlineEmphasis(i *tok.Item, titleCheck bool) {
	p.printToken("have token", i)
	// if i.ID == 104 {
	// p.DumpExit(p.Nodes)
	// }
	// Make sure inline markup is not in a section title
	isInTitle := p.isInlineMarkupInSectionTitle(i)
	// p.DumpExit(isInTitle)
	if titleCheck && isInTitle {
		return
	}
	ni := p.next(1)
	p.printToken("have next token", ni)
	// if i.ID == 100 {
	// p.DumpExit(ni)
	// }
	if len(*p.Nodes) == 0 && !isInTitle {
		np := doc.NewParagraph()
		p.nodeTarget.Append(np)
		p.nodeTarget.SetParent(np)
	}
main:
	for {
		ci := p.next(1)
		p.printToken("have token in loop", ci)
		if ci == nil {
			break
		}

		switch ci.Type {
		case tok.InlineEmphasis:
			ni.Text += "\n" + ci.Text
		case tok.BlankLine:
			continue
		default:
			p.backup()
			break main
		}
		if i.ID == 100 {
			p.DumpExit(ci)
		}
	}
	if ni.Text == "*" {
		// p.DumpExit(ni)
		p.dumpBufferWithContext()
	}
	ni.Length = utf8.RuneCountInString(ni.Text)
	p.nodeTarget.Append(doc.NewInlineEmphasis(ni))
	p.next(1)
}

func (p *Parser) inlineStrong(i *tok.Item, titleCheck bool) {
	// Make sure inline markup is not in a section title
	isInTitle := p.isInlineMarkupInSectionTitle(i)
	if titleCheck && isInTitle {
		return
	}
	ni := p.next(1)
	if len(*p.Nodes) == 0 && !isInTitle {
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
		case tok.InlineStrong:
			ni.Text += "\n" + ci.Text
		case tok.BlankLine:
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

func (p *Parser) inlineLiteral(i *tok.Item, titleCheck bool) {
	// Make sure inline markup is not in a section title
	isInTitle := p.isInlineMarkupInSectionTitle(i)
	if titleCheck && isInTitle {
		return
	}
	ni := p.next(1)
	if len(*p.Nodes) == 0 && !isInTitle {
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
		case tok.InlineLiteral:
			ni.Text += "\n" + ci.Text
		case tok.BlankLine:
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
	n := doc.NewInlineInterpretedText(p.token)
	p.nodeTarget.Append(n)
	p.next(1)
	if p.peek(1).Type == tok.InlineInterpretedTextRoleOpen {
		p.next(2)
		n.NodeList.Append(doc.NewInlineInterpretedTextRole(p.token))
		p.next(1)
	}
}

func (p *Parser) inlineInterpretedTextRole(i *tok.Item) {
	p.next(1)
	p.nodeTarget.Append(doc.NewInlineInterpretedTextRole(p.token))
	p.next(1)
}
