package main

import (
	"io/ioutil"
	"os"

	"github.com/aybabtme/rgbterm"
	"github.com/davecgh/go-spew/spew"
	"github.com/demizer/go-elog"
	"github.com/docopt/docopt.go"
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
	MaxCol1Chars int
	MaxCol2Chars int
	MaxCol3Chars int
}

type TableSection struct {
	Header *TableSectionHeader
	Rows   []*TableRow
	Id     int
}

type TableSectionHeader struct {
	Name     string
	DonePerc int
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
		Table: &Table{Sections: make([]*TableSection, 0)},
		Id:    0,
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
	spd.Dump(s.Table)
}

func (s *State) Walk(z []Item, sec *TableSection, depth int) {
	for _, x := range z {
		if depth == 0 {
			s.Id++
			sec = &TableSection{
				Header: &TableSectionHeader{Name: x.Item},
				Id:     s.Id,
			}
		}
		if x.SubItems != nil {
			name := x.Item
			if depth > 0 {
				s.Id++
				name = sec.Header.Name + ":" + x.Item
			}
			subSec := &TableSection{
				Header: &TableSectionHeader{Name: name},
				Id:     s.Id,
			}
			depth++
			s.Walk(x.SubItems, subSec, depth)
			depth--
			s.Table.Sections = append(s.Table.Sections, subSec)
			continue
		}
		if len(x.Done) > s.Table.MaxCol1Chars {
			s.Table.MaxCol1Chars = len(x.Done)
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

func main() {
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
