package parse

import "testing"

// Basic title, underline, blankline, and paragraph test
func Test_02_00_00_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.00.00-imrr-good-double-underscore")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_01_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.01.00-imrr-good-lots-of-escaping")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_02_00_ParseInlineMarkupRecognitionRulesGood(t *testing.T) {
	testPath := testPathFromName("02.00.02.00-imrr-good-lots-of-escaping-unicode")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_00_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.00.00-strong-good-strong")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_02_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.02.00-strong-good-strong-quoted")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_03_00_ParseInlineMarkupStrongGood(t *testing.T) {
	testPath := testPathFromName("02.01.03.00-strong-good-strong-asterisk")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_01_01_00_ParseInlineMarkupStrongWithApostropheGood(t *testing.T) {
	testPath := testPathFromName("02.01.01.00-strong-good-strong-with-apostrophe")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_02_00_00_ParseInlineMarkupSimpleEmphasisGood(t *testing.T) {
	testPath := testPathFromName("02.02.00.00-emphasis-good-simple-emphasis")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
