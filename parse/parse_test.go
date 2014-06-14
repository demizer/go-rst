// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"
	"strings"
	"testing"

	"code.google.com/p/go.text/unicode/norm"
	"github.com/demizer/go-elog"
)

func init() { SetDebug() }

// SetDebug is typically called from the init() function in a test file.
// SetDebug parses debug flags passed to the test binary and also sets the
// template for logging output.
func SetDebug() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Enable debug output.")
	flag.Parse()

	if debug {
		log.SetLevel(log.LEVEL_DEBUG)
	}

	log.SetTemplate("{{if .Date}}{{.Date}} {{end}}" +
		"{{if .Prefix}}{{.Prefix}} {{end}}" +
		"{{if .LogLabel}}{{.LogLabel}} {{end}}" +
		"{{if .FileName}}{{.FileName}}: {{end}}" +
		"{{if .FunctionName}}{{.FunctionName}}{{end}}" +
		"{{if .LineNumber}}#{{.LineNumber}}: {{end}}" +
		"{{if .Text}}{{.Text}}{{end}}")

	log.SetFlags(log.Lansi | log.LnoPrefix | log.LfunctionName |
		log.LlineNumber)
}

// Contains a single test with data loaded from test files in the testdata
// directory
type Test struct {
	path     string // The path including directory and basename
	data     string // The input data to be parsed
	itemData string // The expected lex items output in json
	nodeData string // The expected parse nodes in json
}

// expectNodes returns the expected parse_tree values from the tests as
// unmarshaled JSON. A panic occurs if there is an error unmarshaling the JSON
// data.
func (l Test) expectNodes() (nl []interface{}) {
	if err := json.Unmarshal([]byte(l.nodeData), &nl); err != nil {
		panic(fmt.Errorf("JSON error: ", err))
	}
	return
}

// expectItems unmarshals the expected lex_items into a silce of items. A panic
// occurs if there is an error decoding the JSON data.
func (l Test) expectItems() (lexItems []item) {
	if err := json.Unmarshal([]byte(l.itemData), &lexItems); err != nil {
		panic(fmt.Errorf("JSON error: ", err))
	}
	return
}

func LoadLexTest(t *testing.T, path string) (test *Test) {
	iDPath := "../testdata/" + path + ".rst"
	inputData, err := ioutil.ReadFile(iDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(inputData) == 0 {
		t.Fatalf("\"%s\" is empty!", iDPath)
	}
	itemFPath := "../testdata/" + path + "_items.json"
	itemData, err := ioutil.ReadFile(itemFPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(itemData) == 0 {
		t.Fatalf("\"%s\" is empty!", itemFPath)
	}
	return &Test{
		path:     path,
		data:     string(inputData[:len(inputData)-1]),
		itemData: string(itemData),
	}
}

func LoadParseTest(t *testing.T, path string) (test *Test) {
	iDPath := "../testdata/" + path + ".rst"
	inputData, err := ioutil.ReadFile(iDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(inputData) == 0 {
		t.Fatalf("\"%s\" is empty!", iDPath)
	}
	nDPath := "../testdata/" + path + "_nodes.json"
	nodeData, err := ioutil.ReadFile(nDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(nodeData) == 0 {
		t.Fatalf("\"%s\" is empty!", nDPath)
	}
	return &Test{
		path:     path,
		data:     string(inputData[:len(inputData)-1]),
		nodeData: string(nodeData),
	}
}

type checkNode struct {
	t          *testing.T
	testPath   string
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
	var got, exp string

	switch c.pFieldVal.(type) {
	case ID:
		got = c.pFieldVal.(ID).String()
		exp = strconv.Itoa(int(c.eFieldVal.(float64)))
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
	eTemp := "(ID: %2d) Got: %s = %q\n\t\t Expect: %s = %q\n\n"
	c.t.Errorf(eTemp, c.id, c.pFieldName, got, c.eFieldName, exp)
}

func (c *checkNode) updateState(eKey string, eVal interface{},
	pVal reflect.Value) bool {

	// Expected parser metadata
	c.eFieldName = eKey
	c.eFieldVal = eVal
	c.eFieldType = reflect.TypeOf(c.eFieldVal)

	// Actual parsed metadata
	c.pNodeName = pVal.Type().Name()
	c.pFieldName = strings.ToUpper(string(c.eFieldName[0]))
	c.pFieldName += c.eFieldName[1:]

	if c.pFieldName == "Id" {
		// Overide for uppercase ID
		c.pFieldName = "ID"
	}

	if !pVal.FieldByName(c.pFieldName).IsValid() {
		eTmp := "Missing field in parser output: %s.%s\n"
		panic(fmt.Errorf(eTmp, c.pNodeName, c.pFieldName))
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
		case "text":
			nExpect := norm.NFC.String(c.eFieldVal.(string))
			if nExpect != c.pFieldVal.(string) {
				c.dError()
			}
		case "type":
			if c.eFieldVal != c.pFieldVal.(NodeType).String() {
				c.dError()
			}
		case "id":
			if c.eFieldVal != float64(c.pFieldVal.(ID)) {
				c.dError()
			}
		case "level", "length":
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
		case "indent", "overLine", "title", "underLine":
			c.checkFields(c.eFieldVal, c.pFieldVal.(Node))
		case "nodeList":
			len1 := len(c.eFieldVal.([]interface{}))
			len2 := len(c.pFieldVal.(NodeList))
			if len1 != len2 {
				iVal := c.eFieldVal.([]interface{})[0]
				id := iVal.(map[string]interface{})["id"]
				spd.Dump(pNode)
				spd.Dump(eNodes)
				eTmp := "Expected NodeList values (len=%d) " +
					"and parsed NodeList values (len=%d) " +
					"do not match beginning at item ID: %d"
				c.t.Fatal(eTmp, len1, len2, id)
			}
			for num, node := range c.eFieldVal.([]interface{}) {
				// Store and reset the parser value, otherwise
				// a panic will occur on the next iteration
				pFieldVal := c.pFieldVal
				c.checkFields(node, c.pFieldVal.(NodeList)[num])
				c.pFieldVal = pFieldVal
			}
		case "rune":
			if c.eFieldVal != string(c.pFieldVal.(rune)) {
				c.dError()
			}
		case "severity":
			pFVal := c.pFieldVal.(systemMessageLevel).String()
			if c.eFieldVal != pFVal {
				c.dError()
			}
		case "enumType":
			if c.eFieldVal != c.pFieldVal.(EnumListType).String() {
				c.dError()
			}
		case "affix":
			if c.eFieldVal != c.pFieldVal.(EnumAffixType).String() {
				c.dError()
			}
		default:
			panic(fmt.Errorf("%s is not implemented!", c.eFieldName))
		}
	}

}

func checkParseNodes(t *testing.T, eTree []interface{}, pNodes []Node,
	testPath string) {

	state := &checkNode{t: t, testPath: testPath}
	for eNum, eNode := range eTree {
		state.checkFields(eNode, pNodes[eNum])
	}

	return
}

func parseTest(t *testing.T, test *Test) (tree *Tree) {
	log.Debugf("Test path: %s\n", test.path)
	log.Debugf("Test Input:\n-----------\n%s\n----------\n", test.data)
	tree, _ = Parse(test.path, test.data)
	return
}

// tokEqualChecker compares the lexed tokens and the expected tokens and
// reports failures.
type tokEqualChecker func(*Tree, reflect.Value, int, string)

// checkTokens checks the lexed tokens against the expected tokens and uses
// isEqual to perform the actual checks and report errors.
func checkTokens(tr *Tree, trExp interface{}, isEqual tokEqualChecker) {
	for i := 0; i < len(tr.token); i++ {
		tokenPos := i - zed
		zedPos := "zed"
		tPi := int(math.Abs(float64(i - zed)))
		tokenPosStr := strconv.Itoa(tPi)
		var fName string
		if tokenPos < 0 {
			fName = "Back" + tokenPosStr + "Tok"
			zedPos = "zed-" + tokenPosStr
		} else if tokenPos == 0 {
			fName = "ZedToken"
		} else {
			fName = "Peek" + tokenPosStr + "Tok"
			zedPos = "zed+" + tokenPosStr
		}
		tokenPos = int(math.Abs(float64(i - zed)))
		tField := reflect.ValueOf(trExp).FieldByName(fName)
		if tField.IsValid() {
			isEqual(tr, tField, i, zedPos)
		}
	}
}

var treeBackupTests = []struct {
	name      string
	input     string
	nextNum   int   // The number of times to call Tree.next().
	backupNum int   // Number of calls to Tree.backup(). Value starts at 1.
	Back4Tok  *item // The fourth backup token.
	Back3Tok  *item // The third backup token.
	Back2Tok  *item // The second backup token.
	Back1Tok  *item // The first backup token.
	ZedToken  *item // The item to expect at Tree.token[zed].
	Peek1Tok  *item // The first peek token.
	Peek2Tok  *item // The second peek token.
	Peek3Tok  *item // The third peek token.
	Peek4Tok  *item // The fourth peek token.
}{
	{
		name:    "Single backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 1,
		ZedToken: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
	},
	{
		name:    "Double backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 2,
		// ZedToken is nil
		Peek1Tok: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
	},
	{
		name:    "Triple backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 3,
		// ZedToken is nil
		Peek2Tok: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
	},
	{
		name:    "Quadruple backup",
		input:   "Title\n=====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 13, backupNum: 4,
		// Back tokens 4 - 1 and ZedToken are nil
		Peek1Tok: &item{ID: 10, Type: itemParagraph, Text: "Four"},
		Peek2Tok: &item{ID: 11, Type: itemBlankLine, Text: "\n"},
		Peek3Tok: &item{ID: 12, Type: itemParagraph, Text: "Five"},
		Peek4Tok: &item{ID: 13, Type: itemEOF},
	},
}

func TestTreeBackup(t *testing.T) {
	var tField reflect.Value
	var tr *Tree
	var zedPos string
	isEqual := func(pos int) {
		val := tField.Interface().(*item)
		if tr.token[pos] == nil {
			t.Fatalf("Test: %q\n\t    "+
				"Got: token[%s] = %v, Expect: %#+v\n\n",
				tr.Name, zedPos, tr.token[pos], val)
		}
		if tr.token[pos] == nil {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s] = %#+v, Expect: %#+v\n\n",
				tr.Name, zedPos, tr.token[pos], val)
		}
		if tr.token[pos].ID != val.ID {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].ID = %d, Expect: %d\n\n",
				tr.Name, zedPos, tr.token[pos].Type, val.ID)
		}
		if tr.token[pos].Type != val.Type {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].Type = %q, Expect: %q\n\n",
				tr.Name, zedPos, tr.token[pos].Type, val.Type)
		}
		if tr.token[pos].Text != val.Text && val.Text != "" {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].Text = %q, Expect: %q\n\n",
				tr.Name, zedPos, tr.token[pos].Text, val.Text)
		}
	}
	for _, tt := range treeBackupTests {
		log.Debugf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name)
		tr = New(tt.name, tt.input)
		tr.lex = lex(tt.name, tt.input)
		for i := 0; i < tt.nextNum; i++ {
			tr.next()
		}
		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}
		for k := 0; k < len(tr.token); k++ {
			tokenPos := k - zed
			zedPos = "zed"
			tPi := int(math.Abs(float64(k - zed)))
			tokenPosStr := strconv.Itoa(tPi)
			var fName string
			if tokenPos < 0 {
				fName = "Back" + tokenPosStr + "Tok"
				zedPos = "zed-" + tokenPosStr
			} else if tokenPos == 0 {
				fName = "ZedToken"
			} else {
				fName = "Peek" + tokenPosStr + "Tok"
				zedPos = "zed+" + tokenPosStr
			}
			tokenPos = int(math.Abs(float64(k - zed)))
			tField = reflect.ValueOf(tt).FieldByName(fName)
			if tField.IsValid() && !tField.IsNil() {
				isEqual(k)
			}
		}
	}
}

var treeNextTests = []struct {
	name     string
	input    string
	nextNum  int   // Number of times to call Tree.next(). Value starts at 1.
	Back4Tok *item // The item to expect at Tree.token[zed-4]
	Back3Tok *item // The item to expect at Tree.token[zed-3]
	Back2Tok *item // The item to expect at Tree.token[zed-2]
	Back1Tok *item // The item to expect at Tree.token[zed-1]
	ZedToken *item // The item to expect at Tree.token[zed]
	Peek1Tok *item // Peek tokens should be blank on next tests.
	Peek2Tok *item
	Peek3Tok *item
	Peek4Tok *item
}{
	{
		name:    "Next no input",
		input:   "",
		nextNum: 1,
		// ZedToken should be nil
	},
	{
		name:     "Single next from start",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  1,
		ZedToken: &item{Type: itemTitle, Text: "Test"},
	},
	{
		name:     "Double next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  2,
		Back1Tok: &item{Type: itemTitle, Text: "Test"},
		ZedToken: &item{Type: itemSectionAdornment, Text: "====="},
	},
	{
		name:     "Triple next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  3,
		Back2Tok: &item{Type: itemTitle, Text: "Test"},
		Back1Tok: &item{Type: itemSectionAdornment, Text: "====="},
		ZedToken: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:     "Quadruple next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  4,
		Back3Tok: &item{Type: itemTitle, Text: "Test"},
		Back2Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Back1Tok: &item{Type: itemBlankLine, Text: "\n"},
		ZedToken: &item{Type: itemParagraph, Text: "Paragraph."},
	},
	{
		name:     "Quintuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  5,
		Back4Tok: &item{Type: itemTitle, Text: "Test"},
		Back3Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Back2Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back1Tok: &item{Type: itemParagraph, Text: "Paragraph."},
		ZedToken: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:     "Sextuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  6,
		Back4Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Back3Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back2Tok: &item{Type: itemParagraph, Text: "Paragraph."},
		Back1Tok: &item{Type: itemBlankLine, Text: "\n"},
		ZedToken: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:     "Septuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  7,
		Back4Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back3Tok: &item{Type: itemParagraph, Text: "Paragraph."},
		Back2Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back1Tok: &item{Type: itemBlankLine, Text: "\n"},
		ZedToken: &item{Type: itemEOF},
	},
	{
		name:     "Two next() on one line of input",
		input:    "Test",
		nextNum:  2,
		Back1Tok: &item{Type: itemParagraph, Text: "Test"},
		ZedToken: &item{Type: itemEOF},
	},
	{
		name:  "Three next() on one line of input; Test channel close.",
		input: "Test",
		// The channel should be closed on the second next(), otherwise
		// a deadlock would occur.
		nextNum: 3,
	},
}

func TestTreeNext(t *testing.T) {
	var tField reflect.Value
	var fName, zedPos string
	var tr *Tree

	isEqual := func(pos int) {
		val := tField.Interface().(*item)
		if tr.token[pos] == nil {
			t.Fatalf("Test: %q\n\t    "+
				"Got: token[%s] = %v, Expect: %#+v\n\n",
				tr.Name, zedPos, tr.token[pos], val)
		}
		if tr.token[pos].Type != val.Type {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].Type = %q, Expect: %q\n\n",
				tr.Name, zedPos, tr.token[pos].Type, val.Type)
		}
		if tr.token[pos].Text != val.Text && val.Text != "" {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].Text = %q, Expect: %q\n\n",
				tr.Name, zedPos, tr.token[pos].Text, val.Text)
		}
	}

	for _, tt := range treeNextTests {
		log.Debugf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name)
		tr = New(tt.name, tt.input)
		tr.lex = lex(tt.name, tt.input)
		for i := 0; i < tt.nextNum; i++ {
			tr.next()
		}
		for k := 0; k < len(tr.token); k++ {
			tokenPos := k - zed
			zedPos = "zed"
			tPi := int(math.Abs(float64(k - zed)))
			tokenPosStr := strconv.Itoa(tPi)
			if tokenPos < 0 {
				fName = "Back" + tokenPosStr + "Tok"
				zedPos = "zed-" + tokenPosStr
			} else if tokenPos == 0 {
				fName = "ZedToken"
			} else {
				fName = "Peek" + tokenPosStr + "Tok"
				zedPos = "zed+" + tokenPosStr
			}
			tokenPos = int(math.Abs(float64(k - zed)))
			tField = reflect.ValueOf(tt).FieldByName(fName)
			if tField.IsValid() && !tField.IsNil() {
				isEqual(k)
			}
		}
	}
}

var treePeekTests = []struct {
	name     string
	input    string
	nextNum  int   // Number of times to call Tree.next() before peek
	peekNum  int   // position argument to Tree.peek()
	Back4Tok *item // The Back tokens should be empty on peek tests.
	Back3Tok *item
	Back2Tok *item
	Back1Tok *item
	ZedToken *item // Should be empty on peek tests.
	Peek1Tok *item
	Peek2Tok *item
	Peek3Tok *item
	Peek4Tok *item
}{
	{
		name:     "Single peek no next",
		input:    "Test\n=====\n\nParagraph.",
		peekNum:  1,
		Peek1Tok: &item{Type: itemTitle, Text: "Test"},
	},
	{
		name:     "Double peek no next",
		input:    "Test\n=====\n\nParagraph.",
		peekNum:  2,
		Peek1Tok: &item{Type: itemTitle, Text: "Test"},
		Peek2Tok: &item{Type: itemSectionAdornment, Text: "====="},
	},
	{
		name:    "Triple peek no next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 0, peekNum: 3,
		Peek1Tok: &item{Type: itemTitle, Text: "Test"},
		Peek2Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Peek3Tok: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:    "Triple peek and double next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 2, peekNum: 3,
		Peek1Tok: &item{Type: itemBlankLine, Text: "\n"},
		Peek2Tok: &item{Type: itemParagraph, Text: "One"},
		Peek3Tok: &item{Type: itemTitle, Text: "Test 2"},
	},
	{
		name:    "Quadruple peek and triple next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 3, peekNum: 4,
		Peek1Tok: &item{Type: itemParagraph, Text: "One"},
		Peek2Tok: &item{Type: itemTitle, Text: "Test 2"},
		Peek3Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Peek4Tok: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:    "Peek on no input",
		peekNum: 1,
	},
}

func TestTreePeek(t *testing.T) {
	var tField reflect.Value
	var fName, zedPos string
	var tr *Tree

	isEqual := func(pos int) {
		val := tField.Interface().(*item)
		if tr.token[pos] == nil {
			t.Fatalf("Test: %q\n\t    "+
				"Got: token[%s] = %v, Expect: %#+v\n\n",
				tr.Name, zedPos, tr.token[pos], val)
		}
		if tr.token[pos].Type != val.Type {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].Type = %q, Expect: %q\n\n",
				tr.Name, zedPos, tr.token[pos].Type, val.Type)
		}
		if tr.token[pos].Text != val.Text && val.Text != "" {
			t.Errorf("Test: %q\n\t    "+
				"Got: token[%s].Text = %q, Expect: %q\n\n",
				tr.Name, zedPos, tr.token[pos].Text, val.Text)
		}
	}

	for _, tt := range treePeekTests {
		log.Debugf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name)
		tr = New(tt.name, tt.input)
		tr.lex = lex(tt.name, tt.input)
		for i := 0; i < tt.nextNum; i++ {
			tr.next()
		}
		tr.peek(tt.peekNum)
		for k := 0; k < len(tr.token); k++ {
			tokenPos := k - zed
			zedPos = "zed"
			tPi := int(math.Abs(float64(k - zed)))
			tokenPosStr := strconv.Itoa(tPi)
			if tokenPos < 0 {
				fName = "Back" + tokenPosStr + "Tok"
				zedPos = "zed-" + tokenPosStr
			} else if tokenPos == 0 {
				fName = "ZedToken"
			} else {
				fName = "Peek" + tokenPosStr + "Tok"
				zedPos = "zed+" + tokenPosStr
			}
			tokenPos = int(math.Abs(float64(k - zed)))
			tField = reflect.ValueOf(tt).FieldByName(fName)
			if tField.IsValid() && !tField.IsNil() {
				isEqual(k)
			}
		}
	}
}

type shortSectionNode struct {
	id    ID
	level int  // SectionNode level
	oRune rune // SectionNode Overline Rune
	uRune rune // SectionNode Underline Rune
}

// The section nodes to add to fill sectionLevels
type testSectionLevelSectionNode struct {
	eMessage parserMessage // Expected parser message
	node     shortSectionNode
}

type testSectionLevelExpectLevels struct {
	rChar    rune
	level    int
	overLine bool
	nodes    []shortSectionNode
}

var testSectionLevelsAdd = []struct {
	name  string
	pSecs []*testSectionLevelSectionNode
	eLvls []*testSectionLevelExpectLevels
}{
	{
		name: "Test two levels with a single SectionNode each",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
		},
		eLvls: []*testSectionLevelExpectLevels{
			{rChar: '=', level: 1, nodes: []shortSectionNode{
				{level: 1, uRune: '='},
			}},
			{rChar: '-', level: 2, nodes: []shortSectionNode{
				{level: 2, uRune: '-'},
			}},
		},
	},
	{
		name: "Test two levels with on level one return",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
		},
		eLvls: []*testSectionLevelExpectLevels{
			{rChar: '=', level: 1, nodes: []shortSectionNode{
				{level: 1, uRune: '='},
				{level: 1, uRune: '='},
			}},
			{rChar: '-', level: 2, nodes: []shortSectionNode{
				{level: 2, uRune: '-'},
				{level: 2, uRune: '-'},
			}},
		},
	},
	{
		name: "Test three levels with one return to level 1",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
			{node: shortSectionNode{level: 3, uRune: '~'}},
			{node: shortSectionNode{level: 1, uRune: '='}},
		},
		eLvls: []*testSectionLevelExpectLevels{
			{rChar: '=', level: 1, nodes: []shortSectionNode{
				{level: 1, uRune: '='},
				{level: 1, uRune: '='},
			}},
			{rChar: '-', level: 2, nodes: []shortSectionNode{
				{level: 2, uRune: '-'},
			}},
			{rChar: '~', level: 3, nodes: []shortSectionNode{
				{level: 3, uRune: '~'},
			}},
		},
	},
	{
		name: "Test three levels with two returns to level 1",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
			{node: shortSectionNode{level: 3, uRune: '~'}},
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
		},
		eLvls: []*testSectionLevelExpectLevels{
			{rChar: '=', level: 1, nodes: []shortSectionNode{
				{level: 1, uRune: '='},
				{level: 1, uRune: '='},
				{level: 1, uRune: '='},
			}},
			{rChar: '-', level: 2, nodes: []shortSectionNode{
				{level: 2, uRune: '-'},
				{level: 2, uRune: '-'},
			}},
			{rChar: '~', level: 3, nodes: []shortSectionNode{
				{level: 3, uRune: '~'},
			}},
		},
	},
	{
		name: "Test inconsistent section level",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
			{node: shortSectionNode{level: 1, uRune: '='}},
			{eMessage: severeTitleLevelInconsistent,
				node: shortSectionNode{level: 2, uRune: '`'}},
		},
	},
	{
		name: "Test inconsistent section level 2",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
			{node: shortSectionNode{level: 3, uRune: '~'}},
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, uRune: '-'}},
			{eMessage: severeTitleLevelInconsistent,
				node: shortSectionNode{level: 3, uRune: '`'}},
		},
	},
	{
		name: "Test level two with overline and all runes similar",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{id: 1, level: 1, uRune: '='}},
			{node: shortSectionNode{
				id: 2, level: 2, oRune: '=', uRune: '=',
			}},
		},
		eLvls: []*testSectionLevelExpectLevels{
			{rChar: '=', level: 1, nodes: []shortSectionNode{
				{level: 1, uRune: '='},
			}},
			{rChar: '=', level: 2, overLine: true,
				nodes: []shortSectionNode{
					{level: 2, uRune: '='},
				},
			},
		},
	},
	{
		name: "Test level two with overline with same rune as level one.",
		pSecs: []*testSectionLevelSectionNode{
			{node: shortSectionNode{id: 1, level: 1, uRune: '='}},
			{node: shortSectionNode{
				id: 2, level: 2, oRune: '=', uRune: '=',
			}},
		},
		eLvls: []*testSectionLevelExpectLevels{
			{rChar: '=', level: 1, nodes: []shortSectionNode{
				{level: 1, uRune: '='},
			}},
			{rChar: '=', level: 2, overLine: true,
				nodes: []shortSectionNode{
					{level: 2, uRune: '='},
				},
			},
		},
	},
}

func testSectionLevelsAddCheckEqual(t *testing.T, testName string,
	pos int, pLvl, eLvl *sectionLevel) {

	if eLvl.level != pLvl.level {
		t.Errorf("Test: %q\n\t    "+
			"Got: sectionLevel.Level = %d, "+
			"Expect: %d\n\n",
			testName, pLvl.level, eLvl.level)
	}
	if eLvl.rChar != pLvl.rChar {
		t.Errorf("Test: %q\n\t    "+
			"Got: sectionLevel.rChar = %#U, "+
			"Expect: %#U\n\n",
			testName, pLvl.rChar, eLvl.rChar)
	}
	if eLvl.overLine != pLvl.overLine {
		t.Errorf("Test: %q\n\t    "+
			"Got: sectionLevel.overLine = %t, "+
			"Expect: %t\n\n",
			testName, pLvl.overLine, eLvl.overLine)
	}
	for eNum, eSec := range eLvl.sections {
		if eSec.ID != pLvl.sections[eNum].ID {
			t.Errorf("Test: %q\n\t    "+
				"Got: level[%d].sections[%d].ID = %d, "+
				"Expect: %d\n\n",
				testName, pos, eNum,
				pLvl.sections[eNum].ID, eSec.ID)
		}
		if eSec.Level != pLvl.sections[eNum].Level {
			t.Errorf("Test: %q\n\t    "+
				"Got: level[%d].sections[%d].Level = %d, "+
				"Expect: %d\n\n",
				testName, pos, eNum,
				pLvl.sections[eNum].Level, eSec.Level)
		}
		eRune := eSec.UnderLine.Rune
		pRune := pLvl.sections[eNum].UnderLine.Rune
		if eRune != pRune {
			t.Errorf("Test: %q\n\t    "+
				"Got: level[%d].section[%d].Rune = %#U, "+
				"Expect: %#U\n\n",
				testName, pos, eNum,
				pLvl.sections[eNum].UnderLine.Rune,
				eSec.UnderLine.Rune)
		}
	}
}

func TestSectionLevelsAdd(t *testing.T) {
	var pSecLvls, eSecLvls sectionLevels
	var testName string

	addSection := func(s *testSectionLevelSectionNode) {
		n := &SectionNode{Level: s.node.level,
			UnderLine: &AdornmentNode{Rune: s.node.uRune}}
		if s.node.oRune != 0 {
			n.OverLine = &AdornmentNode{Rune: s.node.oRune}
		}
		msg := pSecLvls.Add(n)
		if msg > parserMessageNil && msg != s.eMessage {
			t.Fatalf("Test: %q\n\t    Got: parserMessage = %q, "+
				"Expect: %q\n\n", testName, msg, s.eMessage)
		}
	}

	for _, tt := range testSectionLevelsAdd {
		log.Debugf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name)
		pSecLvls = *new(sectionLevels)
		eSecLvls = *new(sectionLevels)
		testName = tt.name

		// pSecLvls := new(sectionLevels)
		for _, secNode := range tt.pSecs {
			addSection(secNode)
		}

		// Initialize the expected sectionLevels
		for _, slvl := range tt.eLvls {
			s := &sectionLevel{rChar: slvl.rChar,
				level: slvl.level, overLine: slvl.overLine,
			}
			for _, sn := range slvl.nodes {
				n := &SectionNode{ID: sn.id, Level: sn.level}
				n.UnderLine = &AdornmentNode{Rune: sn.uRune}
				if sn.oRune != 0 {
					n.OverLine = &AdornmentNode{
						Rune: sn.oRune,
					}
				}
				s.sections = append(s.sections, n)
			}
			eSecLvls.levels = append(eSecLvls.levels, s)
		}

		for i := 0; i < len(eSecLvls.levels); i++ {
			testSectionLevelsAddCheckEqual(t, testName, i,
				pSecLvls.levels[i], eSecLvls.levels[i])
		}
	}
}

var testSectionLevelsLast = []struct {
	name      string
	tLevel    int // The last level to get
	tSections []*SectionNode
	eLevel    sectionLevel // There can be only one
}{
	{
		name:   "Test last section level two",
		tLevel: 2,
		tSections: []*SectionNode{
			{Level: 1, Title: &TitleNode{Text: "Title 1"},
				UnderLine: &AdornmentNode{Rune: '='}},
			{Level: 2, Title: &TitleNode{Text: "Title 2"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 3"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 4"},
				UnderLine: &AdornmentNode{Rune: '-'}},
		},
		eLevel: sectionLevel{
			rChar: '-', level: 2,
			sections: []*SectionNode{
				{Level: 2, Title: &TitleNode{Text: "Title 4"},
					UnderLine: &AdornmentNode{Rune: '~'}},
			},
		},
	},
	{
		name:   "Test last section level one",
		tLevel: 1,
		tSections: []*SectionNode{
			{Level: 1, Title: &TitleNode{Text: "Title 1"},
				UnderLine: &AdornmentNode{Rune: '='}},
			{Level: 2, Title: &TitleNode{Text: "Title 2"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 3"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 4"},
				UnderLine: &AdornmentNode{Rune: '-'}},
		},
		eLevel: sectionLevel{
			rChar: '=', level: 1,
			sections: []*SectionNode{
				{Level: 1, Title: &TitleNode{Text: "Title 1"},
					UnderLine: &AdornmentNode{Rune: '='}},
			},
		},
	},
	{
		name:   "Test last section level three",
		tLevel: 3,
		tSections: []*SectionNode{
			{Level: 1, Title: &TitleNode{Text: "Title 1"},
				UnderLine: &AdornmentNode{Rune: '='}},
			{Level: 2, Title: &TitleNode{Text: "Title 2"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 3"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 4"},
				UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 3, Title: &TitleNode{Text: "Title 5"},
				UnderLine: &AdornmentNode{Rune: '+'}},
		},
		eLevel: sectionLevel{
			rChar: '+', level: 3,
			sections: []*SectionNode{
				{Level: 3, Title: &TitleNode{Text: "Title 5"},
					UnderLine: &AdornmentNode{Rune: '+'}},
			},
		},
	},
}

func TestSectionLevelsLast(t *testing.T) {
	for _, tt := range testSectionLevelsLast {
		log.Debugf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name)
		secLvls := new(sectionLevels)
		for _, secNode := range tt.tSections {
			secLvls.Add(secNode)
		}
		var pSec *SectionNode
		pSec = secLvls.LastSectionByLevel(tt.tLevel)
		if tt.eLevel.level != pSec.Level {
			t.Errorf("Test: %q\n\t    "+
				"Got: sectionLevel.Level = %d, Expect: %d\n\n",
				tt.name, tt.eLevel.level, pSec.Level)
		}
		if tt.eLevel.rChar != pSec.UnderLine.Rune {
			t.Errorf("Test: %q\n\t    "+
				"Got: sectionLevel.rChar = %#U, Expect: %#U\n\n",
				tt.name, tt.eLevel.rChar, pSec.UnderLine.Rune)
		}
		// There can be only one
		if tt.eLevel.sections[0].ID != pSec.ID {
			t.Errorf("Test: %q\n\t    "+
				"Got: level[0].sections[0].ID = %d, "+
				"Expect: %d\n\n",
				tt.name, pSec.ID, tt.eLevel.sections[0].ID)
		}
		if tt.eLevel.sections[0].Title.Text != pSec.Title.Text {
			t.Errorf("Test: %q\n\t    "+
				"Got: level[0].sections[0].Title.Text = %q, "+
				"Expect: %q\n\n",
				tt.name, pSec.Title.Text,
				tt.eLevel.sections[0].Title.Text)
		}
	}
}

func TestSystemMessageLevelFrom(t *testing.T) {
	name := "Test systemMessageLevel with levelInfo"
	test0 := ""
	if -1 != systemMessageLevelFromString(test0) {
		t.Errorf("Test: %q\n\t    "+
			"Got: systemMessageLevel = %q, Expect: %q\n\n",
			name, systemMessageLevelFromString(test0), -1)
	}
	test1 := "INFO"
	if levelInfo != systemMessageLevelFromString(test1) {
		t.Errorf("Test: %q\n\t    "+
			"Got: systemMessageLevel = %q, Expect: %q\n\n",
			name, systemMessageLevelFromString(test1), levelInfo)
	}
	test2 := "SEVERE"
	if levelInfo != systemMessageLevelFromString(test1) {
		t.Errorf("Test: %q\n\t    "+
			"Got: systemMessageLevel = %q, Expect: %q\n\n",
			name, systemMessageLevelFromString(test2), levelSevere)
	}
}
