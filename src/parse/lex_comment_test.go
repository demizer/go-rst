package parse

import "testing"

// A single comment
func TestLexCommentGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A single comment block split with a newline
func TestLexCommentBlockGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-comment-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block that begins on the second line after the comment mark
func TestLexCommentBlockOnSecondLineGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-comment-block-second-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// One comment after another
func TestLexTwoCommentsGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-two-comments")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// One comment not followed by a blank line
func TestLexCommentNoBlankLineBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-comment-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Two comments, no blank line after second comment
func TestLexTwoCommentsNoBlankLineBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-two-comments-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment ending with a literal block mark.
func TestLexCommentWithLiteralMarkGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-comment-with-literal-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with a newline after the comment mark
func TestLexNewlineAfterCommentMarkGood0003(t *testing.T) {
	testPath := testPathFromName("00.03-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with a newline after the comment mark
func TestLexNewlineAfterCommentMarkGood0004(t *testing.T) {
	testPath := testPathFromName("00.04-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with citation syntax in the text
func TestLexCommentNotCitationGood0005(t *testing.T) {
	testPath := testPathFromName("00.05-comment-not-citation")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with substitution definition syntax in the text
func TestLexCommentNotSubstitutionDefinitionGood0005(t *testing.T) {
	testPath := testPathFromName("00.06-comment-not-subs-def")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// An empty comment followed by a blockquote
func TestLexCommentWithBlockquoteGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-empty-comment-with-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A definition list with a comment in the definition
func TestLexCommentInDefinitionGood0400(t *testing.T) {
	testPath := testPathFromName("04.00-comment-in-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment after a definition
func TestLexCommentAfterDefinitionGood0401(t *testing.T) {
	testPath := testPathFromName("04.01-comment-after-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment between two bullet paragraphs
func TestLexCommentBetweenBulletParagrapsGood0500(t *testing.T) {
	testPath := testPathFromName("05.00-comment-between-bullet-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment between two ferns... I mean bullets.
func TestLexCommentBetweenBulletsGood0501(t *testing.T) {
	testPath := testPathFromName("05.01-comment-between-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment trailing a bullet list item
func TestLexCommentTrailingBulletGood0502(t *testing.T) {
	testPath := testPathFromName("05.02-comment-trailing-bullet")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
