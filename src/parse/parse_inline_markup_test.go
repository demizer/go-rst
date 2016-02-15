// go-rst - A reStructuredText parser for Go
// 2015 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import "testing"

func TestParseInlineMarkupGood0000(t *testing.T) {
	// Basic title, underline, blankline, and paragraph test
	testPath := testPathFromName("00.00-double-underscore")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
