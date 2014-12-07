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
