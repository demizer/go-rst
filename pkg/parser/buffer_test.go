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

// checkTokens checks the position of the buffer using three points: the index token, the previous token, and the next token.
func checkTokens(t *testing.T, backToken *tok.Item, indexToken *tok.Item, nextToken *tok.Item, bt bufferTest) {
	assert.Equal(t, bt.previousToken(), backToken, fmt.Sprintf("%s (backToken)", bt.testName()))
	assert.Equal(t, bt.currentToken(), indexToken, fmt.Sprintf("%s (indexToken)", bt.testName()))
	assert.Equal(t, bt.nextToken(), nextToken, fmt.Sprintf("%s (nextToken)", bt.testName()))
}

type parserBackupTest struct {
	name        string
	input       string
	nextNum     int       // The number of times to call Parser.next().
	backupNum   int       // Number of calls to Parser.backup(). Value starts at 1.
	expectIndex int       // The expected token buffer index
	backToken   *tok.Item // The expected token before index token
	indexToken  *tok.Item // The item to expect at Parser.token
	peekToken   *tok.Item // The expected token after index token
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
		expectIndex: 0,
		// backToken should be nil
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1", Line: 1, StartPosition: 1, Length: 7},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=======", Line: 2, StartPosition: 1, Length: 7},
	},
	{
		name:    "Double backup",
		input:   "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		nextNum: 2, backupNum: 2,
		// backToken is nil
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1", Line: 1, StartPosition: 1, Length: 7},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=======", Line: 2, StartPosition: 1, Length: 7},
	},
	{
		name:  "Triple backup",
		input: "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2.",
		// With backupNum = 3, we try to backup past the beginning of the slice
		nextNum: 2, backupNum: 3,
		// BackToken is nil
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1", Line: 1, StartPosition: 1, Length: 7},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=======", Line: 2, StartPosition: 1, Length: 7},
	},
	{
		name:  "Quadruple backup",
		input: "Title\n=====\n\nOne\n\nTwo\n\nThree\n\nFour\n\nFive",
		// cycle next() until the end of the lexing (EOF)
		nextNum: 13, backupNum: 4,
		expectIndex: 8,
		backToken:   &tok.Item{ID: 8, Type: tok.Text, Text: "Three", Line: 8, StartPosition: 1, Length: 5},
		indexToken:  &tok.Item{ID: 9, Type: tok.BlankLine, Text: "\n", Line: 9, StartPosition: 1, Length: 1},
		peekToken:   &tok.Item{ID: 10, Type: tok.Text, Text: "Four", Line: 10, StartPosition: 1, Length: 4},
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

		for j := 0; j < tt.backupNum; j++ {
			tr.backup()
		}

		assert.Equal(t, tt.expectIndex, tr.index, fmt.Sprintf("Expect token buffer index to be the same for test %q", tt.name))

		checkTokens(t, tr.peekBack(1), tr.token, tr.peek(1), tt)

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
	expectIndex int       // The expected token buffer index
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
		indexToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		peekToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
	},
	{
		name:    "Double next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 2, expectIndex: 1,
		backToken:  &tok.Item{ID: 1, Type: tok.Title, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		indexToken: &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
		peekToken:  &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n", Line: 3, StartPosition: 1, Length: 1},
	},
	{
		name:    "Triple next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 3, expectIndex: 2,
		backToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
		indexToken: &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n", Line: 3, StartPosition: 1, Length: 1},
		peekToken:  &tok.Item{ID: 4, Type: tok.Text, Text: "Paragraph.", Line: 4, StartPosition: 1, Length: 10},
	},
	{
		name:    "Quadruple next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 4, expectIndex: 3,
		backToken:  &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n", Line: 3, StartPosition: 1, Length: 1},
		indexToken: &tok.Item{ID: 4, Type: tok.Text, Text: "Paragraph.", Line: 4, StartPosition: 1, Length: 10},
		peekToken:  &tok.Item{ID: 5, Type: tok.EOF, Line: 4, StartPosition: 11, Length: 0},
	},
	{
		name:    "Quintuple next",
		input:   "Test\n=====\n\nParagraph.\n\n",
		nextNum: 5, expectIndex: 4,
		backToken:  &tok.Item{ID: 4, Type: tok.Text, Text: "Paragraph.", Line: 4, StartPosition: 1, Length: 10},
		indexToken: &tok.Item{ID: 5, Type: tok.BlankLine, Text: "\n", Line: 5, StartPosition: 1, Length: 1},
		peekToken:  &tok.Item{ID: 6, Type: tok.BlankLine, Text: "\n", Line: 6, StartPosition: 1, Length: 1},
	},
	{
		name:    "Sextuple next",
		input:   "Test\n=====\n\nParagraph.\n\n",
		nextNum: 6, expectIndex: 5,
		backToken:  &tok.Item{ID: 5, Type: tok.BlankLine, Text: "\n", Line: 5, StartPosition: 1, Length: 1},
		indexToken: &tok.Item{ID: 6, Type: tok.BlankLine, Text: "\n", Line: 6, StartPosition: 1, Length: 1},
		peekToken:  &tok.Item{ID: 7, Type: tok.EOF, Line: 6, StartPosition: 1, Length: 0},
	},
	{
		name:    "Septuple next",
		input:   "Test\n=====\n\nParagraph.\n\n",
		nextNum: 7, expectIndex: 6,
		backToken:  &tok.Item{ID: 6, Type: tok.BlankLine, Text: "\n", Line: 6, StartPosition: 1, Length: 1},
		indexToken: &tok.Item{ID: 7, Type: tok.EOF, Line: 6, StartPosition: 1, Length: 0},
	},
	{
		name:    "Two next() on one line of input",
		input:   "Test",
		nextNum: 2, expectIndex: 1,
		backToken:  &tok.Item{ID: 1, Type: tok.Text, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		indexToken: &tok.Item{ID: 2, Type: tok.EOF, Line: 1, StartPosition: 5, Length: 0},
	},
	{
		name:    "Three next() on one line of input; Test channel close.",
		input:   "Test",
		nextNum: 4, expectIndex: 1, // Index is at EOF
		// The channel should be closed on the second next(), otherwise a deadlock would occur.
		backToken:  &tok.Item{ID: 1, Type: tok.Text, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		indexToken: &tok.Item{ID: 2, Type: tok.EOF, Line: 1, StartPosition: 5, Length: 0},
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

		assert.Equal(t, tt.expectIndex, tr.index, fmt.Sprintf("Expect token buffer index to be the same for test %q", tt.name))

		checkTokens(t, tr.peekBack(1), tr.token, tr.peek(1), tt)

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
	expectIndex int
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
		name:    "Single peek no next",
		input:   "Test\n=====\n\nParagraph.",
		peekNum: 1, expectIndex: -1,
		peekToken: &tok.Item{ID: 1, Type: tok.Title, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
	},
	{
		name:    "Double peek no next",
		input:   "Test\n=====\n\nParagraph.",
		peekNum: 2, expectIndex: -1,
		peekToken: &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
	},
	{
		name:    "Triple peek no next",
		input:   "Test\n=====\n\nParagraph.",
		peekNum: 3, expectIndex: -1,
		peekToken: &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n", Line: 3, StartPosition: 1, Length: 1},
	},
	{
		name:    "Triple peek and double next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 2, peekNum: 3, expectIndex: 1,
		backToken:  &tok.Item{ID: 1, Type: tok.Title, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		indexToken: &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
		peekToken:  &tok.Item{ID: 5, Type: tok.Title, Text: "Test 2", Line: 5, StartPosition: 1, Length: 6},
	},
	{
		name:    "Quadruple peek and triple next",
		input:   "Test\n=====\n\nOne\nTest 2\n=====\n\nTwo",
		nextNum: 3, peekNum: 4, expectIndex: 2,
		backToken:  &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
		indexToken: &tok.Item{ID: 3, Type: tok.BlankLine, Text: "\n", Line: 3, StartPosition: 1, Length: 1},
		peekToken:  &tok.Item{ID: 7, Type: tok.BlankLine, Text: "\n", Line: 7, StartPosition: 1, Length: 1},
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
		pt := tr.peek(tt.peekNum)

		assert.Equal(t, tt.expectIndex, tr.index, fmt.Sprintf("Expect token buffer index to be the same for test %q", tt.name))

		checkTokens(t, tr.peekBack(1), tr.token, pt, tt)

		if !t.Failed() {
			fmt.Printf("PASS %s\n", tt.name)
		}
	}
}

type parserPeekBackTest struct {
	name        string
	input       string
	nextNum     int // Number of times to call Parser.next() before peek
	peekBackNum int // position argument to Parser.peekBack()
	expectError bool
	expectIndex int       // The expected token buffer index
	backToken   *tok.Item // The expected token before index token
	indexToken  *tok.Item // The item to expect at Parser.token
	peekToken   *tok.Item // The expected token after index token
}

func (p parserPeekBackTest) testName() string { return p.name }

func (p parserPeekBackTest) previousToken() *tok.Item { return p.backToken }

func (p parserPeekBackTest) currentToken() *tok.Item { return p.indexToken }

func (p parserPeekBackTest) nextToken() *tok.Item { return p.peekToken }

var parserPeekBackTests = [...]parserPeekBackTest{
	{
		name:        "Single peekBack no next",
		input:       "Test\n=====\n\nParagraph.",
		peekBackNum: 1,
		expectIndex: -1, // Index is initialized to -1 until next() is called
	},
	{
		name:    "Single peekBack with one next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 1, peekBackNum: 1,
		indexToken: &tok.Item{ID: 1, Type: 2, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
	},
	{
		name:    "Single peekBack with two next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 2, peekBackNum: 1,
		expectIndex: 1,
		backToken:   &tok.Item{ID: 1, Type: 2, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		indexToken:  &tok.Item{ID: 2, Type: 3, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
	},
	{
		name:    "Single peekBack with double next",
		input:   "Test\n=====\n\nParagraph.",
		nextNum: 2, peekBackNum: 1,
		expectIndex: 1,
		backToken:   &tok.Item{ID: 1, Type: 2, Text: "Test", Line: 1, StartPosition: 1, Length: 4},
		indexToken:  &tok.Item{ID: 2, Type: 3, Text: "=====", Line: 2, StartPosition: 1, Length: 5},
	},
}

func TestParserPeekBack(t *testing.T) {
	for _, tt := range parserPeekBackTests {
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
		pb := tr.peekBack(tt.peekBackNum)

		assert.Equal(t, tt.expectIndex, tr.index, fmt.Sprintf("Expect token buffer index to be the same for test %q", tt.name))
		assert.Equal(t, tt.backToken, pb, fmt.Sprintf("Expect token from peekBack(1) for test: %q", tt.name))

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

func TestParserAppend(t *testing.T) {
	var input string
	for x := 0; x < 100; x++ {
		input += "\na line\n"
	}
	testutil.Log(fmt.Sprintf("input: %s", input))
	tr, err := NewParser("fillcapacitytest", input, testutil.StdLogger)
	if err != nil {
		t.Errorf("error: %s", err)
		t.Fail()
	}
	tr.next(203)
	assert.Equal(t, 201, tr.index, "expect index to equal 201")
	assert.Equal(t, &tok.Item{ID: 202, Type: tok.EOF, Line: 201, StartPosition: 1, Length: 0}, tr.token, "expect index token")
}

func TestParserPeekSkip(t *testing.T) {
	input := "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2."
	tr, err := NewParser("peekSkip", input, testutil.StdLogger)
	if err != nil {
		t.Errorf("error: %s", err)
		t.Fail()
	}
	ps := tr.peekSkip(tok.Title)
	assert.Equal(t, -1, tr.index, "expect index to equal -1")
	assert.Equal(t, &tok.Item{ID: 2, Type: tok.SectionAdornment, Text: "=======", Line: 2, StartPosition: 1, Length: 7}, ps, "expect peek skip token")
}

func TestParserPeekBackTo(t *testing.T) {
	input := "Title 1\n=======\n\nParagraph 1.\n\nParagraph 2."
	tr, err := NewParser("peekSkip", input, testutil.StdLogger)
	if err != nil {
		t.Errorf("error: %s", err)
		t.Fail()
	}
	tr.next(4)
	pb := tr.peekBackTo(tok.Title)
	assert.Equal(t, 3, tr.index, "expect index to equal 3")
	assert.Equal(t, &tok.Item{ID: 1, Type: tok.Title, Text: "Title 1", Line: 1, StartPosition: 1, Length: 7}, pb, "expect peek back token")
}
