package parser

import (
	"fmt"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"

	tok "github.com/demizer/go-rst/pkg/token"

	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, expect, actual, name)
	}
	if actual != nil && expect == nil {
		assert.Equal(t, expect, actual, name)
	}
	if actual != nil && expect != nil {
		if actual.ID != expect.ID {
			assert.Equal(t, expect.ID, actual.ID, name)
		}
		if actual.Type != expect.Type {
			assert.Equal(t, expect.Type, actual.Type, name)
		}
		if actual.Text != expect.Text && expect.Text != "" {
			assert.Equal(t, expect.Text, actual.Type, name)
		}
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
		fmt.Printf("RUN  %s\n", tt.name)
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

		if !t.Failed() {
			fmt.Printf("PASS %s\n", tt.name)
		}
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
		fmt.Printf("RUN  %s\n", tt.name)
		tr, err := NewParser(tt.name, tt.input, testutil.StdLogger)
		if err != nil && tt.expectError {
			fmt.Printf("PASS %s\n", tt.name)
			continue
		}

		if err != nil && !tt.expectError {
			t.Errorf("lexer error: %s", err)
			t.Fail()
		}

		tr.next(tt.nextNum)
		tr.Msgr("haz index", "index", tr.index)

		checkTokens(t, 1, 1, tr, tt)
		if !t.Failed() {
			fmt.Printf("PASS %s\n", tt.name)
		}
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
		fmt.Printf("RUN  %s\n", tt.name)
		tr, err := NewParser(tt.name, tt.input, testutil.StdLogger)
		if err != nil && tt.expectError {
			fmt.Printf("PASS %s\n", tt.name)
			continue
		}
		if err != nil && !tt.expectError {
			t.Errorf("lexer error: %s", err)
			t.Fail()
		}
		tr.next(tt.nextNum)
		tr.peek(tt.peekNum)

		checkTokens(t, 1, tt.peekNum, tr, tt)

		if !t.Failed() {
			fmt.Printf("PASS %s\n", tt.name)
		}
	}
}

func TestParserNextAfterPeekAtEOF(t *testing.T) {
	input := "Test\n=====\n\nParagraph."

	tr, err := NewParser("nextAfterPeekAtEOF", input, testutil.StdLogger)
	if err != nil {
		t.Errorf("error: %s", err)
		t.Fail()
	}

	n := tr.next(3)
	assert.Equal(t, &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n", Line: 3, StartPosition: 1, Length: 1}, n, "expect token from next(3)")

	pk := tr.peek(2)
	assert.Equal(t, &tok.Item{ID: 5, Type: tok.EOF, Line: 4, StartPosition: 11}, pk, "expect token from peek(2)")

	nt := tr.next(1)
	assert.Equal(t, &tok.Item{ID: 4, Type: tok.Text, Text: "Paragraph.", Line: 4, StartPosition: 1, Length: 10}, nt, "expect token from next() after peek()")
}

func TestParserNextPeekNextInComment(t *testing.T) {
	input := ".. A comment.\n\nParagraph.\n"

	tr, err := NewParser("nextAfterPeekAtEOF", input, testutil.StdLogger)
	if err != nil {
		t.Errorf("error: %s", err)
		t.Fail()
	}

	assert.Equal(t, -1, tr.index, "expect index to equal -1")

	n := tr.next(1)
	assert.Equal(t, &tok.Item{ID: 1, Type: tok.CommentMark, Text: "..", Line: 1, StartPosition: 1, Length: 2}, n, "expect token from next(1)")

	assert.Equal(t, 0, tr.index, "expect index to equal 0")

	pk := tr.peek(2)
	assert.Equal(t, &tok.Item{ID: 3, Type: tok.Text, Text: "A comment.", Line: 1, StartPosition: 4, Length: 10}, pk, "expect token from peek(2)")
	assert.Equal(t, 0, tr.index, "expect index to equal 0")

	nt := tr.next(2)
	assert.Equal(t, &tok.Item{ID: 3, Type: tok.Text, Text: "A comment.", Line: 1, StartPosition: 4, Length: 10}, nt, "expect token from next(2) after peek(2)")
	assert.Equal(t, 2, tr.index, "expect index to equal 2")

}
