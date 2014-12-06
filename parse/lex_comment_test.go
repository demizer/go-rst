// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import "testing"

func TestLexCommentGood0000(t *testing.T) {
	// A single comment
	testPath := testPathFromName("00.00-comment")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexCommentBlockGood0001(t *testing.T) {
	// A single comment block split with a newline
	testPath := testPathFromName("00.01-comment-block")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
