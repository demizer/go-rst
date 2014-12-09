package parse

import (
	"testing"
)

func TestNodeType(t *testing.T) {
	n := NodeSection
	if n.Type() != NodeSection {
		t.Error("node.Type != NodeSection")
	}
}

func TestAdornmentNodeType(t *testing.T) {
	n := &AdornmentNode{Type: NodeAdornment}
	if n.NodeType() != NodeAdornment {
		t.Error("n.Type != NodeAdornment")
	}
}

func TestLiteralBlockNodeType(t *testing.T) {
	n := &LiteralBlockNode{Type: NodeLiteralBlock}
	if n.NodeType() != NodeLiteralBlock {
		t.Error("n.Type != NodeLiteralBlock")
	}
}

func TestTransitionNodeType(t *testing.T) {
	n := &TransitionNode{Type: NodeTransition}
	if n.NodeType() != NodeTransition {
		t.Error("n.Type != NodeTransition")
	}
}

func TestTitleNodeType(t *testing.T) {
	n := &TitleNode{Type: NodeTitle}
	if n.NodeType() != NodeTitle {
		t.Error("n.Type != NodeTitle")
	}
}

func TestCommentNodeType(t *testing.T) {
	n := &CommentNode{Type: NodeComment}
	if n.NodeType() != NodeComment {
		t.Error("n.Type != NodeComment")
	}
}

func TestDefinitionListNodeType(t *testing.T) {
	n := &DefinitionListNode{Type: NodeDefinitionList}
	if n.NodeType() != NodeDefinitionList {
		t.Error("n.Type != NodeDefinitionList")
	}
}

func TestDefinitionListItemNodeType(t *testing.T) {
	n := &DefinitionListItemNode{Type: NodeDefinitionListItem}
	if n.NodeType() != NodeDefinitionListItem {
		t.Error("n.Type != NodeDefinitionListItem")
	}
}

func TestDefinitionTermNodeType(t *testing.T) {
	n := &DefinitionTermNode{Type: NodeDefinitionTerm}
	if n.NodeType() != NodeDefinitionTerm {
		t.Error("n.Type != NodeDefinitionTerm")
	}
}

func TestDefinitionNodeType(t *testing.T) {
	n := &DefinitionNode{Type: NodeDefinition}
	if n.NodeType() != NodeDefinition {
		t.Error("n.Type != NodeDefinition")
	}
}

func TestBulletListType(t *testing.T) {
	n := &BulletListNode{Type: NodeBulletList}
	if n.NodeType() != NodeBulletList {
		t.Error("n.Type != NodeBulletList")
	}
}
