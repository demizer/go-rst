package parse

import "unicode/utf8"

func (t *Tree) paragraph(i *item) Node {
	logp.Log("msg", "Have token", "token", i)
	var np ParagraphNode
	if !t.nodeTarget.isParagraphNode() {
		np := newParagraph()
		t.nodeTarget.append(np)
		t.nodeTarget.setParent(np)
	}
	nt := newText(i)
	t.nodeTarget.append(nt)
outer:
	// Paragraphs can contain many different types of elements, so we'll need to loop until blank line or nil
	for {
		ci := t.next(1)     // current item
		pi := t.peekBack(1) // previous item
		// ni := t.peek(1)     // next item

		logp.Log("msg", "Have token", "token", ci)

		if ci == nil {
			logp.Msg("ci == nil, breaking")
			break
		} else if ci.Type == itemEOF {
			logp.Msg("current item type == itemEOF")
			break
		} else if pi != nil && pi.Type == itemText && ci.Type == itemText {
			logp.Msg("Previous type == itemText, current type == itemText; Concatenating text!")
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
			continue
		}

		logp.Msg("Going into subparser...")

		switch ci.Type {
		case itemSpace:
			if pi != nil && pi.Type == itemEscape {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				logp.Msg("Found escaped space!")
				continue
			}
			// Parse Test 02.00.03.00 :: Emphasis wrapped in unicode spaces
			nt.Text += "\n" + ci.Text
			nt.Length = utf8.RuneCountInString(nt.Text)
		case itemText:
			if pi != nil && pi.Type == itemEscape && pi.StartPosition.Int() > ci.StartPosition.Int() {
				// Parse Test 02.00.01.00 :: Catch escapes at the end of lines
				logp.Msg("Found newline escape!")
				nt.Text += ci.Text
				nt.Length = utf8.RuneCountInString(nt.Text)
			} else {
				nt = newText(ci)
				t.nodeTarget.append(nt)
			}
		case itemInlineEmphasisOpen:
			t.inlineEmphasis(ci)
		case itemInlineStrongOpen:
			t.inlineStrong(ci)
		case itemInlineLiteralOpen:
			t.inlineLiteral(ci)
		case itemInlineInterpretedTextOpen:
			t.inlineInterpretedText(ci)
		case itemInlineInterpretedTextRoleOpen:
			t.inlineInterpretedTextRole(ci)
		case itemCommentMark:
			t.comment(ci)
		case itemEnumListArabic:
			t.nodeTarget.append(t.enumList(ci))
		case itemBlankLine:
			logp.Msg("Found newline, closing paragraph")
			t.backup()
			break outer
		}
		logp.Msg("Continuing...")
	}
	logp.Log("msg", "number of indents", "t.indents.len", t.indents.len())
	if t.indents.len() > 0 {
		t.nodeTarget.setParent(t.indents.topNode())
		logp.Log("msg", "Set node target to t.indents.topNodeList!", "nodePtr", t.nodeTarget)
	} else if len(t.sectionLevels.levels) == 0 {
		logp.Msg("Setting node target to t.nodes!")
		t.nodeTarget.reset()
	}
	return np
}
