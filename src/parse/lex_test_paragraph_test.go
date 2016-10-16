package parse

import "testing"

// A simple one line paragraph
func Test_01_00_00_00_LexParagraphGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.00-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A simple two line paragraph
func Test_01_00_00_01_LexParagraphWithLineBreakGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.01-with-line-break")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A simple three line paragraph
func Test_01_00_00_02_LexParagraphWithThreeLinesGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.02-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Two paragraphs separated by a blank line
func Test_01_00_01_00_LexTwoSeparateParagraphs(t *testing.T) {
	testPath := testPathFromName("01.00.01.00-two-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Two three line paragraphs separated by a blank line
func Test_01_00_01_01_LexTwoSeparateThreeLineParagraphs(t *testing.T) {
	testPath := testPathFromName("01.00.01.01-two-paragraphs-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
