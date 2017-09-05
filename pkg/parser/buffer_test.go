package parser

import (
	"fmt"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"

	tok "github.com/demizer/go-rst/pkg/token"
)

type bufferTest interface {
	testName() string
	previousToken() *tok.Item
	currentToken() *tok.Item
	nextToken() *tok.Item
}

func tokenIsEqual(t *testing.T, name string, actual *tok.Item, expect *tok.Item) {
	if expect == nil {
		return
	}
	if actual == nil && expect != nil {
		t.Fatalf("Test: %q\tGot: %s\tExpect: %s", name, actual, expect)
	}
	if actual != nil && expect == nil {
		t.Fatalf("Test: %q\tGot: %s\tExpect: %s", name, actual, expect)
	}
	if actual.ID != expect.ID {
		t.Errorf("Test: %q\tGot: ID = %d\tExpect: %d", name, actual.ID, expect.ID)
	}
	if actual.Type != expect.Type {
		t.Errorf("Test: %q\tGot: Type = %q\tExpect: %q", name, actual.Type, expect.Type)
	}
	if actual.Text != expect.Text && expect.Text != "" {
		t.Errorf("Test: %q\tGot: Text = %q\tExpect: %q", name, actual.Text, expect.Text)
	}
}

// checkTokens checks the position of the buffer using three points: the index token, the previous token, and the next token.
func checkTokens(t *testing.T, peekBackNum int, peekNextNum int, p *Parser, bt bufferTest) {
	if p.index-peekBackNum > 0 {
		tokenIsEqual(t, fmt.Sprintf("%s (Back %d Token)", bt.testName(), peekBackNum), p.buf[p.index-peekBackNum], bt.previousToken())
	}
	tokenIsEqual(t, fmt.Sprintf("%s (Current Token)", bt.testName()), p.token, bt.currentToken())
	if p.index > 0 {
		tokenIsEqual(t, fmt.Sprintf("%s (With t.buf[t.index])", bt.testName()), p.buf[p.index], bt.currentToken())
	}
	tokenIsEqual(t, fmt.Sprintf("%s (Peek %d Token)", bt.testName(), peekNextNum), p.buf[p.index+peekNextNum], bt.nextToken())
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

func (p parserBackupTest) testName() string { return p.name }

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
		fmt.Printf("RUNNING TEST %q\n", tt.name)
		tr, err := NewParser(tt.name, tt.input, testutil.StdLogger)
		if err != nil {
			t.Errorf("error: %s", err)
			t.Fail()
		}
		tr.next(tt.nextNum)

		// XXX: Remove this before merging to master
		// tr.Dump(tr.buf)

		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}
		checkTokens(t, 1, 1, tr, tt)
	}
}

type parserNextTest struct {
	name        string
	input       string
	nextNum     int // Number of times to call Parser.next(). Value starts at 1.
	expectError bool
	backToken   *tok.Item // The expected token before index token
	indexToken  *tok.Item // The item to expect at Parser.token
	peekToken   *tok.Item // The expected token after index token
}

func (p parserNextTest) testName() string { return p.name }

func (p parserNextTest) previousToken() *tok.Item { return p.backToken }

func (p parserNextTest) currentToken() *tok.Item { return p.indexToken }

func (p parserNextTest) nextToken() *tok.Item { return p.peekToken }

var parserNextTests = [...]parserNextTest{
	{
		name:        "Next no input",
		input:       "",
		nextNum:     1,
		expectError: true,
	},
	{
		name:       "Single next from start",
		input:      "Test\n=====\n\nParagraph.",
		nextNum:    1,
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Test"},
	},
	{
		name:       "Double next",
		input:      "Test\n=====\n\nParagraph.",
		nextNum:    2,
		backToken:  &tok.Item{ID: 1, Type: tok.Title, Text: "Test"},
		indexToken: &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "====="},
	},
	{
		name:       "Triple next",
		input:      "Test\n=====\n\nParagraph.",
		nextNum:    3,
		backToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "====="},
		indexToken: &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n"},
	},
	{
		name:       "Quadruple next",
		input:      "Test\n=====\n\nParagraph.",
		nextNum:    4,
		backToken:  &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n"},
		indexToken: &tok.Item{ID: 4, Type: tok.Text, Text: "Paragraph."},
	},
	{
		name:       "Quintuple next",
		input:      "Test\n=====\n\nParagraph.\n\n",
		nextNum:    5,
		backToken:  &tok.Item{ID: 4, Type: tok.Text, Text: "Paragraph."},
		indexToken: &tok.Item{ID: 5, Type: tok.BlankLine, Text: "\n"},
	},
	{
		name:       "Sextuple next",
		input:      "Test\n=====\n\nParagraph.\n\n",
		nextNum:    6,
		backToken:  &tok.Item{ID: 5, Type: tok.BlankLine, Text: "\n"},
		indexToken: &tok.Item{ID: 6, Type: tok.BlankLine, Text: "\n"},
	},
	{
		name:       "Septuple next",
		input:      "Test\n=====\n\nParagraph.\n\n",
		nextNum:    7,
		backToken:  &tok.Item{ID: 6, Type: tok.BlankLine, Text: "\n"},
		indexToken: &tok.Item{ID: 7, Type: tok.EOF},
	},
	{
		name:       "Two next() on one line of input",
		input:      "Test",
		nextNum:    2,
		backToken:  &tok.Item{ID: 1, Type: tok.Text, Text: "Test"},
		indexToken: &tok.Item{ID: 2, Type: tok.EOF},
	},
	{
		name:    "Three next() on one line of input; Test channel close.",
		input:   "Test",
		nextNum: 4,
		// The channel should be closed on the second next(), otherwise a deadlock would occur.
		backToken:  &tok.Item{ID: 1, Type: tok.Text, Text: "Test"},
		indexToken: &tok.Item{ID: 2, Type: tok.EOF},
	},
}

func TestParserNext(t *testing.T) {
	for _, tt := range parserNextTests {
		fmt.Printf("RUNNING TEST %q\n", tt.name)
		tr, err := NewParser(tt.name, tt.input, testutil.StdLogger)
		if err != nil && tt.expectError {
			continue
		}
		if err != nil && !tt.expectError {
			t.Errorf("lexer error: %s", err)
			t.Fail()
		}
		tr.next(tt.nextNum)
		checkTokens(t, 1, 1, tr, tt)
	}
}

type parserPeekTest struct {
	name        string
	input       string
	nextNum     int // Number of times to call Parser.next() before peek
	peekNum     int // position argument to Parser.peek()
	expectError bool
	backToken   *tok.Item // The expected token before index token
	indexToken  *tok.Item // The item to expect at Parser.token
	peekToken   *tok.Item // The expected token after index token
}

func (p parserPeekTest) testName() string { return p.name }

func (p parserPeekTest) previousToken() *tok.Item { return p.backToken }

func (p parserPeekTest) currentToken() *tok.Item { return p.indexToken }

func (p parserPeekTest) nextToken() *tok.Item { return p.peekToken }

var parserPeekTests = [...]parserPeekTest{
	{
		name:      "Single peek no next",
		input:     "Test\n=====\n\nParagraph.",
		peekNum:   1,
		peekToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Test"},
	},
	{
		name:      "Double peek no next",
		input:     "Test\n=====\n\nParagraph.",
		peekNum:   2,
		peekToken: &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "====="},
	},
	{
		name:      "Triple peek no next",
		input:     "Test\n=====\n\nParagraph.",
		peekNum:   3,
		peekToken: &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n"},
	},
	{
		name:    "Triple peek and double next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 2, peekNum: 3,
		backToken:  &tok.Item{ID: 1, Type: tok.Title, Text: "Test"},
		indexToken: &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "====="},
		peekToken:  &tok.Item{ID: 5, Type: tok.Title, Text: "Test 2"},
	},
	{
		name:    "Quadruple peek and triple next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 3, peekNum: 4,
		backToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "====="},
		indexToken: &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n"},
		peekToken:  &tok.Item{ID: 7, Type: tok.BlankLine, Text: "\n"},
	},
	{
		name:        "Peek on no input",
		peekNum:     1,
		expectError: true,
	},
}

func TestParserPeek(t *testing.T) {
	for _, tt := range parserPeekTests {
		fmt.Printf("RUNNING TEST %q\n", tt.name)
		tr, err := NewParser(tt.name, tt.input, testutil.StdLogger)
		if err != nil && tt.expectError {
			continue
		}
		if err != nil && !tt.expectError {
			t.Errorf("lexer error: %s", err)
			t.Fail()
		}
		tr.next(tt.nextNum)
		tr.peek(tt.peekNum)

		checkTokens(t, 1, tt.peekNum, tr, tt)
	}
}
