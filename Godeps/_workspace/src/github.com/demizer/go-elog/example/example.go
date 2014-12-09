package main

import (
	"bytes"
	"fmt"
	"os"

	"github.com/demizer/go-elog"
)

func main() {
	var buf bytes.Buffer

	logr := log.New(log.LEVEL_DEBUG, os.Stdout)
	logr.SetFlags(log.LdebugFlags | log.Ldate | log.Lprefix)

	logr.Println("\nstdout output:\n")

	logr.SetStreams(os.Stdout, &buf)

	logr.Debugln("Level 0 Output 1")
	lvl3 := func() {
		logr.Debugln("Level 3 Output 1")
	}
	lvl2 := func() {
		logr.Debugln("Level 2 Output 1")
		logr.Debugln("Level 2 Output 2")
		lvl3()
		logr.Debugln("Level 2 Output 3")
	}
	lvl1 := func() {
		logr.Debugln("Level 1 Output 1")
		logr.Debugln("Level 1 Output 2")
		lvl2()
		logr.Debugln("Level 1 Output 3")
	}
	lvl1()

	logr.Debugln("Level 0 Output 2")

	logr.SetStreams(os.Stdout)

	logr.Println("\nShowing output stored in the buffer:\n")

	fmt.Print(buf.String())
}
