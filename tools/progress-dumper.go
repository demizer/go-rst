// progress-dump -- Dumps progress.yml to a reStructuredText Grid Table
//
// progress-dump is used to output a grid table using progress.yml for inclusion
// into README.rst It gives prospective users an idea of how specification
// complete go-rst is compared to the reference docutils parser.
//
// This script is meant to be called from the project root, like so:
//
// go run progress-dump
//
package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"

	"github.com/aybabtme/rgbterm"
	"github.com/davecgh/go-spew/spew"
	"github.com/docopt/docopt-go"
	"gopkg.in/yaml.v2"
)

var appName = rgbterm.String("progress-dump", 255, 255, 135, 0, 0, 0)
var appDesc = rgbterm.String("Dumps progress output", 0, 215, 95, 0, 0, 0)
var appUsage = appName + " - " + appDesc + `

Usage:
  progress-dump [--progress-yml <PATH>] [-h | --help]

Options:
  -h --help                        Show the help message.
  --progress-yml <PATH>  The path to progress.yml [default: tools/progress.yml]
  --readme <PATH>        The path to README.rst
`

type item struct {
	Item      string
	Done      string
	Note      string
	Completed string
	SubItems  []item `yaml:"sub-items"`
}

var spd = spew.ConfigState{Indent: "\t", DisableMethods: true}

type table struct {
	Sections     []*tableSection
	TotalItems   int
	TotalDone    int
	OverAllPerc  float64
	MaxCol1Chars int
	MaxCol2Chars int
	MaxCol3Chars int
}

func (t *table) Len() int {
	return len(t.Sections)
}

func (t *table) Swap(i, j int) {
	t.Sections[i], t.Sections[j] = t.Sections[j], t.Sections[i]
}

func (t *table) Less(i, j int) bool {
	return t.Sections[i].ID < t.Sections[j].ID
}

func (t *table) Dump() {
	// The table sections are in reverse order due to the recursion.
	sort.Sort(t)

	// The addition numbers here compensate for the ascii chars that make up the frame of the ascii table.
	// 1 == space
	// 3 == 2x space and "+"
	tWidth := 1 + t.MaxCol1Chars + 3 + t.MaxCol2Chars + 3 + t.MaxCol3Chars + 1
	tTop := strings.Repeat("-", tWidth)
	topWithEndPoints := fmt.Sprintf("+" + tTop + "+")

	fakeHdr := fmt.Sprintf("| %s | %s | %s |", "**Done**",
		// 8 == (4x asterisks, 2 spaces, 2 frame ascii)
		"**Item**"+strings.Repeat(" ", t.MaxCol2Chars-8),
		"**Note**"+strings.Repeat(" ", t.MaxCol3Chars-8))

	fmt.Printf("\nREADME_STATUS:go-rst implements **%0.0f%%** of the official specification (%d of %d Items)\n\n",
		t.OverAllPerc, t.TotalDone, t.TotalItems)

	sepWithPoints := fmt.Sprintf("+-%s-+-%s-+-%s-+", "--------", strings.Repeat("-", t.MaxCol2Chars),
		strings.Repeat("-", t.MaxCol3Chars))

	fmt.Println(topWithEndPoints)

	totalDoneHdr := fmt.Sprintf("**The go-rst Library Implements %0.0f%% of the Official Specification (%d of %d Items)**",
		t.OverAllPerc, t.TotalDone, t.TotalItems)

	fmt.Printf("| %s |\n", totalDoneHdr+strings.Repeat(" ", tWidth-len(totalDoneHdr)-2))

	fmt.Println(sepWithPoints)

	for x, y := range t.Sections {
		secTitle := fmt.Sprintf("**%0.0f%% Complete -- %s**", y.Header.DonePerc, y.Header.Name)
		fmt.Printf("| %s |\n", secTitle+strings.Repeat(" ", tWidth-len(secTitle)-2))
		if x == 0 {
			fmt.Println(sepWithPoints)
			fmt.Println(fakeHdr)
			fmt.Println(sepWithPoints)
		} else {
			fmt.Println(sepWithPoints)
		}
		for _, z := range y.Rows {
			fmt.Printf("| %s | %s | %s |\n",
				z.Done+strings.Repeat(" ", t.MaxCol1Chars-len(z.Done)),
				z.Item+strings.Repeat(" ", t.MaxCol2Chars-len(z.Item)),
				z.Note+strings.Repeat(" ", t.MaxCol3Chars-len(z.Note)))
			fmt.Println(sepWithPoints)
		}
	}
}

type tableSection struct {
	Header *tableSectionHeader
	Rows   []*tableRow
	ID     int
}

type tableSectionHeader struct {
	Name     string
	Done     string
	DonePerc float64
}

type tableRow struct {
	Done string
	Item string
	Note string
}

type state struct {
	Table *table
	ID    int
	Items []item
}

func newState() *state {
	return &state{
		Table: &table{
			Sections:     make([]*tableSection, 0),
			MaxCol1Chars: 8, // Col1 will always be **Done**
		},
		ID: 0,
	}
}

func (s *state) ReadProgressFile(path string) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
	err = yaml.Unmarshal(f, &s.Items)
	if err != nil {
		fmt.Println("ERROR:", err)
		os.Exit(1)
	}
}

func (s *state) DumpTable() {
	s.Walk(s.Items, nil, 0)
	s.CalcPercentages()
	s.Table.Dump()
}

func (s *state) Walk(z []item, sec *tableSection, depth int) {
	for _, x := range z {
		if depth == 0 {
			s.ID++
			sec = &tableSection{
				Header: &tableSectionHeader{Name: x.Item, Done: x.Done},
				ID:     s.ID,
			}
		}
		if x.SubItems != nil {
			name := x.Item
			if depth > 0 {
				s.ID++
				name = sec.Header.Name + " :: " + x.Item
			}
			subSec := &tableSection{
				Header: &tableSectionHeader{Name: name, Done: x.Done},
				ID:     s.ID,
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
		nRow := &tableRow{x.Done, x.Item, x.Note}
		sec.Rows = append(sec.Rows, nRow)
		if depth == 0 {
			s.Table.Sections = append(s.Table.Sections, sec)
		}
	}
}

func (s *state) CalcPercentages() {
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
	s.Table.OverAllPerc = (float64(s.Table.TotalDone) / float64(s.Table.TotalItems)) * 100
}

func main() {
	args, err := docopt.Parse(appUsage, nil, true, "progress-dump", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	s := newState()
	s.ReadProgressFile(args["--progress-yml"].(string))
	s.DumpTable()
}
