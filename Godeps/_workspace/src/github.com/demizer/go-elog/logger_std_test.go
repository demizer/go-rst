// Copyright 2013,2014 The go-logger Authors. All rights reserved.
// This code is MIT licensed. See the LICENSE file for more info.

// Tests for the default standard logging object

package log

import (
	"bytes"
	"io/ioutil"
	"os"
	"reflect"
	"runtime"
	"testing"
	"time"
)

func TestStdTemplate(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(LdebugFlags)

	SetTemplate("{{.Text}}")
	temp := Template()

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

func TestStdSetTemplate(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(LdebugFlags)

	SetTemplate("{{.Text}}")

	Debugln("Hello, World!")

	expe := "Hello, World!\n"

	if buf.String() != expe {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expe)
	}
}

func TestStdSetTemplateBad(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(LdebugFlags)

	err := SetTemplate("{{.Text")

	Debugln("template: default:1: unclosed action")

	expe := "template: default:1: unclosed action"

	if err.Error() != expe {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expe)
	}
}

func TestStdSetTemplateBadDataObjectPanic(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Lindent)

	SetIndent(1)

	type test struct {
		Test string
	}

	err := SetTemplate("{{.Tes}}")
	if err != nil {
		t.Fatal(err)
	}

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("\nGot:\t%q\nExpect:\tPANIC\n", buf.String())
		}
	}()

	Debugln("Hello, World!")

	// Reset the standard logging object
	SetTemplate(logFmt)
	SetIndent(0)
}

func TestStdDateFormat(t *testing.T) {
	dateFormat := DateFormat()

	expect := "Mon Jan 02 15:04:05 MST 2006"

	if dateFormat != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", dateFormat, expect)
	}
}

func TestStdSetDateFormat(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_PRINT, &buf)

	SetFlags(Ldate)

	SetDateFormat("20060102-15:04:05")

	SetTemplate("{{.Date}}")

	Debugln("Hello")

	expect := time.Now().Format(DateFormat())

	if buf.String() != expect {
		t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
	}

	// Reset the standard logging object
	SetTemplate(logFmt)
}

func TestStdFlags(t *testing.T) {
	SetFlags(LstdFlags)

	flags := Flags()

	expect := LstdFlags

	if flags != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", flags, expect)
	}
}

func TestStdLevel(t *testing.T) {
	SetLevel(LEVEL_DEBUG)

	level := Level()

	expect := "LEVEL_DEBUG"

	if level.String() != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", level, expect)
	}
}

func TestStdPrefix(t *testing.T) {
	SetPrefix("TEST::")

	prefix := Prefix()

	expect := "TEST::"

	if prefix != expect {
		t.Errorf("\nGot:\t%#v\nExpect:\t%#v\n", prefix, expect)
	}
}

func TestStdStreams(t *testing.T) {
	var buf bytes.Buffer

	SetStreams(&buf)

	bufT := Streams()

	if &buf != bufT[0] {
		t.Errorf("\nGot:\t%p\nExpect:\t%p\n", &buf, bufT[0])
	}
}

func TestStdIndent(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Lindent | Llabel)

	SetIndent(0).Debugln("Test 1")
	SetIndent(2).Debugln("Test 2")

	indent := Indent()

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

func TestStdTabStop(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Lindent | Llabel)

	// This SetIndent doesn't have to be on a separate line, but for some
	// reason go test cover wasn't registering its usage when the functions
	// below were chained together.
	SetIndent(1)
	SetTabStop(2).Debugln("Test 1")

	SetIndent(2)
	SetTabStop(4).Debugln("Test 2")

	tabStop := TabStop()

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

// TestStdLnoFileAnsi verifies output sent to os.Stdout contains color codes
// and output sent to a file does not.
func TestStdLnoFileAnsi(t *testing.T) {
	std = New(LEVEL_DEBUG)
	SetFlags(Lprefix | Llabel | Lcolor | LnoFileAnsi)

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
	SetStreams(f, os.Stdout)

	Debugln("Test 1")
	Debugln("Test 2")

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

var stdOutputTests = []struct {
	name   string
	format string
	input  string
	expect string
}{
	{name: "Test 1", format: "%s", input: "Hello, world!", expect: "Hello, world!"},
}

func TestStdOutput(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Llabel)

	SetIndent(0)

	for _, test := range stdOutputTests {

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
			if fName[lenfName-2:] == "ln" {
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

		checkOutput(Print, "PRINT")
		checkOutput(Println, "PRINT")
		checkFormatOutput(Printf, "PRINT")
		checkOutput(Debug, "DEBUG")
		checkOutput(Debugln, "DEBUG")
		checkFormatOutput(Debugf, "DEBUG")
		checkOutput(Info, "INFO")
		checkOutput(Infoln, "INFO")
		checkFormatOutput(Infof, "INFO")
		checkOutput(Warning, "WARNING")
		checkOutput(Warningln, "WARNING")
		checkFormatOutput(Warningf, "WARNING")
		checkOutput(Error, "ERROR")
		checkOutput(Errorln, "ERROR")
		checkFormatOutput(Errorf, "ERROR")
		checkOutput(Critical, "CRITICAL")
		checkOutput(Criticalln, "CRITICAL")
		checkFormatOutput(Criticalf, "CRITICAL")

	}
}

func TestStdPanic(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Llabel)

	SetIndent(0)

	expect := "[CRITICAL] Panic Error!"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should generate panic!")
		}
		if buf.String() != expect {
			t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
		}
	}()

	Panic("Panic Error!")
}

func TestStdPanicln(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Llabel)

	SetIndent(0)

	expect := "[CRITICAL] Panic Error!\n"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should generate panic!")
		}
		if buf.String() != expect {
			t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
		}
	}()

	Panicln("Panic Error!")
}

func TestStdPanicf(t *testing.T) {
	var buf bytes.Buffer

	std = New(LEVEL_DEBUG, &buf)

	SetFlags(Llabel)

	SetIndent(0)

	expect := "[CRITICAL] Panic Error!\n"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("Test should generate panic!")
		}
		if buf.String() != expect {
			t.Errorf("\nGot:\t%q\nExpect:\t%q\n", buf.String(), expect)
		}
	}()

	Panicf("%s\n", "Panic Error!")
}

func TestStdExcludeByHeirarchyID(t *testing.T) {
	var buf bytes.Buffer

	// excludeIDtests is defined in logger_test.go
	for _, test := range excludeIDtests {
		std = New(LEVEL_DEBUG, &buf)

		SetFlags(test.flags)

		ExcludeByHeirarchyID(test.ids...)

		Debugln("Hello!")
		lvl3 := func() {
			Debugln("Almost forgot...")
		}
		lvl2 := func() {
			Debugln("should be suppressed.")
			lvl3()
			Debugln("but we'll find out!")
		}
		lvl1 := func() {
			Debugln("The things")
			lvl2()
			Debugln("that can be suppressed.")
		}
		lvl1()
		Debugln("Goodbye!")

		if buf.String() != test.expect {
			t.Errorf("\nTest: %s\n\nGot:\n\n%s\n%q\n"+
				"\nExpect:\n\n%s\n%q\n\n",
				test.name, buf.String(), buf.String(),
				test.expect, test.expect)
		}
		buf.Reset()
	}
}

func TestStdExcludeByString(t *testing.T) {
	var buf bytes.Buffer

	for _, test := range excludeByStringTests {
		std = New(LEVEL_DEBUG, &buf)

		SetFlags(test.flags)

		ExcludeByString(test.input...)

		Debugln("Hello!")
		lvl3 := func() {
			Debugln("Almost forgot...")
		}
		lvl2 := func() {
			Debugln("should be suppressed.")
			lvl3()
			Debugln("but we'll find out!")
		}
		lvl1 := func() {
			Debugln("The things")
			lvl2()
			Debugln("that can be suppressed.")
		}
		lvl1()
		Debugln("Goodbye!")

		if buf.String() != test.expect {
			t.Errorf("\nTest: %s\n\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
				test.name, buf.String(), buf.String(), test.expect, test.expect)
		}
		buf.Reset()
	}
}

func TestStdExcludeByFuncName(t *testing.T) {
	var buf bytes.Buffer

	for _, test := range excludeByFuncNameTests {
		std = New(LEVEL_DEBUG, &buf)

		SetFlags(test.flags)

		ExcludeByFuncName(test.input...)

		Debugln("Hello!")
		testLvl1(std)
		Debugln("Goodbye!")

		if buf.String() != test.expect {
			t.Errorf("\nTest: %s\n\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
				test.name, buf.String(), buf.String(), test.expect, test.expect)
		}
		buf.Reset()
	}
}

func TestStdWithFlags(t *testing.T) {
	var buf bytes.Buffer
	std = New(LEVEL_DEBUG, &buf)
	SetFlags(Llabel | Lprefix)

	Debugln("Test 1")
	WithFlags(0, Debugln, "Test 2")

	expe := ":: [DEBUG] Test 1\nTest 2\n"

	if buf.String() != expe {
		t.Errorf("%s\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			"Incorrect file output!",
			buf.String(), buf.String(), expe, expe)
	}
}

func TestStdWithFlagsf(t *testing.T) {
	var buf bytes.Buffer
	std = New(LEVEL_DEBUG, &buf)
	SetFlags(Llabel | Lprefix)

	Debugln("Test 1")
	WithFlagsf(0, Debugf, "%s\n", "Test 2")

	expe := ":: [DEBUG] Test 1\nTest 2\n"

	if buf.String() != expe {
		t.Errorf("%s\nGot:\n\n%s\n%q\n\nExpect:\n\n%s\n%q\n\n",
			"Incorrect file output!",
			buf.String(), buf.String(), expe, expe)
	}
}
