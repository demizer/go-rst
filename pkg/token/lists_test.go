package token

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
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

func Test_07_00_00_02_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.02-bullet-list-noblankline-between-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_07_00_00_03_LexBulletListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.00.03-bad-bullet-list-noblankline-at-end")
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

func Test_07_00_02_01_LexBulletListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.02.01-bad-bullet-list-different-bullets-missing-blankline")
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

func Test_07_00_03_01_LexBulletListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("07.00.03.01-bad-bullet-list-empty-item-noblankline")
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

func Test_08_00_00_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.00-enum-list-numbered")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_01_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.01-enum-list-numbered-noblanklines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_02_LexBulletListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.02-enum-list-numbered-indented-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_03_LexBulletListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.03-bad-enum-list-empty-item-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_04_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.04-bad-enum-list-scrambled-items")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_05_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.05-bad-enum-list-skipped-item")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_00_06_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.00.06-bad-enum-list-not-ordinal-1")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_01_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.00-alphabetical-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_01_01_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.01-bad-alphabetical-list-without-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_01_02_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.01.02-alphabetical-list-nbsp-workaround")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_02_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.00-enum-list-items-with-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_02_01_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.02.01-bad-enum-list-unexpected-unindent")
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
	testPath := testutil.TestPathFromName("08.00.04.00-roman-numerals")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_04_01_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.04.01-bad-enum-list-bad-roman-numerals")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_05_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.05.00-enum-list-nested")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_06_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.00-enum-list-sequence-types")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_06_01_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.01-enum-list-ambiguous-sequence-types")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_06_02_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.06.02-bad-enum-list-ambiguous")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_07_00_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.00-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_07_01_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.01-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_07_02_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.02-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_07_03_LexEnumeratedListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.03-enum-list-auto-numbering")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_07_04_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.07.04-bad-enum-list-auto-numbering-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_08_00_08_00_LexEnumeratedListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("08.00.08.00-bad-enum-list-paragraph-not-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_00_00_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.00-def-list")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_00_01_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.01-def-list-with-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_00_02_LexDefinitionListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.02-bad-def-list-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_00_03_LexDefinitionListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.00.03-bad-def-list-not-def-term")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_01_00_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.00-def-list-two-terms")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_01_01_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.01-def-list-two-terms-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_01_02_LexDefinitionListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.01.02-bad-def-list-noblankline-after-two-terms")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_02_00_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.02.00-def-list-nested-terms")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_03_00_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.03.00-def-list-term-with-classifier")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_04_00_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.00-def-list-term-not-classifier")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_04_01_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.01-def-list-term-not-classifier-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_04_02_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.04.02-def-list-two-classifiers")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_05_00_LexDefinitionListGood_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.05.00-def-list-not-literal")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_09_00_06_00_LexDefinitionListBad_NotImplemented(t *testing.T) {
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	testPath := testutil.TestPathFromName("09.00.06.00-bad-def-list-with-inline-markup-errors")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
