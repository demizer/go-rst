package parse

import "testing"

// A single comment
func TestParseCommentGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-comment")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A single comment split with a newline
func TestParseCommentBlockGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-comment-block")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block that begins on the second line after the comment mark
func TestParseCommentBlockOnSecondLineGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-comment-block-second-line")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// One comment after another
func TestParseTwoCommentsGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-two-comments")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// One comment not followed by a blank line
func TestParseCommentNoBlankLineBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-comment-no-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Two comments, no blank line after second comment
func TestParseTwoCommentsNoBlankLineBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-two-comments-no-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment ending with a literal block mark.
func TestParseCommentWithLiteralMarkGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-comment-with-literal-mark")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with a newline after the comment mark
func TestParseNewlineAfterCommentMarkGood0003(t *testing.T) {
	testPath := testPathFromName("00.03-newline-after-comment-mark")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with a newline after the comment mark
func TestParseNewlineAfterCommentMarkGood0004(t *testing.T) {
	testPath := testPathFromName("00.04-newline-after-comment-mark")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with citation syntax in the text
func TestParseCommentNotCitationGood0005(t *testing.T) {
	testPath := testPathFromName("00.05-comment-not-citation")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with substitution definition syntax in the text
func TestParseCommentNotSubstitutionDefinitionGood0005(t *testing.T) {
	testPath := testPathFromName("00.06-comment-not-subs-def")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// An empty comment followed by a blockquote
func TestParseCommentWithBlockquoteGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-empty-comment-with-blockquote")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A definition list with a comment in the definition
func TestParseCommentInDefinitionGood0400(t *testing.T) {
	testPath := testPathFromName("04.00-comment-in-definition")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment after a definition
func TestParseCommentAfterDefinitionGood0401(t *testing.T) {
	testPath := testPathFromName("04.01-comment-after-definition")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment between two paragraphs in a bullet list item
func TestParseCommentBetweenBulletParagraphsGood0500(t *testing.T) {
	testPath := testPathFromName("05.00-comment-between-bullet-paragraphs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment between two ferns... I mean bullets.
func TestParseCommentBetweenBulletsGood0501(t *testing.T) {
	testPath := testPathFromName("05.01-comment-between-bullets")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment trailing a bullet list item
func TestParseCommentTrailingBulletGood0502(t *testing.T) {
	testPath := testPathFromName("05.02-comment-trailing-bullet")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
