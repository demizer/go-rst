package parse

import "testing"

// Basic title, underline, blankline, and paragraph test
func Test_03_00_00_00_ParseSectionTitleGood(t *testing.T) {
	testPath := testPathFromName("03.00.00.00-section-good-title-paragraph")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Basic title, underline, and paragraph with no blankline line after the section.
func Test_03_00_00_01_ParseSectionTitleGood(t *testing.T) {
	testPath := testPathFromName("03.00.00.01-section-good-paragraph-noblankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A title that begins with a combining unicode character \u0301. Tests to make sure the 2 byte unicode does not contribute
// to the underline length calculation.
func Test_03_00_00_02_ParseSectionTitleGood(t *testing.T) {
	testPath := testPathFromName("03.00.00.02-section-good-title-combining-chars")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// A basic section in between paragraphs.
func Test_03_00_01_00_ParseSectionTitleGood(t *testing.T) {
	testPath := testPathFromName("03.00.01.00-section-good-para-head-para")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section parsing on 3 character long title and underline.
func Test_03_00_02_00_ParseSectionTitleGood(t *testing.T) {
	testPath := testPathFromName("03.00.02.00-section-good-short-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests a single section with no other element surrounding it.
func Test_03_00_03_00_ParseSectionTitleGood(t *testing.T) {
	testPath := testPathFromName("03.00.03.00-section-good-empty-section")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for severe system messages when the sections are indented.
func Test_03_00_04_00_ParseSectionTitleBad(t *testing.T) {
	testPath := testPathFromName("03.00.04.00-section-bad-unexpected-titles")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	// eNodes := test.expectNodes()
	spd.Dump(pTree.Nodes)
	// checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for severe system message on short title underline
func Test_03_00_05_00_ParseSectionTitleBad(t *testing.T) {
	testPath := testPathFromName("03.00.05.00-section-bad-short-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for title underlines that are less than three characters.
func Test_03_00_06_00_ParseSectionTitleBad(t *testing.T) {
	testPath := testPathFromName("03.00.06.00-section-bad-short-title-short-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for title overlines and underlines that are less than three characters.
func Test_03_00_06_01_ParseSectionTitleBad(t *testing.T) {
	testPath := testPathFromName("03.00.06.01-section-bad-short-title-short-overline-and-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for short title overline with missing underline when the overline is less than three characters.
func Test_03_00_06_02_ParseSectionTitleBad(t *testing.T) {
	testPath := testPathFromName("03.00.06.02-section-bad-short-title-short-overline-missing-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level return to level one after three subsections.
func Test_03_01_00_00_ParseSectionLevelGood(t *testing.T) {
	testPath := testPathFromName("03.01.00.00-section-good-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level return to level one after 1 subsection. The second level one section has one subsection.
func Test_03_01_00_01_ParseSectionLevelGood(t *testing.T) {
	testPath := testPathFromName("03.01.00.01-section-good-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section level with subsection 4 returning to level two.
func Test_03_01_00_02_ParseSectionLevelGood(t *testing.T) {
	testPath := testPathFromName("03.01.00.02-section-good-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level return with title overlines
func Test_03_01_01_00_ParseSectionLevelGood(t *testing.T) {
	testPath := testPathFromName("03.01.01.00-section-good-section-level-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests section level with two section having the same rune, but the first not having an overline.
func Test_03_01_02_00_ParseSectionLevelGood(t *testing.T) {
	testPath := testPathFromName("03.01.02.00-section-good-two-level-one-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section level return on bad level 2 section adornment
func Test_03_01_03_00_ParseSectionLevelBad(t *testing.T) {
	testPath := testPathFromName("03.01.03.00-section-bad-subsection-order")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section level return with title overlines on bad level 2 section adornment
func Test_03_01_03_01_ParseSectionLevelBad(t *testing.T) {
	testPath := testPathFromName("03.01.03.01-section-bad-subsection-order-with-overlines")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests for a severeTitleLevelInconsistent system message on a bad level two with an overline. Level one does not have an overline.
func Test_03_01_04_00_ParseSectionLevelBad(t *testing.T) {
	testPath := testPathFromName("03.01.04.00-section-bad-two-level-overline-bad-return")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test simple section with title overline.
func Test_03_02_00_00_ParseSectionTitleWithOverlineGood(t *testing.T) {
	testPath := testPathFromName("03.02.00.00-section-good-title-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test simple section with inset title and overline.
func Test_03_02_01_00_ParseSectionTitleWithOverlineGood(t *testing.T) {
	testPath := testPathFromName("03.02.01.00-section-good-inset-title-with-overline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test sections with three character adornments lines.
func Test_03_02_02_00_ParseSectionTitleWithOverlineGood(t *testing.T) {
	testPath := testPathFromName("03.02.02.00-section-good-three-char-section-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section title with overline, but no underline.
func Test_03_02_03_00_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.03.00-section-bad-inset-title-missing-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test inset title with overline but missing underline.
func Test_03_02_03_01_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.03.01-section-bad-inset-title-missing-underline-with-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test inset title with overline but missing underline. The title is followed by a blank line and a paragraph.
func Test_03_02_03_02_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.03.02-section-bad-inset-title-missing-underline-and-para")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test section overline with missmatched underline.
func Test_03_02_03_03_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.03.03-section-bad-inset-title-mismatched-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test overline with really long title.
func Test_03_02_04_00_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.04.00-section-bad-title-too-long")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test overline and underline with blanklines instead of a title.
func Test_03_02_05_00_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.05.00-section-bad-missing-titles-with-blankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test overline and underline with nothing where the title is supposed to be.
func Test_03_02_05_01_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.05.01-section-bad-missing-titles-with-noblankline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test two character overline with no underline.
func Test_03_02_06_00_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.06.00-section-bad-incomplete-section")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Test three character section adornments with no titles or blanklines in between.
func Test_03_02_06_01_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.06.01-section-bad-incomplete-sections-no-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests indented section with overline
func Test_03_02_07_00_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.07.00-section-bad-indented-title-short-overline-and-underline")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests ".." overline (which is a comment element).
func Test_03_02_08_00_ParseSectionTitleWithOverlineBad(t *testing.T) {
	testPath := testPathFromName("03.02.08.00-section-bad-two-char-section-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests lexing a section where the title begins with a number.
func Test_03_03_00_00_ParseSectionTitleNumberedGood(t *testing.T) {
	testPath := testPathFromName("03.03.00.00-section-good-numbered-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

// Tests numbered section lexing with enumerated directly above section.
func Test_03_03_01_00_ParseSectionTitleNumberedGood(t *testing.T) {
	testPath := testPathFromName("03.03.01.00-section-good-enum-list-with-numbered-title")
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
