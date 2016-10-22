package parse

import "testing"

func Test_02_04_00_00_LexInlineReferenceGood(t *testing.T) {
	testPath := testPathFromName("02.04.00.00-ref")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
