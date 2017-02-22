package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"reflect"
	"strconv"
	"testing"

	"github.com/demizer/go-rst/rst/testutil"

	doc "github.com/demizer/go-rst/rst/document"
	tok "github.com/demizer/go-rst/rst/token"
)

// tokEqualChecker compares the lexed tokens and the expected tokens and reports failures.
type tokEqualChecker func(*Parser, reflect.Value, int, string)

// checkTokens checks the lexed tokens against the expected tokens and uses isEqual to perform the actual checks and report
// errors.
func checkTokens(tr *Parser, trExp interface{}, isEqual tokEqualChecker) {
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

func nodeListToInterface(v *doc.NodeList) []interface{} {
	v2 := []doc.Node(*v)
	s := make([]interface{}, len(v2))
	for i, j := range v2 {
		s[i] = j
	}
	return s
}

// checkParseNodes compares the expected parser output (*_nodes.json) against the actual parser output using the jd library.
func checkParseNodes(t *testing.T, eParser []interface{}, pNodes *doc.NodeList, testPath string) {
	pJson, err := json.MarshalIndent(pNodes, "", "    ")
	if err != nil {
		t.Errorf("Error Marshalling JSON: %s", err.Error())
		return
	}

	// Json diff output has a syntax:
	// https://github.com/josephburnett/jd#diff-language
	o, err := testutil.JsonDiff(eParser, nodeListToInterface(pNodes))
	if err != nil {
		fmt.Println(o)
		fmt.Printf("Error diffing JSON: %s", err.Error())
		return
	}

	// There should be no output from the diff
	if len(o) != 0 {
		// Give all other output time to print
		// time.Sleep(time.Second / 2)

		testutil.Log("\nFAIL: parsed nodes do not match expected nodes!")

		testutil.Log("\n[Parsed Nodes JSON]\n\n")
		testutil.Log(string(pJson))

		testutil.Log("\n\n[JSON DIFF]\n\n")
		testutil.Log(o)

		t.FailNow()
	}
}

func LoadParserTest(t *testing.T, path string) (test *testutil.Test) {
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
	return &testutil.Test{
		Path:     path,
		Data:     string(inputData[:len(inputData)-1]),
		NodeData: string(nodeData),
	}
}

// parseTest initiates the parser and parses a test using test.data is input.
func parseTest(t *testing.T, test *testutil.Test) *Parser {
	testutil.Log(fmt.Sprintf("Test path: %s", test.Path))
	testutil.Log(fmt.Sprintf("Test Input:\n\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n%s\n+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+-+\n", test.Data))
	p, _ := Parse(test.Path, test.Data)
	return p
}

var parserBackupTests = []struct {
	name      string
	input     string
	nextNum   int       // The number of times to call Parser.next().
	backupNum int       // Number of calls to Parser.backup(). Value starts at 1.
	Back4Tok  *tok.Item // The fourth backup token.
	Back3Tok  *tok.Item // The third backup token.
	Back2Tok  *tok.Item // The second backup token.
	Back1Tok  *tok.Item // The first backup token.
	ZedToken  *tok.Item // The item to expect at Parser.token[zed].
	Peek1Tok  *tok.Item // The first peek token.
	Peek2Tok  *tok.Item // The second peek token.
	Peek3Tok  *tok.Item // The third peek token.
	Peek4Tok  *tok.Item // The fourth peek token.
}{
	{
		name:    "Single backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 1,
		ZedToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
		Peek1Tok: &tok.Item{ID: 2, Type: tok.SectionAdornment},
	},
	{
		name:    "Double backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 2,
		// ZedToken is nil
		Peek1Tok: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
		Peek2Tok: &tok.Item{ID: 2, Type: tok.SectionAdornment},
	},
	{
		name:    "Triple backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 3,
		// ZedToken is nil
		Peek2Tok: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
		Peek3Tok: &tok.Item{ID: 2, Type: tok.SectionAdornment},
	},
	{
		name:    "Quadruple backup",
		input:   "Title\n=====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 13, backupNum: 4,
		// Back tokens 4 - 1 and ZedToken are nil
		Peek1Tok: &tok.Item{ID: 10, Type: tok.ItemText, Text: "Four"},
		Peek2Tok: &tok.Item{ID: 11, Type: tok.ItemBlankLine, Text: "\n"},
		Peek3Tok: &tok.Item{ID: 12, Type: tok.ItemText, Text: "Five"},
		Peek4Tok: &tok.Item{ID: 13, Type: tok.EOF},
	},
}

func TestParserBackup(t *testing.T) {
	isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*tok.Item)
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
	for _, tt := range parserBackupTests {
		testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = tok.Lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}
		checkTokens(tr, tt, isEqual)
	}
}

var parserNextTests = []struct {
	name     string
	input    string
	nextNum  int       // Number of times to call Parser.next(). Value starts at 1.
	Back4Tok *tok.Item // The item to expect at Parser.token[zed-4]
	Back3Tok *tok.Item // The item to expect at Parser.token[zed-3]
	Back2Tok *tok.Item // The item to expect at Parser.token[zed-2]
	Back1Tok *tok.Item // The item to expect at Parser.token[zed-1]
	ZedToken *tok.Item // The item to expect at Parser.token[zed]
	Peek1Tok *tok.Item // Peek tokens should be blank on next tests.
	Peek2Tok *tok.Item
	Peek3Tok *tok.Item
	Peek4Tok *tok.Item
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
		ZedToken: &tok.Item{Type: tok.Title, Text: "Test"},
	},
	{
		name:     "Double next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  2,
		Back1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		ZedToken: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
	},
	{
		name:     "Triple next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  3,
		Back2Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		Back1Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		ZedToken: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
	},
	{
		name:     "Quadruple next",
		input:    "Test\n=====\n\nParagraph.",
		nextNum:  4,
		Back3Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		Back2Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		Back1Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		ZedToken: &tok.Item{Type: tok.ItemText, Text: "Paragraph."},
	},
	{
		name:     "Quintuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  5,
		Back4Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		Back3Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		Back2Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		Back1Tok: &tok.Item{Type: tok.ItemText, Text: "Paragraph."},
		ZedToken: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
	},
	{
		name:     "Sextuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  6,
		Back4Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		Back3Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		Back2Tok: &tok.Item{Type: tok.ItemText, Text: "Paragraph."},
		Back1Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		ZedToken: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
	},
	{
		name:     "Septuple next",
		input:    "Test\n=====\n\nParagraph.\n\n",
		nextNum:  7,
		Back4Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		Back3Tok: &tok.Item{Type: tok.ItemText, Text: "Paragraph."},
		Back2Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		Back1Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		ZedToken: &tok.Item{Type: tok.EOF},
	},
	{
		name:     "Two next() on one line of input",
		input:    "Test",
		nextNum:  2,
		Back1Tok: &tok.Item{Type: tok.ItemText, Text: "Test"},
		ZedToken: &tok.Item{Type: tok.EOF},
	},
	{
		name:    "Three next() on one line of input; Test channel close.",
		input:   "Test",
		nextNum: 3,
		// The channel should be closed on the second next(), otherwise a deadlock would occur.
		Back2Tok: &tok.Item{Type: tok.ItemText, Text: "Test"},
		Back1Tok: &tok.Item{Type: tok.EOF},
	},
	{
		name:    "Four next() on one line of input; Test channel close.",
		input:   "Test",
		nextNum: 4,
		// The channel should be closed on the second next(), otherwise a deadlock would occur.
		Back3Tok: &tok.Item{Type: tok.ItemText, Text: "Test"},
		Back2Tok: &tok.Item{Type: tok.EOF},
	},
}

func TestParserNext(t *testing.T) {
	isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*tok.Item)
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
	for _, tt := range parserNextTests {
		testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = tok.Lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		checkTokens(tr, tt, isEqual)
	}
}

var parserPeekTests = []struct {
	name     string
	input    string
	nextNum  int // Number of times to call Parser.next() before peek
	peekNum  int // position argument to Parser.peek()
	Back4Tok *tok.Item
	Back3Tok *tok.Item
	Back2Tok *tok.Item
	Back1Tok *tok.Item
	ZedToken *tok.Item
	Peek1Tok *tok.Item
	Peek2Tok *tok.Item
	Peek3Tok *tok.Item
	Peek4Tok *tok.Item
}{
	{
		name:     "Single peek no next",
		input:    "Test\n=====\n\nParagraph.",
		peekNum:  1,
		Peek1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
	},
	{
		name:     "Double peek no next",
		input:    "Test\n=====\n\nParagraph.",
		peekNum:  2,
		Peek1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		Peek2Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
	},
	{
		name:     "Triple peek no next",
		input:    "Test\n=====\n\nParagraph.",
		peekNum:  3,
		Peek1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		Peek2Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		Peek3Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
	},
	{
		name:    "Triple peek and double next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 2, peekNum: 3,
		Back1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		ZedToken: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		Peek1Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		Peek2Tok: &tok.Item{Type: tok.ItemText, Text: "One"},
		Peek3Tok: &tok.Item{Type: tok.Title, Text: "Test 2"},
	},
	{
		name:    "Quadruple peek and triple next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 3, peekNum: 4,
		Back2Tok: &tok.Item{Type: tok.Title, Text: "Test"},
		Back1Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		ZedToken: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
		Peek1Tok: &tok.Item{Type: tok.ItemText, Text: "One"},
		Peek2Tok: &tok.Item{Type: tok.Title, Text: "Test 2"},
		Peek3Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
		Peek4Tok: &tok.Item{Type: tok.ItemBlankLine, Text: "\n"},
	},
	{
		name:    "Peek on no input",
		peekNum: 1,
	},
}

func TestParserPeek(t *testing.T) {
	isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*tok.Item)
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
	for _, tt := range parserPeekTests {
		testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = tok.Lex(tt.name, []byte(tt.input))
		tr.next(tt.nextNum)
		tr.peek(tt.peekNum)
		checkTokens(tr, tt, isEqual)
	}
}

var testParserClearTokensTests = []struct {
	name       string
	input      string
	nextNum    int       // Number of times to call Parser.next() before peek
	peekNum    int       // position argument to Parser.peek()
	clearBegin int       // Passed to Parser.clear() as the begin arg
	clearEnd   int       // Passed to Parser.clear() as the end arg
	Back4Tok   *tok.Item // Use &item{} if the token is not expected to be nil
	Back3Tok   *tok.Item
	Back2Tok   *tok.Item
	Back1Tok   *tok.Item
	ZedToken   *tok.Item
	Peek1Tok   *tok.Item
	Peek2Tok   *tok.Item
	Peek3Tok   *tok.Item
	Peek4Tok   *tok.Item
}{
	{
		name:    "Fill token buffer and clear it.",
		input:   "Parser\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 5, peekNum: 4,
		clearBegin: zed - 4, clearEnd: zed + 4,
	},
	{
		name:    "Fill token buffer and clear back tokens.",
		input:   "Parser\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 5, peekNum: 4,
		clearBegin: zed - 4, clearEnd: zed - 1,
		ZedToken: &tok.Item{},
		Peek1Tok: &tok.Item{},
		Peek2Tok: &tok.Item{},
		Peek3Tok: &tok.Item{},
		Peek4Tok: &tok.Item{},
	},
	{
		name:    "Fill token buffer and clear peek tokens.",
		input:   "Parser\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		nextNum: 5, peekNum: 4,
		clearBegin: zed + 1, clearEnd: zed + 4,
		Back4Tok: &tok.Item{},
		Back3Tok: &tok.Item{},
		Back2Tok: &tok.Item{},
		Back1Tok: &tok.Item{},
		ZedToken: &tok.Item{},
	},
}

func TestParserClearTokens(t *testing.T) {
	isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
		val := tExp.Interface().(*tok.Item)
		if tr.token[tPos] != nil && val == nil {
			t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
		}
	}
	for _, tt := range testParserClearTokensTests {
		testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		tr := New(tt.name, tt.input)
		tr.lex = tok.Lex(tt.name, []byte(tt.input))
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
		n := &doc.SectionNode{Level: s.node.level,
			UnderLine: &doc.AdornmentNode{Rune: s.node.uRune}}
		if s.node.oRune != 0 {
			n.OverLine = &doc.AdornmentNode{Rune: s.node.oRune}
		}
		msg := pSecLvls.Add(n)
		if msg > parserMessageNil && msg != s.eMessage {
			t.Fatalf("Test: %q\n\tGot: parserMessage = %q, "+"Expect: %q", testName, msg, s.eMessage)
		}
	}

	for _, tt := range testSectionLevelsAdd {
		testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
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
				n := &doc.SectionNode{Level: sn.level}
				n.UnderLine = &doc.AdornmentNode{Rune: sn.uRune}
				if sn.oRune != 0 {
					n.OverLine = &doc.AdornmentNode{
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
	tSections []*doc.SectionNode
	eLevel    sectionLevel // There can be only one
}{
	{
		name:   "Test last section level two",
		tLevel: 2,
		tSections: []*doc.SectionNode{
			{Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 2"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 3"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
		},
		eLevel: sectionLevel{
			rChar: '-', level: 2,
			sections: []*doc.SectionNode{
				{Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '~'}},
			},
		},
	},
	{
		name:   "Test last section level one",
		tLevel: 1,
		tSections: []*doc.SectionNode{
			{Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 2"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 3"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
		},
		eLevel: sectionLevel{
			rChar: '=', level: 1,
			sections: []*doc.SectionNode{
				{Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
			},
		},
	},
	{
		name:   "Test last section level three",
		tLevel: 3,
		tSections: []*doc.SectionNode{
			{Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 2"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 3"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
			{Level: 3, Title: &doc.TitleNode{Text: "Title 5"}, UnderLine: &doc.AdornmentNode{Rune: '+'}},
		},
		eLevel: sectionLevel{
			rChar: '+', level: 3,
			sections: []*doc.SectionNode{
				{Level: 3, Title: &doc.TitleNode{Text: "Title 5"}, UnderLine: &doc.AdornmentNode{Rune: '+'}},
			},
		},
	},
}

func TestSectionLevelsLast(t *testing.T) {
	for _, tt := range testSectionLevelsLast {
		testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
		secLvls := new(sectionLevels)
		for _, secNode := range tt.tSections {
			secLvls.Add(secNode)
		}
		var pSec *doc.SectionNode
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
