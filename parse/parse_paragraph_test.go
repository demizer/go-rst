// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import "testing"

func TestParseParagraphGood0000(t *testing.T) {
	// A single paragraph
	testPath := testPathFromName("00.00-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseParagraphWithLineBreakGood0001(t *testing.T) {
	// Parse a single paragraph with a line break in the middle
	testPath := testPathFromName("00.01-para-with-line-break")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseParagraphThreeLinesGood0002(t *testing.T) {
	// Parse a paragraph that is three lines long.
	testPath := testPathFromName("00.02-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseTwoParagraphs0100(t *testing.T) {
	// Parse two paragraps separated by a line
	testPath := testPathFromName("01.00-two-paragraphs")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseTwoParagraphsWithThreeLinesEach0101(t *testing.T) {
	// Parse two paragraps separated by a line. Each paragraph is three
	// lines.
	testPath := testPathFromName("01.01-two-para-three-lines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
