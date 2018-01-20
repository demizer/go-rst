package parser

import (
	"unicode/utf8"

	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) paragraph(i *tok.Item) doc.Node {
	p.printToken("Have token", i)
	// panic("foo")
	var np doc.ParagraphNode
	// pBack := p.peekBack(1)
	// if i.Line == 18 {
	// p.DumpExit(i)
	// }

	tmp := p.peekLine(p.token.Line - 1)
	newParagraph := tmp != nil && tmp.Type == tok.BlankLine
	if newParagraph && p.nodeTarget.IsParagraphNode() { //&& p.bqLevel == nil {
		// p.nodeTarget.Reset()
		// p.nodeTarget.SetParent(
		p.nodeTarget.SetParent(p.sectionLevels.lastSectionNode)
		// p.DumpExit(i)
		// p.DumpExit(p.token)
	}
	// if newParagraph && !p.nodeTarget.IsParagraphNode() { //&& p.bqLevel == nil {
	if !p.nodeTarget.IsParagraphNode() { //&& p.bqLevel == nil {
		// if !p.nodeTarget.IsParagraphNode() { //&& p.bqLevel == nil {
		np := doc.NewParagraph()
		p.nodeTarget.Append(np)
		p.nodeTarget.SetParent(np)
	}
	nt := doc.NewText(i)
	p.nodeTarget.Append(nt)
	// if i.Type == tok.Text && i.Line == 7 {
	// p.DumpExit(p.bqLevel)
	// }
outer:
	// Paragraphs can contain many different types of elements, so we'll need to loop until blank line or nil
	for {
		ci := p.next(1) // current item
		// p.DumpExit(p.index)
		// p.DumpExit(ci)
		// p.DumpExit(p.buf)

		pi := p.peekBack(1) // previous item
		// p.DumpExit(pi)
		ni := p.peek(1) // next item

		// p.DumpExit(p.Messages)
		p.printToken("Previous token", pi)
		p.printToken("Have token", ci)

		if pi.ID == 99 {
			// p.DumpExit(ni)
			p.dumpBufferWithContext()
		}

		// if ci == nil || ci.Type == tok.EOF || ci.Type == tok.Title {
		if ci == nil || ci.Type == tok.EOF {
			p.Msg("current token == nil or current item type == tok.EOF")
			break
		} else if pi != nil && pi.Type == tok.Text && ci.Type == tok.Text {
			p.Msg("Found two sequential tok.Text! Concatenating text!")
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
			continue
		} else if pi != nil && pi.Type == tok.Text && ci.Type == tok.Escape && ni != nil && ni.Type == tok.Text {
			p.Msg("Found escaped newline! Concatenating text!")
			nt.Text += ni.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
			p.next(1)
			continue
		}

		p.Msg("Going into subparser...")

		switch ci.Type {
		case tok.Space:
			if pi != nil && pi.Type == tok.Escape {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				p.Msg("Found escaped space!")
				continue
			} else if ni != nil && ni.Type == tok.Title {
				// Parse test 04.01.03.00 :: Indented section title
				// Need to make sure the space is not before a title
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
		default:
			p.printToken("token not supported in paragraphs", ci)
			return np
		}
		p.Msg("Continuing...")
		// p.DumpExit(p.buf)
		// panic("halt")
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
