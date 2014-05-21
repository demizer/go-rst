// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"fmt"
	"github.com/demizer/go-elog"
	"reflect"
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

type sectionLevels []*SectionNode

func (s *sectionLevels) String() string {
	var out string
	for _, sec := range *s {
		out += fmt.Sprintf("level: %d, rune: %q, overline: %t, length: %d\n",
			sec.Level, sec.UnderLine.Rune, sec.OverLine != nil, sec.Length)
	}
	return out
}

// Returns nil if not found
func (s *sectionLevels) FindByRune(adornChar rune) *SectionNode {
	for _, sec := range *s {
		if sec.UnderLine.Rune == adornChar {
			return sec
		}
	}
	return nil
}

// If exists == true, a section node with the same text and underline has been found in
// sectionLevels, sec is the matching SectionNode. If exists == false, then the sec return value is
// the similarly leveled SectionNode. If exists == false and sec == nil, then the SectionNode added
// to sectionLevels is a new Node.
func (s *sectionLevels) Add(section *SectionNode) (exists bool, sec *SectionNode) {
	sec = s.FindByRune(section.UnderLine.Rune)
	if sec != nil {
		if sec.Text == section.Text {
			return true, sec
		} else if sec.Text != section.Text {
			section.Level = sec.Level
		}
	} else {
		section.Level = len(*s) + 1
	}
	exists = false
	*s = append(*s, section)
	return
}

func (s *sectionLevels) Level() int {
	return len(*s)
}

// Parse is the entry point for the reStructuredText parser.
func Parse(name, text string) (t *Tree, errors []error) {
	t = New(name)
	t.text = text
	_, errors = t.Parse(text, t)
	return
}

func New(name string) *Tree {
	return &Tree{Name: name, Nodes: newList(), nodeTarget: newList(), sectionLevels:
		new(sectionLevels)}
}

type Tree struct {
	Name          string
	Nodes         *NodeList // The root node list
	nodeTarget    *NodeList // Used by the parser to add nodes to a target NodeList
	Errors        []error
	text          string
	lex           *lexer
	peekCount     int
	token         [3]item        // three-token look-ahead for parser.
	sectionLevels *sectionLevels // Encountered section levels
	id            int            // The unique id of the node in the tree
}

func (t *Tree) errorf(format string, args ...interface{}) {
	format = fmt.Sprintf("go-rst: %s:%d: %s\n", t.Name, t.lex.lineNumber(), format)
	t.Errors = append(t.Errors, fmt.Errorf(format, args...))
}

func (t *Tree) error(err error) {
	t.errorf("%s\n", err)
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(lex *lexer) {
	t.lex = lex
}

// stopParse terminates parsing.
func (t *Tree) stopParse() {
	t.Nodes = nil
	t.nodeTarget = nil
	t.lex = nil
}

func (t *Tree) Parse(text string, treeSet *Tree) (tree *Tree, errors []error) {
	log.Debugln("Start")
	t.startParse(lex(t.Name, text))
	t.text = text
	t.parse(treeSet)
	log.Debugln("End")
	return t, t.Errors
}

func (t *Tree) parse(tree *Tree) {
	log.Debugln("Start")

	t.nodeTarget = t.Nodes

	for t.peek().Type != itemEOF {
		var n Node
		token := t.next()
		log.Debugf("Got token: %#+v\n", token)

		switch token.Type {
		case itemTitle: // Section includes overline/underline
			n = t.section(token)
			log.Infof("New Node: %#+v\n", n)
		case itemBlankLine:
			n = newBlankLine(token, &t.id)
			log.Infof("New Node: %#+v\n", n)
		case itemParagraph:
			n = newParagraph(token, &t.id)
			log.Infof("New Node: %#+v\n", n)
		default:
			t.errorf("%q Not implemented!", token.Type)
			continue
		}

		t.nodeTarget.append(n)
		if n.NodeType() == NodeSection {
			t.nodeTarget =
			reflect.ValueOf(n).Elem().FieldByName("NodeList").Addr().Interface().(*NodeList)
		}
	}

	log.Debugln("End")
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
	} else if overline && overAdorn.Text != underAdorn.Text {
		t.errorf("Section title over line does not match section title under line.")
	}

	sec := newSection(title, &t.id, overAdorn, underAdorn)
	exists, eSec := t.sectionLevels.Add(sec)
	if exists && eSec != nil {
		t.errorf("SectionNode using Text \"%s\" and Rune '%s' was previously parsed!",
			sec.Text, string(sec.UnderLine.Rune))
	} else if !exists && eSec != nil {
		// There is a matching level in sectionLevels
		t.nodeTarget = &(*t.sectionLevels)[sec.Level - 2].NodeList
	}

	log.Debugln("End")
	return sec
}
