package parser

func (p *ParserState) comment(i *item) Node {
	logp.Log("msg", "In transition comment", "token", i)
	var n Node

	if p.peek(1).Type == itemBlankLine {
		logp.Msg("Found empty comment block")
		n := newComment(&item{StartPosition: i.StartPosition, Line: i.Line})
		p.nodeTarget.append(n)
		return n
	}

	if nSpace := p.peek(1); nSpace != nil && nSpace.Type != itemSpace {
		// The comment element itself is valid, but we need to add it to the NodeList before the systemMessage.
		logp.Msg("Missing space after comment mark! (warningExplicitMarkupWithUnIndent)")
		n = newComment(&item{Line: i.Line})
		sm := p.systemMessage(warningExplicitMarkupWithUnIndent)
		p.nodeTarget.append(n, sm)
		return n
	}

	nPara := p.peek(2)
	logp.Log("msg", "two peek ahead", "type", nPara.Type)
	if nPara != nil && nPara.Type == itemText {
		// Skip the itemSpace
		p.next(2)
		logp.Log("msg", "have token", "token", p.token[zed])
		// See if next line is indented, if so, it is part of the comment
		if p.peek(1).Type == itemSpace && p.peek(2).Type == itemText {
			logp.Msg("Found NodeComment block")
			p.next(2)
			for {
				nPara.Text += "\n" + p.token[zed].Text
				if p.peek(1).Type == itemSpace && p.peek(2).Type == itemText {
					p.next(2)
				} else {
					break
				}
			}
			nPara.Length = len(nPara.Text)
		} else if z := p.peek(1); z != nil && z.Type != itemBlankLine && z.Type != itemCommentMark && z.Type != itemEOF {
			// A valid comment contains a blank line after the comment block
			logp.Msg("Found warningExplicitMarkupWithUnIndent")
			n = newComment(nPara)
			p.nodeTarget.append(n)
			sm := p.systemMessage(warningExplicitMarkupWithUnIndent)
			p.nodeTarget.append(sm)
			return n
		} else {
			// Just a regular single lined comment
			logp.Msg("Found one-line NodeComment")
		}
		n = newComment(nPara)
	}
	p.nodeTarget.append(n)
	return n
}
