package parser

import (
	"unicode/utf8"

	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) paragraph(i *tok.Item) doc.Node {
	p.Msgr("Have token", "token", i)
	var np doc.ParagraphNode
	if !p.nodeTarget.IsParagraphNode() {
		np := doc.NewParagraph()
		p.nodeTarget.Append(np)
		p.nodeTarget.SetParent(np)
	}
	nt := doc.NewText(i)
	p.nodeTarget.Append(nt)
outer:
	// Paragraphs can contain many different types of elements, so we'll need to loop until blank line or nil
	for {
		ci := p.next(1) // current item
		// p.DumpExit(p.index)
		// p.DumpExit(ci)
		// p.DumpExit(p.buf)

		pi := p.peekBack(1) // previous item
		// p.DumpExit(pi)
		// ni := p.peek(1)     // next item

		// p.DumpExit(p.Messages)
		p.Msgr("Previous token", "token", pi)
		p.Msgr("Have token", "token", ci)

		if ci == nil || ci.Type == tok.EOF {
			p.Msg("current token == nil or current item type == tok.EOF")
			break
			// } else if pi != nil && nt.Text != pi.Text && pi.Type == tok.Text && ci.Type == tok.Text {
		} else if pi != nil && pi.Type == tok.Text && ci.Type == tok.Text {
			// p.DumpExit(ci)
			// p.DumpExit(p.buf)
			p.Msg("Previous type == tok.Text, current type == tok.Text; Concatenating text!")
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
			continue
		}

		p.Msg("Going into subparser...")

		switch ci.Type {
		case tok.Space:
			if pi != nil && pi.Type == tok.Escape {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				p.Msg("Found escaped space!")
				continue
			}
			// Parse Test 02.00.03.00 :: Emphasis wrapped in unicode spaces
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
		case tok.Text:
			if pi != nil && pi.Type == tok.Escape && pi.StartPosition > ci.StartPosition {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				p.Msg("Found newline escape!")
				nt.Text += ci.Text
				nt.Length = utf8.RuneCountInString(nt.Text)
			} else {
				nt = doc.NewText(ci)
				p.nodeTarget.Append(nt)
			}
		case tok.InlineEmphasisOpen:
			p.inlineEmphasis(ci, true)
		case tok.InlineStrongOpen:
			p.inlineStrong(ci, true)
		case tok.InlineLiteralOpen:
			p.inlineLiteral(ci, true)
		case tok.InlineInterpretedTextOpen:
			p.inlineInterpretedText(ci)
		case tok.InlineInterpretedTextRoleOpen:
			p.inlineInterpretedTextRole(ci)
		case tok.CommentMark:
			p.comment(ci)
		case tok.EnumListArabic:
			p.nodeTarget.Append(p.enumList(ci))
		case tok.BlankLine:
			p.Msg("Found newline, closing paragraph")
			p.backup()
			break outer
		}
		p.DumpExit(p.Nodes)
		p.Msg("Continuing...")
	}
	p.Msgr("number of indents", "p.indents.len", p.indents.len())
	if p.indents.len() > 0 {
		p.nodeTarget.SetParent(p.indents.topNode())
		p.Msgr("Set node target to p.indents.topNodeList!", "nodePtr", p.nodeTarget)
	} else if len(p.sectionLevels.levels) == 0 {
		p.Msg("Setting node target to p.nodes!")
		p.nodeTarget.Reset()
	}
	return np
}
