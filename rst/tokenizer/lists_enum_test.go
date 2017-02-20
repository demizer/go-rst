package tokenizer

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/rst/testutil"
)

func Test_08_00_00_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.00-enum-list-numbered")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_01_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.01-enum-list-numbered-noblanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_02_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.02-enum-list-numbered-indented-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_01_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.00-enum-list-items-with-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_02_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.00-enum-list-sequence-types")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_02_01_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.01-enum-list-ambiguous-sequence-types")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_03_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.03.00-enum-list-diff-formats")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_04_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.04.00-enum-list-nested")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_05_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.00-bad-enum-list-unexpected-unindent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_05_01_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.01-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_05_02_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.02-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_05_03_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.03-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.00-bad-enum-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_01_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.00-bad-enum-list-scrambled-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_02_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.00-bad-enum-list-skipped-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_03_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.03.00-bad-enum-list-not-ordinal-1")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_04_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.03.00-bad-enum-list-not-ordinal-1")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_05_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.00-bad-enum-list-ambiguous")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_06_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.00-bad-enum-list-nbsp-workaround")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_07_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.00-bad-enum-list-unexpected-unindent")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_08_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.08.00-bad-enum-list-auto-numbering-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_09_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.09.00-bad-enum-list-paragraph-not-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
