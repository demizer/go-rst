// progress-dump -- Dumps progress.yml to a reStructuredText Grid Table
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

// progress-dump is used to output a grid table using progress.yml for inclusion
// into README.rst It gives prospective users an idea of how specification
// complete go-rst is compared to the reference docutils parser.
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/aybabtme/rgbterm"
	"github.com/davecgh/go-spew/spew"
	"github.com/demizer/go-elog"
	"github.com/docopt/docopt-go"
	"gopkg.in/yaml.v2"
)

var APP_NAME = rgbterm.String("progress-dump", 255, 255, 135)
var APP_DESC = rgbterm.String("Dumps progress output", 0, 215, 95)
var APP_USAGE = APP_NAME + " - " + APP_DESC + `

Usage:
  progress-dump [--progress-yml <PATH>] [-h | --help]

Options:
  -h --help                        Show the help message.
  --progress-yml <PATH>  The path to progress.yml [default: ../../progress.yml]
  --readme <PATH>        The path to README.rst
`

type Item struct {
	Item      string
	Done      string
	Note      string
	Completed string
	SubItems  []Item `yaml:"sub-items"`
}

var spd = spew.ConfigState{Indent: "\t", DisableMethods: true}

type Table struct {
	Sections     []*TableSection
	TotalItems   int
	TotalDone    int
	OverAllPerc  float64
	MaxCol1Chars int
	MaxCol2Chars int
	MaxCol3Chars int
}

func (t *Table) Len() int {
	return len(t.Sections)
}

func (t *Table) Swap(i, j int) {
	t.Sections[i], t.Sections[j] = t.Sections[j], t.Sections[i]
}

func (t *Table) Less(i, j int) bool {
	return t.Sections[i].Id < t.Sections[j].Id
}

func (t *Table) Dump() {
	// The table sections are in reverse order due to the recursion.
	sort.Sort(t)

	// The addition numbers here compensate for the ascii chars that make
	// up the frame of the ascii table.
	// 1 == space
	// 3 == 2x space and "+"
	tWidth := 1 + t.MaxCol1Chars + 3 + t.MaxCol2Chars + 3 + t.MaxCol3Chars + 1
	tTop := strings.Repeat("-", tWidth)
	topWithEndPoints := fmt.Sprintf("+" + tTop + "+")

	fakeHdr := fmt.Sprintf("| %s | %s | %s |", "**Done**",
		// 8 == (4x asterisks, 2 spaces, 2 frame ascii)
		"**Item**"+strings.Repeat(" ", t.MaxCol2Chars-8),
		"**Note**"+strings.Repeat(" ", t.MaxCol3Chars-8))

	sepWithPoints := fmt.Sprintf("+-%s-+-%s-+-%s-+", "--------",
		strings.Repeat("-", t.MaxCol2Chars),
		strings.Repeat("-", t.MaxCol3Chars))

	fmt.Println(topWithEndPoints)

	totalDoneHdr := fmt.Sprintf("**The go-rst Library Implements "+
		"%0.0f%% of the Official Specification (%d of %d Items)**",
		t.OverAllPerc, t.TotalDone, t.TotalItems)

	fmt.Printf("| %s |\n", totalDoneHdr+strings.Repeat(" ",
		tWidth-len(totalDoneHdr)-2))

	fmt.Println(sepWithPoints)

	for x, y := range t.Sections {
		secTitle := fmt.Sprintf("**%0.0f%% Complete -- %s**",
			y.Header.DonePerc, y.Header.Name)
		fmt.Printf("| %s |\n",
			secTitle+strings.Repeat(" ", tWidth-len(secTitle)-2))
		if x == 0 {
			fmt.Println(sepWithPoints)
			fmt.Println(fakeHdr)
			fmt.Println(sepWithPoints)
		} else {
			fmt.Println(sepWithPoints)
		}
		for _, z := range y.Rows {
			fmt.Printf("| %s | %s | %s |\n",
				z.Done+strings.Repeat(" ",
					t.MaxCol1Chars-len(z.Done)),
				z.Item+strings.Repeat(" ",
					t.MaxCol2Chars-len(z.Item)),
				z.Note+strings.Repeat(" ",
					t.MaxCol3Chars-len(z.Note)))
			fmt.Println(sepWithPoints)
		}
	}
}

type TableSection struct {
	Header *TableSectionHeader
	Rows   []*TableRow
	Id     int
}

type TableSectionHeader struct {
	Name     string
	Done     string
	DonePerc float64
}

type TableRow struct {
	Done string
	Item string
	Note string
}

type State struct {
	Table *Table
	Id    int
	Items []Item
}

func NewState() *State {
	return &State{
		Table: &Table{
			Sections:     make([]*TableSection, 0),
			MaxCol1Chars: 8, // Col1 will always be **Done**
		},
		Id: 0,
	}
}

func (s *State) ReadProgressFile(path string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Criticalln(err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(f, &s.Items)
	if err != nil {
		log.Criticalln(err)
		os.Exit(1)
	}
}

func (s *State) DumpTable() {
	s.Walk(s.Items, nil, 0)
	s.CalcPercentages()
	s.Table.Dump()
}

func (s *State) Walk(z []Item, sec *TableSection, depth int) {
	for _, x := range z {
		if depth == 0 {
			s.Id++
			sec = &TableSection{
				Header: &TableSectionHeader{Name: x.Item,
					Done: x.Done},
				Id: s.Id,
			}
		}
		if x.SubItems != nil {
			name := x.Item
			if depth > 0 {
				s.Id++
				name = sec.Header.Name + " :: " + x.Item
			}
			subSec := &TableSection{
				Header: &TableSectionHeader{Name: name,
					Done: x.Done},
				Id: s.Id,
			}
			depth++
			s.Walk(x.SubItems, subSec, depth)
			depth--
			s.Table.Sections = append(s.Table.Sections, subSec)
			continue
		}
		if len(x.Item) > s.Table.MaxCol2Chars {
			s.Table.MaxCol2Chars = len(x.Item)
		}
		if len(x.Note) > s.Table.MaxCol3Chars {
			s.Table.MaxCol3Chars = len(x.Note)
		}
		nRow := &TableRow{x.Done, x.Item, x.Note}
		sec.Rows = append(sec.Rows, nRow)
		if depth == 0 {
			s.Table.Sections = append(s.Table.Sections, sec)
		}
	}
}

func (s *State) CalcPercentages() {
	for _, x := range s.Table.Sections {
		sDone := 0.0
		s.Table.TotalItems++
		if x.Header.Done == "yes" {
			s.Table.TotalDone++
			sDone++
		}
		for _, y := range x.Rows {
			s.Table.TotalItems++
			if y.Done == "yes" {
				s.Table.TotalDone++
				sDone++
			}
		}
		if sDone != 0 {
			x.Header.DonePerc = (sDone / float64(len(x.Rows)+1)) * 100
		}
	}
	s.Table.OverAllPerc = (float64(s.Table.TotalDone) /
		float64(s.Table.TotalItems)) * 100
}

func main() {
	log.SetFlags(0)
	log.SetLevel(log.LEVEL_DEBUG)

	args, err := docopt.Parse(APP_USAGE, nil, true, "progress-dump", false)
	if err != nil {
		log.Criticalln(err)
		os.Exit(1)
	}

	s := NewState()
	s.ReadProgressFile(args["--progress-yml"].(string))
	s.DumpTable()
}
