// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import "testing"

func TestLexCommentGood0000(t *testing.T) {
	// A single comment
	testPath := testPathFromName("00.00-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentBlockGood0001(t *testing.T) {
	// A single comment block split with a newline
	testPath := testPathFromName("00.01-comment-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentBlockOnSecondLineGood0100(t *testing.T) {
	// A comment block that begins on the second line after the comment
	// mark
	testPath := testPathFromName("01.00-comment-block-second-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexTwoCommentsGood0200(t *testing.T) {
	// One comment after another
	testPath := testPathFromName("02.00-two-comments")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentNoBlankLineBad0000(t *testing.T) {
	// One comment not followed by a blank line
	testPath := testPathFromName("00.00-comment-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexTwoCommentsNoBlankLineBad0001(t *testing.T) {
	// Two comments, no blank line after second comment
	testPath := testPathFromName("00.01-two-comments-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentWithLiteralMarkGood0002(t *testing.T) {
	// A comment ending with a literal block mark.
	testPath := testPathFromName("00.02-comment-with-literal-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexNewlineAfterCommentMarkGood0003(t *testing.T) {
	// A comment block with a newline after the comment mark
	testPath := testPathFromName("00.03-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexNewlineAfterCommentMarkGood0004(t *testing.T) {
	// A comment block with a newline after the comment mark
	testPath := testPathFromName("00.04-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentNotCitationGood0005(t *testing.T) {
	// A comment block with citation syntax in the text
	testPath := testPathFromName("00.05-comment-not-citation")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentNotSubstitutionDefinitionGood0005(t *testing.T) {
	// A comment block with substitution definition syntax in the text
	testPath := testPathFromName("00.06-comment-not-subs-def")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentWithBlockquoteGood0300(t *testing.T) {
	// An empty comment followed by a blockquote
	testPath := testPathFromName("03.00-empty-comment-with-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentInDefinitionGood0400(t *testing.T) {
	// A definition list with a comment in the definition
	testPath := testPathFromName("04.00-comment-in-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentAfterDefinitionGood0401(t *testing.T) {
	// A comment after a definition
	testPath := testPathFromName("04.01-comment-after-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
