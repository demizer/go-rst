// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import "testing"

func TestLexParagraphGood0000(t *testing.T) {
	// A simple one line paragraph
	testPath := testPathFromName("00.00-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexParagraphWithLineBreakGood0001(t *testing.T) {
	// A simple two line paragraph
	testPath := testPathFromName("00.01-para-with-line-break")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexParagraphWithThreeLinesGood0002(t *testing.T) {
	// A simple three line paragraph
	testPath := testPathFromName("00.02-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexTwoSeparateParagraphs0100(t *testing.T) {
	// Two paragraphs separated by a blank line
	testPath := testPathFromName("01.00-two-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexTwoSeparateThreeLineParagraphs0101(t *testing.T) {
	// Two three line paragraphs separated by a blank line
	testPath := testPathFromName("01.01-two-para-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
