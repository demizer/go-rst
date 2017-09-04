package parser

type shortSectionNode struct {
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

// var testSectionLevelsAdd = []struct {
// name  string
// pSecs []*testSectionLevelSectionNode
// eLvls []*testSectionLevelExpectLevels
// }{
// {
// name: "Test two levels with a single SectionNode each",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// },
// eLvls: []*testSectionLevelExpectLevels{
// {rChar: '=', level: 1, nodes: []shortSectionNode{
// {level: 1, uRune: '='},
// }},
// {rChar: '-', level: 2, nodes: []shortSectionNode{
// {level: 2, uRune: '-'},
// }},
// },
// },
// {
// name: "Test two levels with on level one return",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// },
// eLvls: []*testSectionLevelExpectLevels{
// {rChar: '=', level: 1, nodes: []shortSectionNode{
// {level: 1, uRune: '='},
// {level: 1, uRune: '='},
// }},
// {rChar: '-', level: 2, nodes: []shortSectionNode{
// {level: 2, uRune: '-'},
// {level: 2, uRune: '-'},
// }},
// },
// },
// {
// name: "Test three levels with one return to level 1",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// {node: shortSectionNode{level: 3, uRune: '~'}},
// {node: shortSectionNode{level: 1, uRune: '='}},
// },
// eLvls: []*testSectionLevelExpectLevels{
// {rChar: '=', level: 1, nodes: []shortSectionNode{
// {level: 1, uRune: '='},
// {level: 1, uRune: '='},
// }},
// {rChar: '-', level: 2, nodes: []shortSectionNode{
// {level: 2, uRune: '-'},
// }},
// {rChar: '~', level: 3, nodes: []shortSectionNode{
// {level: 3, uRune: '~'},
// }},
// },
// },
// {
// name: "Test three levels with two returns to level 1",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// {node: shortSectionNode{level: 3, uRune: '~'}},
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// },
// eLvls: []*testSectionLevelExpectLevels{
// {rChar: '=', level: 1, nodes: []shortSectionNode{
// {level: 1, uRune: '='},
// {level: 1, uRune: '='},
// {level: 1, uRune: '='},
// }},
// {rChar: '-', level: 2, nodes: []shortSectionNode{
// {level: 2, uRune: '-'},
// {level: 2, uRune: '-'},
// }},
// {rChar: '~', level: 3, nodes: []shortSectionNode{
// {level: 3, uRune: '~'},
// }},
// },
// },
// {
// name: "Test inconsistent section level",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// {node: shortSectionNode{level: 1, uRune: '='}},
// {eMessage: severeTitleLevelInconsistent,
// node: shortSectionNode{level: 2, uRune: '`'}},
// },
// },
// {
// name: "Test inconsistent section level 2",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// {node: shortSectionNode{level: 3, uRune: '~'}},
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, uRune: '-'}},
// {eMessage: severeTitleLevelInconsistent,
// node: shortSectionNode{level: 3, uRune: '`'}},
// },
// },
// {
// name: "Test level two with overline and all runes similar",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, oRune: '=', uRune: '='}},
// },
// eLvls: []*testSectionLevelExpectLevels{
// {rChar: '=', level: 1, nodes: []shortSectionNode{
// {level: 1, uRune: '='},
// }},
// {rChar: '=', level: 2, overLine: true,
// nodes: []shortSectionNode{
// {level: 2, uRune: '='},
// },
// },
// },
// },
// {
// name: "Test level two with overline with same rune as level one.",
// pSecs: []*testSectionLevelSectionNode{
// {node: shortSectionNode{level: 1, uRune: '='}},
// {node: shortSectionNode{level: 2, oRune: '=', uRune: '='}},
// },
// eLvls: []*testSectionLevelExpectLevels{
// {rChar: '=', level: 1, nodes: []shortSectionNode{
// {level: 1, uRune: '='},
// }},
// {rChar: '=', level: 2, overLine: true,
// nodes: []shortSectionNode{
// {level: 2, uRune: '='},
// },
// },
// },
// },
// }

// func testSectionLevelsAddCheckEqual(t *testing.T, testName string,
// pos int, pLvl, eLvl *sectionLevel) {

// if eLvl.level != pLvl.level {
// t.Errorf("Test: %q\n\tGot: sectionLevel.Level = %d, "+"Expect: %d", testName, pLvl.level, eLvl.level)
// }
// if eLvl.rChar != pLvl.rChar {
// t.Errorf("Test: %q\n\tGot: sectionLevel.rChar = %#U, "+"Expect: %#U", testName, pLvl.rChar, eLvl.rChar)
// }
// if eLvl.overLine != pLvl.overLine {
// t.Errorf("Test: %q\n\tGot: sectionLevel.overLine = %t, "+"Expect: %t", testName, pLvl.overLine,
// eLvl.overLine)
// }
// for eNum, eSec := range eLvl.sections {
// if eSec.Level != pLvl.sections[eNum].Level {
// t.Errorf("Test: %q\n\tGot: level[%d].sections[%d].Level = %d, "+"Expect: %d", testName, pos,
// eNum, pLvl.sections[eNum].Level, eSec.Level)
// }
// eRune := eSec.UnderLine.Rune
// pRune := pLvl.sections[eNum].UnderLine.Rune
// if eRune != pRune {
// t.Errorf("Test: %q\n\tGot: level[%d].section[%d].Rune = %#U, "+"Expect: %#U", testName, pos,
// eNum, pLvl.sections[eNum].UnderLine.Rune, eSec.UnderLine.Rune)
// }
// }
// }

// func TestSectionLevelsAdd(t *testing.T) {
// var pSecLvls, eSecLvls *sectionLevels
// var testName string

// addSection := func(s *testSectionLevelSectionNode) {
// n := &doc.SectionNode{Level: s.node.level,
// UnderLine: &doc.AdornmentNode{Rune: s.node.uRune}}
// if s.node.oRune != 0 {
// n.OverLine = &doc.AdornmentNode{Rune: s.node.oRune}
// }
// msg := pSecLvls.Add(n)
// if msg > parserMessageNil && msg != s.eMessage {
// t.Fatalf("Test: %q\n\tGot: parserMessage = %q, "+"Expect: %q", testName, msg, s.eMessage)
// }
// }

// for _, tt := range testSectionLevelsAdd {
// testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
// pSecLvls = newSectionLevels(testutil.StdLogger)
// eSecLvls = newSectionLevels(testutil.StdLogger)
// testName = tt.name

// // pSecLvls := newSectionLevels(testutil.StdLogger)
// for _, secNode := range tt.pSecs {
// addSection(secNode)
// }

// // Initialize the expected sectionLevels
// for _, slvl := range tt.eLvls {
// s := &sectionLevel{rChar: slvl.rChar,
// level: slvl.level, overLine: slvl.overLine,
// }
// for _, sn := range slvl.nodes {
// n := &doc.SectionNode{Level: sn.level}
// n.UnderLine = &doc.AdornmentNode{Rune: sn.uRune}
// if sn.oRune != 0 {
// n.OverLine = &doc.AdornmentNode{
// Rune: sn.oRune,
// }
// }
// s.sections = append(s.sections, n)
// }
// eSecLvls.levels = append(eSecLvls.levels, s)
// }

// for i := 0; i < len(eSecLvls.levels); i++ {
// testSectionLevelsAddCheckEqual(t, testName, i,
// pSecLvls.levels[i], eSecLvls.levels[i])
// }
// }
// }

// var testSectionLevelsLast = []struct {
// name      string
// tLevel    int // The last level to get
// tSections []*doc.SectionNode
// eLevel    sectionLevel // There can be only one
// }{
// {
// name:   "Test last section level two",
// tLevel: 2,
// tSections: []*doc.SectionNode{
// {Level: 1, Title: &doc.TitleNode{
// NodeList: doc.NodeList{
// doc.TextNode{Text: "Title 2"},
// },
// }, UnderLine: &doc.AdornmentNode{Rune: '='}},
// // {Level: 2, Title: &doc.TitleNode{NodeList: doc.TextNode{Text: "Title 2"}}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 2, Title: &doc.TitleNode{NodeList: doc.TextNode{Text: "Title 3"}}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 2, Title: &doc.TitleNode{NodeList: doc.TextNode{Text: "Title 4"}}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// },
// // eLevel: sectionLevel{
// // rChar: '-', level: 2,
// // sections: []*doc.SectionNode{
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '~'}},
// // },
// // },
// },
// // {
// // name:   "Test last section level one",
// // tLevel: 1,
// // tSections: []*doc.SectionNode{
// // {Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 2"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 3"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // },
// // eLevel: sectionLevel{
// // rChar: '=', level: 1,
// // sections: []*doc.SectionNode{
// // {Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
// // },
// // },
// // },
// // {
// // name:   "Test last section level three",
// // tLevel: 3,
// // tSections: []*doc.SectionNode{
// // {Level: 1, Title: &doc.TitleNode{Text: "Title 1"}, UnderLine: &doc.AdornmentNode{Rune: '='}},
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 2"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 3"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 2, Title: &doc.TitleNode{Text: "Title 4"}, UnderLine: &doc.AdornmentNode{Rune: '-'}},
// // {Level: 3, Title: &doc.TitleNode{Text: "Title 5"}, UnderLine: &doc.AdornmentNode{Rune: '+'}},
// // },
// // eLevel: sectionLevel{
// // rChar: '+', level: 3,
// // sections: []*doc.SectionNode{
// // {Level: 3, Title: &doc.TitleNode{Text: "Title 5"}, UnderLine: &doc.AdornmentNode{Rune: '+'}},
// // },
// // },
// // },
// }

// func TestSectionLevelsLast(t *testing.T) {
// for _, tt := range testSectionLevelsLast {
// testutil.Log(fmt.Sprintf("\n\n\n\n RUNNING TEST %q \n\n\n\n", tt.name))
// secLvls := newSectionLevels(testutil.StdLogger)
// for _, secNode := range tt.tSections {
// secLvls.Add(secNode)
// }
// var pSec *doc.SectionNode
// pSec = secLvls.LastSectionByLevel(tt.tLevel)
// if tt.eLevel.level != pSec.Level {
// t.Errorf("Test: %q\n\tGot: sectionLevel.Level = %d, Expect: %d", tt.name, tt.eLevel.level,
// pSec.Level)
// }
// if tt.eLevel.rChar != pSec.UnderLine.Rune {
// t.Errorf("Test: %q\n\tGot: sectionLevel.rChar = %#U, Expect: %#U", tt.name, tt.eLevel.rChar,
// pSec.UnderLine.Rune)
// }
// // There can be only one
// // nl := tt.eLevel.sections[0].Title.NodeList[0]
// // if (*doc.TitleNode).nl.Text != pSec.Title.Text {
// // t.Errorf("Test: %q\n\tGot: level[0].sections[0].Title.Text = %q, "+"Expect: %q", tt.name,
// // pSec.Title.Text, tt.eLevel.sections[0].Title.Text)
// // }
// }
// }
