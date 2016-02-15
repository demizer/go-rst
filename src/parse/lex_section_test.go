package parse

import "testing"

func TestLexSectionTitleGood0000(t *testing.T) {
	// Basic title, underline, blankline, and paragraph test
	testPath := testPathFromName("00.00-title-paragraph")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleGood0001(t *testing.T) {
	// Basic title, underline, and paragraph with no blankline line after the
	// section.
	testPath := testPathFromName("00.01-paragraph-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleGood0002(t *testing.T) {
	// A title that begins with a combining unicode character \u0301. Tests to
	// make sure the 2 byte unicode does not contribute to the underline length
	// calculation.
	testPath := testPathFromName("00.02-title-combining-chars")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleGood0100(t *testing.T) {
	// A basic section in between paragraphs.
	testPath := testPathFromName("01.00-para-head-para")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleGood0200(t *testing.T) {
	// Tests section parsing on 3 character long title and underline.
	testPath := testPathFromName("02.00-short-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleGood0300(t *testing.T) {
	// Tests a single section with no other element surrounding it.
	testPath := testPathFromName("03.00-empty-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleBad0000(t *testing.T) {
	// Tests for severe system messages when the sections are indented.
	testPath := testPathFromName("00.00-unexpected-titles")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleBad0100(t *testing.T) {
	// Tests for severe system message on short title underline
	testPath := testPathFromName("01.00-short-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleBad0200(t *testing.T) {
	// Tests for title underlines that are less than three characters.
	testPath := testPathFromName("02.00-short-title-short-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleBad0201(t *testing.T) {
	// Tests for title overlines and underlines that are less than three
	// characters.
	testPath := testPathFromName("02.01-short-title-short-overline-and-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleBad0202(t *testing.T) {
	// Tests for short title overline with missing underline when the
	// overline is less than three characters.
	testPath := testPathFromName("02.02-short-title-short-overline-missing-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelGood0000(t *testing.T) {
	// Tests section level return to level one after three subsections.
	testPath := testPathFromName("00.00-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelGood0001(t *testing.T) {
	// Tests section level return to level one after 1 subsection. The
	// second level one section has one subsection.
	testPath := testPathFromName("00.01-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelGood0002(t *testing.T) {
	// Test section level with subsection 4 returning to level two.
	testPath := testPathFromName("00.02-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelGood0100(t *testing.T) {
	// Tests section level return with title overlines
	testPath := testPathFromName("01.00-section-level-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelGood0200(t *testing.T) {
	// Tests section level with two section having the same rune, but the
	// first not having an overline.
	testPath := testPathFromName("02.00-two-level-one-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelBad0000(t *testing.T) {
	// Test section level return on bad level 2 section adornment
	testPath := testPathFromName("00.00-bad-subsection-order")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelBad0001(t *testing.T) {
	// Test section level return with title overlines on bad level 2
	// section adornment
	testPath := testPathFromName("00.01-bad-subsection-order-with-overlines")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionLevelBad0100(t *testing.T) {
	// Tests for a severeTitleLevelInconsistent system message on a bad
	// level two with an overline. Level one does not have an overline.
	testPath := testPathFromName("01.00-two-level-overline-bad-return")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineGood0000(t *testing.T) {
	// Test simple section with title overline.
	testPath := testPathFromName("00.00-title-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineGood0100(t *testing.T) {
	// Test simple section with inset title and overline.
	testPath := testPathFromName("01.00-inset-title-with-overline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineGood0200(t *testing.T) {
	// Test sections with three character adornments lines.
	testPath := testPathFromName("02.00-three-char-section-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0000(t *testing.T) {
	// Test section title with overline, but no underline.
	testPath := testPathFromName("00.00-inset-title-missing-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0001(t *testing.T) {
	// Test inset title with overline but missing underline.
	testPath := testPathFromName("00.01-inset-title-missing-underline-with-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0002(t *testing.T) {
	// Test inset title with overline but missing underline. The title is
	// followed by a blank line and a paragraph.
	testPath := testPathFromName("00.02-inset-title-missing-underline-and-para")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0003(t *testing.T) {
	// Test section overline with missmatched underline.
	testPath := testPathFromName("00.03-inset-title-mismatched-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0100(t *testing.T) {
	// Test overline with really long title.
	testPath := testPathFromName("01.00-title-too-long")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0200(t *testing.T) {
	// Test overline and underline with blanklines instead of a title.
	testPath := testPathFromName("02.00-missing-titles-with-blankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0201(t *testing.T) {
	// Test overline and underline with nothing where the title is supposed
	// to be.
	testPath := testPathFromName("02.01-missing-titles-with-noblankline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0300(t *testing.T) {
	// Test two character overline with no underline.
	testPath := testPathFromName("03.00-incomplete-section")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0301(t *testing.T) {
	// Test three character section adornments with no titles or blanklines
	// in between.
	testPath := testPathFromName("03.01-incomplete-sections-no-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0400(t *testing.T) {
	// Tests indented section with overline
	testPath := testPathFromName("04.00-indented-title-short-overline-and-underline")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleWithOverlineBad0500(t *testing.T) {
	// Tests ".." overline (which is a comment element).
	testPath := testPathFromName("05.00-two-char-section-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleNumberedGood0000(t *testing.T) {
	// Tests lexing a section where the title begins with a number.
	testPath := testPathFromName("00.00-numbered-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}

func TestLexSectionTitleNumberedGood0100(t *testing.T) {
	// Tests numbered section lexing with enumerated directly above section.
	testPath := testPathFromName("01.00-enum-list-with-numbered-title")
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.expectItems(), items)
}
