package parser

import (
	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/token"
)

func (p *Parser) comment(i *tok.Item) doc.Node {
	log.Log("msg", "In transition comment", "token", i)
	var n doc.Node

	if p.peek(1).Type == tok.BlankLine {
		log.Msg("Found empty comment block")
		n := doc.NewComment(&tok.Item{StartPosition: i.StartPosition, Line: i.Line})
		p.nodeTarget.Append(n)
		return n
	}

	if nSpace := p.peek(1); nSpace != nil && nSpace.Type != tok.Space {
		// The comment element itself is valid, but we need to add it to the NodeList before the systemMessage.
		log.Msg("Missing space after comment mark! (warningExplicitMarkupWithUnIndent)")
		n = doc.NewComment(&tok.Item{Line: i.Line})
		sm := p.systemMessage(warningExplicitMarkupWithUnIndent)
		p.nodeTarget.Append(n, sm)
		return n
	}

	nPara := p.peek(2)
	log.Log("msg", "two peek ahead", "type", nPara.Type)
	if nPara != nil && nPara.Type == tok.Text {
		// Skip the tok.Space
		p.next(2)
		log.Log("msg", "have token", "token", p.token[zed])
		// See if next line is indented, if so, it is part of the comment
		if p.peek(1).Type == tok.Space && p.peek(2).Type == tok.Text {
			log.Msg("Found NodeComment block")
			p.next(2)
			for {
				nPara.Text += "\n" + p.token[zed].Text
				if p.peek(1).Type == tok.Space && p.peek(2).Type == tok.Text {
					p.next(2)
				} else {
					break
				}
			}
			nPara.Length = len(nPara.Text)
		} else if z := p.peek(1); z != nil && z.Type != tok.BlankLine && z.Type != tok.CommentMark && z.Type != tok.EOF {
			// A valid comment contains a blank line after the comment block
			log.Msg("Found warningExplicitMarkupWithUnIndent")
			n = doc.NewComment(nPara)
			p.nodeTarget.Append(n)
			sm := p.systemMessage(warningExplicitMarkupWithUnIndent)
			p.nodeTarget.Append(sm)
			return n
		} else {
			// Just a regular single lined comment
			log.Msg("Found one-line NodeComment")
		}
		n = doc.NewComment(nPara)
	}
	p.nodeTarget.Append(n)
	return n
}
