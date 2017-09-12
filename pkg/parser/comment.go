package parser

import (
	doc "github.com/demizer/go-rst/pkg/document"
	tok "github.com/demizer/go-rst/pkg/token"
)

func (p *Parser) comment(i *tok.Item) doc.Node {
	var n doc.Node

	if p.peek(1).Type == tok.BlankLine {
		p.Msg("Found empty comment block")
		n := doc.NewComment(&tok.Item{StartPosition: i.StartPosition, Line: i.Line})
		p.nodeTarget.Append(n)
		return n
	}

	if nSpace := p.peek(1); nSpace != nil && nSpace.Type != tok.Space {
		// The comment element itself is valid, but we need to add it to the NodeList before the systemMessage.
		p.Msg("Missing space after comment mark! (warningExplicitMarkupWithUnIndent)")
		n = doc.NewComment(&tok.Item{Line: i.Line})
		p.systemMessage(warningExplicitMarkupWithUnIndent)
		return n
	}

	nPara := p.peek(2)
	p.Msgr("two peek ahead", "type", nPara.Type)
	if nPara != nil && nPara.Type == tok.Text {
		// Skip the tok.Space
		p.next(2)
		p.Msgr("have token", "token", p.token)
		// See if next line is indented, if so, it is part of the comment
		if p.peek(1).Type == tok.Space && p.peek(2).Type == tok.Text {
			p.Msg("Found NodeComment block")
			p.next(2)
			for {
				nPara.Text += "\n" + p.token.Text
				if p.peek(1).Type == tok.Space && p.peek(2).Type == tok.Text {
					p.next(2)
				} else {
					break
				}
			}
			nPara.Length = len(nPara.Text)
		} else if z := p.peek(1); z != nil && z.Type != tok.BlankLine && z.Type != tok.CommentMark && z.Type != tok.EOF {
			// A valid comment contains a blank line after the comment block
			p.Msg("Found warningExplicitMarkupWithUnIndent")
			n = doc.NewComment(nPara)
			p.nodeTarget.Append(n)
			p.systemMessage(warningExplicitMarkupWithUnIndent)
			return n
		} else {
			// Just a regular single lined comment
			p.Msg("Found one-line NodeComment")
		}
		n = doc.NewComment(nPara)
	}
	p.nodeTarget.Append(n)
	return n
}
