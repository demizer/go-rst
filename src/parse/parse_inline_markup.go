package parse

import "unicode/utf8"

func (t *Tree) inlineEmphasis(i *item) {
	ni := t.next(1)
	if len(*t.Nodes) == 0 {
		np := newParagraph()
		t.nodeTarget.append(np)
		t.nodeTarget.setParent(np)
	}
main:
	for {
		ci := t.next(1)
		if ci == nil {
			break
		}
		switch ci.Type {
		case itemInlineEmphasis:
			ni.Text += "\n" + ci.Text
		case itemBlankLine:
			continue
		default:
			t.backup()
			break main
		}
	}
	ni.Length = utf8.RuneCountInString(ni.Text)
	t.nodeTarget.append(newInlineEmphasis(ni))
	t.next(1)
}

func (t *Tree) inlineStrong(i *item) {
	t.next(1)
	if len(*t.Nodes) == 0 {
		np := newParagraph()
		t.nodeTarget.append(np)
		t.nodeTarget.setParent(np)
	}
	t.nodeTarget.append(newInlineStrong(t.token[zed]))
	t.next(1)
}

func (t *Tree) inlineLiteral(i *item) {
	ni := t.next(1)
	if len(*t.Nodes) == 0 {
		np := newParagraph()
		t.nodeTarget.append(np)
		t.nodeTarget.setParent(np)
	}
main:
	for {
		ci := t.next(1)
		if ci == nil {
			break
		}
		switch ci.Type {
		case itemInlineLiteral:
			ni.Text += "\n" + ci.Text
		case itemBlankLine:
			continue
		default:
			t.backup()
			break main
		}
	}
	ni.Length = utf8.RuneCountInString(ni.Text)
	t.nodeTarget.append(newInlineLiteral(ni))
	t.next(1)
}

func (t *Tree) inlineInterpretedText(i *item) {
	t.next(1)
	n := newInlineInterpretedText(t.token[zed])
	t.nodeTarget.append(n)
	t.next(1)
	if t.peek(1).Type == itemInlineInterpretedTextRoleOpen {
		t.next(2)
		n.NodeList.append(newInlineInterpretedTextRole(t.token[zed]))
		t.next(1)
	}
}

func (t *Tree) inlineInterpretedTextRole(i *item) {
	t.next(1)
	t.nodeTarget.append(newInlineInterpretedTextRole(t.token[zed]))
	t.next(1)
}
