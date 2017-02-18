package parser

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"testing"

	"github.com/go-kit/kit/log"
	jd "github.com/josephburnett/jd/lib"
)

var debug bool

func init() { SetDebug() }

func tlog(out string) {
	if debug {
		fmt.Println(out)
	}
}

// SetDebug is typically called from the init() function in a test file.  SetDebug parses debug flags passed to the test
// binary and also sets the template for logging output.
func SetDebug() {
	flag.StringVar(&excludeNamedContext, "exclude", "test", "Exclude context from output.")
	flag.BoolVar(&debug, "debug", false, "Enable debug output.")
	flag.Parse()
	if debug {
		LogSetContext(log.NewContext(log.NewLogfmtLogger(os.Stdout)))
	}
}

func nodeListToInterface(v *NodeList) []interface{} {
	v2 := []Node(*v)
	s := make([]interface{}, len(v2))
	for i, j := range v2 {
		s[i] = j
	}
	return s
}

func jsonDiff(expectedItems, parsedItems []interface{}) (string, error) {
	eJson, err := json.Marshal(expectedItems)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal expectedItems: %s", err.Error())
	}

	pJson, err := json.Marshal(parsedItems)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal parsedItems: %s", err.Error())
	}

	a, _ := jd.ReadJsonString(string(eJson))
	b, _ := jd.ReadJsonString(string(pJson))

	return a.Diff(b).Render(), nil
}

// checkParseNodes compares the expected parser output (*_nodes.json) against the actual parser output using the jd library.
func checkParseNodes(t *testing.T, eTree []interface{}, pNodes *NodeList, testPath string) {
	pJson, err := json.MarshalIndent(pNodes, "", "    ")
	if err != nil {
		t.Errorf("Error Marshalling JSON: %s", err.Error())
		return
	}

	// Json diff output has a syntax:
	// https://github.com/josephburnett/jd#diff-language
	o, err := jsonDiff(eTree, nodeListToInterface(pNodes))
	if err != nil {
		fmt.Println(o)
		fmt.Printf("Error diffing JSON: %s", err.Error())
		return
	}

	// There should be no output from the diff
	if len(o) != 0 {
		// Give all other output time to print
		// time.Sleep(time.Second / 2)

		tlog("\nFAIL: parsed nodes do not match expected nodes!")

		tlog("\n[Parsed Nodes JSON]\n\n")
		tlog(string(pJson))

		tlog("\n\n[JSON DIFF]\n\n")
		tlog(o)

		t.FailNow()
	}
}

// Contains a single test with data loaded from test files in the testdata directory
type Test struct {
	path     string // The path including directory and basename
	data     string // The input data to be parsed
	itemData string // The expected lex items output in json
	nodeData string // The expected parse nodes in json
}

// expectNodes returns the expected parse_tree values from the tests as unmarshaled JSON. A panic occurs if there is an error
// unmarshaling the JSON data.
func (l Test) expectNodes() (nl []interface{}) {
	if err := json.Unmarshal([]byte(l.nodeData), &nl); err != nil {
		panic(fmt.Sprintln("JSON error: ", err))
	}
	return
}

// expectItems unmarshals the expected lex_items into a silce of items. A panic occurs if there is an error decoding the JSON
// data.
func (l Test) expectItems() (lexItems []item) {
	if err := json.Unmarshal([]byte(l.itemData), &lexItems); err != nil {
		panic(fmt.Sprintln("JSON error: ", err))
	}
	return
}

// Contains absolute file paths for the test data
var TestDataFiles []string

// testPathsFromDirectory walks through the file tree in the testdata directory containing all of the tests and returns a
// string slice of all the discovered paths.
func testPathsFromDirectory(dir string) ([]string, error) {
	var paths []string
	wFunc := func(p string, info os.FileInfo, err error) error {
		path, err := filepath.Abs(p)
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".rst" {
			paths = append(paths, path[:len(path)-4])
		}
		return nil
	}
	err := filepath.Walk(dir, wFunc)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// testPathFromName loops through TestDataFiles until name is matched.
func testPathFromName(name string) string {
	var err error
	if len(TestDataFiles) < 1 {
		TestDataFiles, err = testPathsFromDirectory("../../testdata")
		if err != nil {
			panic(err)
		}
	}
	for _, p := range TestDataFiles {
		if len(p)-len(name) > 0 {
			if p[len(p)-len(name):] == name {
				return p
			}
		}
	}
	panic(fmt.Sprintf("Could not find test for %q\n", name))
}

func LoadLexTest(t *testing.T, path string) (test *Test) {
	iDPath := path + ".rst"
	inputData, err := ioutil.ReadFile(iDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(inputData) == 0 {
		t.Fatalf("\"%s\" is empty!", iDPath)
	}
	itemFPath := path + "-items.json"
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
	iDPath := path + ".rst"
	inputData, err := ioutil.ReadFile(iDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(inputData) == 0 {
		t.Fatalf("\"%s\" is empty!", iDPath)
	}
	nDPath := path + "-nodes.json"
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

// parseTest initiates the parser and parses a test using test.data is input.
func parseTest(t *testing.T, test *Test) (tree *Tree) {
	tlog(fmt.Sprintf("Test path: %s", test.path))
	tlog(fmt.Sprintf("Test Input:\n\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n%s\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n", test.data))
	tree, _ = Parse(test.path, test.data)
	return
}

// tokEqualChecker compares the lexed tokens and the expected tokens and reports failures.
type tokEqualChecker func(*Tree, reflect.Value, int, string)

// checkTokens checks the lexed tokens against the expected tokens and uses isEqual to perform the actual checks and report
// errors.
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
		Peek1Tok: &item{ID: 2, Type: itemSectionAdornment},
	},
	{
		name:    "Double backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 2,
		// ZedToken is nil
		Peek1Tok: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
		Peek2Tok: &item{ID: 2, Type: itemSectionAdornment},
	},
	{
		name:    "Triple backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 3,
		// ZedToken is nil
		Peek2Tok: &item{ID: 1, Type: itemTitle, Text: "Title 1"},
		Peek3Tok: &item{ID: 2, Type: itemSectionAdornment},
	},
	{
		name:    "Quadruple backup",
		input:   "Title\n=====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 13, backupNum: 4,
		// Back tokens 4 - 1 and ZedToken are nil
		Peek1Tok: &item{ID: 10, Type: itemText, Text: "Four"},
		Peek2Tok: &item{ID: 11, Type: itemBlankLine, Text: "\n"},
		Peek3Tok: &item{ID: 12, Type: itemText, Text: "Five"},
		Peek4Tok: &item{ID: 13, Type: itemEOF},
	},
}

func TestTreeBackup(t *testing.T) {
	isEqual := func(tr *Tree, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*item)
		if val == nil && tr.token[tPos] == nil {
			return
		}
		if val == nil && tr.token[tPos] != nil {
			t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
			return
		}
		if tr.token[tPos].ID != val.ID {
			t.Errorf("Test: %q\n\tGot: token[%s].ID = %d, Expect: %d", tr.Name, tName, tr.token[tPos].Type, val.ID)
		}
		if tr.token[tPos].Type != val.Type {
			t.Errorf("Test: %q\n\tGot: token[%s].Type = %q, Expect: %q", tr.Name, tName, tr.token[tPos].Type, val.Type)
		}
		if tr.token[tPos].Text != val.Text && val.Text != "" {
			t.Errorf("Test: %q\n\tGot: token[%s].Text = %q, Expect: %q", tr.Name, tName, tr.token[tPos].Text, val.Text)
		}
	}
	for _, tt := range treeBackupTests {
		tlog(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}
		checkTokens(tr, tt, isEqual)
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
		ZedToken: &item{Type: itemText, Text: "Paragraph."},
	},
	{
		name:     "Quintuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  5,
		Back4Tok: &item{Type: itemTitle, Text: "Test"},
		Back3Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Back2Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back1Tok: &item{Type: itemText, Text: "Paragraph."},
		ZedToken: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:     "Sextuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  6,
		Back4Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Back3Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back2Tok: &item{Type: itemText, Text: "Paragraph."},
		Back1Tok: &item{Type: itemBlankLine, Text: "\n"},
		ZedToken: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:     "Septuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  7,
		Back4Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back3Tok: &item{Type: itemText, Text: "Paragraph."},
		Back2Tok: &item{Type: itemBlankLine, Text: "\n"},
		Back1Tok: &item{Type: itemBlankLine, Text: "\n"},
		ZedToken: &item{Type: itemEOF},
	},
	{
		name:     "Two next() on one line of input",
		input:    "Test",
		nextNum:  2,
		Back1Tok: &item{Type: itemText, Text: "Test"},
		ZedToken: &item{Type: itemEOF},
	},
	{
		name:    "Three next() on one line of input; Test channel close.",
		input:   "Test",
		nextNum: 3,
		// The channel should be closed on the second next(), otherwise a deadlock would occur.
		Back2Tok: &item{Type: itemText, Text: "Test"},
		Back1Tok: &item{Type: itemEOF},
	},
	{
		name:    "Four next() on one line of input; Test channel close.",
		input:   "Test",
		nextNum: 4,
		// The channel should be closed on the second next(), otherwise a deadlock would occur.
		Back3Tok: &item{Type: itemText, Text: "Test"},
		Back2Tok: &item{Type: itemEOF},
	},
}

func TestTreeNext(t *testing.T) {
	isEqual := func(tr *Tree, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*item)
		if val == nil && tr.token[tPos] == nil {
			return
		}
		if val == nil && tr.token[tPos] != nil {
			t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
			return
		}
		if tr.token[tPos].Type != val.Type {
			t.Errorf("Test: %q\n\tGot: token[%d].Type = %q, Expect: %q", tr.Name, tPos, tr.token[tPos].Type, val.Type)
		}
		if tr.token[tPos].Text != val.Text && val.Text != "" {
			t.Errorf("Test: %q\n\tGot: token[%d].Text = %q, Expect: %q", tr.Name, tPos, tr.token[tPos].Text, val.Text)
		}
	}
	for _, tt := range treeNextTests {
		tlog(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		checkTokens(tr, tt, isEqual)
	}
}

var treePeekTests = []struct {
	name     string
	input    string
	nextNum  int // Number of times to call Tree.next() before peek
	peekNum  int // position argument to Tree.peek()
	Back4Tok *item
	Back3Tok *item
	Back2Tok *item
	Back1Tok *item
	ZedToken *item
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
		name:     "Triple peek no next",
		input:    "Test\n=====\n\nParagraph.",
		peekNum:  3,
		Peek1Tok: &item{Type: itemTitle, Text: "Test"},
		Peek2Tok: &item{Type: itemSectionAdornment, Text: "====="},
		Peek3Tok: &item{Type: itemBlankLine, Text: "\n"},
	},
	{
		name:    "Triple peek and double next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 2, peekNum: 3,
		Back1Tok: &item{Type: itemTitle, Text: "Test"},
		ZedToken: &item{Type: itemSectionAdornment, Text: "====="},
		Peek1Tok: &item{Type: itemBlankLine, Text: "\n"},
		Peek2Tok: &item{Type: itemText, Text: "One"},
		Peek3Tok: &item{Type: itemTitle, Text: "Test 2"},
	},
	{
		name:    "Quadruple peek and triple next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 3, peekNum: 4,
		Back2Tok: &item{Type: itemTitle, Text: "Test"},
		Back1Tok: &item{Type: itemSectionAdornment, Text: "====="},
		ZedToken: &item{Type: itemBlankLine, Text: "\n"},
		Peek1Tok: &item{Type: itemText, Text: "One"},
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
	isEqual := func(tr *Tree, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*item)
		if val == nil && tr.token[tPos] == nil {
			return
		}
		if val == nil && tr.token[tPos] != nil {
			t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
			return
		}
		if tr.token[tPos].Type != val.Type {
			t.Errorf("Test: %q\n\tGot: token[%s].Type = %q, Expect: %q", tr.Name, tName, tr.token[tPos].Type, val.Type)
		}
		if tr.token[tPos].Text != val.Text && val.Text != "" {
			t.Errorf("Test: %q\n\tGot: token[%s].Text = %q, Expect: %q", tr.Name, tName, tr.token[tPos].Text, val.Text)
		}
	}
	for _, tt := range treePeekTests {
		tlog(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		tr.peek(tt.peekNum)
		checkTokens(tr, tt, isEqual)
	}
}

var testTreeClearTokensTests = []struct {
	name       string
	input      string
	nextNum    int   // Number of times to call Tree.next() before peek
	peekNum    int   // position argument to Tree.peek()
	clearBegin int   // Passed to Tree.clear() as the begin arg
	clearEnd   int   // Passed to Tree.clear() as the end arg
	Back4Tok   *item // Use &item{} if the token is not expected to be nil
	Back3Tok   *item
	Back2Tok   *item
	Back1Tok   *item
	ZedToken   *item
	Peek1Tok   *item
	Peek2Tok   *item
	Peek3Tok   *item
	Peek4Tok   *item
}{
	{
		name:    "Fill token buffer and clear it.",
		input:   "Tree\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 5, peekNum: 4,
		clearBegin: zed - 4, clearEnd: zed + 4,
	},
	{
		name:    "Fill token buffer and clear back tokens.",
		input:   "Tree\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 5, peekNum: 4,
		clearBegin: zed - 4, clearEnd: zed - 1,
		ZedToken: &item{},
		Peek1Tok: &item{},
		Peek2Tok: &item{},
		Peek3Tok: &item{},
		Peek4Tok: &item{},
	},
	{
		name:    "Fill token buffer and clear peek tokens.",
		input:   "Tree\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 5, peekNum: 4,
		clearBegin: zed + 1, clearEnd: zed + 4,
		Back4Tok: &item{},
		Back3Tok: &item{},
		Back2Tok: &item{},
		Back1Tok: &item{},
		ZedToken: &item{},
	},
}

func TestTreeClearTokens(t *testing.T) {
	isEqual := func(tr *Tree, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*item)
		if tr.token[tPos] != nil && val == nil {
			t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
		}
	}
	for _, tt := range testTreeClearTokensTests {
		tlog(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		tr.peek(tt.peekNum)
		tr.clearTokens(tt.clearBegin, tt.clearEnd)
		checkTokens(tr, tt, isEqual)
	}
}

type shortSectionNode struct {
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
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, oRune: '=', uRune: '='}},
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
			{node: shortSectionNode{level: 1, uRune: '='}},
			{node: shortSectionNode{level: 2, oRune: '=', uRune: '='}},
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
		t.Errorf("Test: %q\n\tGot: sectionLevel.Level = %d, "+"Expect: %d", testName, pLvl.level, eLvl.level)
	}
	if eLvl.rChar != pLvl.rChar {
		t.Errorf("Test: %q\n\tGot: sectionLevel.rChar = %#U, "+"Expect: %#U", testName, pLvl.rChar, eLvl.rChar)
	}
	if eLvl.overLine != pLvl.overLine {
		t.Errorf("Test: %q\n\tGot: sectionLevel.overLine = %t, "+"Expect: %t", testName, pLvl.overLine,
			eLvl.overLine)
	}
	for eNum, eSec := range eLvl.sections {
		if eSec.Level != pLvl.sections[eNum].Level {
			t.Errorf("Test: %q\n\tGot: level[%d].sections[%d].Level = %d, "+"Expect: %d", testName, pos,
				eNum, pLvl.sections[eNum].Level, eSec.Level)
		}
		eRune := eSec.UnderLine.Rune
		pRune := pLvl.sections[eNum].UnderLine.Rune
		if eRune != pRune {
			t.Errorf("Test: %q\n\tGot: level[%d].section[%d].Rune = %#U, "+"Expect: %#U", testName, pos,
				eNum, pLvl.sections[eNum].UnderLine.Rune, eSec.UnderLine.Rune)
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
			t.Fatalf("Test: %q\n\tGot: parserMessage = %q, "+"Expect: %q", testName, msg, s.eMessage)
		}
	}

	for _, tt := range testSectionLevelsAdd {
		tlog(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
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
				n := &SectionNode{Level: sn.level}
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
		tlog(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		secLvls := new(sectionLevels)
		for _, secNode := range tt.tSections {
			secLvls.Add(secNode)
		}
		var pSec *SectionNode
		pSec = secLvls.LastSectionByLevel(tt.tLevel)
		if tt.eLevel.level != pSec.Level {
			t.Errorf("Test: %q\n\tGot: sectionLevel.Level = %d, Expect: %d", tt.name, tt.eLevel.level,
				pSec.Level)
		}
		if tt.eLevel.rChar != pSec.UnderLine.Rune {
			t.Errorf("Test: %q\n\tGot: sectionLevel.rChar = %#U, Expect: %#U", tt.name, tt.eLevel.rChar,
				pSec.UnderLine.Rune)
		}
		// There can be only one
		if tt.eLevel.sections[0].Title.Text != pSec.Title.Text {
			t.Errorf("Test: %q\n\tGot: level[0].sections[0].Title.Text = %q, "+"Expect: %q", tt.name,
				pSec.Title.Text, tt.eLevel.sections[0].Title.Text)
		}
	}
}

func TestSystemMessageLevelFrom(t *testing.T) {
	name := "Test systemMessageLevel with levelInfo"
	test0 := ""
	if -1 != systemMessageLevelFromString(test0) {
		t.Errorf("Test: %q\n\tGot: systemMessageLevel = %q, Expect: %q", name,
			systemMessageLevelFromString(test0), -1)
	}
	test1 := "INFO"
	if levelInfo != systemMessageLevelFromString(test1) {
		t.Errorf("Test: %q\n\tGot: systemMessageLevel = %q, Expect: %q", name,
			systemMessageLevelFromString(test1), levelInfo)
	}
	test2 := "SEVERE"
	if levelInfo != systemMessageLevelFromString(test1) {
		t.Errorf("Test: %q\n\tGot: systemMessageLevel = %q, Expect: %q", name,
			systemMessageLevelFromString(test2), levelSevere)
	}
}
