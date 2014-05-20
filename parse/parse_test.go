// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import (
	"github.com/demizer/go-elog"
	"reflect"
	"strings"
	"testing"
)

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
	id         int
}

func (c *checkNode) error(args ...interface{}) {
	c.t.Error(args...)
}

func (c *checkNode) errorf(format string, args ...interface{}) {
	c.t.Errorf(format, args...)
}

func (c *checkNode) dError() {
	if c.pFieldName == "Rune" {
		c.t.Errorf("Got: %s.%s = %#v (%v) (%v) (Id: %d),\n\tExpect: %s.%s = %#v (%v)\n",
			c.pNodeName, c.pFieldName, c.pFieldVal, string(c.pFieldVal.(int32)),
			c.pFieldType, c.id, "#parse-tree", c.eFieldName, c.eFieldVal, c.eFieldType)
		return
	}
	c.t.Errorf("Got: %s.%s = %#v (%v) (Id: %d),\n\tExpect: %s.%s = %#v (%v)\n",
		c.pNodeName, c.pFieldName, c.pFieldVal, c.pFieldType, c.id, "#parse-tree",
		c.eFieldName, c.eFieldVal, c.eFieldType)
}

func (c *checkNode) updateState(pVal reflect.Value, eVal interface{}, field int) bool {
	// Actual parsed metadata
	c.pNodeName = pVal.Type().Name()
	c.pFieldName = pVal.Type().Field(field).Name
	c.pFieldVal = pVal.Field(field).Interface()
	c.pFieldType = pVal.Type().Field(field).Type
	c.id = pVal.FieldByName("Id").Interface().(int)

	// Expected parser metadata
	c.eFieldName = strings.ToLower(string(c.pFieldName[0])) + c.pFieldName[1:]
	c.eFieldVal = eVal.(map[string]interface{})[c.eFieldName]
	c.eFieldType = reflect.TypeOf(c.eFieldVal)

	pField := pVal.Field(field)
	if pField.Kind() == reflect.Ptr && pField.IsNil() {
		// Overline adornment nodes can be null
		if c.pFieldName == "OverLine" {
			return false
		} else {
			c.dError()
			return false
		}
	} else if c.pFieldVal == nil {
		c.dError()
		return false
	}

	return true
}

func (c *checkNode) checkFields(expect interface{}, pNode Node) {
	pVal := reflect.ValueOf(pNode).Elem()
	for i := 0; i < pVal.NumField(); i++ {

		if c.updateState(pVal, expect, i) == false {
			continue
		}

		switch c.pFieldName {
		case "Type":
			if c.pFieldVal.(NodeType).String() != c.eFieldVal {
				c.dError()
			}
		case "Id", "Level", "Length":
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
			c.checkFields(c.eFieldVal, c.pFieldVal.(Node))
		case "NodeList":
			for num, node := range c.pFieldVal.(NodeList) {
				// A little hackery to make our recursion easy
				eFieldVal := c.eFieldVal
				c.checkFields(eFieldVal.([]interface{})[num], node.(Node))
				c.eFieldVal = eFieldVal
			}
		case "Rune":
			if string(c.pFieldVal.(rune)) != c.eFieldVal {
				c.dError()
			}
		default:
			if c.pFieldVal != c.eFieldVal {
				c.dError()
			}
		}
	}

}

func checkParseNodes(t *testing.T, eTree []interface{}, pNodes []Node, testName string) {
	state := &checkNode{t: t, testName: testName}
	if len(eTree) != len(pNodes) {
		t.Fatal("len(eTree) != len(pNodes)")
	}
	for eNum, eNode := range eTree {
		state.checkFields(eNode, pNodes[eNum])
	}
	return
}

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

func parseTest(t *testing.T, lexTest *LexTest) (tree *Tree) {
	var errs []error
	log.Debugf("Test Name: %s\n", lexTest.name)
	log.Debugf("Description: %s\n", lexTest.description)
	log.Debugf("Test Input:\n-----------\n%s\n----------\n", lexTest.data)
	tree, errs = Parse(lexTest.name, lexTest.data)
	if errs != nil {
		for _, err := range errs {
			t.Error(err)
		}
	}
	return
}

func TestParseSectionTitlePara(t *testing.T) {
	testName := "SectionTitlePara"
	test := lexTests.testByName(testName)
	pTree := parseTest(t, test)
	eNodes := test.expectJson()
	checkParseNodes(t, eNodes, *pTree.Nodes, testName)
}

func TestParseSectionTitleParaNoBlankLine(t *testing.T) {
	testName := "SectionTitleParaNoBlankLine"
	test := lexTests.testByName(testName)
	pTree := parseTest(t, test)
	eNodes := test.expectJson()
	checkParseNodes(t, eNodes, *pTree.Nodes, testName)
}

func TestParseSectionParaHeadPara(t *testing.T) {
	testName := "SectionParaHeadPara"
	test := lexTests.testByName(testName)
	pTree := parseTest(t, test)
	eNodes := test.expectJson()
	checkParseNodes(t, eNodes, *pTree.Nodes, testName)
}

func TestParseSectionLevelTest1(t *testing.T) {
	testName := "SectionLevelTest1"
	test := lexTests.testByName(testName)
	pTree := parseTest(t, test)
	// spd.Dump(pTree)
	eNodes := test.expectJson()
	checkParseNodes(t, eNodes, *pTree.Nodes, testName)
}

func TestParseSectionUnexpectedTitles(t *testing.T) {
	testName := "SectionUnexpectedTitles"
	test := lexTests.testByName(testName)
	pTree := parseTest(t, test)
	eNodes := test.expectJson()
	checkParseNodes(t, eNodes, *pTree.Nodes, testName)
}
