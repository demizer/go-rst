package parser

import (
	"fmt"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"

	tok "github.com/demizer/go-rst/pkg/token"
)

type bufferTest interface {
	previousToken() *tok.Item
	currentToken() *tok.Item
	nextToken() *tok.Item
}

func tokenIsEqual(t *testing.T, name string, actual *tok.Item, expect *tok.Item) {
	if actual == nil && expect == nil {
		return
	}
	if actual == nil && expect != nil {
		t.Fatalf("Test: %q\tGot: %s\tExpect: %s", name, actual, expect)
	}
	if actual != nil && expect == nil {
		t.Fatalf("Test: %q\tGot: %s\tExpect: %s", name, actual, expect)
	}
	if actual.ID != expect.ID {
		t.Fatalf("Test: %q\tGot: ID = %d\tExpect: %d", name, actual.ID, expect.ID)
	}
	if actual.Type != expect.Type {
		t.Fatalf("Test: %q\tGot: Type = %q\tExpect: %q", name, actual.Type, expect.Type)
	}
	if actual.Text != expect.Text && expect.Text != "" {
		t.Fatalf("Test: %q\tGot: Text = %q\tExpect: %q", name, actual.Text, expect.Text)
	}

}

// checkTokens checks the position of the buffer using three points: the index token, the previous token, and the next token.
func checkTokens(t *testing.T, testName string, p *Parser, bt bufferTest) {
	tokenIsEqual(t, fmt.Sprintf("%s (Previous Token)", testName), p.peekBack(1), bt.previousToken())
	tokenIsEqual(t, fmt.Sprintf("%s (Current Token)", testName), p.token, bt.currentToken())
	tokenIsEqual(t, fmt.Sprintf("%s (Next Token)", testName), p.peek(1), bt.nextToken())
}

type parserBackupTest struct {
	name       string
	input      string
	nextNum    int       // The number of times to call Parser.next().
	backupNum  int       // Number of calls to Parser.backup(). Value starts at 1.
	backToken  *tok.Item // The expected token before index token
	indexToken *tok.Item // The item to expect at Parser.token
	peekToken  *tok.Item // The expected token after index token
}

func (p parserBackupTest) previousToken() *tok.Item { return p.backToken }

func (p parserBackupTest) currentToken() *tok.Item { return p.indexToken }

func (p parserBackupTest) nextToken() *tok.Item { return p.peekToken }

var parserBackupTests = [...]parserBackupTest{
	{
		name:    "Single backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 1,
		// backToken should be nil
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment},
	},
	{
		name:    "Double backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 2,
		// backToken is nil
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment},
	},
	{
		name:  "Triple backup",
		input: "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		// With backupNum = 3, we try to backup past the beginning of the slice
		nextNum: 2, backupNum: 3,
		// BackToken is nil
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1"},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment},
	},
	{
		name:  "Quadruple backup",
		input: "Title\n=====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		// cycle next() until the end of the lexing
		nextNum: 13, backupNum: 4,
		backToken:  &tok.Item{ID: 8, Type: tok.Text, Text: "Three"},
		indexToken: &tok.Item{ID: 9, Type: tok.BlankLine, Text: "\n"},
		peekToken:  &tok.Item{ID: 10, Type: tok.Text, Text: "Four"},
	},
}

func TestParserBackup(t *testing.T) {
	for _, tt := range parserBackupTests {
		tr, err := NewParser(tt.name, tt.input, testutil.StdLogger)
		if err != nil {
			t.Errorf("error: %s", err)
			t.Fail()
		}
		tr.next(tt.nextNum)
		tr.Dump(tr.buf)
		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}
		checkTokens(t, tt.name, tr, tt)
		fmt.Printf("PASSED %q\n", tt.name)
	}
}

// var parserNextTests = []struct {
// name        string
// input       string
// nextNum     int // Number of times to call Parser.next(). Value starts at 1.
// expectError bool
// Back4Tok    *tok.Item // The item to expect at Parser.buf[zed-4]
// Back3Tok    *tok.Item // The item to expect at Parser.buf[zed-3]
// Back2Tok    *tok.Item // The item to expect at Parser.buf[zed-2]
// Back1Tok    *tok.Item // The item to expect at Parser.buf[zed-1]
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
// if val == nil && tr.buf[tPos] == nil {
// return
// }
// if val == nil && tr.buf[tPos] != nil {
// t.Errorf("Test: %q\n\tGot: buf[%s] == %#+v, Expect: nil", tr.Name, tName, tr.buf[tPos])
// return
// }
// if tr.buf[tPos].Type != val.Type {
// t.Errorf("Test: %q\n\tGot: buf[%d].Type = %q, Expect: %q", tr.Name, tPos, tr.buf[tPos].Type, val.Type)
// }
// if tr.buf[tPos].Text != val.Text && val.Text != "" {
// t.Errorf("Test: %q\n\tGot: buf[%d].Text = %q, Expect: %q", tr.Name, tPos, tr.buf[tPos].Text, val.Text)
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
// if val == nil && tr.buf[tPos] == nil {
// return
// }
// if val == nil && tr.buf[tPos] != nil {
// t.Errorf("Test: %q\n\tGot: buf[%s] == %#+v, Expect: nil", tr.Name, tName, tr.buf[tPos])
// return
// }
// if tr.buf[tPos].Type != val.Type {
// t.Errorf("Test: %q\n\tGot: buf[%s].Type = %q, Expect: %q", tr.Name, tName, tr.buf[tPos].Type, val.Type)
// }
// if tr.buf[tPos].Text != val.Text && val.Text != "" {
// t.Errorf("Test: %q\n\tGot: buf[%s].Text = %q, Expect: %q", tr.Name, tName, tr.buf[tPos].Text, val.Text)
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
// if tr.buf[tPos] != nil && val == nil {
// t.Errorf("Test: %q\n\tGot: buf[%s] == %#+v, Expect: nil", tr.Name, tName, tr.buf[tPos])
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
