package parse

import "testing"

func TestParseSectionTitleGood0000(t *testing.T) {
	// Basic title, underline, blankline, and paragraph test
	testPath := testPathFromName("00.00-title-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0001(t *testing.T) {
	// Basic title, underline, and paragraph with no blankline line after
	// the section.
	testPath := testPathFromName("00.01-paragraph-noblankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0002(t *testing.T) {
	// A title that begins with a combining unicode character \u0301. Tests
	// to make sure the 2 byte unicode does not contribute to the underline
	// length calculation.
	testPath := testPathFromName("00.02-title-combining-chars")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0100(t *testing.T) {
	// A basic section in between paragraphs.
	testPath := testPathFromName("01.00-para-head-para")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0200(t *testing.T) {
	// Tests section parsing on 3 character long title and underline.
	testPath := testPathFromName("02.00-short-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0300(t *testing.T) {
	// Tests a single section with no other element surrounding it.
	testPath := testPathFromName("03.00-empty-section")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0000(t *testing.T) {
	// Tests for severe system messages when the sections are indented.
	testPath := testPathFromName("00.00-unexpected-titles")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0100(t *testing.T) {
	// Tests for severe system message on short title underline
	testPath := testPathFromName("01.00-short-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0200(t *testing.T) {
	// Tests for title underlines that are less than three characters.
	testPath := testPathFromName("02.00-short-title-short-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0201(t *testing.T) {
	// Tests for title overlines and underlines that are less than three
	// characters.
	testPath := testPathFromName("02.01-short-title-short-overline-and-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0202(t *testing.T) {
	// Tests for short title overline with missing underline when the
	// overline is less than three characters.
	testPath := testPathFromName("02.02-short-title-short-overline-missing-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0000(t *testing.T) {
	// Tests section level return to level one after three subsections.
	testPath := testPathFromName("00.00-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0001(t *testing.T) {
	// Tests section level return to level one after 1 subsection. The
	// second level one section has one subsection.
	testPath := testPathFromName("00.01-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0002(t *testing.T) {
	// Test section level with subsection 4 returning to level two.
	testPath := testPathFromName("00.02-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0100(t *testing.T) {
	// Tests section level return with title overlines
	testPath := testPathFromName("01.00-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0200(t *testing.T) {
	// Tests section level with two section having the same rune, but the
	// first not having an overline.
	testPath := testPathFromName("02.00-two-level-one-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0000(t *testing.T) {
	// Test section level return on bad level 2 section adornment
	testPath := testPathFromName("00.00-bad-subsection-order")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0001(t *testing.T) {
	// Test section level return with title overlines on bad level 2
	// section adornment
	testPath := testPathFromName("00.01-bad-subsection-order-with-overlines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0100(t *testing.T) {
	// Tests for a severeTitleLevelInconsistent system message on a bad
	// level two with an overline. Level one does not have an overline.
	testPath := testPathFromName("01.00-two-level-overline-bad-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0000(t *testing.T) {
	// Test simple section with title overline.
	testPath := testPathFromName("00.00-title-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0100(t *testing.T) {
	// Test simple section with inset title and overline.
	testPath := testPathFromName("01.00-inset-title-with-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0200(t *testing.T) {
	// Test sections with three character adornments lines.
	testPath := testPathFromName("02.00-three-char-section-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0000(t *testing.T) {
	// Test section title with overline, but no underline.
	testPath := testPathFromName("00.00-inset-title-missing-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0001(t *testing.T) {
	// Test inset title with overline but missing underline.
	testPath := testPathFromName("00.01-inset-title-missing-underline-with-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0002(t *testing.T) {
	// Test inset title with overline but missing underline. The title is
	// followed by a blank line and a paragraph.
	testPath := testPathFromName("00.02-inset-title-missing-underline-and-para")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0003(t *testing.T) {
	// Test section overline with missmatched underline.
	testPath := testPathFromName("00.03-inset-title-mismatched-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0100(t *testing.T) {
	// Test overline with really long title.
	testPath := testPathFromName("01.00-title-too-long")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0200(t *testing.T) {
	// Test overline and underline with blanklines instead of a title.
	testPath := testPathFromName("02.00-missing-titles-with-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0201(t *testing.T) {
	// Test overline and underline with nothing where the title is supposed
	// to be.
	testPath := testPathFromName("02.01-missing-titles-with-noblankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0300(t *testing.T) {
	// Test two character overline with no underline.
	testPath := testPathFromName("03.00-incomplete-section")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0301(t *testing.T) {
	// Test three character section adornments with no titles or blanklines
	// in between.
	testPath := testPathFromName("03.01-incomplete-sections-no-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0400(t *testing.T) {
	// Tests indented section with overline
	testPath := testPathFromName("04.00-indented-title-short-overline-and-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0500(t *testing.T) {
	// Tests ".." overline (which is a comment element).
	testPath := testPathFromName("05.00-two-char-section-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleNumberedGood0000(t *testing.T) {
	// Tests lexing a section where the title begins with a number.
	testPath := testPathFromName("00.00-numbered-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleNumberedGood0100(t *testing.T) {
	// Tests numbered section lexing with enumerated directly above
	// section.
	testPath := testPathFromName("01.00-enum-list-with-numbered-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
