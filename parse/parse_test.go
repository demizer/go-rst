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

// SetDebug is typically called from the init() function in a test file. SetDebug parses debug flags
// passed to the test binary and also sets the template for logging output.
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

// Contains a single test with data loaded from test files in the testdata directory
type Test struct {
	path     string // The path including directory and basename
	data     string // The input data to be parsed
	itemData string // The expected lex items output in json
	nodeData string // The expected parse nodes in json
}

// expectNodes returns the expected parse_tree values from the tests as unmarshaled JSON. A panic
// occurs if there is an error unmarshaling the JSON data.
func (l Test) expectNodes() (nl []interface{}) {
	if err := json.Unmarshal([]byte(l.nodeData), &nl); err != nil {
		panic(fmt.Errorf("JSON error: ", err))
	}
	return
}

// expectItems unmarshals the expected lex_items into a silce of items. A panic occurs if there is
// an error decoding the JSON data.
func (l Test) expectItems() (lexItems []item) {
	if err := json.Unmarshal([]byte(l.itemData), &lexItems); err != nil {
		panic(fmt.Errorf("JSON error: ", err))
	}
	return
}

func LoadTest(path string) (test *Test) {
	inputData, err := ioutil.ReadFile("../testdata/" + path + ".rst")
	if err != nil {
		panic(err)
	}
	itemData, err := ioutil.ReadFile("../testdata/" + path + "_items.json")
	if err != nil {
		panic(err)
	}
	nodeData, err := ioutil.ReadFile("../testdata/" + path + "_nodes.json")
	if err != nil {
		panic(err)
	}
	return &Test{
		path:     path,
		data:     string(inputData[:len(inputData)-1]),
		itemData: string(itemData),
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
		pNumStr := " (" + strconv.Itoa(int(c.pFieldVal.(systemMessageLevel))) + ")"
		got = c.pFieldVal.(systemMessageLevel).String() + pNumStr
		eNumStr := " (" + strconv.Itoa(int(systemMessageLevelFromString(c.eFieldVal.(string)))) + ")"
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
	c.t.Errorf("(ID: %2d) Got: %s = %q\n\t\t Expect: %s = %q\n\n", c.id, c.pFieldName, got,
		c.eFieldName, exp)
}

func (c *checkNode) updateState(eKey string, eVal interface{}, pVal reflect.Value) bool {
	// Expected parser metadata
	c.eFieldName = eKey
	c.eFieldVal = eVal
	c.eFieldType = reflect.TypeOf(c.eFieldVal)

	// Actual parsed metadata
	c.pNodeName = pVal.Type().Name()
	c.pFieldName = strings.ToUpper(string(c.eFieldName[0])) + c.eFieldName[1:]

	if c.pFieldName == "Id" {
		// Overide for uppercase ID
		c.pFieldName = "ID"
	}

	if !pVal.FieldByName(c.pFieldName).IsValid() {
		panic(fmt.Errorf("Missing field in parser output: %s.%s\n", c.pNodeName, c.pFieldName))
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
			if norm.NFC.String(c.eFieldVal.(string)) != c.pFieldVal.(string) {
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
				id := c.eFieldVal.([]interface{})[0].(map[string]interface{})["id"]
				spd.Dump(pNode)
				spd.Dump(eNodes)
				c.t.Fatal("Expected NodeList values ( len =", len1, ") and parsed "+
					"NodeList values ( len =", len2, ") do not match beginning at "+
					"item ID", id)
			}
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

func checkParseNodes(t *testing.T, eTree []interface{}, pNodes []Node, testPath string) {
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

var treeBackupTests = []struct {
	name      string
	input     string
	nextNum   int   // The number of times to call Tree.next().
	backupNum int   // Number of times to call Tree.backup(). Value starts at 1.
	Back3Tok  *item // The third backup token.
	Back2Tok  *item // The second backup token.
	Back1Tok  *item // The first backup token.
	ZedToken  *item // The item to expect at Tree.token[zed].
	Peek1Tok  *item // The first peek token.
	Peek2Tok  *item // The second peek token.
	Peek3Tok  *item // The third peek token.
}{
	{
		name:    "Backup once",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 1,
		ZedToken: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
	},
	{
		name:    "Backup twice",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 2,
		Peek1Tok: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
	},
	{
		name:    "Backup thrice",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 3,
		Peek2Tok: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
	},
}

func TestTreeBackup(t *testing.T) {
	for _, tt := range treeBackupTests {
		tree := New(tt.name, tt.input)
		tree.lex = lex(tt.name, tt.input)
		for i := 0; i < tt.nextNum; i++ {
			tree.next()
		}
		for j := 0; j < tt.backupNum; j++ {
			tree.backup()
		}
		for k := 0; k < len(tree.token); k++ {
			tokenPos := k - zed
			zedPos := "zed"
			tokenPosStr := strconv.Itoa(int(math.Abs(float64(k - zed))))
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
			tField := reflect.ValueOf(tt).FieldByName(fName)
			if tField.IsValid() && !tField.IsNil() {
				val := tField.Interface().(*item)
				if tree.token[k] == nil {
					t.Errorf("Test: %q\n\t    Got: token[%s] = %#+v, Expect: %#+v\n\n",
						tree.Name, zedPos, tree.token[k], val)
				}
				if tree.token[k].ID != val.ID {
					t.Errorf("Test: %q\n\t    Got: token[%s].ID = %d, Expect: %d\n\n",
						tree.Name, zedPos, tree.token[k].Type, val.ID)
				}
				if tree.token[k].Type != val.Type {
					t.Errorf("Test: %q\n\t    Got: token[%s].Type = %q, Expect: %q\n\n",
						tree.Name, zedPos, tree.token[k].Type, val.Type)
				}
				if tree.token[k].Text != val.Text {
					t.Errorf("Test: %q\n\t    Got: token[%s].Text = %q, Expect: %q\n\n",
						tree.Name, zedPos, tree.token[k].Text, val.Text)
				}
			}
		}
	}
}

var treeNextTests = []struct {
	name     string
	input    string
	nextNum  int   // Number of times to call Tree.next(). Value starts at 1.
	back3Tok *item // The item to expect at Tree.token[zed-3]
	back2Tok *item // The item to expect at Tree.token[zed-2]
	back1Tok *item // The item to expect at Tree.token[zed-1]
	zedToken *item // The item to expect at Tree.token[zed]
}{
	{
		name:    "Next no input",
		input:   "",
		nextNum: 1,
		// zedToken should be nil
	},
	{
		name:     "Single next from start",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  1,
		zedToken: &item{Type: itemTitle, Text: "Test"},
	},
	{
		name:     "Double next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  2,
		back1Tok: &item{Type: itemTitle, Text: "Test"},
		zedToken: &item{Type: itemSectionAdornment, Text: "====="},
	},
	{
		name:     "Tripple next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  3,
		back2Tok: &item{Type: itemTitle, Text: "Test"},
		back1Tok: &item{Type: itemSectionAdornment, Text: "====="},
		zedToken: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:     "Quadrupple next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  4,
		back3Tok: &item{Type: itemTitle, Text: "Test"},
		back2Tok: &item{Type: itemSectionAdornment, Text: "====="},
		back1Tok: &item{Type: itemBlankLine, Text: "\n"},
		zedToken: &item{Type: itemParagraph, Text: "Paragraph."},
	},
	{
		name:     "Two next() on one line of input",
		input:    "Test",
		nextNum:  2,
		back1Tok: &item{Type: itemParagraph, Text: "Test"},
		zedToken: &item{Type: itemEOF, Text: ""},
	},
}

func TestTreeNext(t *testing.T) {
	for _, tt := range treeNextTests {
		tree := New(tt.name, tt.input)
		tree.lex = lex(tt.name, tt.input)
		if tt.nextNum > 4 {
			panic("nextNum cannot be greater than 4")
		}
		for i := 0; i < tt.nextNum; i++ {
			tree.next()
		}
		if tree.token[zed] == nil && tt.zedToken == nil {
			continue
		}
		if tt.nextNum > 0 {
			if tree.token[zed].Type != tt.zedToken.Type {
				t.Errorf("Test: %q\n\t    Got: token[zed].Type = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed].Type, tt.zedToken.Type)
			}
			if tree.token[zed].Text != tt.zedToken.Text {
				t.Errorf("Test: %q\n\t    Got: token[zed].Text = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed].Text, tt.zedToken.Text)
			}
		}
		if tt.nextNum > 1 {
			if tree.token[zed-1].Type != tt.back1Tok.Type {
				t.Errorf("Test: %q\n\t    Got: token[zed-1].Type = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed-1].Type, tt.back1Tok.Type)
			}
			if tree.token[zed-1].Text != tt.back1Tok.Text {
				t.Errorf("Test: %q\n\t    Got: token[zed-1].Text = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed-1].Text, tt.back1Tok.Text)
			}
		}
		if tt.nextNum > 2 {
			if tree.token[zed-2].Type != tt.back2Tok.Type {
				t.Errorf("Test: %q\n\t    Got: token[zed-2].Type = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed-2].Type, tt.back2Tok.Type)
			}
			if tree.token[zed-2].Text != tt.back2Tok.Text {
				t.Errorf("Test: %q\n\t    Got: token[zed-2].Text = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed-2].Text, tt.back2Tok.Text)
			}
		}
		if tt.nextNum > 3 {
			if tree.token[zed-3].Type != tt.back3Tok.Type {
				t.Errorf("Test: %q\n\t    Got: token[zed-3].Type = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed-3].Type, tt.back3Tok.Type)
			}
			if tree.token[zed-3].Text != tt.back3Tok.Text {
				t.Errorf("Test: %q\n\t    Got: token[zed-3].Text = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed-3].Text, tt.back3Tok.Text)
			}
		}
		// Calculate the number of open backup tokens and make sure they are
		// empty.
		for j := 0; j < 4-tt.nextNum; j++ {
			if tree.token[j] != nil {
				t.Errorf("Test: %q\n\t    Got: token[%d] = %#+v, Expect: nil\n\n",
					tree.Name, j, tree.token[j])
			}
		}
		// make sure the peek positions are blank
		if tree.token[zed+1] != nil {
			t.Errorf("Test: %q\n\t    Got: token[zed+1] = %#+v, Expect: nil\n\n",
				tree.Name, tree.token[zed+1])
		} else if tree.token[zed+2] != nil {
			t.Errorf("Test: %q\n\t    Got: token[zed+2] = %#+v, Expect: nil\n\n",
				tree.Name, tree.token[zed+2])
		} else if tree.token[zed+3] != nil {
			t.Errorf("Test: %q\n\t    Got: token[zed+3] = %#+v, Expect: nil\n\n",
				tree.Name, tree.token[zed+3])
		}
	}
}

var treePeekTests = []struct {
	name     string
	input    string
	nextNum  int   // Number of times to call Tree.next() before peek
	peekNum  int   // position argument to Tree.peek()
	peek1Tok *item // The item to expect at Tree.token[zed+1]
	peek2Tok *item // The item to expect at Tree.token[zed+2]
	peek3Tok *item // The item to expect at Tree.token[zed+3]
}{
	{
		name:    "Peek from starting position",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 0, peekNum: 1,
		peek1Tok: &item{Type: itemTitle, Text: "Test"},
	},
	{
		name:    "Peek from starting position two times",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 0, peekNum: 2,
		peek1Tok: &item{Type: itemTitle, Text: "Test"},
		peek2Tok: &item{Type: itemSectionAdornment, Text: "====="},
	},
	{
		name:    "Peek three positions",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 0, peekNum: 3,
		peek1Tok: &item{Type: itemTitle, Text: "Test"},
		peek2Tok: &item{Type: itemSectionAdornment, Text: "====="},
		peek3Tok: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:    "Next 2 positions, Peek 3 positions",
		input:   "Test\n=====\n\nParagraph.\nTest 2\n=====\n\nParagraph 2.",
		nextNum: 2, peekNum: 3,
		peek1Tok: &item{Type: itemBlankLine, Text: "\n"},
		peek2Tok: &item{Type: itemParagraph, Text: "Paragraph."},
		peek3Tok: &item{Type: itemTitle, Text: "Test 2"},
	},
}

func TestTreePeek(t *testing.T) {
	for _, tt := range treePeekTests {
		tree := New(tt.name, tt.input)
		tree.lex = lex(tt.name, tt.input)
		for i := 0; i < tt.nextNum; i++ {
			tree.next()
		}
		tree.peek(tt.peekNum)
		if tree.token[zed+1].Type != tt.peek1Tok.Type {
			t.Errorf("Test: %q\n\t    Got: token[zed+1].Type = %q, Expect: %q\n\n",
				tree.Name, tree.token[zed+1].Type, tt.peek1Tok.Type)
		}
		if tree.token[zed+1].Text != tt.peek1Tok.Text {
			t.Errorf("Test: %q\n\t    Got: token[zed+1].Text = %q, Expect: %q\n\n",
				tree.Name, tree.token[zed+1].Text, tt.peek1Tok.Text)
		}
		if tt.peekNum > 1 {
			if tree.token[zed+2].Type != tt.peek2Tok.Type {
				t.Errorf("Test: %q\n\t    Got: token[zed+2].Type = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed+2].Type, tt.peek2Tok.Type)
			}
			if tree.token[zed+2].Text != tt.peek2Tok.Text {
				t.Errorf("Test: %q\n\t    Got: token[zed+2].Text = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed+2].Text, tt.peek2Tok.Text)
			}
		}
		if tt.peekNum > 2 {
			if tree.token[zed+3].Type != tt.peek3Tok.Type {
				t.Errorf("Test: %q\n\t    Got: token[zed+3].Type = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed+3].Type, tt.peek3Tok.Type)
			}
			if tree.token[zed+3].Text != tt.peek3Tok.Text {
				t.Errorf("Test: %q\n\t    Got: token[zed+3].Text = %q, Expect: %q\n\n",
					tree.Name, tree.token[zed+3].Text, tt.peek3Tok.Text)
			}
		}
	}
}

type testSectionLevelSectionNode struct {
	eMessage parserMessage
	node     *SectionNode
}

var testSectionLevelsAdd = []struct {
	name      string
	tSections []*testSectionLevelSectionNode
	eLevels   []*sectionLevel
}{
	{
		name: "Test two levels with a single SectionNode each",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
		},
		eLevels: []*sectionLevel{
			{rChar: '=', level: 1,
				sections: []*SectionNode{
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
				},
			},
			{rChar: '-', level: 2,
				sections: []*SectionNode{
					{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}},
				},
			},
		},
	},
	{
		name: "Test two levels with on level one return",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
		},
		eLevels: []*sectionLevel{
			{rChar: '=', level: 1,
				sections: []*SectionNode{
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
				},
			},
			{rChar: '-', level: 2,
				sections: []*SectionNode{
					{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}},
					{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}},
				},
			},
		},
	},
	{
		name: "Test three levels with one return to level 1",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
			{node: &SectionNode{Level: 3, UnderLine: &AdornmentNode{Rune: '~'}}},
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
		},
		eLevels: []*sectionLevel{
			{rChar: '=', level: 1,
				sections: []*SectionNode{
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
				},
			},
			{rChar: '-', level: 2,
				sections: []*SectionNode{
					{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}},
				},
			},
			{rChar: '~', level: 3,
				sections: []*SectionNode{
					{Level: 3, UnderLine: &AdornmentNode{Rune: '~'}},
				},
			},
		},
	},
	{
		name: "Test three levels with two returns to level 1",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
			{node: &SectionNode{Level: 3, UnderLine: &AdornmentNode{Rune: '~'}}},
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
		},
		eLevels: []*sectionLevel{
			{rChar: '=', level: 1,
				sections: []*SectionNode{
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
					{Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
				},
			},
			{rChar: '-', level: 2,
				sections: []*SectionNode{
					{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}},
					{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}},
				},
			},
			{rChar: '~', level: 3,
				sections: []*SectionNode{
					{Level: 3, UnderLine: &AdornmentNode{Rune: '~'}},
				},
			},
		},
	},
	{
		name: "Test inconsistent section level",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{eMessage: severeTitleLevelInconsistent,
				node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '`'}}},
		},
	},
	{
		name: "Test inconsistent section level 2",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
			{node: &SectionNode{Level: 3, UnderLine: &AdornmentNode{Rune: '~'}}},
			{node: &SectionNode{Level: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{Level: 2, UnderLine: &AdornmentNode{Rune: '-'}}},
			{eMessage: severeTitleLevelInconsistent,
				node: &SectionNode{Level: 3, UnderLine: &AdornmentNode{Rune: '`'}}},
		},
	},
	{
		name: "Test level two with overline with same rune as level one.",
		tSections: []*testSectionLevelSectionNode{
			{node: &SectionNode{ID: 1, UnderLine: &AdornmentNode{Rune: '='}}},
			{node: &SectionNode{ID: 2, OverLine: &AdornmentNode{Rune: '='},
				UnderLine: &AdornmentNode{Rune: '='}}},
		},
		eLevels: []*sectionLevel{
			{rChar: '=', level: 1, overLine: false,
				sections: []*SectionNode{
					{ID: 1, Level: 1, UnderLine: &AdornmentNode{Rune: '='}},
				},
			},
			{rChar: '=', level: 2, overLine: true,
				sections: []*SectionNode{
					{ID: 2, Level: 2, OverLine: &AdornmentNode{Rune: '='},
						UnderLine: &AdornmentNode{Rune: '='}},
				},
			},
		},
	},
}

func TestSectionLevelsAdd(t *testing.T) {
	for _, tt := range testSectionLevelsAdd {
		secLvls := new(sectionLevels)
		for _, secNode := range tt.tSections {
			msg := secLvls.Add(secNode.node)
			if msg > parserMessageNil && msg != secNode.eMessage {
				t.Fatalf("Test: %q\n\t    Got: parserMessage = %q, Expect: %q\n\n",
					tt.name, msg, secNode.eMessage)
			}
		}
		for sNum, secLvl := range tt.eLevels {
			pSecLvl := (secLvls.levels)[sNum]
			if secLvl.level != pSecLvl.level {
				t.Errorf("Test: %q\n\t    Got: sectionLevel.Level = %d, Expect: %d\n\n",
					tt.name, secLvl.level, pSecLvl.level)
			}
			if secLvl.rChar != pSecLvl.rChar {
				t.Errorf("Test: %q\n\t    Got: sectionLevel.rChar = %#U, Expect: %#U\n\n",
					tt.name, secLvl.rChar, pSecLvl.rChar)
			}
			if secLvl.overLine != pSecLvl.overLine {
				t.Errorf("Test: %q\n\t    Got: sectionLevel.overLine = %t, Expect: %t\n\n",
					tt.name, secLvl.overLine, pSecLvl.overLine)
			}
			for eNum, eSec := range secLvl.sections {
				if eSec.ID != pSecLvl.sections[eNum].ID {
					t.Errorf("Test: %q\n\t    Got: level[%d].sections[%d].ID = %d, Expect: %d\n\n",
						tt.name, sNum, eNum, pSecLvl.sections[eNum].ID, eSec.ID)
				}
				if eSec.Level != pSecLvl.sections[eNum].Level {
					t.Errorf("Test: %q\n\t    Got: level[%d].sections[%d].Level = %d, Expect: %d\n\n",
						tt.name, sNum, eNum, pSecLvl.sections[eNum].Level, eSec.Level)
				}
				if eSec.UnderLine.Rune != pSecLvl.sections[eNum].UnderLine.Rune {
					t.Errorf("Test: %q\n\t    Got: level[%d].section[%d].Rune = %#U, Expect: %#U\n\n",
						tt.name, sNum, eNum, pSecLvl.sections[eNum].UnderLine.Rune, eSec.UnderLine.Rune)
				}
			}
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
			{Level: 1, Title: &TitleNode{Text: "Title 1"}, UnderLine: &AdornmentNode{Rune: '='}},
			{Level: 2, Title: &TitleNode{Text: "Title 2"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 3"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 4"}, UnderLine: &AdornmentNode{Rune: '-'}},
		},
		eLevel: sectionLevel{
			rChar: '-', level: 2,
			sections: []*SectionNode{
				{Level: 2, Title: &TitleNode{Text: "Title 4"}, UnderLine: &AdornmentNode{Rune: '~'}},
			},
		},
	},
	{
		name:   "Test last section level one",
		tLevel: 1,
		tSections: []*SectionNode{
			{Level: 1, Title: &TitleNode{Text: "Title 1"}, UnderLine: &AdornmentNode{Rune: '='}},
			{Level: 2, Title: &TitleNode{Text: "Title 2"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 3"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 4"}, UnderLine: &AdornmentNode{Rune: '-'}},
		},
		eLevel: sectionLevel{
			rChar: '=', level: 1,
			sections: []*SectionNode{
				{Level: 1, Title: &TitleNode{Text: "Title 1"}, UnderLine: &AdornmentNode{Rune: '='}},
			},
		},
	},
	{
		name:   "Test last section level three",
		tLevel: 3,
		tSections: []*SectionNode{
			{Level: 1, Title: &TitleNode{Text: "Title 1"}, UnderLine: &AdornmentNode{Rune: '='}},
			{Level: 2, Title: &TitleNode{Text: "Title 2"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 3"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &TitleNode{Text: "Title 4"}, UnderLine: &AdornmentNode{Rune: '-'}},
			{Level: 3, Title: &TitleNode{Text: "Title 5"}, UnderLine: &AdornmentNode{Rune: '+'}},
		},
		eLevel: sectionLevel{
			rChar: '+', level: 3,
			sections: []*SectionNode{
				{Level: 3, Title: &TitleNode{Text: "Title 5"}, UnderLine: &AdornmentNode{Rune: '+'}},
			},
		},
	},
}

func TestSectionLevelsLast(t *testing.T) {
	for _, tt := range testSectionLevelsLast {
		secLvls := new(sectionLevels)
		for _, secNode := range tt.tSections {
			secLvls.Add(secNode)
		}
		var pSec *SectionNode
		pSec = secLvls.LastSectionByLevel(tt.tLevel)
		if tt.eLevel.level != pSec.Level {
			t.Errorf("Test: %q\n\t    Got: sectionLevel.Level = %d, Expect: %d\n\n",
				tt.name, tt.eLevel.level, pSec.Level)
		}
		if tt.eLevel.rChar != pSec.UnderLine.Rune {
			t.Errorf("Test: %q\n\t    Got: sectionLevel.rChar = %#U, Expect: %#U\n\n",
				tt.name, tt.eLevel.rChar, pSec.UnderLine.Rune)
		}
		// There can be only one
		if tt.eLevel.sections[0].ID != pSec.ID {
			t.Errorf("Test: %q\n\t    Got: level[0].sections[0].ID = %d, Expect: %d\n\n",
				tt.name, pSec.ID, tt.eLevel.sections[0].ID)
		}
		if tt.eLevel.sections[0].Title.Text != pSec.Title.Text {
			t.Errorf("Test: %q\n\t    Got: level[0].sections[0].Title.Text = %q, Expect: %q\n\n",
				tt.name, pSec.Title.Text, tt.eLevel.sections[0].Title.Text)
		}
	}
}

func TestSystemMessageLevelFrom(t *testing.T) {
	name := "Test systemMessageLevel with levelInfo"
	test0 := ""
	if -1 != systemMessageLevelFromString(test0) {
		t.Errorf("Test: %q\n\t    Got: systemMessageLevel = %q, Expect: %q\n\n",
			name, systemMessageLevelFromString(test0), -1)
	}
	test1 := "INFO"
	if levelInfo != systemMessageLevelFromString(test1) {
		t.Errorf("Test: %q\n\t    Got: systemMessageLevel = %q, Expect: %q\n\n",
			name, systemMessageLevelFromString(test1), levelInfo)
	}
	test2 := "SEVERE"
	if levelInfo != systemMessageLevelFromString(test1) {
		t.Errorf("Test: %q\n\t    Got: systemMessageLevel = %q, Expect: %q\n\n",
			name, systemMessageLevelFromString(test2), levelSevere)
	}
}

func TestParseSectionTitleGood0000(t *testing.T) {
	// Basic title, underline, blankline, and paragraph test
	testPath := "test_section/01_title_good/00.00_title_paragraph"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0001(t *testing.T) {
	// Basic title, underline, and paragraph with no blankline line after the
	// section.
	testPath := "test_section/01_title_good/00.01_paragraph_noblankline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0002(t *testing.T) {
	// A title that begins with a combining unicode character \u0301. Tests to
	// make sure the 2 byte unicode does not contribute to the underline length
	// calculation.
	testPath := "test_section/01_title_good/00.02_title_combining_chars"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0100(t *testing.T) {
	// A basic section in between paragraphs.
	testPath := "test_section/01_title_good/01.00_para_head_para"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0200(t *testing.T) {
	// Tests section parsing on 3 character long title and underline.
	testPath := "test_section/01_title_good/02.00_short_title"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0300(t *testing.T) {
	// Tests a single section with no other element surrounding it.
	testPath := "test_section/01_title_good/03.00_empty_section"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0000(t *testing.T) {
	// Tests for severe system messages when the sections are indented.
	testPath := "test_section/02_title_bad/00.00_unexpected_titles"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0100(t *testing.T) {
	// Tests for severe system message on short title underline
	testPath := "test_section/02_title_bad/01.00_short_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0200(t *testing.T) {
	// Tests for title underlines that are less than three characters.
	testPath := "test_section/02_title_bad/02.00_short_title_short_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0201(t *testing.T) {
	// Tests for title overlines and underlines that are less than three characters.
	testPath := "test_section/02_title_bad/02.01_short_title_short_overline_and_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0202(t *testing.T) {
	// Tests for short title overline with missing underline when the overline
	// is less than three characters.
	testPath := "test_section/02_title_bad/02.02_short_title_short_overline_missing_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0000(t *testing.T) {
	// Tests section level return to level one after three subsections.
	testPath := "test_section/03_level_good/00.00_section_level_return"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0001(t *testing.T) {
	// Tests section level return to level one after 1 subsection. The second
	// level one section has one subsection.
	testPath := "test_section/03_level_good/00.01_section_level_return"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0002(t *testing.T) {
	// Test section level with subsection 4 returning to level two.
	testPath := "test_section/03_level_good/00.02_section_level_return"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0100(t *testing.T) {
	// Tests section level return with title overlines
	testPath := "test_section/03_level_good/01.00_section_level_return"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0000(t *testing.T) {
	// Test section level return on bad level 2 section adornment
	testPath := "test_section/04_level_bad/00.00_bad_subsection_order"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0001(t *testing.T) {
	// Test section level return with title overlines on bad level 2 section adornment
	testPath := "test_section/04_level_bad/00.01_bad_subsection_order_with_overlines"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0000(t *testing.T) {
	// Test simple section with title overline.
	testPath := "test_section/05_title_with_overline_good/00.00_title_overline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0100(t *testing.T) {
	// Test simple section with inset title and overline.
	testPath := "test_section/05_title_with_overline_good/01.00_inset_title_with_overline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0200(t *testing.T) {
	// Test sections with three character adornments lines.
	testPath := "test_section/05_title_with_overline_good/02.00_three_char_section_title"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0000(t *testing.T) {
	// Test section title with overline, but no underline.
	testPath := "test_section/06_title_with_overline_bad/00.00_inset_title_missing_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0001(t *testing.T) {
	// Test inset title with overline but missing underline.
	testPath := "test_section/06_title_with_overline_bad/00.01_inset_title_missing_underline_with_blankline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0002(t *testing.T) {
	// Test inset title with overline but missing underline. The title is
	// followed by a blank line and a paragraph.
	testPath := "test_section/06_title_with_overline_bad/00.02_inset_title_missing_underline_and_para"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0003(t *testing.T) {
	// Test section overline with missmatched underline.
	testPath := "test_section/06_title_with_overline_bad/00.03_inset_title_mismatched_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0100(t *testing.T) {
	// Test overline with really long title.
	testPath := "test_section/06_title_with_overline_bad/01.00_title_too_long"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0200(t *testing.T) {
	// Test overline and underline with blanklines instead of a title.
	testPath := "test_section/06_title_with_overline_bad/02.00_missing_titles_with_blankline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0201(t *testing.T) {
	// Test overline and underline with nothing where the title is supposed to
	// be.
	testPath := "test_section/06_title_with_overline_bad/02.01_missing_titles_with_noblankline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0300(t *testing.T) {
	// Test two character overline with no underline.
	testPath := "test_section/06_title_with_overline_bad/03.00_incomplete_section"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0301(t *testing.T) {
	// Test three character section adornments with no titles or blanklines in
	// between.
	testPath := "test_section/06_title_with_overline_bad/03.01_incomplete_sections_no_title"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0400(t *testing.T) {
	// Tests indented section with overline
	testPath := "test_section/06_title_with_overline_bad/04.00_indented_title_short_overline_and_underline"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0500(t *testing.T) {
	// Tests ".." overline (which is a comment element).
	testPath := "test_section/06_title_with_overline_bad/05.00_two_char_section_title"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleNumberedGood0000(t *testing.T) {
	// Tests lexing a section where the title begins with a number.
	testPath := "test_section/07_title_numbered_good/00.00_numbered_title"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleNumberedGood0100(t *testing.T) {
	// Tests numbered section lexing with enumerated directly above section.
	testPath := "test_section/07_title_numbered_good/01.00_enum_list_with_numbered_title"
	test := LoadTest(testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
