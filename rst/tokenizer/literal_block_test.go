package tokenizer

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/rst/testutil"
)

func Test_05_00_00_00_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_01_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-literal-block-space-after-colons")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_02_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-double-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_03_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.03-multiline-paragraph-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_04_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.04-literal-block-and-escaped-colon-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_05_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.05-literal-block-and-escaped-colon-blockquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_06_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.06-paragraph-space-double-colon-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_07_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.07-paragraph-colon-newline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_08_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.08-paragraph-colon-newline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_01_00_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.00-wonky-multiline-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_02_00_LexInlineReferenceGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-bad-eof-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_02_01_LexInlineReferenceGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.01-quoted-literal-block-two-blanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_02_02_LexInlineReferenceGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.02-quoted-literal-block-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_02_00_LexLiteralBlockGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
func Test_05_00_00_00_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.00-bad-no-blankline-before-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_01_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.01-bad-no-blankline-after-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_00_02_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.00.02-bad-unindented-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_01_00_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.01.00-bad-section-underline-not-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_02_00_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.02.00-bad-eof-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_03_00_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.03.00-bad-indented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_04_00_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.04.00-bad-unindented-line-after-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_05_00_05_00_LexLiteralBlockBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("05.00.05.00-bad-inconsistent-quoted-literal-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
