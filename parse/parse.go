// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"code.google.com/p/go.text/unicode/norm"
	"fmt"
	"github.com/demizer/go-elog"
	"github.com/demizer/go-spew/spew"
	"reflect"
)

var spd = spew.ConfigState{Indent: "\t", DisableMethods: true}

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

type parserMessage int

const (
	warningShortUnderline parserMessage = iota
	errorUnexpectedSectionTitle
	errorUnexpectedSectionTitleOrTransition
)

var parserErrors = [...]string{
	"warningShortUnderline",
	"errorUnexpectedSectionTitle",
	"errorUnexpectedSectionTitleOrTransition",
}

func (p parserMessage) String() string {
	return parserErrors[p]
}

func (p parserMessage) Message() (s string) {
	switch p {
	case warningShortUnderline:
		s = "Title underline too short."
	case errorUnexpectedSectionTitle:
		s = "Unexpected section title."
	case errorUnexpectedSectionTitleOrTransition:
		s = "Unexpected section title or transition."
	}
	return
}

func (p parserMessage) Level() (s systemMessageLevel) {
	switch p {
	case warningShortUnderline:
		s = levelWarning
	case errorUnexpectedSectionTitle:
		s = levelSevere
	case errorUnexpectedSectionTitleOrTransition:
		s = levelSevere
	}
	return
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
	if !norm.NFC.IsNormalString(text) {
		text = norm.NFC.String(text)
	}
	t.text = text
	_, errors = t.Parse(text, t)
	return
}

func New(name string) *Tree {
	return &Tree{
		Name:          name,
		Nodes:         newList(),
		nodeTarget:    newList(),
		sectionLevels: new(sectionLevels),
		indentWidth:   indentWidth,
	}
}

const (
	zed         = 3
	indentWidth = 4 // Default indent width
)

type Tree struct {
	Name             string
	Nodes            *NodeList // The root node list
	nodeTarget       *NodeList // Used by the parser to add nodes to a target NodeList
	Errors           []error
	text             string
	lex              *lexer
	tokenBackupCount int
	peekCount        int
	token            [7]*item
	sectionLevels    *sectionLevels // Encountered section levels
	id               int            // The unique id of the node in the tree
	indentWidth      int
	indentLevel      int
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(lex *lexer) {
	t.lex = lex
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

	for t.peek(1).Type != itemEOF {
		var n Node

		token := t.next()
		log.Infof("\nParser got token: %#+v\n\n", token)

		switch token.Type {
		case itemSectionAdornment:
			n = t.section(token)
		case itemParagraph:
			n = newParagraph(token, &t.id)
		case itemSpace:
			n = t.indent(token)
			if n == nil {
				continue
			}
		case itemEOF:
			goto exit
		case itemTitle, itemBlankLine:
			// itemTitle is consumed when evaluating itemSectionAdornment
			continue
		}

		t.nodeTarget.append(n)
		switch n.NodeType() {
		case NodeSection, NodeBlockQuote:
			// Set the loop to append items to the NodeList of the new section
			t.nodeTarget = reflect.ValueOf(n).Elem().FieldByName("NodeList").Addr().Interface().(*NodeList)
		}
	}

exit:
	log.Debugln("End")
}

func (t *Tree) peekBack(pos int) *item {
	return t.token[zed-pos]
}

func (t *Tree) peek(pos int) *item {
	// log.Debugln("\n", "Pos:", pos)
	// log.Debugf("##### peek() before #####\n")
	// spd.Dump(t.token)
	nItem := t.token[zed]
	for i := 1; i <= pos; i++ {
		if t.token[zed+i] != nil {
			nItem = t.token[zed+i]
			log.Debugf("Using %#+v\n", nItem)
			continue
		} else {
			log.Debugln("Getting next item")
			t.token[zed+i] = t.lex.nextItem()
			nItem = t.token[zed+i]
		}
	}
	// log.Debugf("\n##### peek() aftermath #####\n")
	// spd.Dump(t.token)
	// log.Debugf("Returning: %#+v\n", nItem)
	return nItem
}

func (t *Tree) peekSkip(iSkip itemElement) *item {
	var nItem *item
	var count int = 1
	for {
		nItem = t.peek(count)
		if nItem.Type != iSkip {
			 break
		}
		count++
	}
	return nItem
}

func (t *Tree) next() *item {
	// log.Debugf("\n##### next() before #####\n")
	// spd.Dump(t.token)
	for x := 0; x < len(t.token)-1; x++ {
		t.token[x] = t.token[x+1]
		t.token[x+1] = nil
	}
	if t.token[zed] == nil {
		t.token[zed] = t.lex.nextItem()
	}
	// log.Debugf("\n##### next() aftermath #####\n\n")
	// spd.Dump(t.token)
	return t.token[zed]
}

func (t *Tree) section(i *item) Node {
	log.Debugln("Start")
	var overAdorn, title, underAdorn *item
	var sysMessage Node

	peekForward := t.peekSkip(1, itemSpace)
	if peekForward != nil && peekForward.Type == itemTitle {
		log.Debugln("FOUND SECTION WITH OVERLINE")
		overAdorn = i
		t.next()
	loop:
		for {
			switch tTok := t.token[zed]; tTok.Type {
			case itemTitle:
				title = tTok
				t.next()
				cur := t.token[zed]
				if cur != nil && cur.Type == itemSectionAdornment {
					continue
				}
			case itemSectionAdornment:
				underAdorn = tTok
				break loop
			}
		}
	} else {
		if peekBack := t.peekBack(1); peekBack != nil && peekBack.Type == itemSpace {
			// Looking back past the white space
			if t.peekBack(2).Type == itemTitle {
				return t.systemMessage(errorUnexpectedSectionTitle)
			}
			return t.systemMessage(errorUnexpectedSectionTitleOrTransition)
		}
		title = t.peekBack(1)
		underAdorn = i
	}

	sec := newSection(title, overAdorn, underAdorn, indent, &t.id)
	exists, eSec := t.sectionLevels.Add(sec)
	if !exists && eSec != nil {
		// There is a matching level in sectionLevels
		t.nodeTarget = &(*t.sectionLevels)[sec.Level-2].NodeList
	}

	// System messages have to be applied after the section is created in order to preserve
	// a consecutive id number.
	if title.Length != underAdorn.Length {
		sysMessage = t.systemMessage(warningShortUnderline)
		sec.NodeList = append(sec.NodeList, sysMessage)
	}

	log.Debugln("End")
	return sec
}

func (t *Tree) systemMessage(err parserMessage) Node {
	var lbText string
	var lbTextLen int
	var backToken int

	s := newSystemMessage(&item{
		Type: itemSystemMessage,
		Line: t.token[zed].Line,
	},
		err.Level(), &t.id)

	msg := newParagraph(&item{
		Text:   err.Message(),
		Length: len(err.Message()),
	}, &t.id)

	log.Debugln("FOUND", err)

	switch err {
	case warningShortUnderline, errorUnexpectedSectionTitle:
		backToken = zed - 1
		if t.peekBack(1).Type == itemSpace {
			backToken = zed - 2
		}
		lbText = t.token[backToken].Text.(string) + "\n" + t.token[zed].Text.(string)
		lbTextLen = len(lbText) + 1
	case errorUnexpectedSectionTitleOrTransition:
		lbText = t.token[zed].Text.(string)
		lbTextLen = len(lbText)
	}

	lb := newLiteralBlock(&item{
		Type:   itemLiteralBlock,
		Text:   lbText,
		Length: lbTextLen, // Add one to account for the backslash
	}, &t.id)

	s.NodeList = append(s.NodeList, msg, lb)
	return s
}

func (t *Tree) indent(i *item) Node {
	level := i.Length / t.indentWidth
	if t.peekBack(1).Type == itemBlankLine {
		if t.indentLevel == level {
			// Append to the current blockquote NodeList
			return nil
		}
		t.indentLevel = level
		return newBlockQuote(&item{Type: itemBlockquote, Line: i.Line}, level, &t.id)
	}
	return nil
}
