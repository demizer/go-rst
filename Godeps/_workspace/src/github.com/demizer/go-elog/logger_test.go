// Copyright 2013,2014 The go-logger Authors. All rights reserved.
// This code is MIT licensed. See the LICENSE file for more info.

package log

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"math/rand"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
	"time"

	"github.com/aybabtme/rgbterm"
)

func TestStream(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_CRITICAL, os.Stdout, &buf)
	if out := logr.Streams()[1]; out != &buf {
		t.Errorf("Stream = %p, want %p", out, &buf)
	}
}

func TestMultiStreams(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	fPath := filepath.Join(os.TempDir(), fmt.Sprint("go_test_",
		rand.Int()))
	file, err := os.Create(fPath)
	if err != nil {
		t.Error("Create(%q) = %v; want: nil", fPath, err)
	}
	defer file.Close()
	var buf bytes.Buffer
	eLen := 22
	logr := New(LEVEL_DEBUG, file, &buf)
	logr.SetFlags(0)
	logr.Debugln("Testing debug output!")
	b := make([]byte, eLen)
	n, err := file.ReadAt(b, 0)
	if n != eLen || err != nil {
		t.Errorf("Read(%d) = %d, %v; want: %d, nil", eLen, n, err,
			eLen)
	}
	if buf.Len() != eLen {
		t.Errorf("buf.Len() = %d; want: %d", buf.Len(), eLen)
	}
}

func TestLongFileFlag(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(LlongFileName | Llabel)
	logr.Debugln("Test long file flag")
	_, file, _, _ := runtime.Caller(0)
	expect := fmt.Sprintf("[DEBUG] %s: Test long file flag\n", file)
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestShortFileFlag(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(LshortFileName | Llabel)

	logr.Debugln("Test short file flag")
	_, file, _, _ := runtime.Caller(0)
	short := file

	for i := len(file) - 1; i > 0; i-- {
		if file[i] == '/' {
			short = file[i+1:]
			break
		}
	}

	file = short
	expect := fmt.Sprintf("[DEBUG] %s: Test short file flag\n", file)
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

var date = "Mon 20060102 15:04:05"

var fprintOutputTests = []struct {
	template   string
	prefix     string
	level      level
	dateFormat string
	flags      int
	text       string
	expect     string
	expectErr  bool
}{
	// Test with color prefix
	{
		template:   logFmt,
		prefix:     rgbterm.String("TEST>", 0, 255, 0),
		level:      LEVEL_PRINT,
		dateFormat: date,
		flags:      LstdFlags,
		text:       "test number 1",
		// The %s format specifier is the placeholder for the date.
		expect:    "%s \x1b[38;5;46mTEST>\x1b[0;00m test number 1",
		expectErr: false,
	},
	// Test output with coloring turned off
	{
		template:   logFmt,
		prefix:     "TEST>",
		level:      LEVEL_PRINT,
		dateFormat: date,
		flags:      Ldate | Lprefix,
		text:       "test number 2",
		expect:     "%s TEST> test number 2",
		expectErr:  false,
	},
	// Test debug output
	{
		template:   logFmt,
		prefix:     rgbterm.String("TEST>", 0, 255, 0),
		level:      LEVEL_DEBUG,
		dateFormat: time.RubyDate,
		flags:      LstdFlags,
		text:       "test number 3",
		expect:     "%s \x1b[38;5;46mTEST>\x1b[0;00m \x1b[38;5;231m[DEBUG]\x1b[0;00m test number 3",
		expectErr:  false,
	},
	// Test info output
	{
		template:   logFmt,
		prefix:     rgbterm.String("TEST>", 0, 255, 0),
		level:      LEVEL_INFO,
		dateFormat: time.RubyDate,
		flags:      LstdFlags,
		text:       "test number 4",
		expect:     "%s \x1b[38;5;46mTEST>\x1b[0;00m \x1b[38;5;41m[INFO]\x1b[0;00m test number 4",
		expectErr:  false,
	},
	// Test warning output
	{
		template:   logFmt,
		prefix:     rgbterm.String("TEST>", 0, 255, 0),
		level:      LEVEL_WARNING,
		dateFormat: time.RubyDate,
		flags:      LstdFlags,
		text:       "test number 5",
		expect:     "%s \x1b[38;5;46mTEST>\x1b[0;00m \x1b[38;5;228m[WARNING]\x1b[0;00m test number 5",
		expectErr:  false,
	},
	// Test error output
	{
		template:   logFmt,
		prefix:     rgbterm.String("TEST>", 0, 255, 0),
		level:      LEVEL_ERROR,
		dateFormat: time.RubyDate,
		flags:      LstdFlags,
		text:       "test number 6",
		expect:     "%s \x1b[38;5;46mTEST>\x1b[0;00m \x1b[38;5;202m[ERROR]\x1b[0;00m test number 6",
		expectErr:  false,
	},
	// Test critical output
	{
		template:   logFmt,
		prefix:     rgbterm.String("TEST>", 0, 255, 0),
		level:      LEVEL_CRITICAL,
		dateFormat: time.RubyDate,
		flags:      LstdFlags,
		text:       "test number 7",
		expect:     "%s \x1b[38;5;46mTEST>\x1b[0;00m \x1b[38;5;196m[CRITICAL]\x1b[0;00m test number 7",
		expectErr:  false,
	},
	// Test date format
	{
		template:   logFmt,
		prefix:     "::",
		level:      LEVEL_PRINT,
		dateFormat: "Mon 20060102 15:04:05",
		flags:      LstdFlags,
		text:       "test number 8",
		expect:     "%s :: test number 8",
		expectErr:  false,
	},
}

func TestFprintOutput(t *testing.T) {
	for i, k := range fprintOutputTests {
		var buf bytes.Buffer
		logr := New(LEVEL_DEBUG, &buf)
		logr.SetPrefix(k.prefix)
		logr.SetDateFormat(k.dateFormat)
		logr.SetFlags(k.flags)
		logr.SetLevel(k.level)
		d := time.Now().Format(logr.DateFormat())
		n, err := logr.Fprint(k.flags, k.level, 1, k.text, &buf)
		if n != buf.Len() {
			t.Error("Error: ", io.ErrShortWrite)
		}
		expect := fmt.Sprintf(k.expect, d)
		if buf.String() != expect || err != nil && !k.expectErr {
			t.Errorf("Test Number %d\nGot:\t%q\nExpect:\t"+"%q",
				i+1, buf.String(), expect)
			continue
		}
	}
}

func TestLevel(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_CRITICAL, &buf)
	logr.Debug("This level should produce no output")
	if buf.Len() != 0 {
		t.Errorf("Debug() produced output at LEVEL_CRITICAL logging level")
	}
	logr.SetLevel(LEVEL_DEBUG)
	logr.Debug("This level should produce output")
	if buf.Len() == 0 {
		t.Errorf("Debug() did not produce output at the LEVEL_DEBUG logging level")
	}
	buf.Reset()
	logr.SetLevel(LEVEL_CRITICAL)
	logr.Println("This level should produce output")
	if buf.Len() == 0 {
		t.Errorf("Debug() did not produce output at the ALL logging level")
	}
	buf.Reset()
	logr.SetLevel(LEVEL_PRINT)
	logr.Debug("This level should produce output")
	if buf.Len() == 0 {
		t.Errorf("Debug() did not produce output at the ALL logging level")
	}

	level := logr.Level()
	expl := LEVEL_PRINT

	if level != expl {
		t.Errorf("\nGot:\t%d\nExpect:\t%d\n", level, expl)
	}
}

func TestLevelString(t *testing.T) {
	var test level
	test = LEVEL_INFO
	if test.String() != "LEVEL_INFO" {
		t.Errorf("\nGot:\t%q\nExpect:\tLEVEL_INFO\n", test.String())
	}
}

func TestPrefixNewline(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(0)
	logr.Print("\n\nThis line should be padded with newlines.\n\n")
	expect := "\n\nThis line should be padded with newlines.\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsLdate(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(0)
	logr.Println("This output should not have a date.")
	expect := "This output should not have a date.\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsLfunctionName(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(LfunctionName)
	logr.Println("This output should have a function name.")
	expect := "TestFlagsLfunctionName: This output should have a function name.\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsLfunctionNameWithFileName(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(LfunctionName | LshortFileName)
	logr.Print("This output should have a file name and a function name.")
	expect := "logger_test.go: TestFlagsLfunctionNameWithFileName" +
		": This output should have a file name and a function name."
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsNoLcolorWithNewlinePadding(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_PRINT, &buf)
	logr.SetFlags(Llabel)
	logr.Debug("\n\nThis output should be padded with newlines and not colored.\n\n")
	expect := "\n\n[DEBUG] This output should be padded with newlines and not colored.\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsLcolorWithNewlinePaddingDebug(t *testing.T) {
	var buf bytes.Buffer
	SetStreams(&buf)
	logr := New(LEVEL_PRINT, &buf)
	logr.SetFlags(Lcolor | Llabel)
	logr.Debug("\n\nThis output should be padded with newlines and colored.\n\n")
	expect := "\n\n\x1b[38;5;231m[DEBUG]\x1b[0;00m This output should be " +
		"padded with newlines and colored.\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsLcolorWithNewlinePaddingDebugf(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_PRINT, &buf)
	logr.SetFlags(Lcolor | Llabel)
	logr.Debugf("\n\nThis output should be padded with newlines and %s.\n\n",
		"colored")
	expect := "\n\n\x1b[38;5;231m[DEBUG]\x1b[0;00m This output should be " +
		"padded with newlines and colored.\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
	buf.Reset()
	logr.Debugf("\n\n##### HELLO %s #####\n\n", "NEWMAN")
	expect = "\n\n\x1b[38;5;231m[DEBUG]\x1b[0;00m ##### HELLO NEWMAN #####\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestFlagsLcolorWithNewlinePaddingDebugln(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_PRINT, &buf)
	logr.SetFlags(Lcolor | Llabel)
	logr.Debugln("\n\nThis output should be padded with newlines and colored.\n\n")
	expect := "\n\n\x1b[38;5;231m[DEBUG]\x1b[0;00m This output should be " +
		"padded with newlines and colored.\n\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
	buf.Reset()
	logr.Debugln("\n\n", "### HELLO", "NEWMAN", "###", "\n\n")
	expect = "\n\n\x1b[38;5;231m[DEBUG]\x1b[0;00m  ### HELLO NEWMAN ### \n\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
	buf.Reset()
	logr.Debugln("\n\n### HELLO", "NEWMAN", "###\n\n")
	expect = "\n\n\x1b[38;5;231m[DEBUG]\x1b[0;00m ### HELLO NEWMAN ###\n\n\n"
	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}
}

func TestTreeDebugln(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_PRINT, &buf)
	logr.SetFlags(Lcolor | Lid | Ltree | Llabel)

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

	expe := "\x1b[38;5;231m[DBUG]\x1b[0;00m [00] Level 0 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [01]     Level 1 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [01]     Level 1 Output 2\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [02]         Level 2 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [02]         Level 2 Output 2\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [03]             Level 3 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [02]         Level 2 Output 3\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [01]     Level 1 Output 3\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m [00] Level 0 Output 2\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestSetIndentDebugln(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(Lcolor | Lindent | Llabel)

	logr.Debugln("Level 0 Output 1")
	logr.SetIndent(1).Debugln("Level 1 Output 1")
	logr.Debugln("Level 1 Output 2")
	logr.SetIndent(0).Debugln("Level 0 Output 1")

	expe := "\x1b[38;5;231m[DEBUG]\x1b[0;00m Level 0 Output 1\n" +
		"\x1b[38;5;231m[DEBUG]\x1b[0;00m     Level 1 Output 1\n" +
		"\x1b[38;5;231m[DEBUG]\x1b[0;00m     Level 1 Output 2\n" +
		"\x1b[38;5;231m[DEBUG]\x1b[0;00m Level 0 Output 1\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestLindentWithLshowIndent(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(Lcolor | Lindent | LshowIndent | Llabel)

	logr.Debugln("Level 0 Output 1")
	logr.SetIndent(1).Debugln("Level 1 Output 1")
	logr.Debugln("Level 1 Output 2")
	logr.SetIndent(0).Debugln("Level 0 Output 1")

	expe := "\x1b[38;5;231m[DEBUG]\x1b[0;00m Level 0 Output 1\n" +
		"\x1b[38;5;231m[DEBUG]\x1b[0;00m \x1b[38;5;31m...|\x1b[0;00mLevel 1 Output 1\n" +
		"\x1b[38;5;231m[DEBUG]\x1b[0;00m \x1b[38;5;31m...|\x1b[0;00mLevel 1 Output 2\n" +
		"\x1b[38;5;231m[DEBUG]\x1b[0;00m Level 0 Output 1\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestSetIndentWithLindentAndLtree(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	// Lindent should have no effect on Lindent
	logr.SetFlags(Lcolor | Lindent | Ltree | LshowIndent | Llabel)

	logr.SetIndent(1).Debugln("Level 0 Output 1")
	lvl3 := func() {
		logr.Debugln("Level 3 Output 1")
	}
	lvl2 := func() {
		logr.Debugln("Level 2 Output 1")
		lvl3()
		logr.Debugln("Level 2 Output 2")
	}
	lvl1 := func() {
		logr.Debugln("Level 1 Output 1")
		lvl2()
		logr.Debugln("Level 1 Output 3")
	}
	lvl1()
	logr.Debugln("Level 0 Output 2")

	expe := "\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|\x1b[0;00mLevel 0 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|...|\x1b[0;00mLevel 1 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|...|...|\x1b[0;00mLevel 2 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|...|...|...|\x1b[0;00mLevel 3 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|...|...|\x1b[0;00mLevel 2 Output 2\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|...|\x1b[0;00mLevel 1 Output 3\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|\x1b[0;00mLevel 0 Output 2\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestSetIndentWithLindentAndLtreeMinus2Indent(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	// Lindent should have no effect on Lindent
	logr.SetFlags(Lcolor | Lindent | Ltree | LshowIndent | Llabel)

	logr.SetIndent(-2).Debugln("Level 0 Output 1")
	lvl3 := func() {
		logr.Debugln("Level 3 Output 1")
	}
	lvl2 := func() {
		logr.Debugln("Level 2 Output 1")
		lvl3()
		logr.Debugln("Level 2 Output 2")
	}
	lvl1 := func() {
		logr.Debugln("Level 1 Output 1")
		lvl2()
		logr.Debugln("Level 1 Output 3")
	}
	lvl1()
	logr.Debugln("Level 0 Output 2")

	expe := "\x1b[38;5;231m[DBUG]\x1b[0;00m Level 0 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 1 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 2 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|\x1b[0;00mLevel 3 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 2 Output 2\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 1 Output 3\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 0 Output 2\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestSetIndentWithLindentAndLtreeMinus4Indent(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	// Lindent should have no effect on Lindent
	logr.SetFlags(Lcolor | Lindent | Ltree | LshowIndent | Llabel)

	logr.SetIndent(-4).Debugln("Level 0 Output 1")
	lvl3 := func() {
		logr.Debugln("Level 3 Output 1")
	}
	lvl2 := func() {
		logr.Debugln("Level 2 Output 1")
		lvl3()
		logr.Debugln("Level 2 Output 2")
	}
	lvl1 := func() {
		logr.Debugln("Level 1 Output 1")
		lvl2()
		logr.Debugln("Level 1 Output 3")
	}
	lvl1()
	logr.Debugln("Level 0 Output 2")

	expe := "\x1b[38;5;231m[DBUG]\x1b[0;00m Level 0 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 1 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 2 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 3 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 2 Output 2\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 1 Output 3\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m Level 0 Output 2\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestTemplate(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(LdebugFlags)

	logr.SetTemplate("{{.Text}}")
	temp := logr.Template()

	type test struct {
		Text string
	}

	err := temp.Execute(&buf, &test{"Hello, World!"})
	if err != nil {
		t.Fatal(err)
	}

	expe := "Hello, World!"

	if buf.String() != expe {
		t.Errorf("\nGot:\t%s\nExpect:\t%s\n", buf.String(), expe)
	}
}

func TestStandardLabelLength(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	// Lindent should have no effect on Lindent
	logr.SetFlags(Lcolor | Lindent | Ltree | LshowIndent | Llabel)

	logr.SetIndent(1).Println("Level 0 Output 1")
	lvl3 := func() {
		logr.Criticalln("Level 3 Output 1")
		logr.Errorln("Level 3 Output 2")
	}
	lvl2 := func() {
		logr.Infoln("Level 2 Output 1")
		lvl3()
		logr.Infoln("Level 2 Output 2")
	}
	lvl1 := func() {
		logr.Debugln("Level 1 Output 1")
		lvl2()
		logr.Warningln("Level 1 Output 3")
	}
	lvl1()
	logr.Println("Level 0 Output 2")

	expe := "       \x1b[38;5;31m...|\x1b[0;00mLevel 0 Output 1\n" +
		"\x1b[38;5;231m[DBUG]\x1b[0;00m \x1b[38;5;31m...|...|\x1b[0;00mLevel 1 Output 1\n" +
		"\x1b[38;5;41m[INFO]\x1b[0;00m \x1b[38;5;31m...|...|...|\x1b[0;00mLevel 2 Output 1\n" +
		"\x1b[38;5;196m[CRIT]\x1b[0;00m \x1b[38;5;31m...|...|...|...|\x1b[0;00mLevel 3 Output 1\n" +
		"\x1b[38;5;202m[EROR]\x1b[0;00m \x1b[38;5;31m...|...|...|...|\x1b[0;00mLevel 3 Output 2\n" +
		"\x1b[38;5;41m[INFO]\x1b[0;00m \x1b[38;5;31m...|...|...|\x1b[0;00mLevel 2 Output 2\n" +
		"\x1b[38;5;228m[WARN]\x1b[0;00m \x1b[38;5;31m...|...|\x1b[0;00mLevel 1 Output 3\n" +
		"       \x1b[38;5;31m...|\x1b[0;00mLevel 0 Output 2\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestStandardLabelLengthNoColor(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	// Lindent should have no effect on Lindent
	logr.SetFlags(Lindent | Ltree | LshowIndent | Llabel)

	logr.SetIndent(1).Println("Level 0 Output 1")
	lvl3 := func() {
		logr.Criticalln("Level 3 Output 1")
		logr.Errorln("Level 3 Output 2")
	}
	lvl2 := func() {
		logr.Infoln("Level 2 Output 1")
		lvl3()
		logr.Infoln("Level 2 Output 2")
	}
	lvl1 := func() {
		logr.Debugln("Level 1 Output 1")
		lvl2()
		logr.Warningln("Level 1 Output 3")
	}
	lvl1()
	logr.Println("Level 0 Output 2")

	expe := "       ...|Level 0 Output 1\n" +
		"[DBUG] ...|...|Level 1 Output 1\n" +
		"[INFO] ...|...|...|Level 2 Output 1\n" +
		"[CRIT] ...|...|...|...|Level 3 Output 1\n" +
		"[EROR] ...|...|...|...|Level 3 Output 2\n" +
		"[INFO] ...|...|...|Level 2 Output 2\n" +
		"[WARN] ...|...|Level 1 Output 3\n" +
		"       ...|Level 0 Output 2\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}
}
func TestSetTemplate(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(LdebugFlags)

	logr.SetTemplate("{{.Text}}")

	logr.Debugln("Hello, World!")

	expe := "Hello, World!\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expe)
	}
}

func TestSetTemplateBad(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(LdebugFlags)

	err := logr.SetTemplate("{{.Text")

	logr.Debugln("template: default:1: unclosed action")

	expe := "template: default:1: unclosed action"

	if err.Error() != expe {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expe)
	}
}

func TestSetTemplateBadDataObjectPanic(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(Lindent)

	logr.SetIndent(1)

	type test struct {
		Test string
	}

	err := logr.SetTemplate("{{.Tes}}")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("\nGot:\t%q\nExpect:\tPANIC\n", buf.String())
		}
	}()

	logr.Debugln("Hello, World!")

}

func TestDateFormat(t *testing.T) {
	logr := New(LEVEL_INFO)

	dateFormat := logr.DateFormat()

	expect := "Mon Jan 02 15:04:05 MST 2006"

	if dateFormat != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", dateFormat, expect)
	}
}

func TestSetDateFormat(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_PRINT, &buf)

	logr.SetFlags(Ldate)

	logr.SetDateFormat("20060102-15:04:05")

	logr.SetTemplate("{{.Date}}")

	logr.Debugln("Hello")

	expect := time.Now().Format(logr.DateFormat())

	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}

	// Reset the standard logging object
	SetTemplate(logFmt)
}

func TestFlags(t *testing.T) {
	logr := New(LEVEL_INFO)

	logr.SetFlags(LstdFlags)

	flags := logr.Flags()

	expect := LstdFlags

	if flags != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", flags, expect)
	}
}

func TestPrefix(t *testing.T) {
	logr := New(LEVEL_INFO)

	logr.SetPrefix("TEST::")

	prefix := logr.Prefix()

	expect := "TEST::"

	if prefix != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", prefix, expect)
	}
}

func TestStreams(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_INFO)

	logr.SetStreams(&buf)

	bufT := logr.Streams()

	if &buf != bufT[0] {
		t.Errorf("\nGot:\t%p\nExpect:\t%p\n", &buf, bufT[0])
	}
}

func TestIndent(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG)

	logr.SetStreams(&buf)

	logr.SetFlags(Lindent | Llabel | Llabel)

	logr.SetIndent(0).Debugln("Test 1")
	logr.SetIndent(2).Debugln("Test 2")

	indent := logr.Indent()

	expe := "[DEBUG] Test 1\n[DEBUG]         Test 2\n"
	expi := 2

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}

	if indent != expi {
		t.Errorf("\nGot:\t%d\nExpect:\t%d\n", indent, expi)
	}
}

func TestTabStop(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(Lindent | Llabel)

	// This SetIndent doesn't have to be on a separate line, but for some
	// reason go test cover wasn't registering its usage when the functions
	// below were chained together.
	logr.SetIndent(1)
	logr.SetTabStop(2).Debugln("Test 1")

	logr.SetIndent(2)
	logr.SetTabStop(4).Debugln("Test 2")

	tabStop := logr.TabStop()

	expe := "[DEBUG]   Test 1\n[DEBUG]         Test 2\n"
	expt := 4

	if buf.String() != expe {
		t.Errorf("\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			buf.String(), buf.String(), expe, expe)
	}

	if tabStop != expt {
		t.Errorf("\nGot:\t%d\nExpect:\t%d\n", tabStop, expt)
	}
}

// TestLnoFileAnsi verifies output sent to os.Stdout contains color codes
// and output sent to a file does not.
func TestLnoFileAnsi(t *testing.T) {
	logr := New(LEVEL_DEBUG)

	logr.SetFlags(Lprefix | Llabel | Lcolor | LnoFileAnsi)

	f, err := ioutil.TempFile("/tmp", "go-elog-test-")
	defer f.Close()
	if err != nil {
		t.Fatal(err)
	}

	r, w, err := os.Pipe()
	if err != nil {
		t.Error(err)
	}
	oStdout := os.Stdout
	os.Stdout = w
	logr.SetStreams(f, os.Stdout)

	logr.Debugln("Test 1")
	logr.Debugln("Test 2")

	os.Stdout = oStdout
	w.Close()

	fOut, _ := ioutil.ReadFile(f.Name())
	stdOut, _ := ioutil.ReadAll(r)

	expe := ":: [DEBUG] Test 1\n:: [DEBUG] Test 2\n"
	expeStdout := "\x1b[38;5;48m::\x1b[0;00m \x1b[38;5;231m[DEBUG]" +
		"\x1b[0;00m Test 1\n\x1b[38;5;48m::\x1b[0;00m \x1b[38;5;231m[DEBUG]" +
		"\x1b[0;00m Test 2\n"

	if string(fOut) != expe {
		t.Errorf("%s\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			"Incorrect file output!",
			string(fOut), string(fOut), expe, expe)
	} else if string(stdOut) != expeStdout {
		t.Errorf("%s\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			"Stdout contained invalid data!",
			string(stdOut), string(stdOut), expeStdout, expeStdout)
	}
}

var printFunctionTests = []struct {
	name   string
	format string
	input  string
	expect string
}{
	{name: "Test 1", format: "%s", input: "Hello, world!", expect: "Hello, world!"},
}

func TestPrintFunctions(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(Llabel)

	for _, test := range printFunctionTests {

		check := func(output, expect, funcName string) {
			if output != expect {
				t.Errorf("\nName: %q\nFunction: %s\nGot: %q\nExpect: %q\n",
					test.name, funcName, output, expect)
			}
		}

		checkOutput := func(pFunc func(...interface{}), lvl string) {
			nl := ""
			pFunc(test.input)
			label := LevelFromString(lvl).Label()
			if len(label) > 1 {
				label = label + " "
			}
			fName := runtime.FuncForPC(reflect.ValueOf(pFunc).Pointer()).Name()
			lenfName := len(fName)
			if fName[lenfName-6:lenfName-4] == "ln" {
				nl = "\n"
			}
			check(buf.String(), label+test.expect+nl, fName)
			buf.Reset()
		}

		checkFormatOutput := func(pFunc func(string, ...interface{}), lvl string) {
			nl := ""
			pFunc(test.format, test.input)
			label := LevelFromString(lvl).Label()
			if len(label) > 1 {
				label = label + " "
			}
			fName := runtime.FuncForPC(reflect.ValueOf(pFunc).Pointer()).Name()
			lenfName := len(fName)
			if fName[lenfName-2:] == "ln" {
				nl = "\n"
			}
			check(buf.String(), label+test.expect+nl, fName)
			buf.Reset()
		}

		checkOutput(logr.Print, "PRINT")
		checkOutput(logr.Println, "PRINT")
		checkFormatOutput(logr.Printf, "PRINT")
		checkOutput(logr.Debug, "DEBUG")
		checkOutput(logr.Debugln, "DEBUG")
		checkFormatOutput(logr.Debugf, "DEBUG")
		checkOutput(logr.Info, "INFO")
		checkOutput(logr.Infoln, "INFO")
		checkFormatOutput(logr.Infof, "INFO")
		checkOutput(logr.Warning, "WARNING")
		checkOutput(logr.Warningln, "WARNING")
		checkFormatOutput(logr.Warningf, "WARNING")
		checkOutput(logr.Error, "ERROR")
		checkOutput(logr.Errorln, "ERROR")
		checkFormatOutput(logr.Errorf, "ERROR")
		checkOutput(logr.Critical, "CRITICAL")
		checkOutput(logr.Criticalln, "CRITICAL")
		checkFormatOutput(logr.Criticalf, "CRITICAL")

	}
}

func TestPanic(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(Llabel)

	expect := "[CRITICAL] Panic Error!"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should generate panic!")
		}
		if buf.String() != expect {
			t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
		}
	}()

	logr.Panic("Panic Error!")
}

func TestPanicln(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(Llabel)

	expect := "[CRITICAL] Panic Error!\n"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should generate panic!")
		}
		if buf.String() != expect {
			t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
		}
	}()

	logr.Panicln("Panic Error!")
}

func TestPanicf(t *testing.T) {
	var buf bytes.Buffer

	logr := New(LEVEL_DEBUG, &buf)

	logr.SetFlags(Llabel)

	expect := "[CRITICAL] Panic Error!\n"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should generate panic!")
		}
		if buf.String() != expect {
			t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
		}
	}()

	logr.Panicf("%s\n", "Panic Error!")
}

var levelFromStringTests = []struct {
	name   string
	input  string
	expect level
}{
	{name: "Test 1", input: "debug", expect: LEVEL_DEBUG},
	{name: "Test 2", input: "Debug", expect: LEVEL_DEBUG},
	{name: "Test 3", input: "_debug", expect: LEVEL_PRINT},
	{name: "Test 4", input: "PRINT", expect: LEVEL_PRINT},
	{name: "Test 5", input: "level_PRINT", expect: LEVEL_PRINT},
	{name: "Test 6", input: "info", expect: LEVEL_INFO},
	{name: "Test 7", input: "_info", expect: LEVEL_PRINT},
	{name: "Test 8", input: "level_info", expect: LEVEL_INFO},
}

func TestLevelFromString(t *testing.T) {
	for _, test := range levelFromStringTests {
		lvl := LevelFromString(test.input)
		if lvl != test.expect {
			t.Errorf("\nTest:\t%q\nGot:\t%q\nExpect:\t%q\n", test.name, lvl, test.expect)
		}
	}
}

var excludeIDtests = []struct {
	name   string
	ids    []int
	flags  int
	expect string
}{
	{name: "Test excluding one ID", ids: []int{1}, flags: Lid | Ltree | LshowIndent,
		expect: "[00] Hello!\n[02] ...|...|should be suppressed.\n" +
			"[03] ...|...|...|Almost forgot...\n" +
			"[02] ...|...|but we'll find out!\n" +
			"[00] Goodbye!\n",
	},
	{name: "Test excluding two IDs", ids: []int{1, 3}, flags: Lid | Ltree | LshowIndent,
		expect: "[00] Hello!\n[02] ...|...|should be suppressed.\n" +
			"[02] ...|...|but we'll find out!\n" +
			"[00] Goodbye!\n",
	},
	{name: "Test excluding two IDs without Lid", ids: []int{1, 3}, flags: Ltree | LshowIndent,
		expect: "Hello!\n" +
			"...|The things\n" +
			"...|...|should be suppressed.\n" +
			"...|...|...|Almost forgot...\n" +
			"...|...|but we'll find out!\n" +
			"...|that can be suppressed.\n" +
			"Goodbye!\n",
	},
	{name: "Test excluding two IDs with only Lid", ids: []int{1, 3}, flags: Lid,
		expect: "[00] Hello!\n" +
			"[02] should be suppressed.\n" +
			"[02] but we'll find out!\n" +
			"[00] Goodbye!\n",
	},
}

func TestExcludeByHeirarchyID(t *testing.T) {
	var buf bytes.Buffer

	for _, test := range excludeIDtests {
		logr := New(LEVEL_DEBUG, &buf)

		logr.SetFlags(test.flags)

		logr.ExcludeByHeirarchyID(test.ids...)

		logr.Debugln("Hello!")
		lvl3 := func() {
			logr.Debugln("Almost forgot...")
		}
		lvl2 := func() {
			logr.Debugln("should be suppressed.")
			lvl3()
			logr.Debugln("but we'll find out!")
		}
		lvl1 := func() {
			logr.Debugln("The things")
			lvl2()
			logr.Debugln("that can be suppressed.")
		}
		lvl1()
		logr.Debugln("Goodbye!")

		if buf.String() != test.expect {
			t.Errorf("\nTest: %s\n\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
				test.name, buf.String(), buf.String(), test.expect, test.expect)
		}
		buf.Reset()
	}
}

var excludeByStringTests = []struct {
	name   string
	flags  int
	input  []string
	expect string
}{
	{name: "Exclude single word", flags: Llabel, input: []string{"Hello"},
		expect: "[DEBUG] The things\n" +
			"[DEBUG] should be suppressed.\n" +
			"[DEBUG] Almost forgot...\n" +
			"[DEBUG] but we'll find out!\n" +
			"[DEBUG] that can be suppressed.\n" +
			"[DEBUG] Goodbye!\n",
	},
	{name: "Exclude single word 2", flags: Llabel, input: []string{"DEBUG"},
		expect: "[DEBUG] Hello!\n" +
			"[DEBUG] The things\n" +
			"[DEBUG] should be suppressed.\n" +
			"[DEBUG] Almost forgot...\n" +
			"[DEBUG] but we'll find out!\n" +
			"[DEBUG] that can be suppressed.\n" +
			"[DEBUG] Goodbye!\n",
	},
	{name: "Exclude non-existing word", input: []string{"Things"},
		expect: "Hello!\n" +
			"The things\n" +
			"should be suppressed.\n" +
			"Almost forgot...\n" +
			"but we'll find out!\n" +
			"that can be suppressed.\n" +
			"Goodbye!\n",
	},
	{name: "Exclude two words", flags: Ltree | Lcolor | Lindent | Lid | LshowIndent,
		input: []string{"forgot", "we'll"},
		expect: "[00] Hello!\n" +
			"[01] \x1b[38;5;31m...|\x1b[0;00mThe things\n" +
			"[02] \x1b[38;5;31m...|...|\x1b[0;00mshould be suppressed.\n" +
			"[01] \x1b[38;5;31m...|\x1b[0;00mthat can be suppressed.\n" +
			"[00] Goodbye!\n",
	},
}

func TestExcludeByString(t *testing.T) {
	var buf bytes.Buffer

	for _, test := range excludeByStringTests {
		logr := New(LEVEL_DEBUG, &buf)

		logr.SetFlags(test.flags)

		logr.ExcludeByString(test.input...)

		logr.Debugln("Hello!")
		lvl3 := func() {
			logr.Debugln("Almost forgot...")
		}
		lvl2 := func() {
			logr.Debugln("should be suppressed.")
			lvl3()
			logr.Debugln("but we'll find out!")
		}
		lvl1 := func() {
			logr.Debugln("The things")
			lvl2()
			logr.Debugln("that can be suppressed.")
		}
		lvl1()
		logr.Debugln("Goodbye!")

		if buf.String() != test.expect {
			t.Errorf("\nTest: %s\n\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
				test.name, buf.String(), buf.String(), test.expect, test.expect)
		}
		buf.Reset()
	}
}

func testLvl3(logr *Logger) {
	logr.Debugln("Almost forgot...")
}

func testLvl2(logr *Logger) {
	logr.Debugln("should be suppressed.")
	testLvl3(logr)
	logr.Debugln("but we'll find out!")
}

func testLvl1(logr *Logger) {
	logr.Debugln("The things")
	testLvl2(logr)
	logr.Debugln("that can be suppressed.")
}

var excludeByFuncNameTests = []struct {
	name   string
	flags  int
	input  []string
	expect string
}{
	{name: "Exclude function", flags: LfunctionName,
		input: []string{"TestExcludeByFuncName", "TestStdExcludeByFuncName"},
		expect: "testLvl1: The things\n" +
			"testLvl2: should be suppressed.\n" +
			"testLvl3: Almost forgot...\n" +
			"testLvl2: but we'll find out!\n" +
			"testLvl1: that can be suppressed.\n",
	},
	{name: "Exclude without LfunctionName", flags: Llabel,
		input: []string{"TestExcludeByFuncName", "TestStdExcludeByFuncName"},
		expect: "[DEBUG] The things\n" +
			"[DEBUG] should be suppressed.\n" +
			"[DEBUG] Almost forgot...\n" +
			"[DEBUG] but we'll find out!\n" +
			"[DEBUG] that can be suppressed.\n",
	},
	{name: "Exclude non-existing name", input: []string{"Imaginary"},
		expect: "Hello!\n" +
			"The things\n" +
			"should be suppressed.\n" +
			"Almost forgot...\n" +
			"but we'll find out!\n" +
			"that can be suppressed.\n" +
			"Goodbye!\n",
	},
	{name: "Exclude two function names", flags: Ltree | LfunctionName | Lcolor | Lindent | Lid | LshowIndent,
		input: []string{"TestExcludeByFuncName", "TestStdExcludeByFuncName", "testLvl3"},
		expect: "[01] \x1b[38;5;31m...|\x1b[0;00mtestLvl1: The things\n" +
			"[02] \x1b[38;5;31m...|...|\x1b[0;00mtestLvl2: should be suppressed.\n" +
			"[02] \x1b[38;5;31m...|...|\x1b[0;00mtestLvl2: but we'll find out!\n" +
			"[01] \x1b[38;5;31m...|\x1b[0;00mtestLvl1: that can be suppressed.\n",
	},
}

func TestExcludeByFuncName(t *testing.T) {
	var buf bytes.Buffer

	for _, test := range excludeByFuncNameTests {
		logr := New(LEVEL_DEBUG, &buf)

		logr.SetFlags(test.flags)

		logr.ExcludeByFuncName(test.input...)

		logr.Debugln("Hello!")
		testLvl1(logr)
		logr.Debugln("Goodbye!")

		if buf.String() != test.expect {
			t.Errorf("\nTest: %s\n\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
				test.name, buf.String(), buf.String(), test.expect, test.expect)
		}
		buf.Reset()
	}
}

func TestWithFlags(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(Llabel | Lprefix)

	logr.Debugln("Test 1")
	logr.WithFlags(0, logr.Debugln, "Test 2")

	expe := ":: [DEBUG] Test 1\nTest 2\n"

	if buf.String() != expe {
		t.Errorf("%s\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			"Incorrect output!",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestWithFlagsf(t *testing.T) {
	var buf bytes.Buffer
	logr := New(LEVEL_DEBUG, &buf)
	logr.SetFlags(Llabel | Lprefix)

	logr.Debugln("Test 1")
	logr.WithFlagsf(0, logr.Debugf, "%s\n", "Test 2")

	expe := ":: [DEBUG] Test 1\nTest 2\n"

	if buf.String() != expe {
		t.Errorf("%s\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			"Incorrect output!",
			buf.String(), buf.String(), expe, expe)
	}
}
