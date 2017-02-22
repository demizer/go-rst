package parser

import (
	"unicode/utf8"

	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/token"
)

func (p *Parser) paragraph(i *tok.Item) doc.Node {
	log.Log("msg", "Have token", "token", i)
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
		ci := p.next(1)     // current item
		pi := p.peekBack(1) // previous item
		// ni := p.peek(1)     // next item

		log.Log("msg", "Have token", "token", ci)

		if ci == nil {
			log.Msg("ci == nil, breaking")
			break
		} else if ci.Type == tok.EOF {
			log.Msg("current item type == tok.EOF")
			break
		} else if pi != nil && pi.Type == tok.Text && ci.Type == tok.ItemText {
			log.Msg("Previous type == tok.Text, current type == tok.ItemText; Concatenating text!")
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
			continue
		}

		log.Msg("Going into subparser...")

		switch ci.Type {
		case tok.Space:
			if pi != nil && pi.Type == tok.Escape {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				log.Msg("Found escaped space!")
				continue
			}
			// Parse Test 02.00.03.00 :: Emphasis wrapped in unicode spaces
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
		case tok.Text:
			if pi != nil && pi.Type == tok.Escape && pi.StartPosition.Int() > ci.StartPosition.Int() {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				log.Msg("Found newline escape!")
				nt.Text += ci.Text
				nt.Length = utf8.RuneCountInString(nt.Text)
			} else {
				nt = doc.NewText(ci)
				p.nodeTarget.Append(nt)
			}
		case tok.InlineEmphasisOpen:
			p.inlineEmphasis(ci)
		case tok.InlineStrongOpen:
			p.inlineStrong(ci)
		case tok.InlineLiteralOpen:
			p.inlineLiteral(ci)
		case tok.InlineInterpretedTextOpen:
			p.inlineInterpretedText(ci)
		case tok.InlineInterpretedTextRoleOpen:
			p.inlineInterpretedTextRole(ci)
		case tok.CommentMark:
			p.comment(ci)
		case tok.EnumListArabic:
			p.nodeTarget.Append(p.enumList(ci))
		case tok.BlankLine:
			log.Msg("Found newline, closing paragraph")
			p.backup()
			break outer
		}
		log.Msg("Continuing...")
	}
	log.Log("msg", "number of indents", "p.indents.len", p.indents.len())
	if p.indents.len() > 0 {
		p.nodeTarget.SetParent(p.indents.topNode())
		log.Log("msg", "Set node target to p.indents.topNodeList!", "nodePtr", p.nodeTarget)
	} else if len(p.sectionLevels.levels) == 0 {
		log.Msg("Setting node target to p.nodes!")
		p.nodeTarget.Reset()
	}
	return np
}
