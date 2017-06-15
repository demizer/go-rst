package token

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_00_00_00_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.00-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_00_01_LexCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.01-bad-comment-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_00_02_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.02-comment-with-literal-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_00_03_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.03-comment-not-reference")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_01_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.00-comment-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_01_01_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.01-comment-block-second-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_02_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.00-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_02_01_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.01-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_02_02_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.02-comment-not-citation")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_02_03_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.03-comment-not-subs-def")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_03_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.03.00-empty-comment-with-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_04_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.00-comment-in-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_04_01_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.01-comment-after-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_05_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.00-comment-between-bullet-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_05_01_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.01-comment-between-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_05_02_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.02-comment-trailing-bullet")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_06_00_LexCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.00-two-comments")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_00_00_06_01_LexCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.01-bad-two-comments-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
