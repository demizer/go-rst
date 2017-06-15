package token

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_01_00_00_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.00-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_00_01_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.01-optional-space-before-colon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_00_02_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.02-bad-target-missing-backquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_00_03_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.03-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_00_04_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.00.04-bad-target-malformed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_01_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.01.00-long-target-names")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_02_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.02.00-target-beginning-with-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_02_01_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.02.01-bad-beginning-with-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_03_00_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.03.00-bad-duplicate-implicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_03_01_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.03.01-bad-duplicate-implicit-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_03_02_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.03.02-bad-duplicate-implicit-directive-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_04_00_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.04.00-bad-duplicate-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_04_01_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.04.01-bad-duplicate-explicit-directive-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_04_02_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.04.02-bad-duplicate-implicit-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_05_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.05.00-escaped-colon-at-the-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_00_05_01_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.00.05.01-bad-unescaped-colon-at-the-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_00_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.00.00-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_00_01_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.00.01-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_01_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.01.00-external-target-mailto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_02_00_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.02.00-bad-duplicate-external-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_02_01_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.02.01-bad-duplicate-external-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_03_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.03.00-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_03_01_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.03.01-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_01_03_02_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.01.03.02-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_02_00_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.00.00-indirect-hyperlink-targets-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_02_00_01_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.00.01-indirect-hyperlink-targets-phrase-references")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_02_01_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.01.00-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_02_01_01_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.01.01-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_02_02_00_LexHyperlinkTargetBad(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.02.00-bad-anon-and-named-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_01_02_03_00_LexHyperlinkTargetGood(t *testing.T) {
	testPath := testutil.TestPathFromName("01.02.03.00-anonymous-indirect-target-multiline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
