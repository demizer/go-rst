package parser

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_00_00_00_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.00-comment")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_00_01_ParseCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.01-bad-comment-no-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_00_02_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.02-comment-with-literal-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_00_03_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.00.03-comment-not-reference")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_01_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.00-comment-block")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_01_01_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.01.01-comment-block-second-line")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_02_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.00-newline-after-comment-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_02_01_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.01-newline-after-comment-mark")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_02_02_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.02-comment-not-citation")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_02_03_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.02.03-comment-not-subs-def")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_03_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.03.00-empty-comment-with-blockquote")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_04_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.00-comment-in-definition")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_04_01_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.04.01-comment-after-definition")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_05_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.00-comment-between-bullet-paragraphs")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_05_01_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.01-comment-between-bullets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_05_02_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.05.02-comment-trailing-bullet")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_06_00_ParseCommentGood(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.00-two-comments")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_00_00_06_01_ParseCommentBad(t *testing.T) {
	testPath := testutil.TestPathFromName("00.00.06.01-bad-two-comments-no-blankline")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
