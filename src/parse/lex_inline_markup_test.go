package parse

import "testing"

// Tests double underscore recognition
func Test_02_00_00_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.00.00-imrr-good-double-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Tests double underscore recognition
func Test_02_00_01_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.01.00-imrr-good-lots-of-escaping")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Tests escaping with unicode literals
func Test_02_00_02_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.02.00-imrr-good-lots-of-escaping-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_00_03_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.03.00-imrr-good-emphasis-wrapped-in-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Bad emphasis start string with unicode literal space after start string. The text that is suppossed to close the emphasis
// is lexed as an emphasis start string without a corresponding end string.
func Test_02_00_03_01_LexInlineMarkupRecognitionRulesBad(t *testing.T) {
	testPath := testPathFromName("02.00.03.01-imrr-bad-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Bad emphasis start string with unicode literal space after start string. The text that is suppossed to close the emphasis
// is lexed as an emphasis start string without a corresponding end string. This test includes more text around the incorrect
// emphasis.
func Test_02_00_03_02_LexInlineMarkupRecognitionRulesBad(t *testing.T) {
	testPath := testPathFromName("02.00.03.02-imrr-bad-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_00_04_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.04.00-imrr-good-openers-and-closers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_00_04_01_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.04.01-imrr-good-strong-and-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_00_05_00_LexInlineMarkupRecognitionRulesGood5(t *testing.T) {
	testPath := testPathFromName("02.00.05.00-imrr-good-emphasis-with-backwards-rule-5")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Single strong
func Test_02_01_00_00_LexStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.00.00-strong-good-strong")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with missing closing markup
func Test_02_01_00_01_LexStrongBad(t *testing.T) {
	testPath := testPathFromName("02.01.00.01-strong-bad-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with missing closing markup across three lines
func Test_02_01_00_02_LexStrongBad(t *testing.T) {
	testPath := testPathFromName("02.01.00.02-strong-bad-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Single strong with apostrophes
func Test_02_01_01_00_LexStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.01.00-strong-good-strong-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with unicode literal quotes
func Test_02_01_02_00_LexStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.02.00-strong-good-strong-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with asterisk and double asterisk
func Test_02_01_03_00_LexStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.03.00-strong-good-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with asterisk and double asterisk and strong asterisk gone to plaid!
func Test_02_01_03_01_LexStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.03.01-strong-good-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Strong with bad kwargs
func Test_02_01_03_02_LexStrongBad(t *testing.T) {
	testPath := testPathFromName("02.01.03.02-strong-bad-strong-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A paragraph containing a single emphasized word
func Test_02_02_00_00_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.00-emphasis-good-simple-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// A document only containing an emphasized word
func Test_02_02_00_01_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.01-emphasis-good-single-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasis across two lines
func Test_02_02_00_02_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.02-emphasis-good-emphasis-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed emphasis
func Test_02_02_00_03_LexEmphasisBad(t *testing.T) {
	testPath := testPathFromName("02.02.00.03-emphasis-bad-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Unclosed emphasis on two lines with paragraphs and stuff.
func Test_02_02_00_04_LexEmphasisBad(t *testing.T) {
	testPath := testPathFromName("02.02.00.04-emphasis-bad-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_02_00_05_LexEmphasisBad(t *testing.T) {
	testPath := testPathFromName("02.02.00.05-emphasis-bad-emphasis-unclosed-surrounded-by-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasis surrounded by apostrophe
func Test_02_02_01_00_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.01.00-emphasis-good-emphasis-with-emphasis-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasis surrounded by quotes from many languages
func Test_02_02_01_01_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.01.01-emphasis-good-emphasis-surrounded-by-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasized asterisk
func Test_02_02_02_00_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.02.00-emphasis-good-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasized asterisk
func Test_02_02_02_01_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.02.01-emphasis-good-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

// Emphasized asterisk
func Test_02_02_02_02_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.02.02-emphasis-good-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_02_03_00_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.03.00-emphasis-good-emphasis-surrounded-by-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_02_02_04_00_LexEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.04.00-emphasis-good-emphasis-closed-with-strong-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
