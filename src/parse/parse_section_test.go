package parse

import "testing"

// Basic title, underline, blankline, and paragraph test
func TestParseSectionTitleGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-title-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Basic title, underline, and paragraph with no blankline line after the section.
func TestParseSectionTitleGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-paragraph-noblankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A title that begins with a combining unicode character \u0301. Tests to make sure the 2 byte unicode does not contribute
// to the underline length calculation.
func TestParseSectionTitleGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-title-combining-chars")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A basic section in between paragraphs.
func TestParseSectionTitleGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-para-head-para")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section parsing on 3 character long title and underline.
func TestParseSectionTitleGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-short-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests a single section with no other element surrounding it.
func TestParseSectionTitleGood0300(t *testing.T) {
	testPath := testPathFromName("03.00-empty-section")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for severe system messages when the sections are indented.
func TestParseSectionTitleBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-unexpected-titles")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for severe system message on short title underline
func TestParseSectionTitleBad0100(t *testing.T) {
	testPath := testPathFromName("01.00-short-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for title underlines that are less than three characters.
func TestParseSectionTitleBad0200(t *testing.T) {
	testPath := testPathFromName("02.00-short-title-short-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for title overlines and underlines that are less than three characters.
func TestParseSectionTitleBad0201(t *testing.T) {
	testPath := testPathFromName("02.01-short-title-short-overline-and-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for short title overline with missing underline when the overline is less than three characters.
func TestParseSectionTitleBad0202(t *testing.T) {
	testPath := testPathFromName("02.02-short-title-short-overline-missing-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level return to level one after three subsections.
func TestParseSectionLevelGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level return to level one after 1 subsection. The second level one section has one subsection.
func TestParseSectionLevelGood0001(t *testing.T) {
	testPath := testPathFromName("00.01-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section level with subsection 4 returning to level two.
func TestParseSectionLevelGood0002(t *testing.T) {
	testPath := testPathFromName("00.02-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level return with title overlines
func TestParseSectionLevelGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level with two section having the same rune, but the first not having an overline.
func TestParseSectionLevelGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-two-level-one-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section level return on bad level 2 section adornment
func TestParseSectionLevelBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-bad-subsection-order")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section level return with title overlines on bad level 2 section adornment
func TestParseSectionLevelBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-bad-subsection-order-with-overlines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for a severeTitleLevelInconsistent system message on a bad level two with an overline. Level one does not have an overline.
func TestParseSectionLevelBad0100(t *testing.T) {
	testPath := testPathFromName("01.00-two-level-overline-bad-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test simple section with title overline.
func TestParseSectionTitleWithOverlineGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-title-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test simple section with inset title and overline.
func TestParseSectionTitleWithOverlineGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-inset-title-with-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test sections with three character adornments lines.
func TestParseSectionTitleWithOverlineGood0200(t *testing.T) {
	testPath := testPathFromName("02.00-three-char-section-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section title with overline, but no underline.
func TestParseSectionTitleWithOverlineBad0000(t *testing.T) {
	testPath := testPathFromName("00.00-inset-title-missing-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test inset title with overline but missing underline.
func TestParseSectionTitleWithOverlineBad0001(t *testing.T) {
	testPath := testPathFromName("00.01-inset-title-missing-underline-with-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test inset title with overline but missing underline. The title is followed by a blank line and a paragraph.
func TestParseSectionTitleWithOverlineBad0002(t *testing.T) {
	testPath := testPathFromName("00.02-inset-title-missing-underline-and-para")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section overline with missmatched underline.
func TestParseSectionTitleWithOverlineBad0003(t *testing.T) {
	testPath := testPathFromName("00.03-inset-title-mismatched-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test overline with really long title.
func TestParseSectionTitleWithOverlineBad0100(t *testing.T) {
	testPath := testPathFromName("01.00-title-too-long")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test overline and underline with blanklines instead of a title.
func TestParseSectionTitleWithOverlineBad0200(t *testing.T) {
	testPath := testPathFromName("02.00-missing-titles-with-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test overline and underline with nothing where the title is supposed to be.
func TestParseSectionTitleWithOverlineBad0201(t *testing.T) {
	testPath := testPathFromName("02.01-missing-titles-with-noblankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test two character overline with no underline.
func TestParseSectionTitleWithOverlineBad0300(t *testing.T) {
	testPath := testPathFromName("03.00-incomplete-section")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test three character section adornments with no titles or blanklines in between.
func TestParseSectionTitleWithOverlineBad0301(t *testing.T) {
	testPath := testPathFromName("03.01-incomplete-sections-no-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests indented section with overline
func TestParseSectionTitleWithOverlineBad0400(t *testing.T) {
	testPath := testPathFromName("04.00-indented-title-short-overline-and-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests ".." overline (which is a comment element).
func TestParseSectionTitleWithOverlineBad0500(t *testing.T) {
	testPath := testPathFromName("05.00-two-char-section-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests lexing a section where the title begins with a number.
func TestParseSectionTitleNumberedGood0000(t *testing.T) {
	testPath := testPathFromName("00.00-numbered-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests numbered section lexing with enumerated directly above section.
func TestParseSectionTitleNumberedGood0100(t *testing.T) {
	testPath := testPathFromName("01.00-enum-list-with-numbered-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
