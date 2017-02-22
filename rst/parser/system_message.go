package parser

import (
	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/token"
)

func (p *Parser) systemMessageSection(s *doc.SystemMessageNode, err parserMessage) *doc.LiteralBlockNode {
	var overLine, indent, title, underLine, newLine string
	var lbText string
	var lbTextLen int

	literalBlock := func() *doc.LiteralBlockNode {
		return doc.NewLiteralBlock(&tok.Item{Type: tok.ItemLiteralBlock, Text: lbText, Length: lbTextLen})
	}

	switch err {
	case infoOverlineTooShortForTitle:
		var inText string
		inText = p.token[zed-1].Text + "\n" + p.token[zed].Text
		s.Line = p.token[zed-1].Line
		if p.token[zed-2] != nil {
			inText = p.token[zed-2].Text + "\n" + p.token[zed-1].Text + "\n" + p.token[zed].Text
			s.Line = p.token[zed-2].Line
			p.token[zed-2] = nil

		}
		infoTextLen := len(inText)
		// Modify the token buffer to change the current token to a tok.ItemText then backup the token buffer so the
		// next loop gets the new paragraph
		p.token[zed-1] = nil
		p.token[zed].Type = tok.ItemText
		p.token[zed].Text = inText
		p.token[zed].Length = infoTextLen
		p.token[zed].Line = s.Line
		p.backup()
	case infoUnexpectedTitleOverlineOrTransition:
		oLin := p.peekBackTo(tok.ItemSectionAdornment)
		titl := p.peekBackTo(tok.Title)
		uLin := p.token[zed]
		inText := oLin.Text + "\n" + titl.Text + "\n" + uLin.Text
		s.Line = oLin.Line
		p.clearTokens(zed-4, zed-1)
		infoTextLen := len(inText)
		// Modify the token buffer to change the current token to a tok.ItemText then backup the token buffer so the
		// next loop gets the new paragraph
		p.token[zed].Type = tok.ItemText
		p.token[zed].Text = inText
		p.token[zed].Length = infoTextLen
		p.token[zed].Line = s.Line
		p.token[zed].StartPosition = oLin.StartPosition
		p.backup()
	case infoUnderlineTooShortForTitle:
		inText := p.token[zed-1].Text + "\n" + p.token[zed].Text
		infoTextLen := len(inText)
		s.Line = p.token[zed-1].Line
		// Modify the token buffer to change the current token to a tok.ItemText then backup the token buffer so the
		// next loop gets the new paragraph
		p.token[zed-1] = nil
		p.token[zed].Type = tok.ItemText
		p.token[zed].Text = inText
		p.token[zed].Length = infoTextLen
		p.token[zed].Line = s.Line
		p.backup()
	case warningShortOverline, severeOverlineUnderlineMismatch:
		backToken := zed - 2
		if p.peekBack(2).Type == tok.ItemSpace {
			backToken = zed - 3
			indent = p.token[zed-2].Text
		}
		overLine = p.token[backToken].Text
		title = p.token[zed-1].Text
		underLine = p.token[zed].Text
		newLine = "\n"
		lbText = overLine + newLine + indent + title + newLine + underLine
		s.Line = p.token[backToken].Line
		lbTextLen = len(lbText)
		return literalBlock()
	case warningShortUnderline, severeUnexpectedSectionTitle:
		backToken := zed - 1
		if p.peekBack(1).Type == tok.ItemSpace {
			backToken = zed - 2
		}
		lbText = p.token[backToken].Text + "\n" + p.token[zed].Text
		lbTextLen = len(lbText)
		s.Line = p.token[zed-1].Line
		return literalBlock()
	case errorInvalidSectionOrTransitionMarker:
		lbText = p.token[zed-1].Text + "\n" + p.token[zed].Text
		s.Line = p.token[zed-1].Line
		lbTextLen = len(lbText)
		return literalBlock()
	case severeIncompleteSectionTitle,
		severeMissingMatchingUnderlineForOverline:
		lbText = p.token[zed-2].Text + "\n" + p.token[zed-1].Text + p.token[zed].Text
		s.Line = p.token[zed-2].Line
		lbTextLen = len(lbText)
		return literalBlock()
	case severeUnexpectedSectionTitleOrTransition:
		lbText = p.token[zed].Text
		lbTextLen = len(lbText)
		s.Line = p.token[zed].Line
		return literalBlock()
	case severeTitleLevelInconsistent:
		if p.peekBack(2).Type == tok.ItemSectionAdornment {
			lbText = p.token[zed-2].Text + "\n" + p.token[zed-1].Text + "\n" + p.token[zed].Text
			lbTextLen = len(lbText)
			s.Line = p.token[zed-2].Line
			return literalBlock()
		}
		lbText = p.token[zed-1].Text + "\n" + p.token[zed].Text
		lbTextLen = len(lbText)
		s.Line = p.token[zed-1].Line
		return literalBlock()
	}
	return nil
}

func (p *Parser) systemMessageInlineMarkup(s *doc.SystemMessageNode, err parserMessage) *doc.LiteralBlockNode {
	switch err {
	case warningExplicitMarkupWithUnIndent:
		s.Line = p.token[zed+1].Line
	}
	return nil
}

// systemMessage generates a Node based on the passed parserMessage. The generated message is returned as a
// SystemMessageNode.
func (p *Parser) systemMessage(err parserMessage) doc.Node {
	s := doc.NewSystemMessage(&tok.Item{Type: tok.ItemSystemMessage, Line: p.token[zed].Line}, err.Message(), err.String())
	msg := doc.NewText(&tok.Item{
		Text:   err.Message(),
		Length: len(err.Message()),
	})

	log.Log("msg", "Adding msg to system message NodeList", "systemMessage", err)
	s.NodeList.Append(msg)

	appendOrDie := func(f func(s2 *doc.SystemMessageNode, err2 parserMessage) *doc.LiteralBlockNode) {
		if lb := f(s, err); lb != nil {
			s.NodeList = append(s.NodeList, lb)
		}
	}

	appendOrDie(p.systemMessageSection)
	appendOrDie(p.systemMessageInlineMarkup)

	return s
}
