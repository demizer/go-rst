package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) systemMessageSection(s *doc.SystemMessageNode, err parserMessage) *doc.LiteralBlockNode {
	var lbText, overLine, indent, title, underLine, newLine string
	var lbLine, lbSPos int

	literalBlock := func() *doc.LiteralBlockNode {
		return doc.NewLiteralBlock(&tok.Item{Type: tok.LiteralBlock, Text: lbText, Length: len(lbText), Line: lbLine, StartPosition: lbSPos})
	}

	switch err {
	case infoOverlineTooShortForTitle:
		if p.index-2 >= 0 && p.buf[p.index-2] != nil && p.buf[p.index-2].Type == tok.SectionAdornment {
			// For title with overline and underline, combine 3 tokens and insert into buffer
			lbText = p.buf[p.index-2].Text + "\n" + p.buf[p.index-1].Text + "\n" + p.token.Text
			lbSPos = p.buf[p.index-2].StartPosition
		} else {
			// For title with only overline, combine two tokens and insert into buffer
			lbText = p.buf[p.index-1].Text + "\n" + p.token.Text
			lbSPos = p.buf[p.index-1].StartPosition
		}
	case infoUnexpectedTitleOverlineOrTransition:
		lbText = p.peekBackTo(tok.SectionAdornment).Text + "\n" + p.peekBackTo(tok.Title).Text + "\n" + p.token.Text
	case infoUnderlineTooShortForTitle:
		lbText = p.buf[p.index-1].Text + "\n" + p.buf[p.index].Text
	case warningShortOverline, severeOverlineUnderlineMismatch:
		backIndex := p.index - 2
		if p.peekBack(2).Type == tok.Space {
			backIndex = p.index - 3
			indent = p.buf[p.index-2].Text
		}
		overLine = p.buf[backIndex].Text
		title = p.buf[p.index-1].Text
		underLine = p.token.Text
		newLine = "\n"
		lbText = overLine + newLine + indent + title + newLine + underLine
	case warningShortUnderline, severeUnexpectedSectionTitle:
		backIndex := p.index - 1
		if p.peekBack(1).Type == tok.Space {
			backIndex = p.index - 2
		}
		lbText = p.buf[backIndex].Text + "\n" + p.token.Text
		s.Line = p.buf[backIndex].Line
	case errorInvalidSectionOrTransitionMarker:
		lbText = p.buf[p.index-1].Text + "\n" + p.token.Text
	case severeIncompleteSectionTitle,
		severeMissingMatchingUnderlineForOverline:
		lbText = p.buf[p.index-2].Text + "\n" + p.buf[p.index-1].Text + p.token.Text
	case severeUnexpectedSectionTitleOrTransition:
		lbText = p.token.Text
	case severeTitleLevelInconsistent:
		if p.peekBack(2).Type == tok.SectionAdornment {
			lbText = p.buf[p.index-2].Text + "\n" + p.buf[p.index-1].Text + "\n" + p.token.Text
			break
		}
		lbText = p.buf[p.index-1].Text + "\n" + p.token.Text
	}
	if len(lbText) > 0 {

		return literalBlock()
	}
	return nil
}

func (p *Parser) systemMessageInlineMarkup(s *doc.SystemMessageNode, err parserMessage) *doc.LiteralBlockNode {
	switch err {
	case warningExplicitMarkupWithUnIndent:
		s.Line = p.peek(1).Line
	}
	return nil
}

// systemMessage generates a Node based on the passed parserMessage. The generated message is returned as a
// SystemMessageNode.
func (p *Parser) systemMessage(err parserMessage) (ok bool) {
	s := doc.NewSystemMessage(&tok.Item{Type: tok.SystemMessage}, err.String(), err.Level())
	// msg := doc.NewText(&tok.Item{
	// Text:   err.Message(),
	// Length: len(err.Message()),
	// })

	mesAppend := func(f func(s2 *doc.SystemMessageNode, err2 parserMessage) *doc.LiteralBlockNode) {
		if lb := f(s, err); lb != nil {
			p.Msgr("Adding msg to system message NodeList", "systemMessage", err)
			p.Messages = append(p.Messages, lb)
		}
	}

	mesAppend(p.systemMessageSection)
	mesAppend(p.systemMessageInlineMarkup)

	return false
}
