package parse

import "testing"

// Basic title, underline, blankline, and paragraph test
func Test_06_00_00_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.00.00-double-underscore")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_00_01_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.01.00-lots-of-escaping")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_00_02_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.02.00-lots-of-escaping-unicode")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_00_03_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.03.00-emphasis-wrapped-in-unicode")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_00_04_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.04.00-openers-and-closers")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_00_04_01_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.04.01-strong-and-kwargs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_00_05_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("06.00.05.00-emphasis-with-backwards-rule-5")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_01_00_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("06.01.00.00-strong")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_01_02_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("06.01.02.00-strong-quoted")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_01_03_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("06.01.03.00-strong-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_01_03_01_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("06.01.03.01-strong-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_01_04_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("06.01.04.00-strong-accross-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_01_01_00_ParseInlineMarkupStrongWithApostropheGood(t *testing.T) {
	testPath := testPathFromName("06.01.01.00-strong-with-apostrophe")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_00_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.00.00-simple-emphasis")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_00_01_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.00.01-single-emphasis")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_00_02_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.00.02-emphasis-across-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_01_01_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.01.01-emphasis-surrounded-by-quotes")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_02_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.02.00-emphasis-with-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_02_01_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.02.01-emphasis-with-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_02_02_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.02.02-emphasis-with-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_03_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.03.00-emphasis-surrounded-by-markup")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_02_04_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("06.02.04.00-emphasis-closed-with-strong-markup")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_00_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.00.00-literal")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_00_01_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.00.01-literal")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_00_02_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.00.02-literal")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_01_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.01.00-literal-with-backslash")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_01_01_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.01.01-literal-with-middle-backslash")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_01_02_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.01.02-literal-with-end-backslash")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_02_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.02.00-literal-with-apostrophe")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_03_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.03.00-literal-quoted")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_03_01_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.03.01-literal-quoted-literal")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_04_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.04.00-literal-interpreted-text")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_05_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.05.00-literal-followed-by-backslash")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_06_03_06_00_ParseInlineMarkupLiteralGood(t *testing.T) {
	testPath := testPathFromName("06.03.06.00-literal-with-tex-quotes")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// func Test_06_04_00_00_ParseInlineReferenceGood(t *testing.T) {
// testPath := testPathFromName("06.04.00.00-ref")
// test := LoadParseTest(t, testPath)
// pTree := parseTest(t, test)
// eNodes := test.expectNodes()
// checkParseNodes(t, eNodes, pTree.Nodes, testPath)
// }
