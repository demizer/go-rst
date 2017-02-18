package parser

type sectionParseSubState struct {
	sectionOverAdorn  *item
	sectionIndent     *item
	sectionTitle      *item
	sectionUnderAdorn *item
	sectionSpace      *item
}

func parseSectionTitle(s *sectionParseSubState, t *Tree, item *item) Node {
	logp.Msg("next type == itemTitle")
	// Section with overline
	pBack := t.peekBack(1)
	tLen := t.token[zed].Length
	// Check for errors
	if tLen < 3 && tLen != s.sectionSpace.Length {
		t.next(2)
		bTok := t.peekBack(1)
		if bTok != nil && bTok.Type == itemSpace {
			t.next(2)
			sm := t.systemMessage(infoUnexpectedTitleOverlineOrTransition)
			t.nodeTarget.append(sm)
			return sm
		}
		sm := t.systemMessage(infoOverlineTooShortForTitle)
		t.nodeTarget.append(sm)
		return sm
	} else if pBack != nil && pBack.Type == itemSpace {
		// Indented section (error) The section title has an indented overline
		sm := t.systemMessage(severeUnexpectedSectionTitleOrTransition)
		t.nodeTarget.append(sm)
		return sm
	}

	s.sectionOverAdorn = item
	t.next(1)

loop:
	for {
		switch tTok := t.token[zed]; tTok.Type {
		case itemTitle:
			s.sectionTitle = tTok
			t.next(1)
		case itemSpace:
			s.sectionIndent = tTok
			t.next(1)
		case itemSectionAdornment:
			s.sectionUnderAdorn = tTok
			break loop
		}
	}

	return nil
}

func parseSectionTitleNoOverline(s *sectionParseSubState, t *Tree, i *item) Node {
	tLen := t.token[zed].Length
	pBack := t.peekBack(1)
	logp.Log("msg", "last item type", "type", pBack.Type)
	// Section with no overline Check for errors
	if pBack.Type == itemSpace {
		pBack := t.peekBack(2)
		if pBack != nil && pBack.Type == itemTitle {
			// The section underline is indented
			sm := t.systemMessage(severeUnexpectedSectionTitle)
			t.nodeTarget.append(sm)
			return sm
		}
	} else if tLen < 3 && tLen != pBack.Length {
		// Short underline
		sm := t.systemMessage(infoUnderlineTooShortForTitle)
		t.nodeTarget.append(sm)
		return sm
	}
	// Section OKAY
	s.sectionTitle = t.peekBack(1)
	s.sectionUnderAdorn = i

	return nil
}

func parseSectionText(s *sectionParseSubState, t *Tree, i *item) Node {
	// If a section contains an itemText, it is because the underline is missing, therefore we generate an error based on
	// what follows the itemText.
	tLen := t.token[zed].Length
	t.next(2) // Move the token buffer past the error tokens
	if tLen < 3 && tLen != s.sectionSpace.Length {
		t.backup()
		sm := t.systemMessage(infoOverlineTooShortForTitle)
		t.nodeTarget.append(sm)
		return sm
	} else if p := t.peek(1); p != nil && p.Type == itemBlankLine {
		sm := t.systemMessage(severeMissingMatchingUnderlineForOverline)
		t.nodeTarget.append(sm)
		return sm
	}
	sm := t.systemMessage(severeIncompleteSectionTitle)
	t.nodeTarget.append(sm)
	return sm
}

func checkSection(s *sectionParseSubState, t *Tree, i *item) Node {
	pBack := t.peekBack(1)

	if s.sectionSpace != nil && s.sectionSpace.Type == itemTitle {
		if sm := parseSectionTitle(s, t, i); sm != nil {
			return sm
		}
	} else if pBack != nil && (pBack.Type == itemTitle || pBack.Type == itemSpace) {
		if sm := parseSectionTitleNoOverline(s, t, i); sm != nil {
			return sm
		}
	} else if s.sectionSpace != nil && s.sectionSpace.Type == itemText {
		if sm := parseSectionText(s, t, i); sm != nil {
			return sm
		}
	} else if s.sectionSpace != nil && s.sectionSpace.Type == itemSectionAdornment {
		// Missing section title
		t.next(1) // Move the token buffer past the error token
		sm := t.systemMessage(errorInvalidSectionOrTransitionMarker)
		t.nodeTarget.append(sm)
		return sm
	} else if s.sectionSpace != nil && s.sectionSpace.Type == itemEOF {
		// Missing underline and at EOF
		sm := t.systemMessage(errorInvalidSectionOrTransitionMarker)
		t.nodeTarget.append(sm)
		return sm
	}

	if s.sectionOverAdorn != nil && s.sectionOverAdorn.Text != s.sectionUnderAdorn.Text {
		sm := t.systemMessage(severeOverlineUnderlineMismatch)
		t.nodeTarget.append(sm)
		return sm
	}
	return nil
}

func checkSectionLevel(s *sectionParseSubState, t *Tree, sec *SectionNode) Node {
	msg := t.sectionLevels.Add(sec)
	logp.Log("msg", "Using section level", "level", len(t.sectionLevels.levels), "rune", string(sec.UnderLine.Rune))
	if msg != parserMessageNil {
		logp.Msg("Found inconsistent section level!")
		sm := t.systemMessage(severeTitleLevelInconsistent)
		// Parse Test 03.01.03.00: add the system message to the last section node's nodelist
		t.sectionLevels.lastSectionNode.NodeList.append(sm)
		t.nodeTarget.setParent(t.sectionLevels.lastSectionNode)
		return sm
	}

	if sec.Level == 1 {
		logp.Msg("Setting nodeTarget to Tree.Nodes!")
		t.nodeTarget.reset()
	} else {
		lSec := t.sectionLevels.lastSectionNode
		logp.Log("msg", "have last section node", "secNode", lSec.Title.Text, "level", lSec.Level)
		if sec.Level > 1 {
			lSec = t.sectionLevels.LastSectionByLevel(sec.Level - 1)
		}
		logp.Log("msg", "setting section node target", "Title", lSec.Title.Text, "level", lSec.Level)
		t.nodeTarget.setParent(lSec)
	}
	return nil
}

func checkSectionLengths(s *sectionParseSubState, t *Tree, sec *SectionNode) {
	// The following checks have to be made after the SectionNode has been initialized so that any parserMessages can be
	// appended to the SectionNode.NodeList.
	oLen := s.sectionTitle.Length
	if s.sectionIndent != nil {
		oLen = s.sectionIndent.Length + s.sectionTitle.Length
	}

	if s.sectionOverAdorn != nil && oLen > s.sectionOverAdorn.Length {
		m := warningShortOverline
		sec.NodeList = append(sec.NodeList, t.systemMessage(m))
	} else if s.sectionOverAdorn == nil && s.sectionTitle.Length != s.sectionUnderAdorn.Length {
		m := warningShortUnderline
		sec.NodeList = append(sec.NodeList, t.systemMessage(m))
	}
}

// section is responsible for parsing the title, overline, and underline tokens returned from the parser. If there are errors
// parsing these elements, than a systemMessage is generated and added to Tree.Nodes.
func (t *Tree) section(i *item) Node {
	logp.Log("msg", "have item", "item", i)

	s := &sectionParseSubState{sectionSpace: t.peekSkip(itemSpace)}

	if sm := checkSection(s, t, i); sm != nil {
		return sm
	}

	// Determine the level of the section and where to append it to in t.Nodes
	sec := newSection(s.sectionTitle, s.sectionOverAdorn, s.sectionUnderAdorn, s.sectionIndent)

	if sm := checkSectionLevel(s, t, sec); sm != nil {
		return sm
	}

	checkSectionLengths(s, t, sec)

	t.nodeTarget.append(sec)
	t.nodeTarget.setParent(sec)

	return sec
}
