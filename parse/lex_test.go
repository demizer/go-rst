// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"code.google.com/p/go.text/unicode/norm"
	"github.com/demizer/go-elog"
	"reflect"
	"testing"
	"fmt"
	"strconv"
)

var (
	tEOF = item{Type: itemEOF, StartPosition: 0, Text: ""}
)

func lexTest(t *testing.T, test *Test) []item {
	log.Debugf("Test Path: %s\n", test.path)
	log.Debugf("Test Input:\n-----------\n%s\n----------\n", test.data)
	var items []item
	l := lex(test.path, test.data)
	for {
		item := l.nextItem()
		items = append(items, *item)
		if item.Type == itemEOF || item.Type == itemError {
			break
		}
	}
	return items
}

// Test equality between items and expected items from unmarshalled json data, field by field.
// Returns error in case of error during json unmarshalling, or mismatch between items and the
// expected output.
func equal(t *testing.T, items []item, expectItems []item) {
	var id int
	var found bool
	var pFieldName, eFieldName string
	var pFieldVal, eFieldVal reflect.Value
	var pFieldValStruct reflect.StructField

	dError := func() {
		var got, exp string
		switch r := pFieldVal.Interface().(type) {
		case Id:
			got = pFieldVal.Interface().(Id).String()
			exp = eFieldVal.Interface().(Id).String()
		case itemElement:
			got = pFieldVal.Interface().(itemElement).String()
			exp = eFieldVal.Interface().(itemElement).String()
		case Line:
			got = pFieldVal.Interface().(Line).String()
			exp = eFieldVal.Interface().(Line).String()
		case StartPosition:
			got = pFieldVal.Interface().(StartPosition).String()
			exp = eFieldVal.Interface().(StartPosition).String()
		case int:
			got = strconv.Itoa(pFieldVal.Interface().(int))
			exp = strconv.Itoa(eFieldVal.Interface().(int))
		default:
			panic(fmt.Errorf("%#v is not implemented!", r))
		}
		t.Errorf("\n(Id: %d) Got: %s = %q, Expect: %s = %q\n", id, pFieldName, got,
			eFieldName, exp)
	}

	for eNum, eItem := range expectItems {
		eVal := reflect.ValueOf(eItem)
		pVal := reflect.ValueOf(items[eNum])
		id = int(pVal.FieldByName("Id").Interface().(Id))
		for x := 0; x < eVal.NumField(); x++ {
			eFieldVal = eVal.Field(x)
			eFieldName = eVal.Type().Field(x).Name
			pFieldVal = pVal.FieldByName(eFieldName)
			pFieldValStruct, found = pVal.Type().FieldByName(eFieldName)
			pFieldName = pFieldValStruct.Name
			if !found {
				t.Errorf("Parsed item (Id: %d) does not contain field %q\n", id,
					eFieldName)
				continue
			} else if eFieldName == "Text" {
				if eFieldVal.Interface() == nil {
					continue
				}
				if pFieldVal.Interface() !=
					norm.NFC.String(eFieldVal.Interface().(string)) {
					dError()
				}
			} else if pFieldVal.Interface() != eFieldVal.Interface() {
				dError()
			}
		}
	}

	return
}

func TestId(t *testing.T) {
	testPath := "test_section/001_title_paragraph"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	if items[0].IdNumber() != 1 {
		t.Error("Id != 1")
	}
	if items[0].Id.String() != "1" {
		t.Error(`String Id != "1"`)
	}
}

func TestLine(t *testing.T) {
	testPath := "test_section/001_title_paragraph"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	if items[0].LineNumber() != 1 {
		t.Error("Line != 1")
	}
	if items[0].Line.String() != "1" {
		t.Error(`String Line != "1"`)
	}
}

func TestStartPosition(t *testing.T) {
	testPath := "test_section/001_title_paragraph"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	if items[0].Position() != 1 {
		t.Error("StartPosition != 1")
	}
	if items[0].StartPosition.String() != "1" {
		t.Error(`String StartPosition != "1"`)
	}
}

func TestLexSection001(t *testing.T) {
	testPath := "test_section/001_title_paragraph"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection002(t *testing.T) {
	testPath := "test_section/002_paragraph_nbl"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection003(t *testing.T) {
	testPath := "test_section/003_para_head_para"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection004(t *testing.T) {
	testPath := "test_section/004_section_level_test"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection005(t *testing.T) {
	testPath := "test_section/005_unexpected_titles"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection006(t *testing.T) {
	testPath := "test_section/006_short_underline"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection007(t *testing.T) {
	testPath := "test_section/007_title_combining_chars"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection008(t *testing.T) {
	testPath := "test_section/008_title_overline"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection009(t *testing.T) {
	testPath := "test_section/009_inset_title_with_overline"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection010(t *testing.T) {
	testPath := "test_section/010_inset_title_missing_underline"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection011(t *testing.T) {
	testPath := "test_section/011_inset_title_missing_underline_with_blankline"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	equal(t, items, test.expectItems())
}

func TestLexSection012(t *testing.T) {
	testPath := "test_section/012_inset_title_missing_underline_and_para"
	test := LoadTest(testPath)
	items := lexTest(t, test)
	// spd.Dump(items, test.expectItems())
	equal(t, items, test.expectItems())
}
