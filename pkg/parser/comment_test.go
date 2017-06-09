package parser

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

// A single comment
func Test_00_00_00_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.00-comment")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// One comment not followed by a blank line
func Test_00_00_00_01_ParseCommentNoBlankLineBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.01-comment-no-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A single comment split with a newline
func Test_00_00_00_02_ParseCommentBlockGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.02-comment-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment ending with a literal block mark.
func Test_00_00_00_03_ParseCommentWithLiteralMarkGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.03-comment-with-literal-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with a newline after the comment mark
func Test_00_00_00_04_ParseNewlineAfterCommentMarkGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.04-newline-after-comment-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with a newline after the comment mark
func Test_00_00_00_05_ParseNewlineAfterCommentMarkGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.05-newline-after-comment-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with citation syntax in the text
func Test_00_00_00_06_ParseCommentNotCitationGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.06-comment-not-citation")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block with substitution definition syntax in the text
func Test_00_00_00_07_ParseCommentNotSubstitutionDefinitionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.07-comment-not-subs-def")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_00_08_ParseCommentIsNotReferenceGood_NotImplemented(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.08-comment-not-reference")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment block that begins on the second line after the comment mark
func Test_00_00_01_00_ParseCommentBlockOnSecondLineGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.00-comment-block-second-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// One comment after another
func Test_00_00_02_00_ParseTwoCommentsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.00-two-comments")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Two comments, no blank line after second comment
func Test_00_00_02_01_ParseTwoCommentsNoBlankLineBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.01-two-comments-no-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// An empty comment followed by a blockquote
func Test_00_00_03_00_ParseCommentWithBlockquoteGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.03.00-empty-comment-with-blockquote")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A definition list with a comment in the definition
func Test_00_00_04_00_ParseCommentInDefinitionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.00-comment-in-definition")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment after a definition
func Test_00_00_04_01_ParseCommentAfterDefinitionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.01-comment-after-definition")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment between two paragraphs in a bullet list item
func Test_00_00_05_00_ParseCommentBetweenBulletParagraphsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.00-comment-between-bullet-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment between two ferns... I mean bullets.
func Test_00_00_05_01_ParseCommentBetweenBulletsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.01-comment-between-bullets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A comment trailing a bullet list item
func Test_00_00_05_02_ParseCommentTrailingBulletGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.02-comment-trailing-bullet")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
