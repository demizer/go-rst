package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	mes "github.com/demizer/go-rst/pkg/messages"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) systemMessageSection(s *doc.SystemMessageNode, err *mes.ParserMessage) {
	// panic("foo")
	st := p.sectionSubState
	switch err.Type {
	case mes.SectionWarningOverlineTooShortForTitle:
		// p.DumpExit(st)
		// For title with only overline (no underline)
		err.LiteralText = st.overline.Text + "\n" + st.title.Text
		ml, sl, el, sp := st.overline.Line, st.overline.Line, st.title.Line, st.overline.StartPosition
		if st.underline != nil && st.underline.Type != tok.BlankLine {
			// For title with overline and underline
			err.LiteralText = st.overline.Text + "\n" + st.title.Text + "\n" + st.underline.Text
			ml, sl, el, sp = st.overline.Line, st.overline.Line, st.underline.Line, st.overline.StartPosition
		}
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = ml, sl, el, sp
		p.nextToLine(err.EndLine)
	case mes.SectionWarningUnexpectedTitleOverlineOrTransition:
		err.LiteralText = st.overline.Text + "\n" + st.title.Text + "\n" + st.underline.Text
		err.StartLine, err.EndLine, err.MessageLine, err.StartPosition = st.overline.Line, st.underline.Line, st.overline.Line, st.overline.StartPosition
		p.nextToLine(err.EndLine)
	case mes.SectionWarningUnderlineTooShortForTitle:
		err.LiteralText = st.title.Text + "\n" + st.underline.Text
		err.StartLine, err.EndLine, err.MessageLine, err.StartPosition = st.title.Line, st.underline.Line, st.underline.Line, st.underline.StartPosition
		// p.DumpExit(p.buf)
	case mes.SectionWarningShortOverline:
		backIndex := p.index - 2
		if p.peekBack(2).Type == tok.Space {
			backIndex = p.index - 3
		}
		overline := p.buf[backIndex]
		underline := p.token
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = overline.Line, overline.Line, underline.Line, overline.StartPosition
	case mes.SectionErrorOverlineUnderlineMismatch:
		err.LiteralText = p.globText(p.index, p.indexFromToken(st.underline)+1)
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = st.overline.Line, st.overline.Line, st.underline.Line, st.overline.StartPosition
		p.nextToLine(err.EndLine)
		// p.DumpExit(err)
	case mes.SectionWarningShortUnderline:
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
	case mes.SectionErrorUnexpectedSectionTitle:
		title := p.peekBackTo(tok.Title)
		underline := p.token
		err.LiteralText = title.Text + "\n" + underline.Text
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = underline.Line, title.Line, underline.Line, underline.StartPosition
		p.next(1) // Next past the underline
		// p.DumpExit(p.buf)
	case mes.SectionErrorInvalidSectionOrTransitionMarker:
		err.LiteralText = st.overline.Text + "\n" + st.title.Text
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = st.overline.Line, st.overline.Line, st.title.Line, st.overline.StartPosition
		p.next(1) // Next past the underline
	case mes.SectionErrorIncompleteSectionTitle, mes.SectionErrorMissingMatchingUnderlineForOverline:
		overline := p.token
		text := p.globTextFromLine(overline.Line + 1)
		err.LiteralText = overline.Text + "\n" + text
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = overline.Line, overline.Line, overline.Line+1, overline.StartPosition
		p.next(2)
	case mes.SectionErrorUnexpectedSectionTitleOrTransition:
		err.MessageLine, err.StartLine, err.EndLine, err.StartPosition = st.overline.Line, st.overline.Line, st.overline.Line, st.overline.StartPosition
		p.token.Type = tok.Text
		p.backup()
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
func (p *Parser) systemMessage(err mes.MessageType) bool {
	nm := mes.NewParserMessage(err)
	s := doc.NewSystemMessage(nm, p.token.Line)
	p.Msgr("Generating system message", "type", err.String())
	// panic("foo")

	// Insert text into the next buffer position to be picked up in the next pass of the parser
	insertText := func() {
		p.insert(&tok.Item{
			Type:          tok.Text,
			Text:          nm.LiteralText,
			Length:        len(nm.LiteralText),
			Line:          nm.StartLine,
			StartPosition: nm.StartPosition,
		}, p.index+1)
		// p.DumpExit(p.buf[:p.index+3])
		p.Msgr("foooooooooooooooooooooooooooooooooooooooooooooooooooooooooo", "line", p.token.Line)
		// if p.token.Line == 9 {
		// // p.DumpExit(p.token)
		// p.DumpExit(p.buf[p.index-2 : p.index+3])
		// }
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
