package parse

import "testing"

// A single paragraph
func Test_01_00_00_00_ParseParagraphGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.00-paragraph-good-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a single paragraph with a line break in the middle
func Test_01_00_00_01_ParseParagraphWithLineBreakGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.01-paragraph-good-with-line-break")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a paragraph that is three lines long.
func Test_01_00_00_02_ParseParagraphThreeLinesGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.02-paragraph-good-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraps separated by a line
func Test_01_00_01_00_ParseTwoParagraphs(t *testing.T) {
	testPath := testPathFromName("01.00.01.00-paragraph-good-two-paragraphs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraps separated by a line. Each paragraph is three lines.
func Test_01_00_01_01_ParseTwoParagraphsWithThreeLinesEach(t *testing.T) {
	testPath := testPathFromName("01.00.01.01-paragraph-good-two-paragraphs-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
