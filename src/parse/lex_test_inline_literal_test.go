package parse

import "testing"

// A single inline literal without anything else.
func Test_02_03_00_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.00.00-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal in the middle of a paragraph.
func Test_02_03_00_01_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.00.01-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal at the beginning of a paragraph.
func Test_02_03_00_02_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.00.02-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed inline literal
func Test_02_03_00_03_LexInlineLiteralBad(t *testing.T) {
	testPath := testPathFromName("02.03.00.03-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed literal across lines
func Test_02_03_00_04_LexInlineLiteralBad(t *testing.T) {
	testPath := testPathFromName("02.03.00.04-literal-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal beginning with a backslash
func Test_02_03_01_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.01.00-literal-with-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal with backlash in the middle
func Test_02_03_01_01_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.01.01-literal-with-middle-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal with backslash at the end
func Test_02_03_01_02_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.01.02-literal-with-end-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal surrounded by apostrophes and unicode literal apostrophes
func Test_02_03_02_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.02.00-literal-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal surrounded by unicode literals
func Test_02_03_03_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.03.00-literal-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal containing apostrophes
func Test_02_03_03_01_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.03.01-literal-quoted-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclose literal with tex quotes
func Test_02_03_03_02_LexInlineLiteralBad(t *testing.T) {
	testPath := testPathFromName("02.03.03.02-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal containing backticks
func Test_02_03_04_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.04.00-literal-interpreted-text")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal followed by '\s'
func Test_02_03_05_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.05.00-literal-followed-by-backslash")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Inline literal containing tex quotes
func Test_02_03_06_00_LexInlineLiteralGood(t *testing.T) {
	testPath := testPathFromName("02.03.06.00-literal-with-tex-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
