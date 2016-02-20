package parse

import "testing"

// Basic title, underline, blankline, and paragraph test
func TestParseInlineMarkupGoodDoubleUnderScore0000(t *testing.T) {
	testPath := testPathFromName("00.00-double-underscore")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseInlineMarkupGoodEscaping0100(t *testing.T) {
	testPath := testPathFromName("01.00-lots-of-escaping")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
