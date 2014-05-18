// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import (
	"encoding/json"
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

type checkNode struct {
	t          *testing.T
	testName   string
	pNodeName  string
	pFieldName string
	pFieldVal  interface{}
	pFieldType reflect.Type
	eFieldName string
	eFieldVal  interface{}
	eFieldType reflect.Type
}

func (c *checkNode) error(args ...interface{}) {
	c.t.Error(args...)
}

func (c *checkNode) errorf(format string, args ...interface{}) {
	c.t.Errorf(format, args...)
}

func (c *checkNode) dError() {
	if c.pFieldName == "Rune" {
		c.t.Errorf("Got: %s.%s = %#v (%#v) (%s),\n\tExpect: %s.%s = %#v (%s)\n", c.pNodeName,
			c.pFieldName, c.pFieldVal, string(c.pFieldVal.(int32)), c.pFieldType,
			c.testName, c.eFieldName, c.eFieldVal, c.eFieldType)
		return
	}
	c.t.Errorf("Got: %s.%s = %#v (%s),\n\tExpect: %s.%s = %#v (%s)\n", c.pNodeName,
		c.pFieldName, c.pFieldVal, c.pFieldType, c.testName, c.eFieldName, c.eFieldVal,
		c.eFieldType)
}

func (c *checkNode) updateState(pVal reflect.Value, eVal interface{}, field int) {
	// Actual parsed metadata
	c.pNodeName = pVal.Type().Name()
	c.pFieldName = pVal.Type().Field(field).Name
	c.pFieldVal = pVal.Field(field).Interface()
	c.pFieldType = pVal.Type().Field(field).Type
	// Expected parser metadata
	c.eFieldName = strings.ToLower(string(c.pFieldName[0])) + c.pFieldName[1:]
	c.eFieldVal = eVal.(map[string]interface{})[c.eFieldName]
	c.eFieldType = reflect.TypeOf(c.eFieldVal)
}

func (c *checkNode) checkSectionNode(pSectionNode *SectionNode,
	expect interface{}, testName string) (errors []error) {

	pSectionNodeVal := reflect.ValueOf(pSectionNode).Elem()
	for i := 0; i < pSectionNodeVal.NumField(); i++ {

		c.updateState(pSectionNodeVal, expect, i)

		// SectionNode.OverLine can be null, if not then we have a missing property
		if c.eFieldVal == nil && c.eFieldName == "overLine" {
			continue
		} else if c.eFieldVal == nil {
			c.errorf("\"%s\" property does not exist in %s.expectTree!\n", c.pFieldName,
				testName)
		}

		switch c.pFieldName {
		case "Type":
			if c.pFieldVal.(NodeType).String() != c.eFieldVal {
				c.dError()
			}
		case "Level", "Length":
			if float64(c.pFieldVal.(int)) != c.eFieldVal {
				c.dError()
			}
		case "Line":
			if float64(c.pFieldVal.(Line)) != c.eFieldVal {
				c.dError()
			}
		case "StartPosition":
			if float64(c.pFieldVal.(StartPosition)) != c.eFieldVal {
				c.dError()
			}
		case "OverLine", "UnderLine":
			c.checkAdornmentNode(c.pFieldVal.(*AdornmentNode), c.eFieldVal, c.testName)
		case "NodeList":
			c.error("NodeList not implemented!")
		default:
			if c.pFieldVal != c.eFieldVal {
				c.dError()
			}
		}
	}
	return
}

func (c *checkNode) checkAdornmentNode(pAdorn *AdornmentNode, expect interface{}, testName string) {
	pAdornVal := reflect.ValueOf(pAdorn).Elem()
	for i := 0; i < pAdornVal.NumField(); i++ {

		c.updateState(pAdornVal, expect, i)

		if c.eFieldVal == nil {
			c.error("\"%s\" does not contain field \"%s\".\n", pAdornVal, c.eFieldName)
			continue
		}

		switch c.pFieldName {
		case "Type":
			if c.pFieldVal.(NodeType).String() != c.eFieldVal {
				c.dError()
			}
		case "Rune":
			if string(c.pFieldVal.(rune)) != c.eFieldVal {
				c.dError()
			}
		case "Length":
			if float64(c.pFieldVal.(int)) != c.eFieldVal {
				c.dError()
			}
		case "Line":
			if float64(c.pFieldVal.(Line)) != c.eFieldVal {
				c.dError()
			}
		case "StartPosition":
			if float64(c.pFieldVal.(StartPosition)) != c.eFieldVal {
				c.dError()
			}
		default:
			if c.pFieldVal != c.eFieldVal {
				c.dError()
			}
		}

	}
	return
}

// checkParseNodes is a recursive function that compares the resulting nodes (pNodes) from the
// parser with the expected output from the testdata (eNodes).
func checkParseNodes(t *testing.T, pNodes *NodeList, eNodes []interface{}, testName string) (errors []error) {
	state := &checkNode{t: t, testName: testName}
	for pNum, pNode := range *pNodes {
		switch node := pNode.(type) {
		case *SectionNode:
			state.checkSectionNode(node, eNodes[pNum], testName)
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
	errors := checkParseNodes(t, tree.Nodes, nodeList, testName)
	if errors != nil {
		for _, err := range errors {
			t.Error(err)
		}
	}
}
