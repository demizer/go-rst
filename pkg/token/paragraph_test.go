package token

import (
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

func Test_02_00_00_00_LexParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.00-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_02_00_00_01_LexParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.01-with-line-break")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_02_00_00_02_LexParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.00.02-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_02_00_01_00_LexParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.00-two-paragraphs")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}

func Test_02_00_01_01_LexParagraphGood(t *testing.T) {
	testPath := testutil.TestPathFromName("02.00.01.01-two-paragraphs-three-lines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
