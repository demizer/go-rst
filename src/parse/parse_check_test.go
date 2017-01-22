package parse

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
	"time"

	"golang.org/x/text/unicode/norm"
)

type checkNode struct {
	t          *testing.T
	parsedNode interface{}
	expectNode interface{}
	testPath   string
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

func (c *checkNode) checkField(parsedValue interface{}, expectedValue interface{}) {
	if parsedValue != expectedValue {
		c.dError()
	}
}

func (c *checkNode) dError() {
	var got, exp string

	switch c.pFieldVal.(type) {
	case NodeType:
		got = c.pFieldVal.(NodeType).String()
		exp = c.eFieldVal.(string)
	case StartPosition:
		got = c.pFieldVal.(StartPosition).String()
		exp = strconv.Itoa(int(c.eFieldVal.(float64)))
	case Line:
		got = c.pFieldVal.(Line).String()
		exp = strconv.Itoa(int(c.eFieldVal.(float64)))
	case systemMessageLevel:
		pNum := int(c.pFieldVal.(systemMessageLevel))
		pNumStr := " (" + strconv.Itoa(pNum) + ")"
		got = c.pFieldVal.(systemMessageLevel).String() + pNumStr
		smsLvl := int(systemMessageLevelFromString(c.eFieldVal.(string)))
		eNumStr := " (" + strconv.Itoa(smsLvl) + ")"
		exp = c.eFieldVal.(string) + eNumStr
	case string:
		got = c.pFieldVal.(string)
		exp = c.eFieldVal.(string)
	case int:
		got = strconv.Itoa(c.pFieldVal.(int))
		exp = strconv.Itoa(int(c.eFieldVal.(float64)))
	case rune:
		got = string(c.pFieldVal.(rune))
		exp = string(c.eFieldVal.(rune))
	}
	lPos := strconv.FormatFloat(c.expectNode.(map[string]interface{})["line"].(float64), 'f', -1, 64)
	sPos := reflect.Indirect(reflect.ValueOf(c.parsedNode)).FieldByName("StartPosition").Int()
	txt := reflect.Indirect(reflect.ValueOf(c.parsedNode)).FieldByName("Text").String()
	c.t.Errorf("[node  text:%q line:%s startpos:%d] Got %s = %q -- Expect %s = %q", txt, lPos, sPos, c.pFieldName, got, c.eFieldName, exp)
}

func (c *checkNode) updateState(eKey string, eVal interface{}, pVal reflect.Value, eNode interface{}, pNode interface{}) bool {
	c.expectNode = eNode
	c.parsedNode = pNode

	// Expected parser metadata
	c.eFieldName = eKey
	c.eFieldVal = eVal
	c.eFieldType = reflect.TypeOf(c.eFieldVal)

	// Actual parsed metadata
	c.pNodeName = pVal.Type().Name()
	c.pFieldName = strings.ToUpper(string(c.eFieldName[0]))
	c.pFieldName += c.eFieldName[1:]

	if !pVal.FieldByName(c.pFieldName).IsValid() {
		return false
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

// checkMatchingFields compares the expected node output retrieved from the nodes.json file and the actual parser NodeList
// output. Returns an error if a mismatch is found.
func (c *checkNode) checkMatchingFields(eNodes interface{}, pNode Node) error {
	if eNodes == nil || pNode == nil {
		panic("arguments must not be nil!")
	}
	// If the value is missing in eNodes and nil in pNode than we can exclude it.
	eFields := eNodes.(map[string]interface{})
	pNodeVal := reflect.Indirect(reflect.ValueOf(pNode))
	// Check expected node to parsed node
	for eName := range eFields {
		var sfName string
		sfName = strings.ToUpper(eName[0:1]) + eName[1:]
		if _, in := pNodeVal.Type().FieldByName(sfName); !in {
			return fmt.Errorf("Parse Node missing field %q:\nParseNode:\n%sExpectNode:\n%s",
				sfName, spd.Sdump(pNode), spd.Sdump(eFields))
		}
	}
	// Compare pNode against eNodes
	for i := 0; i < pNodeVal.NumField(); i++ {
		pName := pNodeVal.Type().Field(i).Tag.Get("json")
		if pName == "" {
			tlog(fmt.Sprintf("Check struct tags! pName = %s", pName))
			os.Exit(1)
		}
		pVal := pNodeVal.Field(i).Interface()
		eFields := eNodes.(map[string]interface{})
		switch pName {
		case "indentLength":
			// Some title nodes aren't indented.
			if pVal == 0 {
				continue
			}
		case "startPosition":
			// Most nodes begin at position one in the line, therefore we can ignore them if it hasn't been
			// specified in the expected nodes.
			if pVal.(StartPosition) == 0 || pVal.(StartPosition) == 1 {
				continue
			}
		case "line":
			// zero, then we ignore it.  systemMessage literal block nodes have no line position.
			if pVal.(Line).LineNumber() == 0 {
				continue
			}
		case "overLine":
			// Some sections don't have overlines
			if eFields[pName] == nil && pVal.(*AdornmentNode) == nil {
				continue
			}
		case "nodeList":
			// Some Nodes don't have child nodes.
			if eFields[pName] == nil && pVal.(NodeList) == nil {
				continue
			}
		case "text":
			// Some Nodes don't have text.
			if eFields[pName] == nil && pVal.(string) == "" {
				continue
			}
		case "length":
			if eFields[pName] == nil && pVal == 0 {
				continue
			}
		}
		eNode := eNodes.(map[string]interface{})
		if eNode[pName] == nil {
			return fmt.Errorf("NodeType: %q Missing field %q -- Parser got: %q == %v -- ExpectNode:\n%s\n",
				eNode["type"], pName, pName, pVal, spd.Sdump(eNode))
		}
	}
	return nil
}

// checkFields is a recursive function that compares the expected node output to the parser output comparing the two objects
// field by field. eNodes is unmarshaled json input and pNode is the parser node to check.
func (c *checkNode) checkFields(eNodes interface{}, pNode Node) error {
	if eNodes == nil || pNode == nil {
		panic("arguments cannot be nil!")
	}
	if err := c.checkMatchingFields(eNodes, pNode); err != nil {
		c.t.Error(err)
	}
	for eKey, eVal := range eNodes.(map[string]interface{}) {
		pVal := reflect.Indirect(reflect.ValueOf(pNode))
		if !c.updateState(eKey, eVal, pVal, eNodes, pNode) {
			continue
		}
		switch c.eFieldName {
		case "text":
			c.checkField(c.pFieldVal.(string), norm.NFC.String(c.eFieldVal.(string)))
		case "type":
			c.checkField(c.pFieldVal.(NodeType).String(), c.eFieldVal)
		case "messageType":
			c.checkField(c.pFieldVal.(parserMessage).String(), c.eFieldVal)
		case "level", "length", "indentLength":
			c.checkField(float64(c.pFieldVal.(int)), c.eFieldVal)
		case "line":
			c.checkField(float64(c.pFieldVal.(Line)), c.eFieldVal)
		case "startPosition":
			c.checkField(float64(c.pFieldVal.(StartPosition)), c.eFieldVal)
		case "indent", "overLine", "title", "underLine":
			if cerr := c.checkFields(c.eFieldVal, c.pFieldVal.(Node)); cerr != nil {
				return cerr
			}
		case "term", "definition":
			if cerr := c.checkFields(c.eFieldVal, c.pFieldVal.(Node)); cerr != nil {
				return cerr
			}
		case "nodeList":
			len1 := len(c.eFieldVal.([]interface{}))
			len2 := len(c.pFieldVal.(NodeList))
			if len1 != len2 {
				return fmt.Errorf("Expected NodeList values (len=%d) and parsed NodeList values (len=%d) "+
					"do not match!", len1, len2)
			}
			for num, node := range c.eFieldVal.([]interface{}) {
				// Store and reset the parser value, otherwise a panic will occur on the next iteration
				pFieldVal := c.pFieldVal
				if cerr := c.checkFields(node, c.pFieldVal.(NodeList)[num]); cerr != nil {
					return cerr
				}
				c.pFieldVal = pFieldVal
			}
		case "rune":
			c.checkField(string(c.pFieldVal.(rune)), c.eFieldVal)
		case "severity":
			c.checkField(c.pFieldVal.(systemMessageLevel).String(), c.eFieldVal)
		case "bullet":
			c.checkField(c.pFieldVal.(string), c.eFieldVal.(string))
		case "enumType":
			c.checkField(c.pFieldVal.(EnumListType).String(), c.eFieldVal)
		case "affix":
			c.checkField(c.pFieldVal.(EnumAffixType).String(), c.eFieldVal)
		default:
			c.t.Errorf("Type %q case is not implemented in checkFields!", c.eFieldName)
		}
	}
	return nil
}

// checkParseNodes compares the expected parser output (*_nodes.json) against the actual parser output node by node.
func checkParseNodes(t *testing.T, eTree []interface{}, pNodes *NodeList, testPath string) {

	state := &checkNode{t: t, testPath: testPath}

	failTest := func(err error) {
		// Give all other output time to print
		time.Sleep(time.Second / 2)
		tlog(fmt.Sprintf("\nFAIL: %s\n", err.Error()))
		tlog("-----------------------------------------------------------------------------")
		tlog("Parse Nodes")
		tlog("-----------------------------------------------------------------------------")
		pnj, err := json.MarshalIndent(pNodes, "", "    ")
		if err != nil {
			tlog(fmt.Sprintf("ERROR: Could not marshal json! Error=%q", err.Error()))
			t.Fail()
		}
		tlog(string(pnj))
		// tlog(spd.Sdump(pNodes))
		tlog("-----------------------------------------------------------------------------")
		tlog("Expected Nodes")
		tlog("-----------------------------------------------------------------------------")
		enj, err := json.MarshalIndent(eTree, "", " ")
		if err != nil {
			tlog(fmt.Sprintf("ERROR: Could not marshal json! Error=%q", err.Error()))
			t.Fail()
		}
		tlog(string(enj))
		// tlog(spd.Sdump(eTree))
		t.FailNow()
	}

	if len(*pNodes) != len(eTree) {
		failTest(errors.New("The number of parsed nodes does not match expected nodes!"))
	}

	for eNum, eNode := range eTree {
		if cerr := state.checkFields(eNode, (*pNodes)[eNum]); cerr != nil {
			failTest(cerr)
		}
	}

	return
}
