package token

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

// Tests double underscore recognition
func Test_06_00_00_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.00.00-double-underscore")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Tests double underscore recognition
func Test_06_00_01_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.01.00-lots-of-escaping")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Tests escaping with unicode literals
func Test_06_00_02_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.02.00-lots-of-escaping-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_03_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.00-emphasis-wrapped-in-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Bad emphasis start string with unicode literal space after start string. The text that is suppossed to close the emphasis
// is lexed as an emphasis start string without a corresponding end string.
func Test_06_00_03_01_LexInlineMarkupRecognitionRulesBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.01-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Bad emphasis start string with unicode literal space after start string. The text that is suppossed to close the emphasis
// is lexed as an emphasis start string without a corresponding end string. This test includes more text around the incorrect
// emphasis.
func Test_06_00_03_02_LexInlineMarkupRecognitionRulesBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.03.02-emphasis-with-unicode-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_04_00_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.00-openers-and-closers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_04_01_LexInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.04.01-strong-and-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_00_05_00_LexInlineMarkupRecognitionRulesGood5(t *testing.T) {
	testPath := testutil.TestPathFromName("06.00.05.00-emphasis-with-backwards-rule-5")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Single strong
func Test_06_01_00_00_LexStrongGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.00-strong")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Strong with missing closing markup
func Test_06_01_00_01_LexStrongBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.01-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Strong with missing closing markup across three lines
func Test_06_01_00_02_LexStrongBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.00.02-strong-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Single strong with apostrophes
func Test_06_01_01_00_LexStrongGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.01.00-strong-with-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Strong with unicode literal quotes
func Test_06_01_02_00_LexStrongGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.02.00-strong-quoted")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Strong with asterisk and double asterisk
func Test_06_01_03_00_LexStrongGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.00-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Strong with asterisk and double asterisk and strong asterisk gone to plaid!
func Test_06_01_03_01_LexStrongGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.01-strong-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Inline strong across two lines
func Test_06_01_04_00_LexStrongGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.04.00-strong-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Strong with bad kwargs
func Test_06_01_03_02_LexStrongBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.01.03.02-strong-kwargs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// A paragraph containing a single emphasized word
func Test_06_02_00_00_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.00-simple-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// A document only containing an emphasized word
func Test_06_02_00_01_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.01-single-emphasis")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Emphasis across two lines
func Test_06_02_00_02_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.02-emphasis-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Unclosed emphasis
func Test_06_02_00_03_LexEmphasisBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.03-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Unclosed emphasis on two lines with paragraphs and stuff.
func Test_06_02_00_04_LexEmphasisBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.04-emphasis-unclosed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_00_05_LexEmphasisBad(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.00.05-emphasis-unclosed-surrounded-by-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Emphasis surrounded by apostrophe
func Test_06_02_01_00_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.00-emphasis-with-emphasis-apostrophe")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Emphasis surrounded by quotes from many languages
func Test_06_02_01_01_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.01.01-emphasis-surrounded-by-quotes")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Emphasized asterisk
func Test_06_02_02_00_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.00-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Emphasized asterisk
func Test_06_02_02_01_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.01-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

// Emphasized asterisk
func Test_06_02_02_02_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.02.02-emphasis-with-asterisk")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_03_00_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.03.00-emphasis-surrounded-by-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_06_02_04_00_LexEmphasisGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.02.04.00-emphasis-closed-with-strong-markup")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
