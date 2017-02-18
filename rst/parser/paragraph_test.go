package parser

import "testing"

// A single paragraph
func Test_02_00_00_00_ParseParagraphGood(t *testing.T) {
	testPath := testPathFromName("02.00.00.00-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a single paragraph with a line break in the middle
func Test_02_00_00_01_ParseParagraphWithLineBreakGood(t *testing.T) {
	testPath := testPathFromName("02.00.00.01-with-line-break")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a paragraph that is three lines long.
func Test_02_00_00_02_ParseParagraphThreeLinesGood(t *testing.T) {
	testPath := testPathFromName("02.00.00.02-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraphs separated by a line
func Test_02_00_01_00_ParseTwoParagraphs(t *testing.T) {
	testPath := testPathFromName("02.00.01.00-two-paragraphs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraphs separated by a line. Each paragraph is three lines.
func Test_02_00_01_01_ParseTwoParagraphsWithThreeLinesEach(t *testing.T) {
	testPath := testPathFromName("02.00.01.01-two-paragraphs-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
