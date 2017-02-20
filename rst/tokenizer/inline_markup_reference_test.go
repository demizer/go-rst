package tokenizer

import (
	"testing"

	"github.com/demizer/go-rst/rst/testutil"
)

func Test_06_04_00_00_LexInlineReferenceGood(t *testing.T) {
	testPath := testutil.TestPathFromName("06.04.00.00-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
}
