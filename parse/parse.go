// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"fmt"
	"github.com/demizer/go-elog"
)

type systemMessageLevel int

const (
	levelInfo systemMessageLevel = iota
	levelWarning
	levelError
	levelSevere
)

var systemMessageLevels = [...]string{
	"INFO",
	"WARNING",
	"ERROR",
	"SEVERE",
}

func (s systemMessageLevel) String() string {
	return systemMessageLevels[s]
}

type systemMessage struct {
	level  systemMessageLevel
	line   int
	source string
	items  []item
}

type sectionLevel struct {
	char     rune // The adornment character used to describe the section
	overline bool // The section contains an overline
	length   int  // The length of the adornment lines
}

type sectionLevels []sectionLevel

func (s *sectionLevels) String() string {
	var out string
	for lvl, sec := range *s {
		out += fmt.Sprintf("level: %d, rune: %q, overline: %t, length: %d\n",
			lvl+1, sec.char, sec.overline, sec.length)
	}
	return out
}

func (s *sectionLevels) Add(adornChar rune, overline bool, length int) int {
	lvl := s.Find(adornChar)
	if lvl > 0 {
		return lvl
	}
	*s = append(*s, sectionLevel{char: adornChar, overline: overline, length: length})
	return len(*s)
}

// Returns -1 if not found
func (s *sectionLevels) Find(adornChar rune) int {
	for lvl, sec := range *s {
		if sec.char == adornChar {
			return lvl + 1
		}
	}
	return -1
}

func (s *sectionLevels) Level() int {
	return len(*s)
}

func Parse(name, text string) (t *Tree, err error) {
	t = New(name)
	t.text = text
	_, err = t.Parse(text, t)
	return
}

func New(name string) *Tree {
	return &Tree{Name: name, sectionLevels: new(sectionLevels)}
}

type Tree struct {
	Name          string
	Nodes         *NodeList // The root node list
	text          string
	branch        *NodeList // The current branch to add nodes to
	lex           *lexer
	peekCount     int
	token         [3]item        // three-token look-ahead for parser.
	sectionLevel  int            // The current section level of parsing
	sectionLevels *sectionLevels // Encountered section levels
	id            int            // The unique id of the node in the tree
}

func (t *Tree) errorf(format string, args ...interface{}) {
	t.Nodes = nil
	format = fmt.Sprintf("go-rst: %s:%d: %s", t.Name, t.lex.lineNumber(), format)
	panic(fmt.Errorf(format, args...))
}

func (t *Tree) error(err error) {
	t.errorf("%s", err)
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(lex *lexer) {
	t.Nodes = nil
	t.lex = lex
}

// stopParse terminates parsing.
func (t *Tree) stopParse() {
	t.lex = nil
}

func (t *Tree) Parse(text string, treeSet *Tree) (tree *Tree, err error) {
	log.Debugln("Start")
	t.startParse(lex(t.Name, text))
	t.text = text
	t.parse(treeSet)
	log.Debugln("End")
	return t, nil
}

func (t *Tree) parse(tree *Tree) (next Node) {
	log.Debugln("Start")
	t.Nodes = newList()

	for t.peek().Type != itemEOF {
		var n Node
		switch token := t.next(); token.Type {
		case itemTitle: // Section includes overline/underline
			n = t.section(token)
		case itemBlankLine:
			n = newBlankLine(token, &t.id)
		case itemParagraph:
			n = newParagraph(token, &t.id)
		}

		if len([]Node(*t.Nodes)) == 0 {
			t.Nodes.append(n)
		} else {
			t.branch.append(n)
		}
	}

	log.Debugln("End")
	return nil
}

func (t *Tree) backup() {
	t.peekCount++
}

// peekBack returns the last item sent from the lexer.
func (t *Tree) peekBack() item {
	return *t.lex.lastItem
}

// peek returns but does not consume the next token.
func (t *Tree) peek() item {
	if t.peekCount > 0 {
		return t.token[t.peekCount-1]

	}
	t.peekCount = 1
	t.token[0] = t.lex.nextItem()
	return t.token[0]
}

func (t *Tree) next() item {
	if t.peekCount > 0 {
		t.peekCount--
	} else {
		t.token[0] = t.lex.nextItem()
	}
	return t.token[t.peekCount]
}

func (t *Tree) section(i item) Node {
	log.Debugln("Start")
	var overAdorn, title, underAdorn item
	var overline bool

	if t.peekBack().Type == itemSectionAdornment {
		overline = true
		overAdorn = t.peekBack()
	}

	title = i
	underAdorn = t.next() // Grab the section underline

	// Check adornment for proper syntax
	if title.Length != underAdorn.Length {
		t.errorf("Section under line  not equal to title length!")
	} else if overline && title.Length != overAdorn.Length {
		t.errorf("Section over line not equal to title length!")
	} else if overline && overAdorn.Value != underAdorn.Value {
		t.errorf("Section title over line does not match section title under line.")
	}

	// Check section levels to make sure the order of sections seen has not been violated
	if level := t.sectionLevels.Find(rune(underAdorn.Value.(string)[0])); level > 0 {
		if t.sectionLevel == t.sectionLevels.Level() {
			t.sectionLevel++
		} else {
			// The current section level of the parser does not match the previously
			// found section level. This means the user has used incorrect section
			// syntax.
			t.errorf("Incorrect section adornment \"%q\" for section level %d",
				underAdorn.Value.(string)[0], t.sectionLevel)
		}
	} else {
		t.sectionLevel++
	}

	t.sectionLevels.Add(rune(underAdorn.Value.(string)[0]), overline, len(underAdorn.Value.(string)))
	ret := newSection(title, &t.id, t.sectionLevel, overAdorn, underAdorn)
	t.branch = &ret.NodeList

	log.Debugln("End")
	return ret
}
