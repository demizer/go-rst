// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import (
	"fmt"
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

func (c *checkNode) updateState(eKey string, eVal interface{}, pVal reflect.Value) bool {
	// Expected parser metadata
	c.eFieldName = eKey
	c.eFieldVal = eVal
	c.eFieldType = reflect.TypeOf(c.eFieldVal)

	// Actual parsed metadata
	c.pNodeName = pVal.Type().Name()
	c.pFieldName = strings.ToUpper(string(c.eFieldName[0])) + c.eFieldName[1:]
	if !pVal.FieldByName(c.pFieldName).IsValid() {
		panic(fmt.Errorf("Missing field: %s.%s\n", c.pNodeName, c.pFieldName))
	}
	c.pFieldVal = pVal.FieldByName(c.pFieldName).Interface()
	c.pFieldType = pVal.FieldByName(c.pFieldName).Type()

	// Overline adornment nodes can be null
	if c.eFieldName == "overLine" && c.eFieldVal == nil {
		return false
	} else if c.eFieldVal == nil {
		c.dError()
		return false
	}

	return true
}

func (c *checkNode) checkFields(eNodes interface{}, pNode Node) {
	c.id = int(eNodes.(map[string]interface{})["id"].(float64))
	for eKey, eVal := range eNodes.(map[string]interface{}) {
		pVal := reflect.Indirect(reflect.ValueOf(pNode))
		if c.updateState(eKey, eVal, pVal) == false {
			continue
		}
		switch c.eFieldName {
		case "type":
			if c.eFieldVal != c.pFieldVal.(NodeType).String() {
				c.dError()
			}
		case "id", "level", "length":
			if c.eFieldVal != float64(c.pFieldVal.(int)) {
				c.dError()
			}
		case "line":
			if c.eFieldVal != float64(c.pFieldVal.(Line)) {
				c.dError()
			}
		case "startPosition":
			if c.eFieldVal != float64(c.pFieldVal.(StartPosition)) {
				c.dError()
			}
		case "overLine", "underLine":
			c.checkFields(c.eFieldVal, c.pFieldVal.(Node))
		case "nodeList":
			for num, node := range c.eFieldVal.([]interface{}) {
				// Store and reset the parser value, otherwise a panic will occur on
				// the next iteration
				pFieldVal := c.pFieldVal
				c.checkFields(node, c.pFieldVal.(NodeList)[num])
				c.pFieldVal = pFieldVal
			}
		case "rune":
			if c.eFieldVal != string(c.pFieldVal.(rune)) {
				c.dError()
			}
		case "severity":
			if c.eFieldVal != c.pFieldVal.(systemMessageLevel).String() {
				c.dError()
			}
		default:
			if c.eFieldVal != c.pFieldVal {
				c.dError()
			}
		}
	}

}

func checkParseNodes(t *testing.T, eTree []interface{}, pNodes []Node, testName string) {
	state := &checkNode{t: t, testName: testName}
	for eNum, eNode := range eTree {
		state.checkFields(eNode, pNodes[eNum])
	}
	return
}

var testSectionLevel = [...]SectionNode{
	SectionNode{
		OverLine:  &AdornmentNode{Rune: '='},
		UnderLine: &AdornmentNode{Rune: '='},
		Level:     1,
		Length:    5,
	},
	SectionNode{
		OverLine:  &AdornmentNode{Rune: '-'},
		UnderLine: &AdornmentNode{Rune: '-'},
		Level:     2,
		Length:    8,
	},
	SectionNode{
		UnderLine: &AdornmentNode{Rune: '~'},
		Level:     3,
		Length:    6,
	},
}

func TestSectionLevelsAdd(t *testing.T) {
	x := new(sectionLevels)
	v := &testSectionLevel[0]
	x.Add(v)
	if v.Level != 1 {
		t.Errorf("Improper level on first add, Got level: %d, expected: %d", v.Level, 1)
	}
}

func TestSectionLevelsString(t *testing.T) {
	var p sectionLevels
	p.Add(&testSectionLevel[0])
	p.Add(&testSectionLevel[1])
	p.Add(&testSectionLevel[2])
	out := p.String()
	expect := "level: 1, rune: '=', overline: true, length: 5\nlevel: 2, rune: '-', " +
		"overline: true, length: 8\nlevel: 3, rune: '~', overline: false, length: 6\n"
	if out != expect {
		t.Errorf("Expect:\t%q,\nGot:\t%q\n", expect, out)
	}
}

func TestSectionLevelsFind(t *testing.T) {
	var p sectionLevels
	p.Add(&testSectionLevel[0])
	p.Add(&testSectionLevel[1])
	p.Add(&testSectionLevel[2])
	sec := p.FindByRune('-')
	if sec.Level == -1 {
		t.Errorf("Expect:\t%d\nGot:\t%t\n", 2, sec.Level)
	}
	if sec.Level != 2 {
		t.Errorf("Expect:\t%d\nGot:\t%d\n", 2, sec.Level)
	}
}

func TestSectionLevelsFindNoResult(t *testing.T) {
	var p sectionLevels
	p.Add(&testSectionLevel[0])
	sec := p.FindByRune('-')
	if sec != nil {
		t.Error("Expect no return value!")
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
	p.Add(&testSectionLevel[0])
	p.Add(&testSectionLevel[1])
	p.Add(&testSectionLevel[2])
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

func TestParseSectionTitleParagraph(t *testing.T) {
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
