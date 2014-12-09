// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import "testing"

func TestParseCommentGood0000(t *testing.T) {
	// A single comment
	testPath := testPathFromName("00.00-comment")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentBlockGood0001(t *testing.T) {
	// A single comment split with a newline
	testPath := testPathFromName("00.01-comment-block")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentBlockOnSecondLineGood0100(t *testing.T) {
	// A comment block that begins on the second line after the comment
	// mark
	testPath := testPathFromName("01.00-comment-block-second-line")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseTwoCommentsGood0200(t *testing.T) {
	// One comment after another
	testPath := testPathFromName("02.00-two-comments")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentNoBlankLineBad0000(t *testing.T) {
	// One comment not followed by a blank line
	testPath := testPathFromName("00.00-comment-no-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseTwoCommentsNoBlankLineBad0001(t *testing.T) {
	// Two comments, no blank line after second comment
	testPath := testPathFromName("00.01-two-comments-no-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentWithLiteralMarkGood0002(t *testing.T) {
	// A comment ending with a literal block mark.
	testPath := testPathFromName("00.02-comment-with-literal-mark")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseNewlineAfterCommentMarkGood0003(t *testing.T) {
	// A comment block with a newline after the comment mark
	testPath := testPathFromName("00.03-newline-after-comment-mark")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseNewlineAfterCommentMarkGood0004(t *testing.T) {
	// A comment block with a newline after the comment mark
	testPath := testPathFromName("00.04-newline-after-comment-mark")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentNotCitationGood0005(t *testing.T) {
	// A comment block with citation syntax in the text
	testPath := testPathFromName("00.05-comment-not-citation")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentNotSubstitutionDefinitionGood0005(t *testing.T) {
	// A comment block with substitution definition syntax in the text
	testPath := testPathFromName("00.06-comment-not-subs-def")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentWithBlockquoteGood0300(t *testing.T) {
	// An empty comment followed by a blockquote
	testPath := testPathFromName("03.00-empty-comment-with-blockquote")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentInDefinitionGood0400(t *testing.T) {
	// A definition list with a comment in the definition
	testPath := testPathFromName("04.00-comment-in-definition")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentAfterDefinitionGood0401(t *testing.T) {
	// A comment after a definition
	testPath := testPathFromName("04.01-comment-after-definition")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentBetweenBulletParagraphsGood0500(t *testing.T) {
	// A comment between two paragraphs in a bullet list item
	testPath := testPathFromName("05.00-comment-between-bullet-paragraphs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentBetweenBulletsGood0501(t *testing.T) {
	// A comment between two ferns... I mean bullets.
	testPath := testPathFromName("05.01-comment-between-bullets")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseCommentTrailingBulletGood0502(t *testing.T) {
	// A comment trailing a bullet list item
	testPath := testPathFromName("05.02-comment-trailing-bullet")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
