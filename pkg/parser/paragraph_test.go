package parser

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_02_00_00_00_ParseParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.00-paragraph")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_00_01_ParseParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.01-with-line-break")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_00_02_ParseParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.02-three-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_01_00_ParseParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.00-two-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_02_00_01_01_ParseParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.01-two-paragraphs-three-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
