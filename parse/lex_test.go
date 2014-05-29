// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"code.google.com/p/go.text/unicode/norm"
	"github.com/demizer/go-elog"
	"reflect"
	"testing"
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
	var pFieldType, eFieldType reflect.Type
	var pFieldVal, eFieldVal reflect.Value
	var pFieldValStruct reflect.StructField

	dError := func() {
		t.Errorf("Got: %s = %#v (%T) (Id: %d)\n\tExpect: %s = %#v (%T)\n", pFieldName,
			pFieldVal.Interface(), pFieldType.Name(), id, eFieldName,
			eFieldVal.Interface(), eFieldType.Name())
	}

	for eNum, eItem := range expectItems {
		eVal := reflect.ValueOf(eItem)
		pVal := reflect.ValueOf(items[eNum])
		id = int(pVal.FieldByName("Id").Interface().(Id))
		for x := 0; x < eVal.NumField(); x++ {
			eFieldVal = eVal.Field(x)
			eFieldType = eFieldVal.Type()
			eFieldName = eVal.Type().Field(x).Name
			pFieldVal = pVal.FieldByName(eFieldName)
			pFieldType = pFieldVal.Type()
			pFieldValStruct, found = pVal.Type().FieldByName(eFieldName)
			pFieldName = pFieldValStruct.Name
			if !found {
				t.Errorf("Parsed item (Id: %d) does not contain field %q\n", id,
					eFieldName)
				continue
			} else if eFieldName == "Text" {
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
	// spd.Dump(items, test.expectItems())
	equal(t, items, test.expectItems())
}
