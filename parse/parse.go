// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// Package parse is a reStructuredText parser implemented in Go!
//
// This package is only meant for lexing and parsing reStructuredText. See the
// top level package documentation for details on using the go-rst package
// package API.
package parse

import (
	"reflect"

	"code.google.com/p/go.text/unicode/norm"
	"github.com/demizer/go-elog"
	"github.com/demizer/go-spew/spew"
)

// Used for debugging only
var spd = spew.ConfigState{Indent: "\t", DisableMethods: true}

// systemMessageLevel implements four levels for messages and is used in
// conjunction with the parserMessage type.
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

// String implments Stringer and return a string of the systemMessageLevel.
func (s systemMessageLevel) String() string {
	return systemMessageLevels[s]
}

// FromString returns the systemMessageLevel converted from the string name.
func systemMessageLevelFromString(name string) systemMessageLevel {
	for num, sLvl := range systemMessageLevels {
		if name == sLvl {
			return systemMessageLevel(num)
		}
	}
	return -1
}

// parserMessage implements messages generated by the parser. Parser messages
// are leveled using systemMessageLevels.
type parserMessage int

const (
	parserMessageNil parserMessage = iota
	infoOverlineTooShortForTitle
	infoUnexpectedTitleOverlineOrTransition
	infoUnderlineTooShortForTitle
	warningShortOverline
	warningShortUnderline
	warningExplicitMarkupWithUnIndent
	errorInvalidSectionOrTransitionMarker
	severeUnexpectedSectionTitle
	severeUnexpectedSectionTitleOrTransition
	severeIncompleteSectionTitle
	severeMissingMatchingUnderlineForOverline
	severeOverlineUnderlineMismatch
	severeTitleLevelInconsistent
)

var parserErrors = [...]string{
	"parserMessageNil",
	"infoOverlineTooShortForTitle",
	"infoUnexpectedTitleOverlineOrTransition",
	"infoUnderlineTooShortForTitle",
	"warningShortOverline",
	"warningShortUnderline",
	"warningExplicitMarkupWithUnIndent",
	"errorInvalidSectionOrTransitionMarker",
	"severeUnexpectedSectionTitle",
	"severeUnexpectedSectionTitleOrTransition",
	"severeIncompleteSectionTitle",
	"severeMissingMatchingUnderlineForOverline",
	"severeOverlineUnderlineMismatch",
	"severeTitleLevelInconsistent",
}

// String implements Stringer and returns the parserMessage as a string. The
// returned string is the parserMessage name, not the message itself.
func (p parserMessage) String() string {
	return parserErrors[p]
}

// Message returns the message of the parserMessage as a string.
func (p parserMessage) Message() (s string) {
	switch p {
	case infoOverlineTooShortForTitle:
		s = "Possible incomplete section title.\n" +
			"Treating the overline as ordinary text because it's so short."
	case infoUnexpectedTitleOverlineOrTransition:
		s = "Unexpected possible title overline or transition.\n" +
			"Treating it as ordinary text because it's so short."
	case infoUnderlineTooShortForTitle:
		s = "Possible title underline, too short for the title.\n" +
			"Treating it as ordinary text because it's so short."
	case warningShortOverline:
		s = "Title overline too short."
	case warningShortUnderline:
		s = "Title underline too short."
	case warningExplicitMarkupWithUnIndent:
		s = "Explicit markup ends without a blank line; unexpected unindent."
	case errorInvalidSectionOrTransitionMarker:
		s = "Invalid section title or transition marker."
	case severeUnexpectedSectionTitle:
		s = "Unexpected section title."
	case severeUnexpectedSectionTitleOrTransition:
		s = "Unexpected section title or transition."
	case severeIncompleteSectionTitle:
		s = "Incomplete section title."
	case severeMissingMatchingUnderlineForOverline:
		s = "Missing matching underline for section title overline."
	case severeOverlineUnderlineMismatch:
		s = "Title overline & underline mismatch."
	case severeTitleLevelInconsistent:
		s = "Title level inconsistent"
	}
	return
}

// Level returns the parserMessage level.
func (p parserMessage) Level() (s systemMessageLevel) {
	lvl := int(p)
	switch {
	case lvl > 0 && lvl <= 3:
		s = levelInfo
	case lvl <= 6:
		s = levelWarning
	case lvl == 7:
		s = levelError
	case lvl >= 8:
		s = levelSevere
	}
	return
}

// sectionLevel is a single section level. sections containes a list of
// pointers to SectionNode that are dertermined to be a section of the level
// indicated by level. rChar is the rune character that denotes the section
// level.
type sectionLevel struct {
	rChar    rune
	level    int
	overLine bool           // If true, the section level should have an overline adornment.
	sections []*SectionNode // New sections matching level are appended here
}

// sectionLevels contains the encountered section levels in order by level.
// levels[0] is section level 1 and levels[1] is section level 2.
// lastSectionNode is a pointer to the lastSectionNode added to levels.
type sectionLevels struct {
	lastSectionNode *SectionNode
	levels          []*sectionLevel
}

// FindByRune loops through the sectionLevels to find a section using a Rune as
// the key. If the section is found, a pointer to the SectionNode is returned.
func (s *sectionLevels) FindByRune(rChar rune) *sectionLevel {
	for _, sec := range s.levels {
		if sec.rChar == rChar {
			return sec
		}
	}
	return nil
}

// Add determines if the underline rune in the sec argument matches any
// existing sectionLevel in sectionLevels. Add also checks the section level
// ordering is correct and returns a severeTitleLevelInconsistent parserMessage
// if inconsistencies are found.
func (s *sectionLevels) Add(sec *SectionNode) (err parserMessage) {
	level := 1
	secLvl := s.FindByRune(sec.UnderLine.Rune)

	// Local function for creating a sectionLevel
	var newSectionLevel = func() {
		var oLine bool
		if sec.OverLine != nil {
			oLine = true
		}
		log.Debugln("Creating new sectionLevel:", level)
		secLvl = &sectionLevel{rChar: sec.UnderLine.Rune, level: level, overLine: oLine}
		s.levels = append(s.levels, secLvl)
		// secLvl.sections = append(secLvl.sections, sec)
	}

	if secLvl == nil {
		if s.lastSectionNode != nil {
			// Check if the provisional level of sec is already in
			// sectionLevels; if it is and the adornment characters don't
			// match, then we have an inconsistent level error.
			level = s.lastSectionNode.Level + 1
			nextLevel := s.SectionLevelByLevel(level)
			if nextLevel != nil && nextLevel.rChar != sec.UnderLine.Rune {
				return severeTitleLevelInconsistent
			}
		} else {
			level = len(s.levels) + 1
		}
		newSectionLevel()
	} else {
		if secLvl.overLine && sec.OverLine == nil ||
			!secLvl.overLine && sec.OverLine != nil {
			// If sec has an OverLine, but the matching sectionLevel with
			// the same Rune as sec does not have an OverLine, then they
			// are not in the same sectionLevel, and visa versa.
			level = len(s.levels) + 1
			newSectionLevel()
		} else {
			log.Debugln("Using existing sectionLevel:", secLvl.level)
			level = secLvl.level
		}
	}

	secLvl.sections = append(secLvl.sections, sec)
	sec.Level = level
	s.lastSectionNode = sec
	return
}

// SectionLevelByLevel returns a pointer to a sectionLevel of level level. Nil
// is returned if l is greater than the number of section levels encountered.
func (s *sectionLevels) SectionLevelByLevel(level int) *sectionLevel {
	if level > len(s.levels) {
		return nil
	}
	return (s.levels)[level-1]
}

// LastSectionByLevel returns a pointer to the last section encountered by
// level.
func (s *sectionLevels) LastSectionByLevel(level int) (sec *SectionNode) {
exit:
	for i := len(s.levels) - 1; i >= 0; i-- {
		if (s.levels)[i].level != level {
			continue
		}
		for j := len((s.levels)[i].sections) - 1; j >= 0; j-- {
			sec = (s.levels)[i].sections[j]
			if sec.Level == level {
				log.Debugln("Found section with level", sec.Level)
				break exit
			}
		}
	}
	return
}

// Parse is the entry point for the reStructuredText parser. Errors generated
// by the parser are returned as a NodeList.
func Parse(name, text string) (t *Tree, errors *NodeList) {
	t = New(name, text)
	if !norm.NFC.IsNormalString(text) {
		text = norm.NFC.String(text)
	}
	t.Parse(text, t)
	errors = t.Messages
	return
}

// New returns a fresh parser tree.
func New(name, text string) *Tree {
	return &Tree{
		Name:          name,
		Nodes:         newList(),
		Messages:      newList(),
		text:          text,
		nodeTarget:    newList(),
		sectionLevels: new(sectionLevels),
		indentWidth:   indentWidth,
	}
}

const (
	// The middle of the Tree.token buffer so that there are three possible
	// "backup" token positions and three forward "peek" positions.
	zed = 4

	// Default indent width
	indentWidth = 4
)

// Tree contains the parser tree. The Nodes field contains the parsed nodes of
// the input input data.
type Tree struct {
	Name          string    // The name of the current parser input
	Nodes         *NodeList // The root node list
	Messages      *NodeList // Messages generated by the parser
	nodeTarget    *NodeList // Used by the parser to add nodes to a target NodeList
	text          string    // The input text
	lex           *lexer
	token         [9]*item
	sectionLevels *sectionLevels // Encountered section levels
	sections      []*SectionNode // Pointers to encountered sections
	id            int            // The consecutive id of the node in the tree
	indentWidth   int
	indentLevel   int
}

// startParse initializes the parser, using the lexer.
func (t *Tree) startParse(lex *lexer) {
	t.lex = lex
}

// Parse activates the parser using text as input data. A parse tree is
// returned on success or failure. Users of the Parse package should use the
// Top level Parse function.
func (t *Tree) Parse(text string, treeSet *Tree) (tree *Tree) {
	log.Debugln("Start")
	t.startParse(lex(t.Name, text))
	t.text = text
	t.parse(treeSet)
	log.Debugln("End")
	return t
}

// parse is where items are retrieved from the parser and dispatched according
// to the itemElement type.
func (t *Tree) parse(tree *Tree) {
	log.Debugln("Start")

	t.nodeTarget = t.Nodes

	for t.peek(1).Type != itemEOF {
		var n Node

		token := t.next()
		log.Infof("\nParser got token: %#+v\n\n", token)

		switch token.Type {
		case itemParagraph:
			n = newParagraph(token, &t.id)
		case itemTransition:
			n = newTransition(token, &t.id)
		case itemComment:
			n = t.comment(token)
		case itemSectionAdornment:
			n = t.section(token)
		case itemSpace:
			n = t.indent(token)
			if n == nil {
				continue
			}
		case itemTitle, itemBlankLine:
			// itemTitle is consumed when evaluating itemSectionAdornment
			continue
		}

		t.nodeTarget.append(n)
		switch n.NodeType() {
		case NodeSection, NodeBlockQuote:
			// Set the loop to append items to the NodeList of the new section
			// FIXME: Remove this reflection somehow
			t.nodeTarget = reflect.ValueOf(n).Elem().FieldByName("NodeList").Addr().Interface().(*NodeList)
		}

	}

	log.Debugln("End")
}

// backup shifts the token buffer right one position.
func (t *Tree) backup() {
	t.token[0] = nil
	for x := len(t.token) - 1; x > 0; x-- {
		t.token[x] = t.token[x-1]
		t.token[x-1] = nil
	}
}

// peekBack uses the token buffer to "look back" a number of positions (pos).
// Looking back more positions than the Tree.token buffer allows (3) will
// generate a panic.
func (t *Tree) peekBack(pos int) *item {
	return t.token[zed-pos]
}

func (t *Tree) peekBackTo(item itemElement) (tok *item) {
	for i := zed - 1; i >= 0; i-- {
		if t.token[i] != nil && t.token[i].Type == item {
			return t.token[i]
		}
	}
	return
}

// peek looks ahead in the token stream a number of positions (pos) and gets
// the next token from the lexer. A pointer to the token is kept in the
// Tree.token buffer. If a token pointer already exists in the buffer, that
// token is used instead and no tokens are received the the lexer stream
// (channel).
func (t *Tree) peek(pos int) *item {
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
	return nItem
}

// peekSkip looks ahead one position skipiing a specified itemElement. If that
// element is found, a pointer is returned, otherwise nil is returned.
func (t *Tree) peekSkip(iSkip itemElement) *item {
	var nItem *item
	count := 1
	for {
		nItem = t.peek(count)
		if nItem.Type != iSkip {
			break
		}
		count++
	}
	return nItem
}

// next is the workhorse of the parser. It is repsonsible for getting the next
// token from the lexer stream (channel). If the next token already exists in
// the token buffer, than the token buffer is shifted left and the pointer to
// the "zed" token is returned.
func (t *Tree) next() *item {
	for x := 0; x < len(t.token)-1; x++ {
		t.token[x] = t.token[x+1]
		t.token[x+1] = nil
	}
	if t.token[zed] == nil && t.lex != nil {
		t.token[zed] = t.lex.nextItem()
	}
	return t.token[zed]
}

// section is responsible for parsing the title, overline, and underline tokens
// returned from the parser. If there are errors parsing these elements, than a
// systemMessage is generated and added to Tree.Nodes.
func (t *Tree) section(i *item) Node {
	log.Debugln("Start")
	var overAdorn, indent, title, underAdorn *item

	if pFor := t.peekSkip(itemSpace); pFor != nil && pFor.Type == itemTitle {
		// Section with overline
		// Check for errors
		if t.token[zed].Length < 3 && t.token[zed].Length != pFor.Length {
			t.next()
			t.next()
			if bTok := t.peekBack(1); bTok != nil && bTok.Type == itemSpace {
				t.next()
				t.next()
				return t.systemMessage(infoUnexpectedTitleOverlineOrTransition)
			}
			return t.systemMessage(infoOverlineTooShortForTitle)
		} else if pBack := t.peekBack(1); pBack != nil && pBack.Type == itemSpace {
			// Indented section (error)
			// The section title has an indented overline
			return t.systemMessage(severeUnexpectedSectionTitleOrTransition)
		}

		overAdorn = i
		t.next()

	loop:
		for {
			switch tTok := t.token[zed]; tTok.Type {
			case itemTitle:
				title = tTok
				t.next()
			case itemSpace:
				indent = tTok
				t.next()
			case itemSectionAdornment:
				underAdorn = tTok
				break loop
			}
		}
	} else if pBack := t.peekBack(1); pBack != nil &&
		(pBack.Type == itemTitle || pBack.Type == itemSpace) {
		// Section with no overline
		// Check for errors
		if pBack.Type == itemSpace {
			pBack := t.peekBack(2)
			if pBack != nil && pBack.Type == itemTitle {
				// The section underline is indented
				return t.systemMessage(severeUnexpectedSectionTitle)
			}
		} else if t.token[zed].Length < 3 && t.token[zed].Length != pBack.Length {
			// Short underline
			return t.systemMessage(infoUnderlineTooShortForTitle)
		}

		// Section OKAY
		title = t.peekBack(1)
		underAdorn = i

	} else if pFor := t.peekSkip(itemSpace); pFor != nil && pFor.Type == itemParagraph {
		// If a section contains an itemParagraph, it is because the underline
		// is missing, therefore we generate an error based on what follows the
		// itemParagraph.
		t.next() // Move the token buffer past the error tokens
		t.next()
		if t.token[zed].Length < 3 && t.token[zed].Length != pFor.Length {
			t.backup()
			return t.systemMessage(infoOverlineTooShortForTitle)
		} else if p := t.peek(1); p != nil && p.Type == itemBlankLine {
			return t.systemMessage(severeMissingMatchingUnderlineForOverline)
		}
		return t.systemMessage(severeIncompleteSectionTitle)
	} else if pFor := t.peekSkip(itemSpace); pFor != nil && pFor.Type == itemSectionAdornment {
		// Missing section title
		t.next() // Move the token buffer past the error token
		return t.systemMessage(errorInvalidSectionOrTransitionMarker)
	} else if pFor := t.peekSkip(itemSpace); pFor != nil && pFor.Type == itemEOF {
		// Missing underline and at EOF
		return t.systemMessage(errorInvalidSectionOrTransitionMarker)
	}

	if overAdorn != nil && overAdorn.Text.(string) != underAdorn.Text.(string) {
		return t.systemMessage(severeOverlineUnderlineMismatch)
	}

	// Determine the level of the section and where to append it to in t.Nodes
	undoID := t.id
	sec := newSection(title, overAdorn, underAdorn, indent, &t.id)
	log.Debugf("Adding  %#U to sectionLevels\n", sec.UnderLine.Rune)
	msg := t.sectionLevels.Add(sec)
	if msg != parserMessageNil {
		log.Debugln("Found inconsistent section level!")
		t.id = undoID
		return t.systemMessage(severeTitleLevelInconsistent)
	}
	sec.Level = t.sectionLevels.lastSectionNode.Level
	if sec.Level == 1 {
		log.Debugln("Setting nodeTarget to Tree.Nodes!")
		t.nodeTarget = t.Nodes
	} else {
		lSec := t.sectionLevels.lastSectionNode
		if sec.Level > 1 {
			lSec = t.sectionLevels.LastSectionByLevel(sec.Level - 1)
		}
		t.nodeTarget = &lSec.NodeList
		log.Debugln("Setting nodeTarget to section ID", lSec.ID.String())
	}

	// The following checks have to be made after the SectionNode has been
	// initialized so that any parserMessages can be appended to the
	// SectionNode.NodeList.
	oLen := title.Length
	if indent != nil {
		oLen = indent.Length + title.Length
	}

	if overAdorn != nil && oLen > overAdorn.Length {
		sec.NodeList = append(sec.NodeList, t.systemMessage(warningShortOverline))
	} else if overAdorn == nil && title.Length != underAdorn.Length {
		sec.NodeList = append(sec.NodeList, t.systemMessage(warningShortUnderline))
	}

	log.Debugln("End")
	return sec
}

func (t *Tree) comment(i *item) Node {
	n := newComment(i, &t.id)
	nTok := t.peek(1)
	if nTok != nil && nTok.Type != itemSpace {
		// The comment element itself is valid, but we need to add it to the
		// NodeList before the systemMessage.
		t.nodeTarget.append(n)
		return t.systemMessage(warningExplicitMarkupWithUnIndent)
	}
	return n
}

// systemMessage generates a Node based on the passed parserMessage. The
// generated message is returned as a SystemMessageNode.
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
	s.NodeList = append(s.NodeList, msg)

	log.Debugln("FOUND", err)
	// if t.token[zed].Line == 9 {
	// spd.Dump(t.token)
	// os.Exit(1)
	// }
	var overLine, indent, title, underLine, newLine string

	switch err {
	case infoOverlineTooShortForTitle:
		var infoText string
		if t.token[zed-2] != nil {
			infoText = t.token[zed-2].Text.(string) + "\n" + t.token[zed-1].Text.(string) + "\n" + t.token[zed].Text.(string)
			s.Line = t.token[zed-2].Line
			t.token[zed-2] = nil
		} else {
			infoText = t.token[zed-1].Text.(string) + "\n" + t.token[zed].Text.(string)
			s.Line = t.token[zed-1].Line
		}
		infoTextLen := len(infoText)
		// Modify the token buffer to change the current token to a
		// itemParagraph then backup the token buffer so the next loop gets the
		// new paragraph
		t.token[zed-1] = nil
		t.token[zed].Type = itemParagraph
		t.token[zed].Text = infoText
		t.token[zed].Length = infoTextLen
		t.token[zed].Line = s.Line
		t.backup()
	case infoUnexpectedTitleOverlineOrTransition:
		oLin := t.peekBackTo(itemSectionAdornment)
		titl := t.peekBackTo(itemTitle)
		uLin := t.token[zed]
		infoText := oLin.Text.(string) + "\n" + titl.Text.(string) + "\n" + uLin.Text.(string)
		s.Line = oLin.Line
		// FIXME: DRY
		t.token[zed-4] = nil
		t.token[zed-3] = nil
		t.token[zed-2] = nil
		t.token[zed-1] = nil
		infoTextLen := len(infoText)
		// Modify the token buffer to change the current token to a
		// itemParagraph then backup the token buffer so the next loop gets the
		// new paragraph
		t.token[zed].Type = itemParagraph
		t.token[zed].Text = infoText
		t.token[zed].Length = infoTextLen
		t.token[zed].Line = s.Line
		t.token[zed].StartPosition = oLin.StartPosition
		t.backup()
	case infoUnderlineTooShortForTitle:
		infoText := t.token[zed-1].Text.(string) + "\n" + t.token[zed].Text.(string)
		infoTextLen := len(infoText)
		s.Line = t.token[zed-1].Line
		// Modify the token buffer to change the current token to a
		// itemParagraph then backup the token buffer so the next loop gets the
		// new paragraph
		t.token[zed-1] = nil
		t.token[zed].Type = itemParagraph
		t.token[zed].Text = infoText
		t.token[zed].Length = infoTextLen
		t.token[zed].Line = s.Line
		t.backup()
	case warningShortOverline, severeOverlineUnderlineMismatch:
		backToken = zed - 2
		if t.peekBack(2).Type == itemSpace {
			backToken = zed - 3
			indent = t.token[zed-2].Text.(string)
		}
		overLine = t.token[backToken].Text.(string)
		title = t.token[zed-1].Text.(string)
		underLine = t.token[zed].Text.(string)
		newLine = "\n"
		lbText = overLine + newLine + indent + title + newLine + underLine
		s.Line = t.token[backToken].Line
		lbTextLen = len(lbText)
	case warningShortUnderline, severeUnexpectedSectionTitle:
		backToken = zed - 1
		if t.peekBack(1).Type == itemSpace {
			backToken = zed - 2
		}
		lbText = t.token[backToken].Text.(string) + "\n" + t.token[zed].Text.(string)
		lbTextLen = len(lbText)
		s.Line = t.token[zed-1].Line
	case warningExplicitMarkupWithUnIndent:
		s.Line = t.token[zed+1].Line
	case errorInvalidSectionOrTransitionMarker:
		lbText = t.token[zed-1].Text.(string) + "\n" + t.token[zed].Text.(string)
		s.Line = t.token[zed-1].Line
		lbTextLen = len(lbText)
	case severeIncompleteSectionTitle, severeMissingMatchingUnderlineForOverline:
		lbText = t.token[zed-2].Text.(string) + "\n" +
			t.token[zed-1].Text.(string) + t.token[zed].Text.(string)
		s.Line = t.token[zed-2].Line
		lbTextLen = len(lbText)
	case severeUnexpectedSectionTitleOrTransition:
		lbText = t.token[zed].Text.(string)
		lbTextLen = len(lbText)
		s.Line = t.token[zed].Line
	case severeTitleLevelInconsistent:
		if t.peekBack(2).Type == itemSectionAdornment {
			lbText = t.token[zed-2].Text.(string) + "\n" +
				t.token[zed-1].Text.(string) + "\n" + t.token[zed].Text.(string)
			lbTextLen = len(lbText)
			s.Line = t.token[zed-2].Line
		} else {
			lbText = t.token[zed-1].Text.(string) + "\n" + t.token[zed].Text.(string)
			lbTextLen = len(lbText)
			s.Line = t.token[zed-1].Line
		}
	}

	if lbTextLen > 0 {
		lb := newLiteralBlock(&item{
			Type:   itemLiteralBlock,
			Text:   lbText,
			Length: lbTextLen, // Add one to account for the backslash
		}, &t.id)
		s.NodeList = append(s.NodeList, lb)
	}

	t.Messages.append(s)

	return s
}

// indent parses IndentNode's returned from the lexer and returns a
// BlockQuoteNode.
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
