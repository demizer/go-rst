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
		case itemTitle, itemBlankLine:
			// itemTitle is consumed when evaluating itemSectionAdornment
			continue
		case itemEOF:
			goto exit
		default:
			panic(fmt.Errorf("%q Not implemented!", token.Type))
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

func (t *Tree) backup() *item {
	t.tokenBackupCount++
	// log.Debugln("t.tokenBackupCount:", t.peekCount)
	for i := len(t.token) - 1; i > 0; i-- {
		t.token[i] = t.token[i-1]
		t.token[i-1] = nil
	}
	// log.Debugf("\n##### backup() aftermath #####\n\n")
	// spd.Dump(t.token)
	return t.token[zed-t.tokenBackupCount]
}

func (t *Tree) peekBack(pos int) *item {
	return t.token[zed-pos]
}

func (t *Tree) peek(pos int) *item {
	// log.Debugln("t.peekCount:", t.peekCount, "Pos:", pos)
	if pos < 1 {
		panic("pos cannot be < 1")
	}
	var nItem *item
	for i := 0; i < pos; i++ {
		// log.Debugln("i:", i, "peekCount:", t.peekCount, "pos:", pos)
		if t.peekCount > i {
			nItem = t.token[zed+i]
			log.Debugf("Using %#+v\n", nItem)
			continue
		}
		log.Debugln(zed + t.peekCount + i)
		if t.token[zed+t.peekCount+i+1] == nil {
			t.peekCount++
			// log.Debugln("Getting next item")
			t.token[zed+t.peekCount+i] = t.lex.nextItem()
			nItem = t.token[zed+t.peekCount+i]
		} else {
			nItem = t.token[zed+t.peekCount+i]
		}
	}
	// log.Debugf("\n##### peek() aftermath #####\n\n")
	// spd.Dump(t.token)
	// log.Debugf("Returning: %#+v\n", nItem)
	return nItem
}

func (t *Tree) peekSkip(pos int, iSkip itemElement) *item {
	var nItem *item
outer:
	for i := 1; i <= pos; i++ {
		for {
			nItem = t.peek(i)
			if nItem.Type == iSkip {
				continue
			} else {
				break outer
			}
		}
	}
	return nItem
}

func (t *Tree) next() *item {
	// log.Debugln("t.peekCount:", t.peekCount)
	// skip shifts the pointers left in t.token, pos is the amount to shift
	skip := func(num int) {
		for i := num; i > 0; i-- {
			for x := 0; x < len(t.token)-1; x++ {
				t.token[x] = t.token[x+1]
				t.token[x+1] = nil
			}
		}
	}
	if t.peekCount > 0 {
		skip(t.peekCount)
	} else {
		skip(1)
		t.token[zed] = t.lex.nextItem()
	}
	t.tokenBackupCount, t.peekCount = 0, 0
	// log.Debugf("\n##### next() aftermath #####\n\n")
	// spd.Dump(t.token)
	return t.token[zed]
}

func (t *Tree) section(i *item) Node {
	log.Debugln("Start")
	var overAdorn, title, underAdorn *item
	var overline bool
	var sysMessage Node

	peekForward := t.peekSkip(1, itemSpace)
	if peekForward != nil && peekForward.Type == itemTitle {
		log.Debugln("FOUND SECTION WITH OVERLINE")
		if peekBack := t.peekBack(1); peekBack != nil && peekBack.Type == itemSpace {
			return t.systemMessage(errorUnexpectedSectionTitleOrTransition)
		}
		overAdorn = i
		t.next()
	loop:
		for {
			switch tTok := t.token[zed]; tTok.Type {
			case itemSpace:
				t.next()
			case itemTitle:
				title = tTok
				t.next()
				cur := t.token[zed]
				if cur != nil && cur.Type == itemSectionAdornment {
					continue
				}
				if pNext := t.peek(1); pNext != nil && pNext.Type != itemSectionAdornment {
					panic("Missing section underline!")
				}
			case itemSectionAdornment:
				underAdorn = tTok
				break loop
			}
		}
	} else {
		peekBack := t.peekBack(1)
		if peekBack != nil {
			if peekBack.Type == itemSpace {
				// Looking back past the white space
				if t.peekBack(2).Type == itemTitle {
					return t.systemMessage(errorUnexpectedSectionTitle)
				}
				return t.systemMessage(errorUnexpectedSectionTitleOrTransition)
			} else if peekBack.Type == itemTitle {
				if t.peekBack(2) != nil && t.peekBack(2).Type ==
					itemSectionAdornment {
					// The overline of the section
					overline = true
					overAdorn = peekBack
				}
			}
		}
		title = t.peekBack(1)
		underAdorn = i
	}

	// TODO: Change these into proper error messages!
	// Check adornment for proper syntax
	if underAdorn.Type == itemSpace {
		t.backup() // Put the parser back on the title
		return t.systemMessage(errorUnexpectedSectionTitle)
	} else if overline && title.Length != overAdorn.Length {
		panic("Section over line not equal to title length!")
	} else if overline && overAdorn.Text != underAdorn.Text {
		panic("Section title over line does not match section title under line.")
	}

	sec := newSection(title, overAdorn, underAdorn, &t.id)
	exists, eSec := t.sectionLevels.Add(sec)
	if exists && eSec != nil {
		panic(fmt.Errorf("SectionNode using Text \"%s\" and Rune '%s' was previously parsed!",
			sec.Text, string(sec.UnderLine.Rune)))
	} else if !exists && eSec != nil {
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
