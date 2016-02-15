package parse

import "testing"

// A single paragraph
func TestParseParagraphGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a single paragraph with a line break in the middle
func TestParseParagraphWithLineBreakGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-para-with-line-break")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a paragraph that is three lines long.
func TestParseParagraphThreeLinesGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraps separated by a line
func TestParseTwoParagraphs0100(t *testing.T) {
	testPath := testPathFromName("01.00-two-paragraphs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraps separated by a line. Each paragraph is three lines.
func TestParseTwoParagraphsWithThreeLinesEach0101(t *testing.T) {
	testPath := testPathFromName("01.01-two-para-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
