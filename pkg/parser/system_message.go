package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	mes "github.com/demizer/go-rst/pkg/messages"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) systemMessageSection(s *doc.SystemMessageNode, err *mes.ParserMessage) {
	// panic("foo")
	switch err.Type {
	case mes.SectionWarningOverlineTooShortForTitle:
		var ml, sl, el, sp int
		var overline *tok.Item

		overline = p.token
		title := p.peek(1)
		underline := p.peek(2)

		// if underline != nil && underline.Type == tok.SectionAdornment {
		// // there is an underline
		// }
		// p.DumpExit(p.token)
		// p.DumpExit(underline)

		// if p.peekBackTo(tok.SectionAdornment) != nil {
		// overline = p.buf[p.index-2]
		// }

		if overline.Type == tok.SectionAdornment && underline.Type != tok.SectionAdornment {
			// For title with only overline (no underline)
			err.LiteralText = overline.Text + "\n" + title.Text
			ml, sl, el, sp = overline.Line, overline.Line, title.Line, overline.StartPosition
			p.next(1)
		} else if overline.Type == tok.SectionAdornment && underline.Type == tok.SectionAdornment {
			// For title with overline and underline
			err.LiteralText = overline.Text + "\n" + title.Text + "\n" + underline.Text
			ml, sl, el, sp = overline.Line, overline.Line, underline.Line, overline.StartPosition
			// } else {
			// // For title with underline
			// err.LiteralText = title.Text + "\n" + underline.Text
			// ml, sl, el, sp = underline.Line, title.Line, underline.Line, title.StartPosition
		}

		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = ml, sl, el, sp
		// panic("BLAH")
		// p.DumpExit(buf)
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
		// p.DumpExit(p.buf)
	case mes.SectionWarningShortOverline, mes.SectionErrorOverlineUnderlineMismatch:
		backIndex := p.index - 2
		if p.peekBack(2).Type == tok.Space {
			backIndex = p.index - 3
		}
		overline := p.buf[backIndex]
		underline := p.token
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = overline.Line, overline.Line, underline.Line, overline.StartPosition
	case mes.SectionWarningShortUnderline, mes.SectionErrorUnexpectedSectionTitle:
		backIndex := p.index - 1
		if p.peekBack(1).Type == tok.Space {
			backIndex = p.index - 2
		}
		err.LiteralText = p.buf[backIndex].Text + "\n" + p.token.Text
		err.MessageLine = p.buf[backIndex+1].Line
		err.StartLine = p.buf[backIndex].Line
		err.EndLine = p.buf[backIndex+1].Line
		err.StartPosition = p.buf[backIndex].StartPosition
		s.Line = p.buf[backIndex].Line
	case mes.SectionErrorInvalidSectionOrTransitionMarker:
		err.LiteralText = p.buf[p.index-1].Text + "\n" + p.token.Text
	case mes.SectionErrorIncompleteSectionTitle, mes.SectionErrorMissingMatchingUnderlineForOverline:
		overline := p.token
		var text string
		// p.DumpExit(overline)
		for {
			p.next(1)
			if p.token.Type == tok.EOF {
				break
			}
			text += p.token.Text
			if p.token.Line > overline.Line+1 {
				break
			}
		}
		p.backup()
		// title := p.peekSkip(tok.Space)
		// p.DumpExit(text)
		err.LiteralText = overline.Text + "\n" + text
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = overline.Line, overline.Line, overline.Line+1, overline.StartPosition
		// p.DumpExit(err)
		// panic("foo")
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
		err.MessageLine = tok.Line
		err.StartLine = tok.Line - 1
		err.EndLine = tok.Line
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
		// p.DumpExit(p.buf)
	}

	if mes.IsSectionMessage(err) {
		p.systemMessageSection(s, nm)
		if len(nm.LiteralText) > 0 {
			insertText()
		}
	} else if mes.IsInlineMarkupMessage(err) {
		p.systemMessageInlineMarkup(s, nm)
	}

	s.Line = nm.MessageLine
	s.StartPosition = nm.StartPosition
	s.StartLine = nm.StartLine
	s.EndLine = nm.EndLine
	p.Messages.Append(s)

	// p.DumpExit(p.Messages)
	return false
}
