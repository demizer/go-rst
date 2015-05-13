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

func TestLexInlineMarkupRecognitionRulesGood0401(t *testing.T) {
	testPath := testPathFromName("04.01-strong-and-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexInlineMarkupRecognitionRulesGood5000(t *testing.T) {
	testPath := testPathFromName("05.00-emphasis-with-backwards-rule-5")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Bad emphasis start string with unicode literal space after start string. The
// text that is suppossed to close the emphasis is lexed as an emphasis start
// string without a corresponding end string.
func TestLexInlineMarkupRecognitionOpenersAndClosersBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Bad emphasis start string with unicode literal space after start string. The
// text that is suppossed to close the emphasis is lexed as an emphasis start
// string without a corresponding end string. This test includes more text
// around the incorrect emphasis.
func TestLexInlineMarkupRecognitionOpenersAndClosersBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-emphasis-with-unicode-literal")
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

func TestLexEmphasisGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-emphasis-surrounded-by-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexEmphasisGood0400(t *testing.T) {
	testPath := testPathFromName("04.00-emphasis-closed-with-strong-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed emphasis
func TestLexEmphasisUnclosedBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed emphasis on two lines with paragraphs and stuff.
func TestLexEmphasisUnclosedBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexEmphasisUnclosedBad0100(t *testing.T) {
	testPath := testPathFromName("01.00-emphasis-unclosed-surrounded-by-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Single strong
func TestLexStrongGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-strong")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Single strong with apostrophes
func TestLexStrongGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-strong-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with unicode literal quotes
func TestLexStrongGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-strong-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with asterisk and double asterisk
func TestLexStrongGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with asterisk and double asterisk and strong asterisk gone to plaid!
func TestLexStrongGood0301(t *testing.T) {
	testPath := testPathFromName("03.01-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with bad kwargs
func TestLexStrongBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-strong-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with missing closing markup
func TestLexStrongBad0100(t *testing.T) {
	testPath := testPathFromName("01.00-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with missing closing markup across three lines
func TestLexStrongBad0101(t *testing.T) {
	testPath := testPathFromName("01.01-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
