package parser

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

// tokEqualChecker compares the lexed tokens and the expected tokens and reports failures.
type tokEqualChecker func(*Parser, reflect.Value, int, string)

// checkTokens checks the lexed tokens against the expected tokens and uses isEqual to perform the actual checks and report
// errors.
func checkTokens(tr *Parser, trExp interface{}, isEqual tokEqualChecker) {
	for i := 0; i < len(tr.tokens); i++ {
		tr.Msgr("have token", "token", tr.tokens[i])
		// tokenPos := i - zed
		// zedPos := "zed"
		// tPi := int(math.Abs(float64(i - zed)))
		// tokenPosStr := strconv.Itoa(tPi)
		// var fName string
		// if tokenPos < 0 {
		// fName = "Back" + tokenPosStr + "Tok"
		// zedPos = "zed-" + tokenPosStr
		// } else if tokenPos == 0 {
		// fName = "ZedToken"
		// } else {
		// fName = "Peek" + tokenPosStr + "Tok"
		// zedPos = "zed+" + tokenPosStr
		// }
		// tField := reflect.ValueOf(trExp).FieldByName(fName)
		// if tField.IsValid() {
		// isEqual(tr, tField, i, zedPos)
		// }
	}
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
	ZedToken  *tok.Item // The item to expect at Parser.token.
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
	// {
	// name:    "Double backup",
	// input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
	// nextNum: 2, backupNum: 2,
	// // ZedToken is nil
	// Peek1Tok: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
	// Peek2Tok: &tok.Item{ID: 2, Type: tok.SectionAdornment},
	// },
	// {
	// name:    "Triple backup",
	// input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
	// nextNum: 2, backupNum: 3,
	// // ZedToken is nil
	// Peek2Tok: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
	// Peek3Tok: &tok.Item{ID: 2, Type: tok.SectionAdornment},
	// },
	// {
	// name:    "Quadruple backup",
	// input:   "Title\n=====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
	// nextNum: 13, backupNum: 4,
	// // Back tokens 4 - 1 and ZedToken are nil
	// Peek1Tok: &tok.Item{ID: 10, Type: tok.Text, Text: "Four"},
	// Peek2Tok: &tok.Item{ID: 11, Type: tok.BlankLine, Text: "\n"},
	// Peek3Tok: &tok.Item{ID: 12, Type: tok.Text, Text: "Five"},
	// Peek4Tok: &tok.Item{ID: 13, Type: tok.EOF},
	// },
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
		tr := New(tt.name, tt.input, testutil.StdLogger)
		var err error
		tr.lex, err = tok.Lex(tt.name, []byte(tt.input), testutil.StdLogger)
		if err != nil {
			t.Errorf("lexer error: %s", err)
			t.Fail()
		}
		tr.next(tt.nextNum)
		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}
		checkTokens(tr, tt, isEqual)
	}
}

// var parserNextTests = []struct {
// name        string
// input       string
// nextNum     int // Number of times to call Parser.next(). Value starts at 1.
// expectError bool
// Back4Tok    *tok.Item // The item to expect at Parser.token[zed-4]
// Back3Tok    *tok.Item // The item to expect at Parser.token[zed-3]
// Back2Tok    *tok.Item // The item to expect at Parser.token[zed-2]
// Back1Tok    *tok.Item // The item to expect at Parser.token[zed-1]
// ZedToken    *tok.Item // The item to expect at Parser.token
// Peek1Tok    *tok.Item // Peek tokens should be blank on next tests.
// Peek2Tok    *tok.Item
// Peek3Tok    *tok.Item
// Peek4Tok    *tok.Item
// }{
// {
// name:        "Next no input",
// input:       "",
// nextNum:     1,
// expectError: true,
// },
// {
// name:     "Single next from start",
// input:    "Test\n=====\n\nParagraph.",
// nextNum:  1,
// ZedToken: &tok.Item{Type: tok.Title, Text: "Test"},
// },
// {
// name:     "Double next",
// input:    "Test\n=====\n\nParagraph.",
// nextNum:  2,
// Back1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// ZedToken: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// },
// {
// name:     "Triple next",
// input:    "Test\n=====\n\nParagraph.",
// nextNum:  3,
// Back2Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// Back1Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// ZedToken: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// },
// {
// name:     "Quadruple next",
// input:    "Test\n=====\n\nParagraph.",
// nextNum:  4,
// Back3Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// Back2Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// Back1Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// ZedToken: &tok.Item{Type: tok.Text, Text: "Paragraph."},
// },
// {
// name:     "Quintuple next",
// input:    "Test\n=====\n\nParagraph.\n\n",
// nextNum:  5,
// Back4Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// Back3Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// Back2Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// Back1Tok: &tok.Item{Type: tok.Text, Text: "Paragraph."},
// ZedToken: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// },
// {
// name:     "Sextuple next",
// input:    "Test\n=====\n\nParagraph.\n\n",
// nextNum:  6,
// Back4Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// Back3Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// Back2Tok: &tok.Item{Type: tok.Text, Text: "Paragraph."},
// Back1Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// ZedToken: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// },
// {
// name:     "Septuple next",
// input:    "Test\n=====\n\nParagraph.\n\n",
// nextNum:  7,
// Back4Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// Back3Tok: &tok.Item{Type: tok.Text, Text: "Paragraph."},
// Back2Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// Back1Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// ZedToken: &tok.Item{Type: tok.EOF},
// },
// {
// name:     "Two next() on one line of input",
// input:    "Test",
// nextNum:  2,
// Back1Tok: &tok.Item{Type: tok.Text, Text: "Test"},
// ZedToken: &tok.Item{Type: tok.EOF},
// },
// {
// name:    "Three next() on one line of input; Test channel close.",
// input:   "Test",
// nextNum: 3,
// // The channel should be closed on the second next(), otherwise a deadlock would occur.
// Back2Tok: &tok.Item{Type: tok.Text, Text: "Test"},
// Back1Tok: &tok.Item{Type: tok.EOF},
// },
// {
// name:    "Four next() on one line of input; Test channel close.",
// input:   "Test",
// nextNum: 4,
// // The channel should be closed on the second next(), otherwise a deadlock would occur.
// Back3Tok: &tok.Item{Type: tok.Text, Text: "Test"},
// Back2Tok: &tok.Item{Type: tok.EOF},
// },
// }

// func TestParserNext(t *testing.T) {
// isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
// val := tExp.Interface().(*tok.Item)
// if val == nil && tr.token[tPos] == nil {
// return
// }
// if val == nil && tr.token[tPos] != nil {
// t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
// return
// }
// if tr.token[tPos].Type != val.Type {
// t.Errorf("Test: %q\n\tGot: token[%d].Type = %q, Expect: %q", tr.Name, tPos, tr.token[tPos].Type, val.Type)
// }
// if tr.token[tPos].Text != val.Text && val.Text != "" {
// t.Errorf("Test: %q\n\tGot: token[%d].Text = %q, Expect: %q", tr.Name, tPos, tr.token[tPos].Text, val.Text)
// }
// }
// for _, tt := range parserNextTests {
// testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
// tr := New(tt.name, tt.input, testutil.StdLogger)
// var err error
// tr.lex, err = tok.Lex(tt.name, []byte(tt.input), testutil.StdLogger)
// if err != nil && !tt.expectError {
// t.Errorf("lexer error: %s", err)
// t.Fail()
// }
// tr.next(tt.nextNum)
// checkTokens(tr, tt, isEqual)
// }
// }

// var parserPeekTests = []struct {
// name        string
// input       string
// nextNum     int // Number of times to call Parser.next() before peek
// peekNum     int // position argument to Parser.peek()
// expectError bool
// Back4Tok    *tok.Item
// Back3Tok    *tok.Item
// Back2Tok    *tok.Item
// Back1Tok    *tok.Item
// ZedToken    *tok.Item
// Peek1Tok    *tok.Item
// Peek2Tok    *tok.Item
// Peek3Tok    *tok.Item
// Peek4Tok    *tok.Item
// }{
// {
// name:     "Single peek no next",
// input:    "Test\n=====\n\nParagraph.",
// peekNum:  1,
// Peek1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// },
// {
// name:     "Double peek no next",
// input:    "Test\n=====\n\nParagraph.",
// peekNum:  2,
// Peek1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// Peek2Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// },
// {
// name:     "Triple peek no next",
// input:    "Test\n=====\n\nParagraph.",
// peekNum:  3,
// Peek1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// Peek2Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// Peek3Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// },
// {
// name:    "Triple peek and double next",
// input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
// nextNum: 2, peekNum: 3,
// Back1Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// ZedToken: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// Peek1Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// Peek2Tok: &tok.Item{Type: tok.Text, Text: "One"},
// Peek3Tok: &tok.Item{Type: tok.Title, Text: "Test 2"},
// },
// {
// name:    "Quadruple peek and triple next",
// input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
// nextNum: 3, peekNum: 4,
// Back2Tok: &tok.Item{Type: tok.Title, Text: "Test"},
// Back1Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// ZedToken: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// Peek1Tok: &tok.Item{Type: tok.Text, Text: "One"},
// Peek2Tok: &tok.Item{Type: tok.Title, Text: "Test 2"},
// Peek3Tok: &tok.Item{Type: tok.SectionAdornment, Text: "====="},
// Peek4Tok: &tok.Item{Type: tok.BlankLine, Text: "\n"},
// },
// {
// name:        "Peek on no input",
// peekNum:     1,
// expectError: true,
// },
// }

// func TestParserPeek(t *testing.T) {
// isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
// val := tExp.Interface().(*tok.Item)
// if val == nil && tr.token[tPos] == nil {
// return
// }
// if val == nil && tr.token[tPos] != nil {
// t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
// return
// }
// if tr.token[tPos].Type != val.Type {
// t.Errorf("Test: %q\n\tGot: token[%s].Type = %q, Expect: %q", tr.Name, tName, tr.token[tPos].Type, val.Type)
// }
// if tr.token[tPos].Text != val.Text && val.Text != "" {
// t.Errorf("Test: %q\n\tGot: token[%s].Text = %q, Expect: %q", tr.Name, tName, tr.token[tPos].Text, val.Text)
// }
// }
// for _, tt := range parserPeekTests {
// testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
// tr := New(tt.name, tt.input, testutil.StdLogger)
// var err error
// tr.lex, err = tok.Lex(tt.name, []byte(tt.input), testutil.StdLogger)
// if err != nil && !tt.expectError {
// t.Errorf("lexer error: %s", err)
// t.Fail()
// }
// tr.next(tt.nextNum)
// tr.peek(tt.peekNum)
// checkTokens(tr, tt, isEqual)
// }
// }

// var testParserClearTokensTests = []struct {
// name       string
// input      string
// nextNum    int       // Number of times to call Parser.next() before peek
// peekNum    int       // position argument to Parser.peek()
// clearBegin int       // Passed to Parser.clear() as the begin arg
// clearEnd   int       // Passed to Parser.clear() as the end arg
// Back4Tok   *tok.Item // Use &item{} if the token is not expected to be nil
// Back3Tok   *tok.Item
// Back2Tok   *tok.Item
// Back1Tok   *tok.Item
// ZedToken   *tok.Item
// Peek1Tok   *tok.Item
// Peek2Tok   *tok.Item
// Peek3Tok   *tok.Item
// Peek4Tok   *tok.Item
// }{
// {
// name:    "Fill token buffer and clear it.",
// input:   "Parser\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
// nextNum: 5, peekNum: 4,
// clearBegin: zed - 4, clearEnd: zed + 4,
// },
// {
// name:    "Fill token buffer and clear back tokens.",
// input:   "Parser\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
// nextNum: 5, peekNum: 4,
// clearBegin: zed - 4, clearEnd: zed - 1,
// ZedToken: &tok.Item{},
// Peek1Tok: &tok.Item{},
// Peek2Tok: &tok.Item{},
// Peek3Tok: &tok.Item{},
// Peek4Tok: &tok.Item{},
// },
// {
// name:    "Fill token buffer and clear peek tokens.",
// input:   "Parser\n====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
// nextNum: 5, peekNum: 4,
// clearBegin: zed + 1, clearEnd: zed + 4,
// Back4Tok: &tok.Item{},
// Back3Tok: &tok.Item{},
// Back2Tok: &tok.Item{},
// Back1Tok: &tok.Item{},
// ZedToken: &tok.Item{},
// },
// }

// func TestParserClearTokens(t *testing.T) {
// isEqual := func(tr *Parser, tExp reflect.Value, tPos int, tName string) {
// val := tExp.Interface().(*tok.Item)
// if tr.token[tPos] != nil && val == nil {
// t.Errorf("Test: %q\n\tGot: token[%s] == %#+v, Expect: nil", tr.Name, tName, tr.token[tPos])
// }
// }
// for _, tt := range testParserClearTokensTests {
// testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
// tr := New(tt.name, tt.input, testutil.StdLogger)
// var err error
// tr.lex, err = tok.Lex(tt.name, []byte(tt.input), testutil.StdLogger)
// if err != nil {
// t.Errorf("lexer error: %s", err)
// t.Fail()
// }
// tr.next(tt.nextNum)
// tr.peek(tt.peekNum)
// tr.clearTokens(tt.clearBegin, tt.clearEnd)
// checkTokens(tr, tt, isEqual)
// }
// }

// func TestSystemMessageLevelFrom(t *testing.T) {
// name := "Test systemMessageLevel with levelInfo"
// test0 := ""
// if -1 != systemMessageLevelFromString(test0) {
// t.Errorf("Test: %q\n\tGot: systemMessageLevel = %q, Expect: %q", name,
// systemMessageLevelFromString(test0), -1)
// }
// test1 := "INFO"
// if levelInfo != systemMessageLevelFromString(test1) {
// t.Errorf("Test: %q\n\tGot: systemMessageLevel = %q, Expect: %q", name,
// systemMessageLevelFromString(test1), levelInfo)
// }
// test2 := "SEVERE"
// if levelInfo != systemMessageLevelFromString(test1) {
// t.Errorf("Test: %q\n\tGot: systemMessageLevel = %q, Expect: %q", name,
// systemMessageLevelFromString(test2), levelSevere)
// }
// }
