package parse

import "testing"

// A simple one line paragraph
func TestLexParagraphGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A simple two line paragraph
func TestLexParagraphWithLineBreakGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-para-with-line-break")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A simple three line paragraph
func TestLexParagraphWithThreeLinesGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Two paragraphs separated by a blank line
func TestLexTwoSeparateParagraphs0100(t *testing.T) {
	testPath := testPathFromName("01.00-two-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Two three line paragraphs separated by a blank line
func TestLexTwoSeparateThreeLineParagraphs0101(t *testing.T) {
	testPath := testPathFromName("01.01-two-para-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
