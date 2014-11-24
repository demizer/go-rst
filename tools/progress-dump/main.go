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
  progress-dump  -p <PATH> | --progress-yml <PATH> [-h | --help]

Options:
  -h --help                        Show the help message.
  -p <PATH> --progress-yml <PATH>  The path to progress.yml
  -r <PATH> --readme <PATH>        The path to README.rst
`

func readProgressFile(path string) interface{} {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		log.Criticalln(err)
		os.Exit(1)
	}
	obj := new(interface{})
	err = yaml.Unmarshal(f, obj)
	if err != nil {
		log.Criticalln(err)
		os.Exit(1)
	}
	return obj
}

var spd = spew.ConfigState{Indent: "\t", DisableMethods: true}

func main() {
	log.SetLevel(log.LEVEL_DEBUG)

	args, err := docopt.Parse(APP_USAGE, nil, true, "progress-dump", false)
	if err != nil {
		log.Criticalln(err)
		os.Exit(1)
	}

	pYml := readProgressFile(args["--progress-yml"].(string))
	spd.Dump(pYml)
}
