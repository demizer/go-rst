// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"encoding/json"
	"github.com/demizer/go-elog"
	"github.com/demizer/go-spew/spew"
	"testing"
)

var (
	tEOF = item{Type: itemEOF, StartPosition: 0, Value: ""}
)

var spd = spew.ConfigState{Indent: "\t", DisableMethods: true}

var lexSectionTests LexTests

// collect gathers the emitted items into a slice.
func collect(t *LexTest) (items []item) {
	l := lex(t.name, t.data)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.Type == itemEOF || item.Type == itemError {
			break
		}
	}
	return
}

func lexSectionTest(t *testing.T, testName string) []item {
	var err error
	if lexSectionTests == nil {
		lexSectionTests, err = ParseTestData("../testdata/test_lex_sections.dat")
		if err != nil {
			t.Fatal(err)
		}
	}
	test := lexSectionTests.SearchByName(testName)
	if test != nil {
		log.Debugf("Test Name: \t%s\n", test.name)
		log.Debugf("Description: \t%s\n", test.description)
		log.Debugf("Test Input:\n-----------\n%s\n----------\n", test.data)
		items := collect(test)
		return items
	}
	return nil
}

// Test equality between items and expected items from unmarshalled json data, field by field.
// Returns error in case of error during json unmarshalling, or mismatch between items and the
// expected output.
func equal(t *testing.T, items []item, testName string) []error {
	test := lexSectionTests.SearchByName(testName)
	eItems, err := JsonToItems([]byte(test.items))
	if err != nil {
		t.Fatal("JSON error: ", err)
	}
	if len(items) != len(eItems) {
		t.Fatalf("Collected items is not the same length as eItems!\n"+
			"\nGot items (%d): -------------------------------\n\n%s\n"+
			"Expect items (%d): ------------------------------\n\n%s\n"+
			"-------------------------------------------------\n",
			len(items), spd.Sdump(items), len(eItems), spd.Sdump(eItems))
	}
	for i, item := range items {
		if item.ElementType != eItems[i].ElementType {
			t.Errorf("\n\nItem:\t%d\nElement Name:\t%q\nLine:\t%d\nValue:\t%q\n\n"+
				"Got ElementType:\t%s\nExpect ElementType:\t%s\n\n",
				i, item.ElementName, item.Line, item.Value, item.ElementType,
				eItems[i].ElementType)
		}
		if item.Line != eItems[i].Line {
			t.Errorf("\n\nItem:\t%d\nElement Name:\t%q\nValue:\t%q\n\n"+
				"Got Line Number:\t%d\nExpect Line Number:\t%d\n\n",
				i, item.ElementName, item.Value, item.Line, eItems[i].Line)
		}
		if item.StartPosition != eItems[i].StartPosition {
			t.Errorf("\n\nItem:\t%d\nElement Name:\t%q\nLine:\t%d\nValue:\t%q\n\n"+
				"Got Position:\t\t%d\nExpect Position:\t%d\n\n",
				i, item.ElementName, item.Line, item.Value, item.Position,
				eItems[i].Position)
		}
		if item.Value != eItems[i].Value {
			t.Errorf("\n\nItem:\t%d\nElement Name:\t%q\n\n"+
				"Got Value:\n\t%q\nExpect Value:\n\t%q\n\n",
				i, item.ElementName, item.Value, eItems[i].Value)
		}
	}
	return nil
}

func TestSectionTitlePara(t *testing.T) {
	testName := "SectionTitlePara"
	items := lexSectionTest(t, testName)
	errors := equal(t, items, testName)
	if errors != nil {
		for err := range errors {
			t.Error(err)
		}
	}
}

func TestSectionTitleParaNoBlankline(t *testing.T) {
	testName := "SectionTitleParaNoBlankline"
	items := lexSectionTest(t, testName)
	errors := equal(t, items, testName)
	if errors != nil {
		for err := range errors {
			t.Error(err)
		}
	}
}

func TestSectionParaHeadPara(t *testing.T) {
	testName := "SectionParaHeadPara"
	items := lexSectionTest(t, testName)
	errors := equal(t, items, testName)
	if errors != nil {
		for err := range errors {
			t.Error(err)
		}
	}
}

func TestSectionUnexpectedTitles(t *testing.T) {
	testName := "SectionUnexpectedTitles"
	items := lexSectionTest(t, testName)
	errors := equal(t, items, testName)
	if errors != nil {
		for err := range errors {
			t.Error(err)
		}
	}
}
