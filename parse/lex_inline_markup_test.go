// go-rst - A reStructuredText parser for Go
// 2015 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import "testing"

// Tests double underscore recognition
func TestLexInlineMarkupRecognitionRulesGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-double-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Tests double underscore recognition
func TestLexInlineMarkupRecognitionRulesEscapingGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-lots-of-escaping")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Tests escaping with unicode literals
func TestLexInlineMarkupRecognitionRulesEscapingUnicodeLiteralsGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-lots-of-escaping-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexInlineMarkupRecognitionEmphasisWrappedWithUnicodeLiteralsGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-emphasis-wrapped-in-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexInlineMarkupRecognitionOpenersAndClosersGood0400(t *testing.T) {
	testPath := testPathFromName("04.00-openers-and-closers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A paragraph containing a single emphasized word
func TestLexSingleEmphasisGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-simple-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A document only containing an emphasized word
func TestLexSingleEmphasisGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-single-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasis across two lines
func TestLexEmphasisAcrossLinesGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-emphasis-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasis surrounded by apostrophe
func TestLexEmphasisSurrondedByApostropheGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-emphasis-with-emphasis-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasis surrounded by quotes from many languages
func TestLexEmphasisSurrondedByQuotesGood0101(t *testing.T) {
	testPath := testPathFromName("01.01-emphasis-surrounded-by-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasized asterisk
func TestLexEmphasizedAsteriskGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasized asterisk
func TestLexEmphasizedAsteriskGood0201(t *testing.T) {
	testPath := testPathFromName("02.01-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasized asterisk
func TestLexEmphasizedAsteriskGood0202(t *testing.T) {
	testPath := testPathFromName("02.02-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
