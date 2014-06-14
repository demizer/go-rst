// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import "testing"

func TestLexSectionTitleGood0000(t *testing.T) {
	// Basic title, underline, blankline, and paragraph test
	testPath := "test_section/01_title_good/00.00_title_paragraph"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleGood0001(t *testing.T) {
	// Basic title, underline, and paragraph with no blankline line after the
	// section.
	testPath := "test_section/01_title_good/00.01_paragraph_noblankline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleGood0002(t *testing.T) {
	// A title that begins with a combining unicode character \u0301. Tests to
	// make sure the 2 byte unicode does not contribute to the underline length
	// calculation.
	testPath := "test_section/01_title_good/00.02_title_combining_chars"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleGood0100(t *testing.T) {
	// A basic section in between paragraphs.
	testPath := "test_section/01_title_good/01.00_para_head_para"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleGood0200(t *testing.T) {
	// Tests section parsing on 3 character long title and underline.
	testPath := "test_section/01_title_good/02.00_short_title"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleGood0300(t *testing.T) {
	// Tests a single section with no other element surrounding it.
	testPath := "test_section/01_title_good/03.00_empty_section"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleBad0000(t *testing.T) {
	// Tests for severe system messages when the sections are indented.
	testPath := "test_section/02_title_bad/00.00_unexpected_titles"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleBad0100(t *testing.T) {
	// Tests for severe system message on short title underline
	testPath := "test_section/02_title_bad/01.00_short_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleBad0200(t *testing.T) {
	// Tests for title underlines that are less than three characters.
	testPath := "test_section/02_title_bad/" +
		"02.00_short_title_short_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleBad0201(t *testing.T) {
	// Tests for title overlines and underlines that are less than three
	// characters.
	testPath := "test_section/02_title_bad/" +
		"02.01_short_title_short_overline_and_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleBad0202(t *testing.T) {
	// Tests for short title overline with missing underline when the
	// overline is less than three characters.
	testPath := "test_section/02_title_bad/0" +
		"2.02_short_title_short_overline_missing_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelGood0000(t *testing.T) {
	// Tests section level return to level one after three subsections.
	testPath := "test_section/03_level_good/00.00_section_level_return"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelGood0001(t *testing.T) {
	// Tests section level return to level one after 1 subsection. The
	// second level one section has one subsection.
	testPath := "test_section/03_level_good/00.01_section_level_return"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelGood0002(t *testing.T) {
	// Test section level with subsection 4 returning to level two.
	testPath := "test_section/03_level_good/00.02_section_level_return"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelGood0100(t *testing.T) {
	// Tests section level return with title overlines
	testPath := "test_section/03_level_good/01.00_section_level_return"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelGood0200(t *testing.T) {
	// Tests section level with two section having the same rune, but the
	// first not having an overline.
	testPath := "test_section/03_level_good/02.00_two_level_one_overline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelBad0000(t *testing.T) {
	// Test section level return on bad level 2 section adornment
	testPath := "test_section/04_level_bad/00.00_bad_subsection_order"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelBad0001(t *testing.T) {
	// Test section level return with title overlines on bad level 2
	// section adornment
	testPath := "test_section/04_level_bad/" +
		"00.01_bad_subsection_order_with_overlines"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionLevelBad0100(t *testing.T) {
	// Tests for a severeTitleLevelInconsistent system message on a bad
	// level two with an overline. Level one does not have an overline.
	testPath := "test_section/04_level_bad/" +
		"01.00_two_level_overline_bad_return"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineGood0000(t *testing.T) {
	// Test simple section with title overline.
	testPath := "test_section/05_title_with_overline_good/" +
		"00.00_title_overline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineGood0100(t *testing.T) {
	// Test simple section with inset title and overline.
	testPath := "test_section/05_title_with_overline_good/" +
		"01.00_inset_title_with_overline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineGood0200(t *testing.T) {
	// Test sections with three character adornments lines.
	testPath := "test_section/05_title_with_overline_good/" +
		"02.00_three_char_section_title"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0000(t *testing.T) {
	// Test section title with overline, but no underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.00_inset_title_missing_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0001(t *testing.T) {
	// Test inset title with overline but missing underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.01_inset_title_missing_underline_with_blankline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0002(t *testing.T) {
	// Test inset title with overline but missing underline. The title is
	// followed by a blank line and a paragraph.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.02_inset_title_missing_underline_and_para"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0003(t *testing.T) {
	// Test section overline with missmatched underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.03_inset_title_mismatched_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0100(t *testing.T) {
	// Test overline with really long title.
	testPath := "test_section/06_title_with_overline_bad/" +
		"01.00_title_too_long"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0200(t *testing.T) {
	// Test overline and underline with blanklines instead of a title.
	testPath := "test_section/06_title_with_overline_bad/" +
		"02.00_missing_titles_with_blankline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0201(t *testing.T) {
	// Test overline and underline with nothing where the title is supposed
	// to be.
	testPath := "test_section/06_title_with_overline_bad/" +
		"02.01_missing_titles_with_noblankline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0300(t *testing.T) {
	// Test two character overline with no underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"03.00_incomplete_section"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0301(t *testing.T) {
	// Test three character section adornments with no titles or blanklines
	// in between.
	testPath := "test_section/06_title_with_overline_bad/" +
		"03.01_incomplete_sections_no_title"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0400(t *testing.T) {
	// Tests indented section with overline
	testPath := "test_section/06_title_with_overline_bad/" +
		"04.00_indented_title_short_overline_and_underline"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleWithOverlineBad0500(t *testing.T) {
	// Tests ".." overline (which is a comment element).
	testPath := "test_section/06_title_with_overline_bad/" +
		"05.00_two_char_section_title"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleNumberedGood0000(t *testing.T) {
	// Tests lexing a section where the title begins with a number.
	testPath := "test_section/07_title_numbered_good/00.00_numbered_title"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSectionTitleNumberedGood0100(t *testing.T) {
	// Tests numbered section lexing with enumerated directly above section.
	testPath := "test_section/07_title_numbered_good/" +
		"01.00_enum_list_with_numbered_title"
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}
