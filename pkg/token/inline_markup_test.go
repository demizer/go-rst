package token

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_06_00_00_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.00.00-double-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_01_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.01.00-lots-of-escaping")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_02_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.02.00-lots-of-escaping-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_03_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.00-emphasis-wrapped-in-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_03_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.01-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_03_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.02-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_04_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.00-openers-and-closers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_04_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.01-strong-and-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_05_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.05.00-emphasis-with-backwards-rule-5")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_00_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.00-strong")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_00_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.01-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_00_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.02-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_01_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.01.00-strong-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_02_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.02.00-strong-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_03_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.00-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_03_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.01-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_03_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.02-strong-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_01_04_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.04.00-strong-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.00-simple-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.01-single-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.02-emphasis-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_03_LexInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.03-bad-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_04_LexInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.04-bad-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_05_LexInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.05-bad-emphasis-unclosed-surrounded-by-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_01_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.00-emphasis-with-emphasis-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_01_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.01-emphasis-surrounded-by-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_02_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.00-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_02_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.01-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_02_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.02-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_03_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.03.00-emphasis-surrounded-by-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_04_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.04.00-emphasis-closed-with-strong-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_00_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.00-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_00_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.01-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_00_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.02-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_00_03_LexInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.03-bad-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_00_04_LexInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.00.04-bad-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_01_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.00-literal-with-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_01_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.01-literal-with-middle-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_01_02_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.01.02-literal-with-end-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_02_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.02.00-literal-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_03_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.00-literal-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_03_01_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.01-literal-quoted-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_03_02_LexInlineMarkupBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.03.02-bad-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_04_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.04.00-literal-interpreted-text")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_05_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.05.00-literal-followed-by-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_03_06_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.03.06.00-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_00_00_LexInlineMarkupGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.04.00.00-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_00_01_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.01-bad-phrase-ref-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_00_02_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.00.02-bad-phrase-ref-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.01.00-ref-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.02.00-ref-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_03_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.03.00-ref-anon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_04_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.04.00-ref-anon-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_05_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.05.00-ref-anon-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_06_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.06.00-ref-with-anon-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_07_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.07.00-phrase-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_07_01_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.07.01-bad-phrase-ref-missing-backtick")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_08_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.08.00-phrase-ref-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_09_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.09.00-phrase-ref-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_09_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.09.01-phrase-ref-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_10_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.10.00-phrase-ref-anon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_10_01_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.10.01-bad-phrase-ref-anon-missing-backtick")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_11_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.11.00-phrase-ref-anon-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_12_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.12.00-phrase-ref-anon-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_12_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.12.01-phrase-ref-anon-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_13_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.13.00-phrase-ref-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_04_14_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.04.14.00-phrase-ref-literal-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.00.00-phrase-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.01.00-anon-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.00-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_02_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.01-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_02_02_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.02-across-lines-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_02_03_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.03-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_02_04_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.02.04-lots-of-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_03_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.03.00-relative-no-text")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_05_04_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.05.04.00-escaped-low-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_06_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.00.00-alias-phrase-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_06_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.01.00-alias-anon-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_06_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.00-alias-multi-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_06_02_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.01-alias-multi-line")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_06_02_02_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.02-alias-multi-line-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_06_02_03_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.06.02.03-alias-lots-of-whitespace")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_07_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.00-inline-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_07_00_01_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.01-bad-invalid")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_07_00_02_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.00.02-bad-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_07_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.01.00-inline-target-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_07_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.02.00-inline-target-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_07_03_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.07.03.01-inline-target-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_08_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.00.00-footnote-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_08_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.01.00-footnote-ref-auto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_08_01_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.01.01-footnote-ref-auto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_08_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.02.00-footnote-ref-auto-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_08_03_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.08.03.00-footnote-ref-adjacent-refs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_09_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.00.00-citation-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_09_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.01.00-citation-ref-multiple")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_09_01_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.09.01.01-citation-ref-adjacent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_10_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.00-subs-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_10_00_01_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.01-subs-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_10_00_02_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.02-bad-subs-ref-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_10_00_03_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.00.03-bad-subs-ref-is-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_10_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.01.00-subs-ref-multiple")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_10_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.10.02.00-subs-ref-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_11_00_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.00-standalone-hyperlink")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_11_00_01_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.01-bad-invalid-hyperlinks")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_11_00_02_LexInlineMarkupBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.00.02-bad-escaped-email-addresses")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_11_01_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.01.00-urls-with-escaped-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_11_02_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.02.00-urls-in-angle-brackets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_11_03_00_LexInlineMarkupGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("06.11.03.00-urls-with-interesting-endings")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
