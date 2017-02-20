package parser

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/rst/testutil"
)

func Test_01_00_00_00_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.00-internal-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_00_01_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.01-internal-target-space-before-colon")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_00_02_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.02-internal-target-space-before-colon")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_00_03_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.03-internal-target-across-lines")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_01_00_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.01.00-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_01_01_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.01.00-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_01_02_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.01.02-external-target-mailto")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_02_00_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.02.00-indirect-hyperlink-targets-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_03_00_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.00-anonymous-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_03_01_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.01-anonymous-external-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_03_02_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.00-anonymous-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
func Test_01_00_04_00_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.00-anonymous-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_04_01_ParseReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.01-anonymous-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_00_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.00-bad-target-malformed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_00_01_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.01-bad-target-malformed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_00_02_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.00.02-bad-target-malformed")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_01_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.01.00-bad-duplicate-internal-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_02_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.02.00-bad-duplicate-external-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_03_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.03.00-bad-duplicate-implicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_04_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.00-bad-duplicate-implicit-explicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_04_01_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.04.01-bad-duplicate-implicit-explicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_05_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.05.00-bad-duplicate-implicit-directive-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_06_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.06.00-bad-duplicate-explicit-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_07_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.07.00-bad-duplicate-explicit-directive-targets")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func Test_01_00_08_00_ParseReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("01.00.08.00-bad-anon-and-named-indirect-target")
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.ExpectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
