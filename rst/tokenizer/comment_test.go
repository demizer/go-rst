package tokenizer

import (
	"os"
	"testing"
)

// A single comment
func Test_00_00_00_00_LexCommentGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.00-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// One comment not followed by a blank line
func Test_00_00_00_01_LexCommentNoBlankLineBad(t *testing.T) {
	testPath := testPathFromName("00.00.00.01-comment-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A single comment block split with a newline
func Test_00_00_00_02_LexCommentBlockGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.02-comment-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment ending with a literal block mark.
func Test_00_00_00_03_LexCommentWithLiteralMarkGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.03-comment-with-literal-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with a newline after the comment mark
func Test_00_00_00_04_LexNewlineAfterCommentMarkGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.04-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with a newline after the comment mark
func Test_00_00_00_05_LexNewlineAfterCommentMarkGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.05-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with citation syntax in the text
func Test_00_00_00_06_LexCommentNotCitationGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.06-comment-not-citation")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block with substitution definition syntax in the text
func Test_00_00_00_07_LexCommentNotSubstitutionDefinitionGood(t *testing.T) {
	testPath := testPathFromName("00.00.00.07-comment-not-subs-def")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_00_00_00_08_LexCommentIsNotReferenceGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("00.00.00.08-comment-not-reference")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment block that begins on the second line after the comment mark
func Test_00_00_01_00_LexCommentBlockOnSecondLineGood(t *testing.T) {
	testPath := testPathFromName("00.00.01.00-comment-block-second-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// One comment after another
func Test_00_00_02_00_LexTwoCommentsGood(t *testing.T) {
	testPath := testPathFromName("00.00.02.00-two-comments")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Two comments, no blank line after second comment
func Test_00_00_02_01_LexTwoCommentsNoBlankLineBad(t *testing.T) {
	testPath := testPathFromName("00.00.02.01-two-comments-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// An empty comment followed by a blockquote
func Test_00_00_03_00_LexCommentWithBlockquoteGood(t *testing.T) {
	testPath := testPathFromName("00.00.03.00-empty-comment-with-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A definition list with a comment in the definition
func Test_00_00_04_00_LexCommentInDefinitionGood(t *testing.T) {
	testPath := testPathFromName("00.00.04.00-comment-in-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment after a definition
func Test_00_00_04_01_LexCommentAfterDefinitionGood(t *testing.T) {
	testPath := testPathFromName("00.00.04.01-comment-after-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment between two bullet paragraphs
func Test_00_00_05_00_LexCommentBetweenBulletParagrapsGood(t *testing.T) {
	testPath := testPathFromName("00.00.05.00-comment-between-bullet-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment between two ferns... I mean bullets.
func Test_00_00_05_01_LexCommentBetweenBulletsGood(t *testing.T) {
	testPath := testPathFromName("00.00.05.01-comment-between-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A comment trailing a bullet list item
func Test_00_00_05_02_LexCommentTrailingBulletGood(t *testing.T) {
	testPath := testPathFromName("00.00.05.02-comment-trailing-bullet")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
