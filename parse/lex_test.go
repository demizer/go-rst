// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"encoding/json"
	"github.com/demizer/go-elog"
	"github.com/demizer/go-spew/spew"
	"reflect"
	"testing"
)

var (
	tEOF = item{Type: itemEOF, StartPosition: 0, Text: ""}
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
		log.Debugf("Test Name: %s\n", test.name)
		log.Debugf("Description: %s\n", test.description)
		log.Debugf("Test Input:\n-----------\n%s\n----------\n", test.data)
		items := collect(test)
		return items
	} else {
		t.Fatalf("%s could not be found.\n", testName)
	}
	return nil
}

// Test equality between items and expected items from unmarshalled json data, field by field.
// Returns error in case of error during json unmarshalling, or mismatch between items and the
// expected output.
func equal(t *testing.T, items []item, testName string) {
	test := lexSectionTests.SearchByName(testName)
	var exp []item
	err := json.Unmarshal([]byte(test.items), &exp)
	if err != nil {
		t.Fatal("JSON error: ", err)
	}

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

	for eNum, eItem := range exp {
		eVal := reflect.ValueOf(eItem)
		pVal := reflect.ValueOf(items[eNum])
		id = pVal.FieldByName("Id").Interface().(int)
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
			}
			if pFieldVal.Interface() != eFieldVal.Interface() {
				// log.Debugln(eFieldType.Name())
				dError()
			}
		}
	}

	return
}

func TestLexSectionTitlePara(t *testing.T) {
	testName := "SectionTitlePara"
	items := lexSectionTest(t, testName)
	equal(t, items, testName)
}

func TestLexSectionTitleParaNoBlankline(t *testing.T) {
	testName := "SectionTitleParaNoBlankLine"
	items := lexSectionTest(t, testName)
	equal(t, items, testName)
}

func TestLexSectionParaHeadPara(t *testing.T) {
	testName := "SectionParaHeadPara"
	items := lexSectionTest(t, testName)
	equal(t, items, testName)
}

func TestLexSectionUnexpectedTitles(t *testing.T) {
	testName := "SectionUnexpectedTitles"
	items := lexSectionTest(t, testName)
	equal(t, items, testName)
}
