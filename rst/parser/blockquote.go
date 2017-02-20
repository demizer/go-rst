package parser

func (p *ParserState) emptyblockquote(i *item) {
	//
	//  FIXME: Blockquote parsing is NOT fully implemented.
	//
	sec := newEmptyBlockQuote(i)
	p.nodeTarget.append(sec)
	p.nodeTarget.setParent(sec)
	p.bqLevel = sec
}

func (p *ParserState) blockquote(i *item) {
	//
	//  FIXME: Blockquote parsing is NOT fully implemented.
	//
	if p.bqLevel != nil {
		// Parser Test 03.02.07.00
		logp.Msg("Adding blockquote text as NodeText to existing blockquote")
		p.bqLevel.NodeList.append(newParagraphWithNodeText(i))
		return
	}
	logp.Msg("Creating blockquote")
	sec := newBlockQuote(i)
	p.nodeTarget.append(sec)
	p.nodeTarget.setParent(sec)
	p.bqLevel = sec
}
