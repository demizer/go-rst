package parse

import (
	"os"
	"testing"
)

func Test_01_00_00_00_LexReferenceHyperlinkTargetGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.00-internal-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_00_01_LexReferenceHyperlinkTargetGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.01-internal-target-space-before-colon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_00_02_LexReferenceHyperlinkTargetGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.02-internal-target-space-before-colon")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_00_03_LexReferenceHyperlinkTargetGood(t *testing.T) {
	testPath := testPathFromName("01.00.00.03-internal-target-across-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_01_00_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.01.00-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	// str, _ := json.MarshalIndent(items, "", "    ")
	// fmt.Println(string(str))
	// os.Exit(1)
	equal(t, test.expectItems(), items)
}

func Test_01_00_01_01_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.01.00-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_01_02_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.01.02-external-target-mailto")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_02_00_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.02.00-indirect-hyperlink-targets-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_03_00_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.03.00-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_03_01_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.03.01-anonymous-external-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_03_02_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.04.00-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_04_00_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.04.00-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_04_01_LexReferenceHyperlinkTargetGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.04.01-anonymous-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_00_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.00.00-bad-target-missing-backquote")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_00_01_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.00.01-bad-target-malformed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_00_02_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.00.02-bad-target-malformed")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_01_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.01.00-bad-duplicate-internal-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_02_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.02.00-bad-duplicate-external-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_03_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.03.00-bad-duplicate-implicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_04_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.04.00-bad-duplicate-implicit-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_04_01_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.04.01-bad-duplicate-implicit-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_05_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.05.00-bad-duplicate-implicit-directive-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_06_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.06.00-bad-duplicate-explicit-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_07_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.07.00-bad-duplicate-explicit-directive-targets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func Test_01_00_08_00_LexReferenceHyperlinkTargetBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testPathFromName("01.00.08.00-bad-anon-and-named-indirect-target")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
