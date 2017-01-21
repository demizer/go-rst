package parse

func (t *Tree) emptyblockquote(i *item) {
	//
	//  FIXME: Blockquote parsing is NOT fully implemented.
	//
	sec := newEmptyBlockQuote(i)
	t.nodeTarget.append(sec)
	t.nodeTarget.setParent(sec)
	t.bqLevel = sec
}

func (t *Tree) blockquote(i *item) {
	//
	//  FIXME: Blockquote parsing is NOT fully implemented.
	//
	if t.bqLevel != nil {
		// Parser Test 03.02.07.00
		logp.Msg("Adding blockquote text as NodeText to existing blockquote")
		t.bqLevel.NodeList.append(newParagraphWithNodeText(i))
		return
	}
	logp.Msg("Creating blockquote")
	sec := newBlockQuote(i)
	t.nodeTarget.append(sec)
	t.nodeTarget.setParent(sec)
	t.bqLevel = sec
}
