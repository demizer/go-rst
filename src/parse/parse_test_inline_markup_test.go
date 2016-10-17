package parse

import "testing"

// Basic title, underline, blankline, and paragraph test
func Test_02_00_00_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.00.00-double-underscore")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_01_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.01.00-lots-of-escaping")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_02_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.02.00-lots-of-escaping-unicode")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_03_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.03.00-emphasis-wrapped-in-unicode")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_04_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.04.00-openers-and-closers")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_04_01_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.04.01-strong-and-kwargs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_05_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.05.00-emphasis-with-backwards-rule-5")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_00_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.00.00-strong")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_02_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.02.00-strong-quoted")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_03_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.03.00-strong-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_03_01_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.03.01-strong-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_01_00_ParseInlineMarkupStrongWithApostropheGood(t *testing.T) {
	testPath := testPathFromName("02.01.01.00-strong-with-apostrophe")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_00_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.00-simple-emphasis")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_00_01_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.01-single-emphasis")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_00_02_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.02-emphasis-across-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_01_01_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.01.01-emphasis-surrounded-by-quotes")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_02_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.02.00-emphasis-with-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_02_01_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.02.01-emphasis-with-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_02_02_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.02.02-emphasis-with-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_03_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.03.00-emphasis-surrounded-by-markup")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_04_00_ParseInlineMarkupEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.04.00-emphasis-closed-with-strong-markup")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
