// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

import (
	"testing"
	"github.com/demizer/go-elog"
)

type lexTest struct {
	name  string
	input string
	items []item
}

var (
	tEOF = item{itemEOF, 0, ""}
)

// collect gathers the emitted items into a slice.
func collect(t *lexTest) (items []item) {
	l := lex(t.name, t.input)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.typ == itemEOF || item.typ == itemError {
			break
		}
	}
	return
}

func equal(i1, i2 []item, checkPos bool) bool {
	if len(i1) != len(i2) {
		return false
	}
	for k := range i1 {
		if i1[k].typ != i2[k].typ {
			return false
		}
		if i1[k].val != i2[k].val {
			return false
		}
		if checkPos && i1[k].pos != i2[k].pos {
			return false
		}
	}
	return true
}

func setDebug() {
	log.SetLevel(log.LEVEL_DEBUG)
	log.SetFlags(log.Lansi | log.LnoFileAnsi | log.LnoPrefix)
}

func checkLexTest(t *testing.T, test *lexTest) {
	items := collect(test)
	if !equal(items, test.items, false) {
		t.Errorf("Test Name: %s\nGot:\n\t%+v\nExpected:\n\t%+v", test.name, items,
			test.items)
	}
}

func TestParagraphLex(t *testing.T) {
	var lexParagraphTests = []lexTest{
		{"Empty", "", []item{tEOF}},
		{"paragraph", "A paragraph.", []item{
			{itemParagraph, 0, "A paragraph."},
			tEOF,
		}},
	}
	for _, test := range lexParagraphTests {
		checkLexTest(t, &test)
	}
}

func TestSectionNoBlankLine(t *testing.T) {
	test := &lexTest{"section header, no blank line",
		"Title\n=====\nParagraph (no blank line).", []item{
		{itemTitle, 0, "Title"},
		{itemSectionAdornment, 0, "====="},
		{itemParagraph, 0, "Paragraph (no blank line)."},
		tEOF,
	}}
	checkLexTest(t, test)
}

func TestSectionWithBlankLine(t *testing.T) {
	test := &lexTest{"Section and paragraph", "Title\n=====\n\nParagraph.", []item{
		{itemTitle, 0, "Title"},
		{itemSectionAdornment, 0, "====="},
		{itemParagraph, 0, "Paragraph."},
		tEOF,
	}}
	checkLexTest(t, test)
}

func TestSectionWithOverline(t *testing.T) {
	test := &lexTest{"Section and paragraph (overline)", "=====\nTitle\n=====\nParagraph.", []item{
		{itemSectionAdornment, 0, "====="},
		{itemTitle, 0, "Title"},
		{itemSectionAdornment, 0, "====="},
		{itemParagraph, 0, "Paragraph."},
		tEOF,
	}}
	checkLexTest(t, test)
}

func TestParagraphSectionParagraph(t *testing.T) {
	test := &lexTest{"Paragraph section paragraph", "Paragraph.\nTitle\n=====\n\nParagraph.", []item{
		{itemParagraph, 0, "Paragraph."},
		{itemTitle, 0, "Title"},
		{itemSectionAdornment, 0, "====="},
		{itemParagraph, 0, "Paragraph."},
		tEOF,
	}}
	checkLexTest(t, test)
}
