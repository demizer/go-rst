package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	mes "github.com/demizer/go-rst/pkg/messages"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) systemMessageSection(s *doc.SystemMessageNode, err *mes.ParserMessage) {
	switch err.Type {
	case mes.SectionWarningOverlineTooShortForTitle:
		overline := p.buf[p.index-2]
		title := p.buf[p.index-1]
		// For title with only overline, combine two tokens and insert into buffer
		err.LiteralText = overline.Text + "\n" + p.token.Text
		// For title with overline and underline, combine 3 tokens and insert into buffer
		if p.index-2 >= 0 && overline != nil && overline.Type == tok.SectionAdornment {
			err.LiteralText = overline.Text + "\n" + title.Text + "\n" + p.token.Text
		}
		err.StartLine = overline.Line
		err.EndLine = title.Line
		err.MessageLine = overline.Line
		err.StartPosition = overline.StartPosition
	case mes.SectionWarningUnexpectedTitleOverlineOrTransition:
		err.LiteralText = p.peekBackTo(tok.SectionAdornment).Text + "\n" + p.peekBackTo(tok.Title).Text + "\n" + p.token.Text
	case mes.SectionWarningUnderlineTooShortForTitle:
		title := p.buf[p.index-1]
		underline := p.buf[p.index]
		err.LiteralText = title.Text + "\n" + underline.Text
		err.StartLine = title.Line
		err.EndLine = underline.Line
		err.MessageLine = underline.Line
		err.StartPosition = underline.StartPosition
	case mes.SectionWarningShortOverline, mes.SectionErrorOverlineUnderlineMismatch:
		var indent string
		backIndex := p.index - 2
		if p.peekBack(2).Type == tok.Space {
			backIndex = p.index - 3
			indent = p.buf[p.index-2].Text
		}
		overLine := p.buf[backIndex].Text
		title := p.buf[p.index-1].Text
		underLine := p.token.Text
		newLine := "\n"
		err.LiteralText = overLine + newLine + indent + title + newLine + underLine
	case mes.SectionWarningShortUnderline, mes.SectionErrorUnexpectedSectionTitle:
		backIndex := p.index - 1
		if p.peekBack(1).Type == tok.Space {
			backIndex = p.index - 2
		}
		err.LiteralText = p.buf[backIndex].Text + "\n" + p.token.Text
		s.Line = p.buf[backIndex].Line
	case mes.SectionErrorInvalidSectionOrTransitionMarker:
		err.LiteralText = p.buf[p.index-1].Text + "\n" + p.token.Text
	case mes.SectionErrorIncompleteSectionTitle,
		mes.SectionErrorMissingMatchingUnderlineForOverline:
		err.LiteralText = p.buf[p.index-2].Text + "\n" + p.buf[p.index-1].Text + p.token.Text
	case mes.SectionErrorUnexpectedSectionTitleOrTransition:
		err.LiteralText = p.token.Text
	case mes.SectionErrorTitleLevelInconsistent:
		if p.peekBack(2).Type == tok.SectionAdornment {
			err.LiteralText = p.buf[p.index-2].Text + "\n" + p.buf[p.index-1].Text + "\n" + p.token.Text
			break
		}
		err.LiteralText = p.buf[p.index-1].Text + "\n" + p.token.Text
	}
}

func (p *Parser) systemMessageInlineMarkup(s *doc.SystemMessageNode, err *mes.ParserMessage) {
	switch err.Type {
	case mes.InlineMarkupWarningExplicitMarkupWithUnIndent:
		tok := p.peek(1)
		err.StartLine = tok.Line
		err.EndLine = tok.Line
		err.MessageLine = tok.Line
		err.StartPosition = tok.StartPosition
	}
}

// systemMessage generates a Node based on the passed mes.ParserMessage. The generated message is returned as a
// SystemMessageNode.
func (p *Parser) systemMessage(err mes.MessageType) (ok bool) {
	nm := mes.NewParserMessage(err)
	s := doc.NewSystemMessage(nm, p.token.Line)

	// Insert text into the next buffer position to be picked up in the next pass of the parser
	insertText := func() {
		p.insert(&tok.Item{
			Type:          tok.Text,
			Text:          nm.LiteralText,
			Length:        len(nm.LiteralText),
			Line:          nm.StartLine,
			StartPosition: nm.StartPosition,
		}, p.index+1)
	}

	if mes.IsSectionMessage(err) {
		p.systemMessageSection(s, nm)
		insertText()
	} else if mes.IsInlineMarkupMessage(err) {
		p.systemMessageInlineMarkup(s, nm)
	}

	s.Line = nm.MessageLine
	s.StartPosition = nm.StartPosition
	p.Messages.Append(s)

	// p.DumpExit(p.Messages)
	return false
}
