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

func checkAdornmentNode(pAdorn *AdornmentNode, expect interface{}, testName string) (errors []error) {
	var pNodeName string
	var pFieldName, eFieldName string
	var pFieldVal, eFieldVal interface{}
	var pFieldType, eFieldType reflect.Type

	gError := func() {
		temp := "Got: %s.%s = %#v (%s),\n\tExpect: %s.%s = %#v (%s)\n"
		errors = append(errors, fmt.Errorf(temp, pNodeName, pFieldName, pFieldVal,
			pFieldType, testName, eFieldName, eFieldVal, eFieldType))
	}

	pAdornVal := reflect.ValueOf(pAdorn).Elem()

	for i := 0; i < pAdornVal.NumField(); i++ {

		// Actual parsed metadata
		pNodeName = pAdornVal.Type().Name()
		pFieldName = pAdornVal.Type().Field(i).Name
		pFieldVal = pAdornVal.Field(i).Interface()
		pFieldType = pAdornVal.Type().Field(i).Type

		// Expected parser metadata
		eFieldName = strings.ToLower(string(pFieldName[0])) + pFieldName[1:]
		eFieldVal = expect.(map[string]interface{})[eFieldName]
		eFieldType = reflect.TypeOf(eFieldVal)

		if eFieldVal == nil {
			errors = append(errors,
				fmt.Errorf("\"%s\" does not contain field \"%s\".\n", pAdornVal,
					eFieldName))
			continue
		}

		switch pFieldName {
		case "Type":
			if pFieldVal.(NodeType).String() != eFieldVal {
				gError()
			}
		case "Rune":
			if string(pFieldVal.(rune)) != eFieldVal {
				gError()
			}
		case "Length":
			if float64(pFieldVal.(int)) != eFieldVal {
				gError()
			}
		case "Line":
			if float64(pFieldVal.(Line)) != eFieldVal {
				gError()
			}
		case "StartPosition":
			if float64(pFieldVal.(StartPosition)) != eFieldVal {
				gError()
			}
		default:
			if pFieldVal != eFieldVal {
				gError()
			}
		}

	}
	return
}

func checkSectionNode(pSectionNode *SectionNode, expect interface{}, testName string) (errors []error) {
	var pNodeName string
	var pFieldName, eFieldName string
	var pFieldVal, eFieldVal interface{}
	var pFieldType, eFieldType reflect.Type

	gError := func() {
		temp := "Got: %s.%s = %#v (%s),\n\tExpect: %s.%s = %#v (%s)\n"
		errors = append(errors, fmt.Errorf(temp, pNodeName, pFieldName, pFieldVal,
			pFieldType, testName, eFieldName, eFieldVal, eFieldType))
	}

	pSectionNodeVal := reflect.ValueOf(pSectionNode).Elem()

	for i := 0; i < pSectionNodeVal.NumField(); i++ {

		// Actual parser metadata
		pNodeName = pSectionNodeVal.Type().Name()
		pFieldName = pSectionNodeVal.Type().Field(i).Name
		pFieldVal = pSectionNodeVal.Field(i).Interface()
		pFieldType = pSectionNodeVal.Type().Field(i).Type

		// Expected parser metadata
		eFieldName = strings.ToLower(string(pFieldName[0])) + pFieldName[1:]
		eFieldVal = expect.(map[string]interface{})[eFieldName]
		eFieldType = reflect.TypeOf(eFieldVal)

		// SectionNode.OverLine can be null, if not then we have a missing property
		if eFieldVal == nil && eFieldName == "overLine" {
			continue
		} else if eFieldVal == nil {
			errors = append(errors,
				fmt.Errorf("\"%s\" property does not exist in %s.expectTree!\n",
					pFieldName, testName))
			continue
		}

		switch pFieldName {
		case "Type":
			if pFieldVal.(NodeType).String() != eFieldVal {
				gError()
			}
		case "Level", "Length":
			if float64(pFieldVal.(int)) != eFieldVal {
				gError()
			}
		case "Line":
			if float64(pFieldVal.(Line)) != eFieldVal {
				gError()
			}
		case "StartPosition":
			if float64(pFieldVal.(StartPosition)) != eFieldVal {
				gError()
			}
		case "OverLine", "UnderLine":
			errs := checkAdornmentNode(pFieldVal.(*AdornmentNode), eFieldVal, testName)
			if errs != nil {
				errors = append(errors, errs...)
			}
		case "NodeList":
			errors = append(errors, fmt.Errorf("%s.%s not implemented yet!", pNodeName,
				pFieldName))
		default:
			if pFieldVal != eFieldVal {
				gError()
			}
		}

	}

	return
}

// checkParseNodes is a recursive function that compares the resulting nodes (pNodes) from the
// parser with the expected output from the testdata (eNodes).
func checkParseNodes(pNodes *NodeList, eNodes []interface{}, testName string) (errors []error) {
	// spd.Dump(pNodes)
	// os.Exit(1)
	for pNum, pNode := range *pNodes {
		switch node := pNode.(type) {
		case *SectionNode:
			errs := checkSectionNode(node, eNodes[pNum], testName)
			if errs != nil {
				errors = append(errors, errs...)
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
	errors := checkParseNodes(tree.Nodes, nodeList, testName)
	if errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}
