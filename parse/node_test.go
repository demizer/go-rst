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

func TestIndentNodeType(t *testing.T) {
	n := &IndentNode{Type: NodeIndent}
	if n.NodeType() != NodeIndent {
		t.Error("n.Type != NodeIndent")
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
