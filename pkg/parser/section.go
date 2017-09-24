package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	mes "github.com/demizer/go-rst/pkg/messages"
	tok "github.com/demizer/go-rst/pkg/token"
)

type sectionParseSubState struct {
	overline  *tok.Item      // Section overline
	titleNode *doc.TitleNode // Titles can contain inline markup, therefore it must be fully parsed
	title     *tok.Item      // Unparsed title token
	underline *tok.Item      // Section underline
	padding   *tok.Item      // The whitespace after a section title and section adornment
	indent    *tok.Item      // Section indent level
	indented  bool
}

func (p *Parser) sectionTitle() bool {
	p.Msg("Parsing section title with overline and underline")

	// Section with overline
	overline := p.sectionSubState.overline
	title := p.sectionSubState.title
	underline := p.sectionSubState.underline
	indented := title.StartPosition != 1

	if indented {
		p.sectionSubState.indent = p.peekLine(title.Line)
		p.sectionSubState.indented = indented
	}

	overlineIndented := indented && (overline != nil && overline.StartPosition != 1)
	missingUnderline := underline != nil && underline.Type != tok.SectionAdornment
	overlineUnderlineMismatch := overline != nil && underline != nil && overline.Text != underline.Text
	underlineMissingOrWrongType := underline == nil || underline.Type != tok.SectionAdornment

	hazTitle := title != nil && title.Type == tok.Title
	hazTextTitle := title != nil && title.Type == tok.Text
	hazOverline := overline != nil && overline.Type == tok.SectionAdornment
	hazUnderline := underline != nil && underline.Type == tok.SectionAdornment

	// p.DumpExit(overline)
	// p.DumpExit(title)
	// p.DumpExit(underline)

	if hazOverline && overline.Length > 3 && (hazTitle || hazTextTitle) && missingUnderline {
		// Missing underline is nil
		return p.systemMessage(mes.SectionErrorMissingMatchingUnderlineForOverline)
	} else if hazOverline && overline.Length > 3 && (hazTitle || hazTextTitle) && underlineMissingOrWrongType {
		// underline is not a section underline
		return p.systemMessage(mes.SectionErrorIncompleteSectionTitle)
	} else if hazOverline && overlineIndented && overline.Length < 4 {
		// Indented line containing section adornment chars
		return p.systemMessage(mes.SectionWarningUnexpectedTitleOverlineOrTransition)
	} else if hazOverline && overlineIndented {
		// Overline is indented
		return p.systemMessage(mes.SectionErrorUnexpectedSectionTitleOrTransition)
	} else if hazOverline && overline.Length < 4 && overline.Length < title.Length {
		// Overline is less than three chars. In this case the title may be a text token
		return p.systemMessage(mes.SectionWarningOverlineTooShortForTitle)
	} else if hazUnderline && underline.Length < 4 && underline.Length < title.Length {
		// Underline is less than three chars. In this case the title may be a text token
		return p.systemMessage(mes.SectionWarningUnderlineTooShortForTitle)
	} else if hazOverline && hazUnderline && overlineUnderlineMismatch {
		// Overline and underline do not match
		return p.systemMessage(mes.SectionErrorOverlineUnderlineMismatch)
	} else if hazOverline && !hazTitle {
		// Line after overline is not a title
		return p.systemMessage(mes.SectionErrorInvalidSectionOrTransitionMarker)
	}

	p.next(1)

loop:
	for {
		switch cTok := p.token; cTok.Type {
		case tok.Title:
			p.sectionSubState.titleNode = doc.NewTitleNodeWithText(cTok)
			p.next(1)
		case tok.Space:
			p.sectionSubState.indent = cTok
			p.next(1)
		case tok.SectionAdornment:
			p.sectionSubState.underline = cTok
			break loop
		}
	}

	return true
}

func (p *Parser) sectionTitleNoOverline() bool {
	p.Msg("Parsing section with no overline")
	s := p.sectionSubState

	if s.indented {
		// The section underline is indented
		return p.systemMessage(mes.SectionErrorUnexpectedSectionTitle)
	} else if s.underline.Length < 4 && s.underline.Length < s.title.Length {
		// Underline too short
		return p.systemMessage(mes.SectionWarningUnderlineTooShortForTitle)
	}

	// Section OKAY
	s.titleNode = doc.NewTitleNodeWithText(s.title)

	return true
}

func (p *Parser) sectionTitleWithInlineMarkupAndNoOverline() bool {
	p.Msg("Parsing section with inline markup and no overline")

	var titleLen int
	s := p.sectionSubState

	titleToks := p.peekLineAllTokens(p.token.Line - 1)
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

	if p.token.Length < 4 && p.token.Length != tn.Length {
		// Short underline
		return p.systemMessage(mes.SectionWarningUnderlineTooShortForTitle)
	}

	// Section OKAY
	s.titleNode = tn
	s.underline = p.token

	return true
}

func (p *Parser) parseSection() bool {
	var title, overline, underline *tok.Item
	var indented bool

	title = p.peekLineSkipSpace(p.token.Line - 1)
	underline = p.token

	if title == nil || title.Type != tok.Title {
		// Section adornment is overline
		overline = p.token
		title = p.peekLineSkipSpace(overline.Line + 1)
		underline = p.peekLineSkipSpace(title.Line + 1)
	}

	// Basic checks
	indented = title.StartPosition != 1
	titleWithUnderline := title != nil && underline != nil
	titleWithIndentedUnderline := titleWithUnderline && indented
	titleWithOverline := overline != nil && title != nil
	titleWithIndentedOverline := titleWithOverline && indented
	titleWithInlineMarkup := p.isInlineMarkupInSectionTitle(title)

	// if i.Line == 7 && i.Type == tok.SectionAdornment {
	// p.DumpExit(indented)
	// }

	// p.DumpExit(title)
	// p.DumpExit(underline)
	// p.DumpExit(indented)

	s := p.sectionSubState
	s.overline, s.title, s.underline = overline, title, underline
	if indented {
		s.indent = p.peekLine(title.Line)
		s.indented = indented
	}

	if titleWithOverline || titleWithIndentedOverline {
		// section title with an overline and maybe an underline
		return p.sectionTitle()
	} else if (titleWithUnderline || titleWithIndentedUnderline) && !titleWithInlineMarkup {
		// Title with underline and definitely no overline
		return p.sectionTitleNoOverline()
		// } else if titleWithOverline || title.Type == tok.Text { //&& underline.Type == tok.SectionAdornment {
	} else if titleWithInlineMarkup {
		return p.sectionTitleWithInlineMarkupAndNoOverline()
	}

	p.Msg("Found invalid section")
	return false
}

func checkSectionLevel(p *Parser, sec *doc.SectionNode) bool {
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

func checkSectionLengths(p *Parser, sec *doc.SectionNode) bool {
	// The following checks have to be made after the doc.SectionNode has been initialized so that any parserMessages can be
	// appended to the doc.SectionNode.NodeList.
	s := p.sectionSubState
	titleLen := s.titleNode.Length
	if s.indent != nil {
		titleLen = s.indent.Length + s.titleNode.Length
	}
	// p.DumpExit(s)
	if s.overline != nil && titleLen > s.overline.Length {
		return p.systemMessage(mes.SectionWarningShortOverline)
	} else if s.overline == nil && s.titleNode.Length != s.underline.Length {
		return p.systemMessage(mes.SectionWarningShortUnderline)
	}
	return true
}

// section is responsible for parsing the title, overline, and underline tokens returned from the parser. If there are errors
// parsing these elements, than a systemMessage is generated and added to Tree.Nodes.
func (p *Parser) section(i *tok.Item) {
	p.Msgr("have item", "item", i)

	if !p.parseSection() {
		p.Msg("Failed to parse section!")
		return
	}

	s := p.sectionSubState

	// Determine the level of the section and where to append it to in p.Nodes
	sec := doc.NewSection(s.titleNode, s.overline, s.underline, s.indent)

	if !checkSectionLevel(p, sec) {
		p.Msg("Section checks failed!")
		return
	}
	checkSectionLengths(p, sec)

	p.nodeTarget.Append(sec)
	p.nodeTarget.SetParent(sec)

	// p.DumpExit(p.Nodes)
}

func (p *Parser) isInlineMarkupInSectionTitle(i *tok.Item) bool {
	foundInlineMarkup := false
	lineToks := p.peekLineAllTokens(i.Line)
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
