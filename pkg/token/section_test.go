package token

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_04_00_00_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.00-title-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_00_01_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.01-paragraph-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_00_02_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.02-title-combining-chars")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_00_03_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.03-bad-short-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_00_04_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.00.04-bad-short-title-short-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_01_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.01.00-paragraph-head-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_02_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.02.00-short-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_03_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.03.00-empty-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_04_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.04.00-numbered-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_04_01_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.04.01-bad-enum-list-with-numbered-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_00_05_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.00.05.00-title-with-imu")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_00_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.00-title-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_00_01_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.01-bad-title-too-long")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_00_02_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.02-bad-short-title-short-overline-and-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_00_03_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.00.03-bad-short-title-short-overline-missing-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_01_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.00-inset-title-with-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_01_01_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.01-bad-inset-title-missing-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_01_02_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.02-bad-inset-title-mismatched-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_01_03_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.03-bad-inset-title-missing-underline-with-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_01_04_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.01.04-bad-inset-title-missing-underline-and-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_02_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.02.00-three-char-section-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_03_00_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.03.00-bad-unexpected-titles")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_04_00_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.04.00-bad-missing-titles-with-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_04_01_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.04.01-bad-missing-titles-with-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_05_00_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.05.00-bad-incomplete-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_05_01_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.05.01-bad-incomplete-sections-no-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_06_00_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.06.00-bad-indented-title-short-overline-and-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_01_07_00_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.01.07.00-bad-two-char-section-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_00_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.00-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_00_01_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.01-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_00_02_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.02-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_00_03_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.00.03-bad-subsection-order")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_01_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.00-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_01_01_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.01-bad-two-level-overline-bad-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_01_02_LexSectionBad(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.01.02-bad-subsection-order-with-overlines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_04_02_02_00_LexSectionGood(t *testing.T) {
	testPath := testutil.TestPathFromName("04.02.02.00-two-level-one-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
