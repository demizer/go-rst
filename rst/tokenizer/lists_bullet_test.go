package tokenizer

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/rst/testutil"
)

func Test_07_00_00_00_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bullet-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_00_01_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.01-bullet-list-with-two-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_00_03_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bullet-list-noblankline-between-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_01_00_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_01_01_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.01-bullet-list-item-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_02_00_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.00-bullet-list-different-bullets")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_03_00_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.00-bullet-list-empty-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_04_00_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.04.00-bullet-list-unicode")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_00_00_LexBulletListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.00-bad-bullet-list-noblankline-at-end")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_01_00_LexBulletListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.01.00-bullet-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
