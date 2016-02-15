package parse

import "testing"

// A single inline literal without anything else.
func TestLexInlineLiteralGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal in the middle of a paragraph.
func TestLexInlineLiteralGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal at the beginning of a paragraph.
func TestLexInlineLiteralGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal beginning with a backslash
func TestLexInlineLiteralGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-literal-with-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal with backlash in the middle
func TestLexInlineLiteralGood0101(t *testing.T) {
	testPath := testPathFromName("01.01-literal-with-middle-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal with backslash at the end
func TestLexInlineLiteralGood0102(t *testing.T) {
	testPath := testPathFromName("01.02-literal-with-end-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal surrounded by apostrophes and unicode literal apostrophes
func TestLexInlineLiteralGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-literal-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal surrounded by unicode literals
func TestLexInlineLiteralGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-literal-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal containing apostrophes
func TestLexInlineLiteralGood0301(t *testing.T) {
	testPath := testPathFromName("03.01-literal-quoted-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal containing backticks
func TestLexInlineLiteralGood0400(t *testing.T) {
	testPath := testPathFromName("04.00-literal-interpreted-text")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal followed by '\s'
func TestLexInlineLiteralGood0500(t *testing.T) {
	testPath := testPathFromName("05.00-literal-followed-by-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal containing tex quotes
func TestLexInlineLiteralGood0600(t *testing.T) {
	testPath := testPathFromName("06.00-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed inline literal
func TestLexInlineLiteralBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed literal across lines
func TestLexInlineLiteralBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclose literal with tex quotes
func TestLexInlineLiteralBad0100(t *testing.T) {
	testPath := testPathFromName("01.00-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
