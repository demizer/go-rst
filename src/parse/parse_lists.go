package parse

import (
	"errors"
	"fmt"
	"os"
)

func (t *Tree) definitionTerm(i *item) Node {
	//
	//  FIXME: Definition list parsing is NOT fully implemented.
	//
	dl := newDefinitionList(&item{Line: i.Line})
	t.nodeTarget.append(dl)
	t.nodeTarget.setParent(dl)
	t.next(1)

	// Container for definition items
	dli := newDefinitionListItem(i, t.peek(1))
	t.nodeTarget.append(dli)
	t.nodeTarget.setParent(dli.Definition)

	// Gather definitions and body elements
	for {
		ni := t.next(1)
		if ni == nil {
			break
		}
		logp.Log("msg", "Have token", "token", ni)
		pb := t.peekBack(1)
		if ni.Type == itemSpace {
			logp.Msg("continue; ni.Type == itemSpace")
			continue
		} else if ni.Type == itemEOF {
			logp.Msg("break; ni.Type == itemEOF")
			break
		} else if ni.Type == itemBlankLine {
			logp.Msg("Setting nodeTarget to dli")
			t.nodeTarget.setParent(dli.Definition)
		} else if ni.Type == itemCommentMark && (pb != nil && pb.Type != itemSpace) {
			// Comment at start of the line breaks current definition list
			logp.Msg("Have itemCommentMark at start of the line!")
			t.nodeTarget.reset()
			t.backup()
			break
		} else if ni.Type == itemDefinitionText {
			// FIXME: This function is COMPLETELY not final. It is only setup for passing section test TitleNumberedGood0100.
			np := newParagraphWithNodeText(ni)
			t.nodeTarget.append(np)
			t.nodeTarget.setParent(np)
			logp.Msg("continue; ni.Type == itemDefinitionText")
			continue
		} else if ni.Type == itemDefinitionTerm {
			dli2 := newDefinitionListItem(ni, t.peek(2))
			t.nodeTarget.setParent(dl)
			t.nodeTarget.append(dli2)
			t.nodeTarget.setParent(dli2.Definition)
			logp.Msg("continue; ni.Type == itemDefinitionTerm")
			continue
		}
		t.subParseBodyElements(ni)
	}
	return dl
}

func (t *Tree) bulletList(i *item) {
	//
	// FIXME: Bullet List Parsing is NOT fully implemented
	//
	bl := newBulletListNode(i)
	t.openList = bl
	t.nodeTarget.append(bl)
	t.nodeTarget.setParent(bl)

	// Get the bullet list paragraph
	t.next(1)
	bli := newBulletListItemNode(i)

	// Set the node target to the bullet list paragraph
	t.nodeTarget.append(bli)
	t.nodeTarget.setParent(bli)
	t.indents.add(i, bli)

	// Capture all bullet items until un-indent
	for {
		ni := t.next(1)
		logp.Log("msg", "Have token", "token", fmt.Sprintf("%+#v", ni))
		if ni == nil {
			logp.Log("break next item == nil")
			break
		} else if ni.Type == itemEOF {
			logp.Log("break itemEOF")
			break
		} else if t.indents.len() > 0 && len(*t.indents.topNodeList()) > 0 && t.peekBack(1).Type == itemSpace &&
			t.peekBack(2).Type != itemCommentMark {
			logp.Log("msg", "Have indents",
				"lastStartPosition", t.indents.lastStartPosition(),
				"ni.StartPosition", ni.StartPosition)
			if t.indents.lastStartPosition() != ni.StartPosition {
				// FIXME: WE SHOULD NEVER EXIT IN LIBRARY !! This is just debug code, but we need to add
				// proper handling for this ...
				logp.Log(errors.New("Unexpected un-indent!"))
				spd.Dump(t.indents)
				os.Exit(1)
			}
		}

		t.subParseBodyElements(ni)
	}
	t.indents.pop()
}

var lastEnum *EnumListNode

func (t *Tree) enumList(i *item) (n Node) {
	var eNode *EnumListNode
	var affix *item
	if lastEnum == nil {
		t.next(1)
		affix = t.token[zed]
		t.next(1)
		eNode = newEnumListNode(i, affix)
		t.next(1)
		eNode.NodeList.append(newParagraphWithNodeText(t.token[zed]))
	} else {
		t.next(3)
		lastEnum.NodeList.append(newParagraphWithNodeText(t.token[zed]))
		return nil
	}
	lastEnum = eNode
	return eNode
}
