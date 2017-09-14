package token

//
// AUTO-GENERATED using tools/gentests.go
//

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_00_00_00_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.00-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_00_01_LexerCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.01-bad-comment-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_00_02_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.02-comment-with-literal-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_00_03_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.03-comment-not-reference")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_01_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.00-comment-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_01_01_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.01-comment-block-second-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_02_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.00-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_02_01_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.01-newline-after-comment-mark")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_02_02_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.02-comment-not-citation")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_02_03_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.03-comment-not-subs-def")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_03_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.03.00-empty-comment-with-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_04_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.00-comment-in-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_04_01_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.01-comment-after-definition")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_05_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.00-comment-between-bullet-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_05_01_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.01-comment-between-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_05_02_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.02-comment-trailing-bullet")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_06_00_LexerCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.00-two-comments")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_00_00_06_01_LexerCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.01-bad-two-comments-no-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_00_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.00-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_00_01_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.01-optional-space-before-colon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_00_02_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.02-bad-target-missing-backquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_00_03_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.03-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_00_04_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.04-bad-target-malformed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_01_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.01.00-long-target-names")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_02_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.02.00-target-beginning-with-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_02_01_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.02.01-bad-beginning-with-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_03_00_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.00-bad-duplicate-implicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_03_01_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.01-bad-duplicate-implicit-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_03_02_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.02-bad-duplicate-implicit-directive-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_04_00_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.00-bad-duplicate-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_04_01_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.01-bad-duplicate-explicit-directive-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_04_02_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.02-bad-duplicate-implicit-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_05_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.05.00-escaped-colon-at-the-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_00_05_01_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.05.01-bad-unescaped-colon-at-the-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_00_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.00.00-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_00_01_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.00.01-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_01_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.01.00-external-target-mailto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_02_00_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.02.00-bad-duplicate-external-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_02_01_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.02.01-bad-duplicate-external-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_03_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.03.00-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_03_01_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.03.01-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_01_03_02_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.03.02-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_02_00_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.00.00-indirect-hyperlink-targets-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_02_00_01_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.00.01-indirect-hyperlink-targets-phrase-references")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_02_01_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.01.00-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_02_01_01_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.01.01-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_02_02_00_LexerReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.02.00-bad-anon-and-named-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_01_02_03_00_LexerReferenceHyperlinkTargetsGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.03.00-anonymous-indirect-target-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_02_00_00_00_LexerParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.00-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_02_00_00_01_LexerParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.01-with-line-break")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_02_00_00_02_LexerParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.02-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_02_00_01_00_LexerParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.00-two-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_02_00_01_01_LexerParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.01-two-paragraphs-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_00_00_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.00-paragraph-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_00_01_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.01-paragraph-blockquote-short-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_00_02_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.02-paragraph-blockquote-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_00_03_LexerBlockquoteBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.03-bad-no-blank-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_00_04_LexerBlockquoteBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.04-bad-unexpected-indent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_01_00_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.01.00-two-levels")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_02_00_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.02.00-unicode-em-dash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_03_00_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.03.00-uneven-indents")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_04_00_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.00-paragraph-blockquote-attrib")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_04_01_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.01-paragraph-blockquote-two-line-attrib")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_04_02_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.02-paragraph-blockquote-attrib-no-space")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_04_03_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.03-paragraph-blockquote-one-attrib")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_04_04_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.04-paragraph-blockquote-attrib-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_03_00_04_05_LexerBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.05-paragraph-blockquote-attrib-with-invalid-attrib")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_00_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.00-title-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_00_01_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.01-paragraph-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_00_02_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.02-title-combining-chars")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_00_03_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.03-bad-short-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_00_04_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.04-bad-short-title-short-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_01_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.01.00-paragraph-head-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_02_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.02.00-short-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_03_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.03.00-empty-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_04_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.04.00-numbered-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_04_01_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.04.01-bad-enum-list-with-numbered-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_00_05_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.05.00-title-with-imu")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_00_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.00-title-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_00_01_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.01-bad-title-too-long")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_00_02_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.02-bad-short-title-short-overline-and-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_00_03_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.03-bad-short-title-short-overline-missing-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_01_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.00-inset-title-with-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_01_01_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.01-bad-inset-title-missing-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_01_02_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.02-bad-inset-title-mismatched-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_01_03_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.03-bad-inset-title-missing-underline-with-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_01_04_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.04-bad-inset-title-missing-underline-and-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_02_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.02.00-three-char-section-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_03_00_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.03.00-bad-unexpected-titles")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_04_00_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.04.00-bad-missing-titles-with-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_04_01_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.04.01-bad-missing-titles-with-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_05_00_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.05.00-bad-incomplete-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_05_01_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.05.01-bad-incomplete-sections-no-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_06_00_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.06.00-bad-indented-title-short-overline-and-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_01_07_00_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.07.00-bad-two-char-section-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_00_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.00-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_00_01_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.01-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_00_02_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.02-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_00_03_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.03-bad-subsection-order")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_01_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.00-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_01_01_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.01-bad-two-level-overline-bad-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_01_02_LexerSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.02-bad-subsection-order-with-overlines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_04_02_02_00_LexerSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.02.00-two-level-one-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-literal-block-space-after-colons")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-literal-block-space-after-colons")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-literal-block-space-after-colons")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-literal-block-space-after-colons")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-bad-unindented-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-bad-unindented-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-bad-unindented-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-bad-unindented-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_03_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.03-bad-no-blankline-after-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_03_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.03-bad-no-blankline-after-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_03_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.03-bad-no-blankline-after-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_03_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.03-bad-no-blankline-after-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_04_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.04-multiline-paragraph-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_04_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.04-multiline-paragraph-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_04_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.04-multiline-paragraph-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_04_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.04-multiline-paragraph-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_05_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.05-bad-no-blankline-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_05_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.05-bad-no-blankline-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_05_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.05-bad-no-blankline-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_05_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.05-bad-no-blankline-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_06_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.06-paragraph-space-double-colon-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_06_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.06-paragraph-space-double-colon-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_06_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.06-paragraph-space-double-colon-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_06_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.06-paragraph-space-double-colon-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_07_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.07-paragraph-colon-newline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_07_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.07-paragraph-colon-newline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_07_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.07-paragraph-colon-newline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_07_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.07-paragraph-colon-newline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_08_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.08-bad-section-underline-not-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_08_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.08-bad-section-underline-not-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_08_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.08-bad-section-underline-not-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_08_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.08-bad-section-underline-not-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_09_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.09-bad-eof-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_09_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.09-bad-eof-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_09_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.09-bad-eof-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_00_09_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.09-bad-eof-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_01_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.00-multiline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_01_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.01-wonky-multiline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_01_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.01-wonky-multiline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_01_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.01-wonky-multiline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_01_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.01-wonky-multiline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-double-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-double-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-double-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-double-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.01-literal-block-and-escaped-colon-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.01-literal-block-and-escaped-colon-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.01-literal-block-and-escaped-colon-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_00_02_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.01-literal-block-and-escaped-colon-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.00-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.00-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.00-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.00-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.01-quoted-literal-block-two-blanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.01-quoted-literal-block-two-blanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.01-quoted-literal-block-two-blanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_01_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.01-quoted-literal-block-two-blanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.02-bad-inconsistent-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.02-bad-inconsistent-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.02-bad-inconsistent-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_00_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.02-bad-inconsistent-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.00-quoted-literal-block-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.00-quoted-literal-block-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.00-quoted-literal-block-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_00_LexerLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.00-quoted-literal-block-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_01_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.01-bad-indented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_01_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.01-bad-indented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_01_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.01-bad-indented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_01_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.01-bad-indented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.02-bad-unindented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.02-bad-unindented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.02-bad-unindented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_05_01_01_02_LexerLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.02-bad-unindented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_00_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.00.00-double-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_01_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.01.00-lots-of-escaping")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_02_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.02.00-lots-of-escaping-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_03_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.00-emphasis-wrapped-in-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_03_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.01-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_03_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.02-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_04_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.00-openers-and-closers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_04_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.01-strong-and-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_00_05_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.05.00-emphasis-with-backwards-rule-5")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_00_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.00-strong")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_00_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.01-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_00_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.02-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_01_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.01.00-strong-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_02_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.02.00-strong-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_03_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.00-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_03_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.01-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_03_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.02-strong-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_01_04_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.04.00-strong-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_00_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.00-simple-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_00_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.01-single-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_00_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.02-emphasis-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_00_03_LexerInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.03-bad-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_00_04_LexerInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.04-bad-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_00_05_LexerInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.05-bad-emphasis-unclosed-surrounded-by-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_01_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.00-emphasis-with-emphasis-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_01_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.01-emphasis-surrounded-by-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_02_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.00-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_02_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.01-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_02_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.02-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_03_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.03.00-emphasis-surrounded-by-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_02_04_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.04.00-emphasis-closed-with-strong-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_00_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.00-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_00_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.01-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_00_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.02-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_00_03_LexerInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.03-bad-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_00_04_LexerInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.04-bad-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_01_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.00-literal-with-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_01_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.01-literal-with-middle-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_01_02_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.02-literal-with-end-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_02_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.02.00-literal-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_03_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.00-literal-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_03_01_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.01-literal-quoted-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_03_02_LexerInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.02-bad-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_04_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.04.00-literal-interpreted-text")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_05_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.05.00-literal-followed-by-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_03_06_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.06.00-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_00_00_LexerInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.04.00.00-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_00_01_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.01-bad-phrase-ref-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_00_02_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.02-bad-phrase-ref-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.01.00-ref-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.02.00-ref-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_03_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.03.00-ref-anon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_04_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.04.00-ref-anon-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_05_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.05.00-ref-anon-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_06_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.06.00-ref-with-anon-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_07_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.07.00-phrase-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_07_01_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.07.01-bad-phrase-ref-missing-backtick")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_08_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.08.00-phrase-ref-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_09_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.09.00-phrase-ref-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_09_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.09.01-phrase-ref-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_10_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.10.00-phrase-ref-anon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_10_01_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.10.01-bad-phrase-ref-anon-missing-backtick")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_11_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.11.00-phrase-ref-anon-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_12_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.12.00-phrase-ref-anon-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_12_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.12.01-phrase-ref-anon-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_13_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.13.00-phrase-ref-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_04_14_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.14.00-phrase-ref-literal-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.00.00-phrase-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.01.00-anon-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.00-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_02_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.01-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_02_02_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.02-across-lines-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_02_03_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.03-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_02_04_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.04-lots-of-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_03_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.03.00-relative-no-text")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_05_04_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.04.00-escaped-low-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_06_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.00.00-alias-phrase-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_06_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.01.00-alias-anon-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_06_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.00-alias-multi-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_06_02_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.01-alias-multi-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_06_02_02_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.02-alias-multi-line-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_06_02_03_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.03-alias-lots-of-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_07_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.00-inline-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_07_00_01_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.01-bad-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_07_00_02_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.02-bad-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_07_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.01.00-inline-target-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_07_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.02.00-inline-target-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_07_03_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.03.01-inline-target-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_08_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.00.00-footnote-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_08_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.01.00-footnote-ref-auto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_08_01_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.01.01-footnote-ref-auto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_08_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.02.00-footnote-ref-auto-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_08_03_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.03.00-footnote-ref-adjacent-refs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_09_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.00.00-citation-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_09_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.01.00-citation-ref-multiple")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_09_01_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.01.01-citation-ref-adjacent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_10_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.00-subs-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_10_00_01_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.01-subs-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_10_00_02_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.02-bad-subs-ref-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_10_00_03_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.03-bad-subs-ref-is-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_10_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.01.00-subs-ref-multiple")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_10_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.02.00-subs-ref-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_11_00_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.00-standalone-hyperlink")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_11_00_01_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.01-bad-invalid-hyperlinks")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_11_00_02_LexerInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.02-bad-escaped-email-addresses")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_11_01_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.01.00-urls-with-escaped-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_11_02_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.02.00-urls-in-angle-brackets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_06_11_03_00_LexerInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.03.00-urls-with-interesting-endings")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bullet-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bullet-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bullet-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bullet-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.01-bullet-list-with-two-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.01-bullet-list-with-two-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.01-bullet-list-with-two-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.01-bullet-list-with-two-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_02_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.02-bullet-list-noblankline-between-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_02_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.02-bullet-list-noblankline-between-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_02_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.02-bullet-list-noblankline-between-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_02_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.02-bullet-list-noblankline-between-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_03_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bad-bullet-list-noblankline-at-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_03_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bad-bullet-list-noblankline-at-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_03_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bad-bullet-list-noblankline-at-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_00_03_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bad-bullet-list-noblankline-at-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.01-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.01-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.01-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_01_01_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.01-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_02_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.00-bullet-list-different-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_02_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.00-bullet-list-different-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_02_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.00-bullet-list-different-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_02_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.00-bullet-list-different-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_02_01_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.01-bad-bullet-list-different-bullets-missing-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_02_01_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.01-bad-bullet-list-different-bullets-missing-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_03_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.00-bullet-list-empty-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_03_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.00-bullet-list-empty-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_03_01_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.01-bad-bullet-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_03_01_LexerListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.01-bad-bullet-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_04_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.04.00-bullet-list-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_07_00_04_00_LexerListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.04.00-bullet-list-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.00-numbered")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_01_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.01-numbered-noblanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_02_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.02-numbered-indented-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_03_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.03-bad-enum-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_03_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.03-bad-enum-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_04_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.04-bad-enum-list-scrambled-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_04_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.04-bad-enum-list-scrambled-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_05_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.05-bad-enum-list-skipped-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_05_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.05-bad-enum-list-skipped-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_06_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.06-bad-enum-list-not-ordinal-1")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_00_06_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.06-bad-enum-list-not-ordinal-1")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_01_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.00-alphabetical-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_01_01_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.01-bad-alphabetical-list-without-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_01_01_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.01-bad-alphabetical-list-without-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_01_02_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.02-alphabetical-list-nbsp-workaround")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_02_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.00-items-with-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_02_01_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.01-bad-enum-list-unexpected-unindent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_02_01_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.01-bad-enum-list-unexpected-unindent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_03_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.03.00-diff-formats")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_04_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.04.00-roman-numerals")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_04_01_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.04.01-bad-enum-list-bad-roman-numerals")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_05_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.00-nested")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_06_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.00-sequence-types")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_06_01_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.01-ambiguous-sequence-types")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_06_02_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.02-bad-enum-list-ambiguous")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_07_00_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.00-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_07_01_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.01-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_07_02_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.02-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_07_03_LexerListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.03-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_07_04_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.04-bad-enum-list-auto-numbering-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_08_00_08_00_LexerListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.08.00-bad-enum-list-paragraph-not-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_00_01_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.01-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_00_02_LexerListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.02-bad-def-list-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_00_03_LexerListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.03-bad-def-list-not-def-term")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_01_00_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.00-two-terms")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_01_01_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.01-two-terms-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_01_02_LexerListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.02-bad-def-list-noblankline-after-two-terms")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_02_00_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.02.00-nested-terms")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_03_00_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.03.00-term-with-classifier")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_04_00_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.00-term-not-classifier")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_04_01_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.01-term-not-classifier-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_04_02_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.02-two-classifiers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_05_00_LexerListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.05.00-not-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_09_00_06_00_LexerListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.06.00-bad-def-list-with-inline-markup-errors")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_10_00_00_00_LexerListOptionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("10.00.00.00-three-short-options")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

func Test_11_00_00_00_LexerListOptionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("11.00.00.00-three-short-options")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItemData, items)
}

