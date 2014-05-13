// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import (
	"encoding/json"
	"fmt"
	"github.com/demizer/go-elog"
	"reflect"
	"strings"
	"testing"
)

var lexParseTests LexTests

func TestSectionLevelsAdd(t *testing.T) {
	var p sectionLevels
	lvl := p.Add('=', true, 5)
	if lvl != 1 {
		t.Errorf("Improper level on first add, Got level: %d, expected: %d", lvl, 1)
	}
}

func TestSectionLevelsString(t *testing.T) {
	var p sectionLevels
	p.Add('=', true, 5)
	p.Add('-', true, 8)
	p.Add('~', false, 6)
	out := p.String()
	expect := "level: 1, rune: '=', overline: true, length: 5\nlevel: 2, rune: '-', " +
		"overline: true, length: 8\nlevel: 3, rune: '~', overline: false, length: 6\n"
	if out != expect {
		t.Errorf("String output mismatch!\nExpect:\n\n\t%q,\nGot:\n\n\t%q\n", expect, out)
	}
}

func TestSectionLevelsFind(t *testing.T) {
	var p sectionLevels
	p.Add('=', true, 5)
	p.Add('-', true, 8)
	p.Add('~', false, 6)
	lvl := p.Find('-')
	if lvl == -1 {
		t.Errorf("Level not found!\nExpect:\n\n\t%d\nGot:\n\n\t%t\n", 2, lvl)
	}
	if lvl != 2 {
		t.Errorf("Level not correct!\nExpect:\n\n\t%d\nGot:\n\n\t%d\n", 2, lvl)
	}
}

func TestSectionLevelsFindNoResult(t *testing.T) {
	var p sectionLevels
	lvl := p.Find('-')
	if lvl > 0 {
		t.Errorf("Should not find any levels!\nExpect:\n\n\t%d\nGot:\n\n\t%d\n", -1, lvl)
	}
}

func TestSectionLevelsLevelEmpty(t *testing.T) {
	var p sectionLevels
	lvl := p.Level()
	expect := 0
	if lvl != expect {
		t.Errorf("Empty sectionLevels should return \"%d\"!\nExpect:\t%d\nGot:\t%d\n",
			expect, expect, lvl)
	}

}

func TestSectionLevelsLevel(t *testing.T) {
	var p sectionLevels
	p.Add('=', true, 5)
	p.Add('-', true, 8)
	p.Add('~', false, 6)
	lvl := p.Level()
	if lvl != 3 {
		t.Errorf("Level() returned incorrect level!\nExpect:\n\n\t%d\nGot:\n\n\t%d\n", 3, lvl)
	}
}

func parseTest(t *testing.T, testName string) (tree *Tree, err error) {
	if lexParseTests == nil {
		lexParseTests, err = ParseTestData("../testdata/test_lex_sections.dat")
		if err != nil {
			t.Fatal(err)
		}
	}
	test := lexParseTests.SearchByName(testName)
	if test != nil {
		log.Debugf("Test Name: \t%s\n", test.name)
		log.Debugf("Description: \t%s\n", test.description)
		log.Debugf("Test Input:\n-----------\n%s\n----------\n", test.data)
		tree, err = Parse(test.name, test.data)
		if err != nil {
			t.Fatal(err)
		}
	}
	return
}

// compareNodes is a recursive function that compares the resulting nodes (pNodes) from the parser
// with the expected output from the testdata (eNodes).
func compareNodes(pNodes *NodeList, eNodes []interface{}, testName string) (errors []error) {
	for pNum, pNode := range *pNodes {
		pVal := reflect.ValueOf(pNode).Elem()
		pType := pVal.Type()
		for i := 0; i < pVal.NumField(); i++ {
			pStructName := pType.Name()
			pFieldName := pType.Field(i).Name
			eName := strings.ToLower(string(pFieldName[0])) + pFieldName[1:]
			gVal := pVal.Field(i)
			eVal := eNodes[pNum].(map[string]interface{})[eName]
			eType := reflect.TypeOf(eVal)

			// SectionNode.OverLine and SectionNode.UnderLine can be null
			if eVal == nil && eName != "overLine" && eName != "underLine" {
				errors = append(errors, fmt.Errorf("\"%s\" property does not exist "+
				"in %s.expectTree!\n", pFieldName, testName))
				continue
			}

			var val interface{}
			var cErr error

			match := false

			// log.Println(gVal.Kind())

			switch gVal.Kind() {
			case reflect.String:
				if gVal.String() != eVal || gVal.Kind() != reflect.String ||
					eType.Kind() != reflect.String {
					cErr = fmt.Errorf("Got Type: %s.%s = %#v (%s),\n\t"+
						"Expect Type: %s.%s = %#v (%s)\n", pStructName,
						pFieldName, gVal.String(), gVal.Type(), testName, eName,
						eVal, eType.Kind())
					break
				}
				match = true
			case reflect.Int:
				// Check for the "Type" field name, this requires a conversion from
				// string to int.
				if pFieldName == "Type" {
					gNodeType := NodeTypeFromString(eVal.(string))
					if gVal.Int() == int64(gNodeType) {
						match = true
						break
					} else {
						cErr = fmt.Errorf(
							"Got: %s.%s = %#v,\n\tExpect: %s.%s = %#v\n",
							pStructName, pFieldName, val, testName, eName,
							eVal)
						break
					}
				}
				// Check for matching types betwen the parsed value and expected
				// value.
				val = gVal.Int()
				if eType.Kind() != reflect.Float64 || gVal.Kind() != reflect.Int ||
					eType.Kind() != reflect.Float64 {
					cErr = fmt.Errorf("Got Type: %s.%s = %#v (%s),\n\t"+
						"Expect Type: %s.%s = %f (%s)\n", pStructName,
						pFieldName, gVal.Int(), gVal.Type(), testName, eName,
						eVal, eType.Kind())
					break
				}
				// Finally, check the actual values
				match = int(val.(int64)) == int(eVal.(float64))
			case reflect.Ptr:
				// A pointer in the struct is most likely another struct such as
				// an AdornmentNode.
				cErr = fmt.Errorf("reflect.Ptr for %s.%s not implemented yet!",
					pStructName, pFieldName)
			case reflect.Slice:
				// A silce in the struct is probably a NodeList.
				cErr = fmt.Errorf("reflect.Slice for %s.%s not implemented yet!",
					pStructName, pFieldName)
			}

			if match == false {
				errors = append(errors, cErr)
			}
		}
	}
	return
}

func TestParseSectionTitlePara(t *testing.T) {
	testName := "SectionTitlePara"
	tree, err := parseTest(t, testName)
	if err != nil {
		t.Error(err)
	}
	test := lexParseTests.SearchByName(testName)
	var nodeList []interface{}
	err = json.Unmarshal([]byte(test.expectTree), &nodeList)
	if err != nil {
		t.Error(err)
	}
	errors := compareNodes(tree.Nodes, nodeList, testName)
	if errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}
