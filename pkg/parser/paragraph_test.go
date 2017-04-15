package parser

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

// A single paragraph
func Test_02_00_00_00_ParseParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.00-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a single paragraph with a line break in the middle
func Test_02_00_00_01_ParseParagraphWithLineBreakGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.01-with-line-break")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse a paragraph that is three lines long.
func Test_02_00_00_02_ParseParagraphThreeLinesGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.02-three-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraphs separated by a line
func Test_02_00_01_00_ParseTwoParagraphs(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.00-two-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Parse two paragraphs separated by a line. Each paragraph is three lines.
func Test_02_00_01_01_ParseTwoParagraphsWithThreeLinesEach(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.01-two-paragraphs-three-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
