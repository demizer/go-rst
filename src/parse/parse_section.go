package parse

// section is responsible for parsing the title, overline, and underline tokens returned from the parser. If there are errors
// parsing these elements, than a systemMessage is generated and added to Tree.Nodes.
func (t *Tree) section(i *item) Node {
	var overAdorn, indent, title, underAdorn *item
	logp.Log("msg", "have item", "item", i)

	pBack := t.peekBack(1)
	pFor := t.peekSkip(itemSpace)
	tZedLen := t.token[zed].Length

	if pFor != nil && pFor.Type == itemTitle {
		logp.Msg("next type == itemTitle")
		// Section with overline
		pBack := t.peekBack(1)
		// Check for errors
		if tZedLen < 3 && tZedLen != pFor.Length {
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

		overAdorn = i
		t.next(1)

	loop:
		for {
			switch tTok := t.token[zed]; tTok.Type {
			case itemTitle:
				title = tTok
				t.next(1)
			case itemSpace:
				indent = tTok
				t.next(1)
			case itemSectionAdornment:
				underAdorn = tTok
				break loop
			}
		}
	} else if pBack != nil && (pBack.Type == itemTitle || pBack.Type == itemSpace) {
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
		} else if tZedLen < 3 && tZedLen != pBack.Length {
			// Short underline
			sm := t.systemMessage(infoUnderlineTooShortForTitle)
			t.nodeTarget.append(sm)
			return sm
		}
		// Section OKAY
		title = t.peekBack(1)
		underAdorn = i

	} else if pFor != nil && pFor.Type == itemText {
		// If a section contains an itemText, it is because the underline is missing, therefore we generate an
		// error based on what follows the itemText.
		t.next(2) // Move the token buffer past the error tokens
		if tZedLen < 3 && tZedLen != pFor.Length {
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
	} else if pFor != nil && pFor.Type == itemSectionAdornment {
		// Missing section title
		t.next(1) // Move the token buffer past the error token
		sm := t.systemMessage(errorInvalidSectionOrTransitionMarker)
		t.nodeTarget.append(sm)
		return sm
	} else if pFor != nil && pFor.Type == itemEOF {
		// Missing underline and at EOF
		sm := t.systemMessage(errorInvalidSectionOrTransitionMarker)
		t.nodeTarget.append(sm)
		return sm
	}

	if overAdorn != nil && overAdorn.Text != underAdorn.Text {
		sm := t.systemMessage(severeOverlineUnderlineMismatch)
		t.nodeTarget.append(sm)
		return sm
	}

	// Determine the level of the section and where to append it to in t.Nodes
	sec := newSection(title, overAdorn, underAdorn, indent)

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
		logp.Log("msg", "setting section node target", "sectionTitle", lSec.Title.Text, "level", lSec.Level)
		t.nodeTarget.setParent(lSec)
	}

	// The following checks have to be made after the SectionNode has been initialized so that any parserMessages can be
	// appended to the SectionNode.NodeList.
	oLen := title.Length
	if indent != nil {
		oLen = indent.Length + title.Length
	}

	if overAdorn != nil && oLen > overAdorn.Length {
		m := warningShortOverline
		sec.NodeList = append(sec.NodeList, t.systemMessage(m))
	} else if overAdorn == nil && title.Length != underAdorn.Length {
		m := warningShortUnderline
		sec.NodeList = append(sec.NodeList, t.systemMessage(m))
	}

	t.nodeTarget.append(sec)
	t.nodeTarget.setParent(sec)

	return sec
}
