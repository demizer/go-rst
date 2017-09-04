package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

type sectionParseSubState struct {
	sectionOverAdorn  *tok.Item      // Section overline
	sectionIndent     *tok.Item      // Section indent level
	sectionTitle      *doc.TitleNode // Titles can contain inline markup, therefore it must be fully parsed
	sectionUnderAdorn *tok.Item      // Section underline
	sectionSpace      *tok.Item      // The whitespace after a section title and section adornment
}

func parseSectionTitle(s *sectionParseSubState, p *Parser, item *tok.Item) doc.Node {
	p.Msg("next type == tok.Title")
	// Section with overline
	pBack := p.peekBack(1)
	tLen := p.token[zed].Length
	// Check for errors
	if tLen < 3 && tLen != s.sectionSpace.Length {
		p.next(2)
		bTok := p.peekBack(1)
		if bTok != nil && bTok.Type == tok.Space {
			p.next(2)
			sm := p.systemMessage(infoUnexpectedTitleOverlineOrTransition)
			p.nodeTarget.Append(sm)
			return sm
		}
		sm := p.systemMessage(infoOverlineTooShortForTitle)
		p.nodeTarget.Append(sm)
		return sm
	} else if pBack != nil && pBack.Type == tok.Space {
		// Indented section (error) The section title has an indented overline
		sm := p.systemMessage(severeUnexpectedSectionTitleOrTransition)
		p.nodeTarget.Append(sm)
		return sm
	}

	s.sectionOverAdorn = item
	p.next(1)

loop:
	for {
		switch tTok := p.token[zed]; tTok.Type {
		case tok.Title:
			// s.sectionTitle = tTok
			// p.DumpExit(p.nodeTarget)
			// p.nodeTarget.Append(s.sectionTitle)
			// p.nodeTarget.SetParent(s.sectionTitle)
			p.next(1)
		case tok.Space:
			s.sectionIndent = tTok
			p.next(1)
		case tok.SectionAdornment:
			s.sectionUnderAdorn = tTok
			break loop
		default:
			p.subParseInlineMarkup(tTok)
		}
	}

	return nil
}

func parseSectionTitleNoOverline(s *sectionParseSubState, p *Parser, i *tok.Item) doc.Node {
	tLen := p.token[zed].Length
	pBack := p.peekBack(1)
	p.Msgr("last item type", "type", pBack.Type)
	// Section with no overline Check for errors
	if pBack.Type == tok.Space {
		pBack := p.peekBack(2)
		if pBack != nil && pBack.Type == tok.Title {
			// The section underline is indented
			sm := p.systemMessage(severeUnexpectedSectionTitle)
			p.nodeTarget.Append(sm)
			return sm
		}
	} else if tLen < 3 && tLen != pBack.Length {
		// Short underline
		sm := p.systemMessage(infoUnderlineTooShortForTitle)
		p.nodeTarget.Append(sm)
		return sm
	}
	// Section OKAY
	// s.sectionTitle = p.peekBack(1)
	s.sectionUnderAdorn = i

	return nil
}

func parseSectionTitleWithInlineMarkupAndNoOverline(s *sectionParseSubState, p *Parser, i *tok.Item) doc.Node {
	// Go back to before the first occurrence of tok.Title
	var pb *tok.Item
	for {
		p.backup()
		if p.token[zed] == nil {
			p.next(1)
			break
		}
		pb = p.peekBack(1)
		if p.token[zed].Type == tok.Title && (pb == nil || (pb.Type == tok.SectionAdornment || pb.Type == tok.BlankLine)) {
			break
		}
	}
	title := p.token[zed]
	// tLen := p.token[zed].Length
	p.Msgr("have title", "exists", title != nil)
	// p.Msgr("item before title", "exists", beforeTitle != nil)
	// if beforeTitle != nil {
	// p.Msgr("item before title", "item", beforeTitle)
	// }
	p.DumpExit(p.token)
	// panic("SHOW ME THE STACKS!")
	// if title.Type == tok.Space {
	// title := p.peekBack(2)
	// if title != nil && title.Type == tok.Title {
	// // The section underline is indented
	// sm := p.systemMessage(severeUnexpectedSectionTitle)
	// p.nodeTarget.Append(sm)
	// return sm
	// }
	// if tLen < 3 && tLen != title.Length {
	// // Short underline
	// sm := p.systemMessage(infoUnderlineTooShortForTitle)
	// p.nodeTarget.Append(sm)
	// return sm
	// }
	// // Section OKAY
	// // s.sectionTitle = p.peekBack(1)
	// s.sectionUnderAdorn = i

	return nil
}

func parseSectionText(s *sectionParseSubState, p *Parser, i *tok.Item) doc.Node {
	// If a section contains an tok.Text, it is because the underline is missing, therefore we generate an error based on
	// what follows the tok.Text.
	tLen := p.token[zed].Length
	p.next(2) // Move the token buffer past the error tokens
	if tLen < 3 && tLen != s.sectionSpace.Length {
		p.backup()
		sm := p.systemMessage(infoOverlineTooShortForTitle)
		p.nodeTarget.Append(sm)
		return sm
	} else if t := p.peek(1); t != nil && t.Type == tok.BlankLine {
		sm := p.systemMessage(severeMissingMatchingUnderlineForOverline)
		p.nodeTarget.Append(sm)
		return sm
	}
	sm := p.systemMessage(severeIncompleteSectionTitle)
	p.nodeTarget.Append(sm)
	return sm
}

func checkSection(s *sectionParseSubState, p *Parser, i *tok.Item) doc.Node {
	pBack := p.peekBack(1)
	pBackTitle := p.peekBackTo(tok.Title)
	pBackIsInlineMarkup := pBack.Type >= tok.InlineStrongOpen && pBack.Type <= tok.InlineLiteralClose

	// p.DumpExit(p.token)
	// p.DumpExit(pBack)
	p.Msgr("poop", "Plop", pBackTitle != nil)
	// p.Msgr("poop", "Plop", pBackIsInlineMarkup)
	p.Msgr("poop", "Plop", pBack.Type >= tok.InlineStrongOpen && pBack.Type <= tok.InlineLiteralClose)

	if s.sectionSpace != nil && s.sectionSpace.Type == tok.Title {
		// TODO: why would s.sectionSpace.Type == tok.Title ?
		if sm := parseSectionTitle(s, p, i); sm != nil {
			return sm
		}
	} else if pBackTitle != nil && pBackIsInlineMarkup {
		// Token before underline is inline markup and looking backwards, there is a section title
		if sm := parseSectionTitleWithInlineMarkupAndNoOverline(s, p, i); sm != nil {
			return sm
		}
	} else if pBack != nil && (pBack.Type == tok.Title || pBack.Type == tok.Space) {
		// TODO: why would s.sectionSpace.Type == tok.Text ?
		if sm := parseSectionTitleNoOverline(s, p, i); sm != nil {
			return sm
		}
	} else if s.sectionSpace != nil && s.sectionSpace.Type == tok.Text {
		if sm := parseSectionText(s, p, i); sm != nil {
			return sm
		}
	} else if s.sectionSpace != nil && s.sectionSpace.Type == tok.SectionAdornment {
		// TODO: why would s.sectionSpace.Type == tok.SectionAdornment ?
		// panic("SHOW ME THE STACKS!")
		// Missing section title
		p.next(1) // Move the token buffer past the error token
		sm := p.systemMessage(errorInvalidSectionOrTransitionMarker)
		p.nodeTarget.Append(sm)
		return sm
	} else if s.sectionSpace != nil && s.sectionSpace.Type == tok.EOF {
		// Missing underline and at EOF
		sm := p.systemMessage(errorInvalidSectionOrTransitionMarker)
		p.nodeTarget.Append(sm)
		return sm
	}

	if s.sectionOverAdorn != nil && s.sectionOverAdorn.Text != s.sectionUnderAdorn.Text {
		sm := p.systemMessage(severeOverlineUnderlineMismatch)
		p.nodeTarget.Append(sm)
		return sm
	}
	// panic("SHOW ME THE STACKS!")
	return nil
}

func checkSectionLevel(s *sectionParseSubState, p *Parser, sec *doc.SectionNode) doc.Node {
	msg := p.sectionLevels.Add(sec)
	p.Msgr("Using section level", "level", len(p.sectionLevels.levels), "rune", string(sec.UnderLine.Rune))
	if msg != parserMessageNil {
		p.Msg("Found inconsistent section level!")
		sm := p.systemMessage(severeTitleLevelInconsistent)
		// Parse Test 03.01.03.00: add the system message to the last section node's nodelist
		p.sectionLevels.lastSectionNode.NodeList.Append(sm)
		p.nodeTarget.SetParent(p.sectionLevels.lastSectionNode)
		return sm
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
	return nil
}

func checkSectionLengths(s *sectionParseSubState, p *Parser, sec *doc.SectionNode) {
	// The following checks have to be made after the doc.SectionNode has been initialized so that any parserMessages can be
	// appended to the doc.SectionNode.NodeList.
	oLen := s.sectionTitle.Length
	if s.sectionIndent != nil {
		oLen = s.sectionIndent.Length + s.sectionTitle.Length
	}

	if s.sectionOverAdorn != nil && oLen > s.sectionOverAdorn.Length {
		m := warningShortOverline
		sec.NodeList = append(sec.NodeList, p.systemMessage(m))
	} else if s.sectionOverAdorn == nil && s.sectionTitle.Length != s.sectionUnderAdorn.Length {
		m := warningShortUnderline
		sec.NodeList = append(sec.NodeList, p.systemMessage(m))
	}
}

// section is responsible for parsing the title, overline, and underline tokens returned from the parser. If there are errors
// parsing these elements, than a systemMessage is generated and added to Tree.Nodes.
func (p *Parser) section(i *tok.Item) doc.Node {
	p.Msgr("have item", "item", i)

	s := &sectionParseSubState{sectionSpace: p.peekSkip(tok.Space)}

	if sm := checkSection(s, p, i); sm != nil {
		return sm
	}

	// Determine the level of the section and where to append it to in p.Nodes
	sec := doc.NewSection(s.sectionTitle, s.sectionOverAdorn, s.sectionUnderAdorn, s.sectionIndent)

	if sm := checkSectionLevel(s, p, sec); sm != nil {
		return sm
	}

	checkSectionLengths(s, p, sec)

	p.nodeTarget.Append(sec)
	p.nodeTarget.SetParent(sec)

	return sec
}
