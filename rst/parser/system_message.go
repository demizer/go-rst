package parser

func (t *Tree) systemMessageSection(s *SystemMessageNode, err ParserMessage) *LiteralBlockNode {
	var overLine, indent, title, underLine, newLine string
	var lbText string
	var lbTextLen int

	literalBlock := func() *LiteralBlockNode {
		return newLiteralBlock(&item{Type: itemLiteralBlock, Text: lbText, Length: lbTextLen})
	}

	switch err {
	case infoOverlineTooShortForTitle:
		var inText string
		inText = t.token[zed-1].Text + "\n" + t.token[zed].Text
		s.Line = t.token[zed-1].Line
		if t.token[zed-2] != nil {
			inText = t.token[zed-2].Text + "\n" + t.token[zed-1].Text + "\n" + t.token[zed].Text
			s.Line = t.token[zed-2].Line
			t.token[zed-2] = nil

		}
		infoTextLen := len(inText)
		// Modify the token buffer to change the current token to a itemText then backup the token buffer so the
		// next loop gets the new paragraph
		t.token[zed-1] = nil
		t.token[zed].Type = itemText
		t.token[zed].Text = inText
		t.token[zed].Length = infoTextLen
		t.token[zed].Line = s.Line
		t.backup()
	case infoUnexpectedTitleOverlineOrTransition:
		oLin := t.peekBackTo(itemSectionAdornment)
		titl := t.peekBackTo(itemTitle)
		uLin := t.token[zed]
		inText := oLin.Text + "\n" + titl.Text + "\n" + uLin.Text
		s.Line = oLin.Line
		t.clearTokens(zed-4, zed-1)
		infoTextLen := len(inText)
		// Modify the token buffer to change the current token to a itemText then backup the token buffer so the
		// next loop gets the new paragraph
		t.token[zed].Type = itemText
		t.token[zed].Text = inText
		t.token[zed].Length = infoTextLen
		t.token[zed].Line = s.Line
		t.token[zed].StartPosition = oLin.StartPosition
		t.backup()
	case infoUnderlineTooShortForTitle:
		inText := t.token[zed-1].Text + "\n" + t.token[zed].Text
		infoTextLen := len(inText)
		s.Line = t.token[zed-1].Line
		// Modify the token buffer to change the current token to a itemText then backup the token buffer so the
		// next loop gets the new paragraph
		t.token[zed-1] = nil
		t.token[zed].Type = itemText
		t.token[zed].Text = inText
		t.token[zed].Length = infoTextLen
		t.token[zed].Line = s.Line
		t.backup()
	case warningShortOverline, severeOverlineUnderlineMismatch:
		backToken := zed - 2
		if t.peekBack(2).Type == itemSpace {
			backToken = zed - 3
			indent = t.token[zed-2].Text
		}
		overLine = t.token[backToken].Text
		title = t.token[zed-1].Text
		underLine = t.token[zed].Text
		newLine = "\n"
		lbText = overLine + newLine + indent + title + newLine + underLine
		s.Line = t.token[backToken].Line
		lbTextLen = len(lbText)
		return literalBlock()
	case warningShortUnderline, severeUnexpectedSectionTitle:
		backToken := zed - 1
		if t.peekBack(1).Type == itemSpace {
			backToken = zed - 2
		}
		lbText = t.token[backToken].Text + "\n" + t.token[zed].Text
		lbTextLen = len(lbText)
		s.Line = t.token[zed-1].Line
		return literalBlock()
	case errorInvalidSectionOrTransitionMarker:
		lbText = t.token[zed-1].Text + "\n" + t.token[zed].Text
		s.Line = t.token[zed-1].Line
		lbTextLen = len(lbText)
		return literalBlock()
	case severeIncompleteSectionTitle,
		severeMissingMatchingUnderlineForOverline:
		lbText = t.token[zed-2].Text + "\n" + t.token[zed-1].Text + t.token[zed].Text
		s.Line = t.token[zed-2].Line
		lbTextLen = len(lbText)
		return literalBlock()
	case severeUnexpectedSectionTitleOrTransition:
		lbText = t.token[zed].Text
		lbTextLen = len(lbText)
		s.Line = t.token[zed].Line
		return literalBlock()
	case severeTitleLevelInconsistent:
		if t.peekBack(2).Type == itemSectionAdornment {
			lbText = t.token[zed-2].Text + "\n" + t.token[zed-1].Text + "\n" + t.token[zed].Text
			lbTextLen = len(lbText)
			s.Line = t.token[zed-2].Line
			return literalBlock()
		}
		lbText = t.token[zed-1].Text + "\n" + t.token[zed].Text
		lbTextLen = len(lbText)
		s.Line = t.token[zed-1].Line
		return literalBlock()
	}
	return nil
}

func (t *Tree) systemMessageInlineMarkup(s *SystemMessageNode, err ParserMessage) *LiteralBlockNode {
	switch err {
	case warningExplicitMarkupWithUnIndent:
		s.Line = t.token[zed+1].Line
	}
	return nil
}

// systemMessage generates a Node based on the passed ParserMessage. The generated message is returned as a
// SystemMessageNode.
func (t *Tree) systemMessage(err ParserMessage) Node {
	s := newSystemMessage(&item{Type: itemSystemMessage, Line: t.token[zed].Line}, err)
	msg := newText(&item{
		Text:   err.Message(),
		Length: len(err.Message()),
	})

	logp.Log("msg", "Adding msg to system message NodeList", "systemMessage", err)
	s.NodeList.append(msg)

	appendOrDie := func(f func(s2 *SystemMessageNode, err2 ParserMessage) *LiteralBlockNode) {
		if lb := f(s, err); lb != nil {
			s.NodeList = append(s.NodeList, lb)
		}
	}

	appendOrDie(t.systemMessageSection)
	appendOrDie(t.systemMessageInlineMarkup)

	return s
}
