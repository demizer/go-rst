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
