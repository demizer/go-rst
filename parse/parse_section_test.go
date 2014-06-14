// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// To enable debug output when testing, use "go test -debug"

package parse

import (
	"math"
	"reflect"
	"strconv"
	"testing"

	"github.com/demizer/go-elog"
)

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
		if tr.token[pos].Text != val.Text && val.Text != nil {
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
		if tr.token[pos].Text != val.Text && val.Text != nil {
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
		if tr.token[pos].Text != val.Text && val.Text != nil {
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

func TestParseSectionTitleGood0000(t *testing.T) {
	// Basic title, underline, blankline, and paragraph test
	testPath := "test_section/01_title_good/00.00_title_paragraph"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0001(t *testing.T) {
	// Basic title, underline, and paragraph with no blankline line after
	// the section.
	testPath := "test_section/01_title_good/00.01_paragraph_noblankline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0002(t *testing.T) {
	// A title that begins with a combining unicode character \u0301. Tests
	// to make sure the 2 byte unicode does not contribute to the underline
	// length calculation.
	testPath := "test_section/01_title_good/00.02_title_combining_chars"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0100(t *testing.T) {
	// A basic section in between paragraphs.
	testPath := "test_section/01_title_good/01.00_para_head_para"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0200(t *testing.T) {
	// Tests section parsing on 3 character long title and underline.
	testPath := "test_section/01_title_good/02.00_short_title"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleGood0300(t *testing.T) {
	// Tests a single section with no other element surrounding it.
	testPath := "test_section/01_title_good/03.00_empty_section"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0000(t *testing.T) {
	// Tests for severe system messages when the sections are indented.
	testPath := "test_section/02_title_bad/00.00_unexpected_titles"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0100(t *testing.T) {
	// Tests for severe system message on short title underline
	testPath := "test_section/02_title_bad/01.00_short_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0200(t *testing.T) {
	// Tests for title underlines that are less than three characters.
	testPath := "test_section/02_title_bad/02.00_short_title_short_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0201(t *testing.T) {
	// Tests for title overlines and underlines that are less than three
	// characters.
	testPath := "test_section/02_title_bad/" +
		"02.01_short_title_short_overline_and_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleBad0202(t *testing.T) {
	// Tests for short title overline with missing underline when the
	// overline is less than three characters.
	testPath := "test_section/02_title_bad/" +
		"02.02_short_title_short_overline_missing_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0000(t *testing.T) {
	// Tests section level return to level one after three subsections.
	testPath := "test_section/03_level_good/" +
		"00.00_section_level_return"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0001(t *testing.T) {
	// Tests section level return to level one after 1 subsection. The
	// second level one section has one subsection.
	testPath := "test_section/03_level_good/" +
		"00.01_section_level_return"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0002(t *testing.T) {
	// Test section level with subsection 4 returning to level two.
	testPath := "test_section/03_level_good/00.02_section_level_return"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0100(t *testing.T) {
	// Tests section level return with title overlines
	testPath := "test_section/03_level_good/01.00_section_level_return"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelGood0200(t *testing.T) {
	// Tests section level with two section having the same rune, but the
	// first not having an overline.
	testPath := "test_section/03_level_good/02.00_two_level_one_overline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0000(t *testing.T) {
	// Test section level return on bad level 2 section adornment
	testPath := "test_section/04_level_bad/00.00_bad_subsection_order"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0001(t *testing.T) {
	// Test section level return with title overlines on bad level 2
	// section adornment
	testPath := "test_section/04_level_bad/" +
		"00.01_bad_subsection_order_with_overlines"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionLevelBad0100(t *testing.T) {
	// Tests for a severeTitleLevelInconsistent system message on a bad
	// level two with an overline. Level one does not have an overline.
	testPath := "test_section/04_level_bad/" +
		"01.00_two_level_overline_bad_return"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0000(t *testing.T) {
	// Test simple section with title overline.
	testPath := "test_section/05_title_with_overline_good/" +
		"00.00_title_overline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0100(t *testing.T) {
	// Test simple section with inset title and overline.
	testPath := "test_section/05_title_with_overline_good/" +
		"01.00_inset_title_with_overline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineGood0200(t *testing.T) {
	// Test sections with three character adornments lines.
	testPath := "test_section/05_title_with_overline_good/" +
		"02.00_three_char_section_title"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0000(t *testing.T) {
	// Test section title with overline, but no underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.00_inset_title_missing_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0001(t *testing.T) {
	// Test inset title with overline but missing underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.01_inset_title_missing_underline_with_blankline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0002(t *testing.T) {
	// Test inset title with overline but missing underline. The title is
	// followed by a blank line and a paragraph.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.02_inset_title_missing_underline_and_para"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0003(t *testing.T) {
	// Test section overline with missmatched underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"00.03_inset_title_mismatched_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0100(t *testing.T) {
	// Test overline with really long title.
	testPath := "test_section/06_title_with_overline_bad/" +
		"01.00_title_too_long"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0200(t *testing.T) {
	// Test overline and underline with blanklines instead of a title.
	testPath := "test_section/06_title_with_overline_bad/" +
		"02.00_missing_titles_with_blankline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0201(t *testing.T) {
	// Test overline and underline with nothing where the title is supposed
	// to be.
	testPath := "test_section/06_title_with_overline_bad/" +
		"02.01_missing_titles_with_noblankline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0300(t *testing.T) {
	// Test two character overline with no underline.
	testPath := "test_section/06_title_with_overline_bad/" +
		"03.00_incomplete_section"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0301(t *testing.T) {
	// Test three character section adornments with no titles or blanklines
	// in between.
	testPath := "test_section/06_title_with_overline_bad/" +
		"03.01_incomplete_sections_no_title"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0400(t *testing.T) {
	// Tests indented section with overline
	testPath := "test_section/06_title_with_overline_bad/" +
		"04.00_indented_title_short_overline_and_underline"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleWithOverlineBad0500(t *testing.T) {
	// Tests ".." overline (which is a comment element).
	testPath := "test_section/06_title_with_overline_bad/" +
		"05.00_two_char_section_title"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleNumberedGood0000(t *testing.T) {
	// Tests lexing a section where the title begins with a number.
	testPath := "test_section/07_title_numbered_good/" +
		"00.00_numbered_title"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}

func TestParseSectionTitleNumberedGood0100(t *testing.T) {
	// Tests numbered section lexing with enumerated directly above
	// section.
	testPath := "test_section/07_title_numbered_good/" +
		"01.00_enum_list_with_numbered_title"
	test := LoadParseTest(t, testPath)
	pTree := parseTest(t, test)
	eNodes := test.expectNodes()
	checkParseNodes(t, eNodes, pTree.Nodes, testPath)
}
