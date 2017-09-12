package parser

//
// AUTO-GENERATED using tools/gentests.go
//

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_00_00_00_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.00-comment")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_00_01_ParserCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.01-bad-comment-no-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_00_02_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.02-comment-with-literal-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_00_03_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.03-comment-not-reference")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_01_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.00-comment-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_01_01_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.01-comment-block-second-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_02_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.00-newline-after-comment-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_02_01_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.01-newline-after-comment-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_02_02_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.02-comment-not-citation")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_02_03_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.03-comment-not-subs-def")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_03_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.03.00-empty-comment-with-blockquote")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_04_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.00-comment-in-definition")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_04_01_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.01-comment-after-definition")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_05_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.00-comment-between-bullet-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_05_01_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.01-comment-between-bullets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_05_02_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.02-comment-trailing-bullet")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_06_00_ParserCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.00-two-comments")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_00_00_06_01_ParserCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.01-bad-two-comments-no-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_00_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.00-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_00_01_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.01-optional-space-before-colon")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_00_02_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.02-bad-target-missing-backquote")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_00_03_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.03-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_00_04_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.04-bad-target-malformed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_01_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.01.00-long-target-names")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_02_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.02.00-target-beginning-with-underscore")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_02_01_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.02.01-bad-beginning-with-underscore")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_03_00_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.00-bad-duplicate-implicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_03_01_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.01-bad-duplicate-implicit-explicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_03_02_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.02-bad-duplicate-implicit-directive-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_04_00_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.00-bad-duplicate-explicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_04_01_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.01-bad-duplicate-explicit-directive-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_04_02_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.02-bad-duplicate-implicit-explicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_05_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.05.00-escaped-colon-at-the-end")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_00_05_01_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.05.01-bad-unescaped-colon-at-the-end")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_00_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.00.00-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_00_01_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.00.01-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_01_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.01.00-external-target-mailto")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_02_00_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.02.00-bad-duplicate-external-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_02_01_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.02.01-bad-duplicate-external-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_03_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.03.00-anonymous-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_03_01_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.03.01-anonymous-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_01_03_02_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.01.03.02-anonymous-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_02_00_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.00.00-indirect-hyperlink-targets-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_02_00_01_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.00.01-indirect-hyperlink-targets-phrase-references")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_02_01_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.01.00-anonymous-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_02_01_01_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.01.01-anonymous-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_02_02_00_ParserReferenceHyperlinkTargetsBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.02.00-bad-anon-and-named-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_01_02_03_00_ParserReferenceHyperlinkTargetsGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.02.03.00-anonymous-indirect-target-multiline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_02_00_00_00_ParserParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.00-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_02_00_00_01_ParserParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.01-with-line-break")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_02_00_00_02_ParserParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.02-three-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_02_00_01_00_ParserParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.00-two-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_02_00_01_01_ParserParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.01-two-paragraphs-three-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_00_00_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.00-paragraph-blockquote")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_00_01_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.01-paragraph-blockquote-short-section")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_00_02_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.02-paragraph-blockquote-comment")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_00_03_ParserBlockquoteBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.03-bad-no-blank-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_00_04_ParserBlockquoteBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.00.04-bad-unexpected-indent")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_01_00_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.01.00-two-levels")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_02_00_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.02.00-unicode-em-dash")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_03_00_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.03.00-uneven-indents")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_04_00_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.00-paragraph-blockquote-attrib")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_04_01_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.01-paragraph-blockquote-two-line-attrib")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_04_02_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.02-paragraph-blockquote-attrib-no-space")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_04_03_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.03-paragraph-blockquote-one-attrib")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_04_04_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.04-paragraph-blockquote-attrib-invalid")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_03_00_04_05_ParserBlockquoteGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("03.00.04.05-paragraph-blockquote-attrib-with-invalid-attrib")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_00_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.00-title-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_00_01_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.01-paragraph-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_00_02_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.02-title-combining-chars")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_00_03_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.03-bad-short-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_00_04_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.04-bad-short-title-short-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_01_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.01.00-paragraph-head-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_02_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.02.00-short-title")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_03_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.03.00-empty-section")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_04_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.04.00-numbered-title")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_04_01_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.04.01-bad-enum-list-with-numbered-title")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_00_05_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.05.00-title-with-imu")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_00_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.00-title-overline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_00_01_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.01-bad-title-too-long")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_00_02_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.02-bad-short-title-short-overline-and-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_00_03_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.03-bad-short-title-short-overline-missing-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_01_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.00-inset-title-with-overline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_01_01_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.01-bad-inset-title-missing-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_01_02_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.02-bad-inset-title-mismatched-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_01_03_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.03-bad-inset-title-missing-underline-with-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_01_04_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.04-bad-inset-title-missing-underline-and-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_02_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.02.00-three-char-section-title")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_03_00_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.03.00-bad-unexpected-titles")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_04_00_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.04.00-bad-missing-titles-with-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_04_01_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.04.01-bad-missing-titles-with-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_05_00_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.05.00-bad-incomplete-section")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_05_01_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.05.01-bad-incomplete-sections-no-title")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_06_00_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.06.00-bad-indented-title-short-overline-and-underline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_01_07_00_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.07.00-bad-two-char-section-title")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_00_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.00-section-level-return")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_00_01_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.01-section-level-return")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_00_02_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.02-section-level-return")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_00_03_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.03-bad-subsection-order")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_01_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.00-section-level-return")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_01_01_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.01-bad-two-level-overline-bad-return")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_01_02_ParserSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.02-bad-subsection-order-with-overlines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_04_02_02_00_ParserSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.02.00-two-level-one-overline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_00_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_01_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-literal-block-space-after-colons")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_02_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-bad-unindented-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_03_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.03-bad-no-blankline-after-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_04_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.04-multiline-paragraph-before-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_05_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.05-bad-no-blankline-before-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_06_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.06-paragraph-space-double-colon-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_07_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.07-paragraph-colon-newline-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_08_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.08-bad-section-underline-not-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_00_09_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.09-bad-eof-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_01_00_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.00-multiline-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_01_01_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.01-wonky-multiline-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_02_00_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-double-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_00_02_01_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.01-literal-block-and-escaped-colon-blockquote")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_01_00_00_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.00-quoted-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_01_00_01_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.01-quoted-literal-block-two-blanklines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_01_00_02_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.00.02-bad-inconsistent-quoted-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_01_01_00_ParserLiteralBlockGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.00-quoted-literal-block-multiline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_01_01_01_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.01-bad-indented-line-after-quoted-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_05_01_01_02_ParserLiteralBlockBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.01.01.02-bad-unindented-line-after-quoted-literal-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_00_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.00.00-double-underscore")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_01_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.01.00-lots-of-escaping")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_02_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.02.00-lots-of-escaping-unicode")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_03_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.00-emphasis-wrapped-in-unicode")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_03_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.00.03.01-emphasis-with-unicode-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_03_02_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.00.03.02-emphasis-with-unicode-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_04_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.00-openers-and-closers")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_04_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.01-strong-and-kwargs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_00_05_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.05.00-emphasis-with-backwards-rule-5")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_00_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.00-strong")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_00_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.01.00.01-strong-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_00_02_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.01.00.02-strong-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_01_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.01.00-strong-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_02_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.02.00-strong-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_03_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.00-strong-asterisk")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_03_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.01-strong-asterisk")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_03_02_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.01.03.02-strong-kwargs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_01_04_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.04.00-strong-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_00_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.00-simple-emphasis")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_00_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.01-single-emphasis")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_00_02_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.02-emphasis-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_00_03_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.02.00.03-bad-emphasis-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_00_04_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.02.00.04-bad-emphasis-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_00_05_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.02.00.05-bad-emphasis-unclosed-surrounded-by-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_01_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.00-emphasis-with-emphasis-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_01_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.01-emphasis-surrounded-by-quotes")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_02_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.00-emphasis-with-asterisk")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_02_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.01-emphasis-with-asterisk")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_02_02_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.02-emphasis-with-asterisk")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_03_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.03.00-emphasis-surrounded-by-markup")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_02_04_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.04.00-emphasis-closed-with-strong-markup")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_00_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.00-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_00_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.01-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_00_02_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.02-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_00_03_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.03.00.03-bad-literal-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_00_04_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.03.00.04-bad-literal-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_01_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.00-literal-with-backslash")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_01_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.01-literal-with-middle-backslash")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_01_02_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.02-literal-with-end-backslash")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_02_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.02.00-literal-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_03_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.00-literal-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_03_01_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.01-literal-quoted-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_03_02_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.03.03.02-bad-literal-with-tex-quotes")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_04_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.04.00-literal-interpreted-text")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_05_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.05.00-literal-followed-by-backslash")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_03_06_00_ParserInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.06.00-literal-with-tex-quotes")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.00-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_00_01_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.01-bad-phrase-ref-invalid")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_00_02_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.02-bad-phrase-ref-invalid")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.01.00-ref-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.02.00-ref-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_03_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.03.00-ref-anon")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_04_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.04.00-ref-anon-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_05_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.05.00-ref-anon-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_06_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.06.00-ref-with-anon-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_07_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.07.00-phrase-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_07_01_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.07.01-bad-phrase-ref-missing-backtick")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_08_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.08.00-phrase-ref-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_09_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.09.00-phrase-ref-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_09_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.09.01-phrase-ref-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_10_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.10.00-phrase-ref-anon")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_10_01_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.10.01-bad-phrase-ref-anon-missing-backtick")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_11_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.11.00-phrase-ref-anon-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_12_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.12.00-phrase-ref-anon-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_12_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.12.01-phrase-ref-anon-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_13_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.13.00-phrase-ref-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_04_14_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.14.00-phrase-ref-literal-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.00.00-phrase-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.01.00-anon-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.00-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_02_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.01-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_02_02_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.02-across-lines-whitespace")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_02_03_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.03-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_02_04_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.04-lots-of-whitespace")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_03_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.03.00-relative-no-text")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_05_04_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.04.00-escaped-low-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_06_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.00.00-alias-phrase-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_06_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.01.00-alias-anon-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_06_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.00-alias-multi-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_06_02_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.01-alias-multi-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_06_02_02_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.02-alias-multi-line-whitespace")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_06_02_03_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.03-alias-lots-of-whitespace")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_07_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.00-inline-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_07_00_01_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.01-bad-invalid")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_07_00_02_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.02-bad-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_07_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.01.00-inline-target-with-apostrophe")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_07_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.02.00-inline-target-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_07_03_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.03.01-inline-target-quoted")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_08_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.00.00-footnote-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_08_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.01.00-footnote-ref-auto")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_08_01_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.01.01-footnote-ref-auto")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_08_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.02.00-footnote-ref-auto-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_08_03_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.03.00-footnote-ref-adjacent-refs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_09_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.00.00-citation-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_09_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.01.00-citation-ref-multiple")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_09_01_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.01.01-citation-ref-adjacent")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_10_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.00-subs-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_10_00_01_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.01-subs-ref")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_10_00_02_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.02-bad-subs-ref-unclosed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_10_00_03_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.03-bad-subs-ref-is-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_10_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.01.00-subs-ref-multiple")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_10_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.02.00-subs-ref-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_11_00_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.00-standalone-hyperlink")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_11_00_01_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.01-bad-invalid-hyperlinks")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_11_00_02_ParserInlineMarkupBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.02-bad-escaped-email-addresses")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_11_01_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.01.00-urls-with-escaped-markup")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_11_02_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.02.00-urls-in-angle-brackets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_06_11_03_00_ParserInlineMarkupGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.03.00-urls-with-interesting-endings")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_00_00_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bullet-list")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_00_01_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.01-bullet-list-with-two-items")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_00_02_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.02-bullet-list-noblankline-between-items")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_00_03_ParserListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bad-bullet-list-noblankline-at-end")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_01_00_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-item-with-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_01_01_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.01-bullet-list-item-with-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_02_00_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.00-bullet-list-different-bullets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_02_01_ParserListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.01-bad-bullet-list-different-bullets-missing-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_03_00_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.00-bullet-list-empty-item")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_03_01_ParserListBulletBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.01-bad-bullet-list-empty-item-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_07_00_04_00_ParserListBulletGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.04.00-bullet-list-unicode")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.00-numbered")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_01_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.01-numbered-noblanklines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_02_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.02-numbered-indented-items")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_03_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.03-bad-enum-list-empty-item-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_04_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.04-bad-enum-list-scrambled-items")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_05_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.05-bad-enum-list-skipped-item")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_00_06_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.06-bad-enum-list-not-ordinal-1")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_01_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.00-alphabetical-list")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_01_01_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.01-bad-alphabetical-list-without-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_01_02_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.02-alphabetical-list-nbsp-workaround")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_02_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.00-items-with-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_02_01_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.01-bad-enum-list-unexpected-unindent")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_03_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.03.00-diff-formats")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_04_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.04.00-roman-numerals")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_04_01_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.04.01-bad-enum-list-bad-roman-numerals")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_05_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.00-nested")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_06_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.00-sequence-types")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_06_01_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.01-ambiguous-sequence-types")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_06_02_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.02-bad-enum-list-ambiguous")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_07_00_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.00-auto-numbering")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_07_01_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.01-auto-numbering")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_07_02_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.02-auto-numbering")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_07_03_ParserListEnumeratedGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.03-auto-numbering")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_07_04_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.04-bad-enum-list-auto-numbering-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_08_00_08_00_ParserListEnumeratedBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.08.00-bad-enum-list-paragraph-not-list")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_00_01_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.01-with-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_00_02_ParserListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.02-bad-def-list-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_00_03_ParserListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.03-bad-def-list-not-def-term")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_01_00_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.00-two-terms")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_01_01_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.01-two-terms-noblankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_01_02_ParserListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.02-bad-def-list-noblankline-after-two-terms")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_02_00_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.02.00-nested-terms")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_03_00_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.03.00-term-with-classifier")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_04_00_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.00-term-not-classifier")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_04_01_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.01-term-not-classifier-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_04_02_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.02-two-classifiers")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_05_00_ParserListDefinitionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.05.00-not-literal")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_09_00_06_00_ParserListDefinitionBad(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.06.00-bad-def-list-with-inline-markup-errors")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

func Test_10_00_00_00_ParserListOptionGood(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("10.00.00.00-three-short-options")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
}

