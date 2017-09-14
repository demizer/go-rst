package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	mes "github.com/demizer/go-rst/pkg/messages"
	tok "github.com/demizer/go-rst/pkg/token"
)

type sectionParseSubState struct {
	sectionOverAdorn  *tok.Item      // Section overline
	sectionIndent     *tok.Item      // Section indent level
	sectionTitle      *doc.TitleNode // Titles can contain inline markup, therefore it must be fully parsed
	sectionUnderAdorn *tok.Item      // Section underline
	sectionSpace      *tok.Item      // The whitespace after a section title and section adornment
}

func parseSectionTitle(s *sectionParseSubState, p *Parser, item *tok.Item) bool {
	p.Msg("next type == tok.Title")

	// Section with overline
	pBack := p.peekBack(1)
	tLen := p.token.Length

	if tLen < 3 && tLen != s.sectionSpace.Length {
		p.next(2)
		bTok := p.peekBack(1)
		if bTok != nil && bTok.Type == tok.Space {
			p.next(2)
			return p.systemMessage(mes.SectionWarningUnexpectedTitleOverlineOrTransition)
		}
		return p.systemMessage(mes.SectionWarningOverlineTooShortForTitle)
	} else if pBack != nil && pBack.Type == tok.Space {
		// Indented section (error) The section title has an indented overline
		return p.systemMessage(mes.SectionErrorUnexpectedSectionTitleOrTransition)
	}

	s.sectionOverAdorn = item
	p.next(1)

loop:
	for {
		switch tTok := p.token; tTok.Type {
		case tok.Title:
			s.sectionTitle = doc.NewTitleNode()
			s.sectionTitle.Append(doc.NewText(tTok))
			s.sectionTitle.Length = tTok.Length
			s.sectionTitle.StartPosition = tTok.StartPosition
			s.sectionTitle.Line = tTok.Line
			p.next(1)
		case tok.Space:
			s.sectionIndent = tTok
			p.next(1)
		case tok.SectionAdornment:
			s.sectionUnderAdorn = tTok
			break loop
		}
	}
	return true
}

func parseSectionTitleNoOverline(s *sectionParseSubState, p *Parser, i *tok.Item) (ok bool) {
	tLen := p.token.Length
	pBack := p.peekBack(1)
	p.Msgr("last item type", "type", pBack.Type)
	// Section with no overline Check for errors
	if pBack.Type == tok.Space {
		pBack := p.peekBack(2)
		if pBack != nil && pBack.Type == tok.Title {
			// The section underline is indented
			return p.systemMessage(mes.SectionErrorUnexpectedSectionTitle)
		}
	} else if tLen < 3 && tLen != pBack.Length {
		// Short underline
		return p.systemMessage(mes.SectionWarningUnderlineTooShortForTitle)
	}

	// Section OKAY
	s.sectionTitle = doc.NewTitleNode()
	s.sectionTitle.Append(doc.NewText(pBack))
	s.sectionUnderAdorn = i

	return true
}

func parseSectionTitleWithInlineMarkupAndNoOverline(s *sectionParseSubState, p *Parser, i *tok.Item) (ok bool) {
	var titleLen int
	titleToks := p.peekLine(p.token.Line - 1)
	secAdorn := p.token

	tn := doc.NewTitleNode()
	p.nodeTarget.SetParent(tn)
	for _, v := range titleToks {
		titleLen += v.Length
		switch v.Type {
		case tok.Title:
			tn.Append(doc.NewText(v))
		default:
			p.index = p.indexFromToken(v)
			p.subParseInlineMarkup(v)
		}
	}
	p.nodeTarget.Reset()

	tn.StartPosition = titleToks[0].StartPosition
	tn.Line = titleToks[0].Line
	tn.Length = titleLen

	p.index = p.indexFromToken(secAdorn)
	p.token = p.buf[p.index]

	// Section with no overline Check for errors
	// if pBack.Type == tok.Space {
	// pBack := p.peekBack(2)
	// if pBack != nil && pBack.Type == tok.Title {
	// // The section underline is indented
	// return p.systemMessage(mes.SectionErrorUnexpectedSectionTitle)
	// }

	if p.token.Length < 3 && p.token.Length != tn.Length {
		// Short underline
		return p.systemMessage(mes.SectionWarningUnderlineTooShortForTitle)
	}

	// Section OKAY
	s.sectionTitle = tn
	s.sectionUnderAdorn = p.token

	return true
}

func sectionOK(s *sectionParseSubState, p *Parser, i *tok.Item) (ok bool) {
	if s.sectionSpace != nil && s.sectionSpace.Type == tok.Text {
		//
		// If sectionSpace is set to a tok.Text,
		// * The underline is missing, therefore we generate an error based on what follows the tok.Text
		// * There is no blankline after the underline
		//
		// Check for an underline...
		if und := p.peekLine(s.sectionSpace.Line - 1); len(und) == 1 && und[0].Type == tok.SectionAdornment {
			// There is no blankline after the underline, this is a good section
			return true
		}
		p.Msg("IN HERE 4")
		tLen := p.token.Length
		p.next(2) // Move the token buffer past the error tokens

		if tLen < 3 && tLen != s.sectionSpace.Length {
			p.backup()
			return p.systemMessage(mes.SectionWarningOverlineTooShortForTitle)
		} else if t := p.peek(1); t != nil && t.Type == tok.BlankLine {
			return p.systemMessage(mes.SectionErrorMissingMatchingUnderlineForOverline)
		}

		return p.systemMessage(mes.SectionErrorIncompleteSectionTitle)
	} else if s.sectionSpace != nil && s.sectionSpace.Type == tok.SectionAdornment {
		p.Msg("IN HERE 5")
		// Missing section title
		p.next(1) // Move the token buffer past the error token
		return p.systemMessage(mes.SectionErrorInvalidSectionOrTransitionMarker)
	} else if s.sectionSpace != nil && s.sectionSpace.Type == tok.EOF {
		p.Msg("IN HERE 6")
		// Missing underline and at EOF
		return p.systemMessage(mes.SectionErrorInvalidSectionOrTransitionMarker)
	} else if s.sectionOverAdorn != nil && s.sectionOverAdorn.Text != s.sectionUnderAdorn.Text {
		p.Msg("IN HERE 7")
		return p.systemMessage(mes.SectionErrorOverlineUnderlineMismatch)
	}

	p.Msg("IN HERE OMEGA")
	return true
}

func parseSection(s *sectionParseSubState, p *Parser, i *tok.Item) (ok bool) {
	pBack := p.peekBack(1)
	pBackTitle := p.peekBackTo(tok.Title)

	if !sectionOK(s, p, i) {
		return false
	}

	if s.sectionSpace != nil && s.sectionSpace.Type == tok.Title {
		p.Msg("IN HERE 1")
		return parseSectionTitle(s, p, i)
	} else if pBackTitle != nil && p.isInlineMarkupInSectionTitle(pBackTitle) {
		p.Msg("IN HERE 2")
		return parseSectionTitleWithInlineMarkupAndNoOverline(s, p, i)
	} else if pBack != nil && (pBack.Type == tok.Title || pBack.Type == tok.Space) {
		p.Msg("IN HERE 3")
		return parseSectionTitleNoOverline(s, p, i)
	}

	p.Msg("IN HERE 8")
	return
}

func checkSectionLevel(s *sectionParseSubState, p *Parser, sec *doc.SectionNode) (ok bool) {
	msg := p.sectionLevels.Add(sec)
	p.Msgr("Using section level", "level", len(p.sectionLevels.levels), "rune", string(sec.UnderLine.Rune))
	if msg != mes.ParserMessageNil {
		p.Msg("Found inconsistent section level!")
		return p.systemMessage(mes.SectionErrorTitleLevelInconsistent)
		// Parse Test 03.01.03.00: add the system message to the last section node's nodelist
		// p.sectionLevels.lastSectionNode.NodeList.Append(sm)
		// p.nodeTarget.SetParent(p.sectionLevels.lastSectionNode)
		// return
	}

	if sec.Level == 1 {
		p.Msg("Setting nodeTarget to Tree.Nodes!")
		p.nodeTarget.Reset()
	} else {
		lSec := p.sectionLevels.lastSectionNode
		// p.Msgr("have last section node", "secNode", lSec.Title.Nod Text, "level", lSec.Level)
		if sec.Level > 1 {
			lSec = p.sectionLevels.LastSectionByLevel(sec.Level - 1)
		}
		// p.Msgr("setting section node target", "Title", lSec.Title.Text, "level", lSec.Level)
		p.nodeTarget.SetParent(lSec)
	}
	return true
}

func checkSectionLengths(s *sectionParseSubState, p *Parser, sec *doc.SectionNode) (ok bool) {
	// The following checks have to be made after the doc.SectionNode has been initialized so that any parserMessages can be
	// appended to the doc.SectionNode.NodeList.
	oLen := s.sectionTitle.Length
	if s.sectionIndent != nil {
		oLen = s.sectionIndent.Length + s.sectionTitle.Length
	}
	if s.sectionOverAdorn != nil && oLen > s.sectionOverAdorn.Length {
		return p.systemMessage(mes.SectionWarningShortOverline)
	} else if s.sectionOverAdorn == nil && s.sectionTitle.Length != s.sectionUnderAdorn.Length {
		return p.systemMessage(mes.SectionWarningShortUnderline)
	}
	return true
}

// section is responsible for parsing the title, overline, and underline tokens returned from the parser. If there are errors
// parsing these elements, than a systemMessage is generated and added to Tree.Nodes.
func (p *Parser) section(i *tok.Item) {
	p.Msgr("have item", "item", i)

	s := &sectionParseSubState{sectionSpace: p.peekSkip(tok.Space)}
	if !parseSection(s, p, i) {
		return
	}

	// p.DumpExit(p.buf)
	// Determine the level of the section and where to append it to in p.Nodes
	sec := doc.NewSection(s.sectionTitle, s.sectionOverAdorn, s.sectionUnderAdorn, s.sectionIndent)

	checkSectionLevel(s, p, sec)
	checkSectionLengths(s, p, sec)

	p.nodeTarget.Append(sec)
	p.nodeTarget.SetParent(sec)
}

func (p *Parser) isInlineMarkupInSectionTitle(i *tok.Item) bool {
	foundInlineMarkup := false
	lineToks := p.peekLine(i.Line)
	for _, v := range lineToks {
		if v.Type >= tok.InlineStrongOpen && v.Type <= tok.InlineLiteralClose {
			foundInlineMarkup = true
		}
		if v.Type == tok.Title && foundInlineMarkup {
			return true
		}
	}
	return false
}
