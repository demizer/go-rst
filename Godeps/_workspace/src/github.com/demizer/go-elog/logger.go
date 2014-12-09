// Copyright 2013,2014 The go-elog Authors. All rights reserved.
// This code is MIT licensed. See the LICENSE file for more info.

// The go-elog package is a drop in replacement for the Go standard log package
// that provides a number of enhancements. Including colored output, logging
// levels, custom log formatting, and multiple simultaneous output streams like
// os.Stdout or a File.
package log

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"sync"
	"text/template"
	"time"

	"github.com/aybabtme/rgbterm"
)

// Label contains the name of a label as well as the short name and RGB color
// values.
type Label struct {
	level level

	name string

	// StdLengthNames are used for output when the Ltree flag is set. Short
	// names of all labels should be the same length.
	StdLengthName string

	colorRGB [3]uint8
}

// StdLenName() is used for output that must be aligned accross labels.
// Specifically for output when the Ltree flag is set.
func (l Label) StdLenName() string { return l.StdLengthName }

// String satisfies the Stringer interface.
func (l Label) String() string { return l.name }

// Colorized returns the colorized label for console output using ANSI escape
// sequences.
func (l Label) Colorized() string {
	if l.level == LEVEL_PRINT {
		return l.name
	}
	return rgbterm.String(l.name, l.colorRGB[0], l.colorRGB[1], l.colorRGB[2])
}

// StdLenColorized returns the colorized standard length label for console
// output using ANSI escape sequences.
func (l Label) StdLenColorized() string {
	if l.level == LEVEL_PRINT {
		return l.StdLenName()
	}
	return rgbterm.String(l.StdLenName(), l.colorRGB[0], l.colorRGB[1],
		l.colorRGB[2])
}

// Labels are prefixed to the beginning of a string on output. Labels can be
// colored. A special shortened case is used when the Ltree flag is set so that
// ouput is properly aligned.
var Labels = [6]Label{
	Label{LEVEL_DEBUG, "[DEBUG]", "[DBUG]",
		[3]uint8{255, 255, 255}, // White
	},

	Label{LEVEL_INFO, "[INFO]", "[INFO]",
		[3]uint8{0, 215, 95}, // Green
	},

	Label{LEVEL_WARNING, "[WARNING]", "[WARN]",
		[3]uint8{255, 255, 135}, // Yellow
	},

	Label{LEVEL_ERROR, "[ERROR]", "[EROR]",
		[3]uint8{255, 99, 0}, // Orange
	},

	Label{LEVEL_CRITICAL, "[CRITICAL]", "[CRIT]",
		[3]uint8{255, 0, 0}, // Red
	},

	Label{level: LEVEL_PRINT, StdLengthName: "      "}, // LEVEL_PRINT requires no label
}

type level int

// Used for string output of the logging object
var levels = [6]string{
	"LEVEL_DEBUG",
	"LEVEL_INFO",
	"LEVEL_WARNING",
	"LEVEL_ERROR",
	"LEVEL_CRITICAL",
	"LEVEL_PRINT",
}

// Returns the string representation of the level
func (l level) String() string { return levels[l] }

// Returns the label for the level
func (l level) Label() string { return Labels[l].String() }

// Returns the standard length label.
func (l level) StdLenLabel() string { return Labels[l].StdLengthName }

// Returns the ansi colorized label for the level
func (l level) AnsiLabel() string { return Labels[l].Colorized() }

// Returns the ansi colorized stdand length label for the level
func (l level) AnsiStdLenLabel() string { return Labels[l].StdLenColorized() }

// Returns the level using string input. lvl must be the name of the level in
// the form of "debug", "DEBUG", "level_debug", or "LEVEL_DEBUG". Returns
// LEVEL_PRINT if the level is not found.
func LevelFromString(lvl string) level {
	// Determine if lvl includes "level"
	lvl = strings.ToLower(lvl)
	if len(lvl) > 4 && lvl[0:5] != "level" {
		lvl = "level_" + lvl
	} else if len(lvl) < 5 {
		lvl = "level_" + lvl
	}
	for num, llvl := range levels {
		if lvl == strings.ToLower(llvl) {
			return level(num)
		}
	}
	return LEVEL_PRINT
}

const (
	// LEVEL_DEBUG level messages should be used for development logging
	// instead of Printf calls. When used in this manner, instead of
	// sprinkling Printf calls everywhere and then having to remove them
	// once the bug is fixed, the developer can simply change to a higher
	// logging level and the debug messages will not be sent to the output
	// stream.
	LEVEL_DEBUG level = iota

	// LEVEL_INFO level messages should be used to convey more informative
	// output than debug that could be used by a user.
	LEVEL_INFO

	// LEVEL_WARNING messages should be used to notify the user that
	// something worked, but the expected value was not the result.
	LEVEL_WARNING

	// LEVEL_ERROR messages should be used when something just did not work
	// at all.
	LEVEL_ERROR

	// LEVEL_CRITICAL messages are used when something is completely broken
	// and unrecoverable. Critical messages are usually followed by
	// os.Exit().
	LEVEL_CRITICAL

	// LEVEL_PRINT shows output for the standard Print functions and above.
	LEVEL_PRINT
)

var (
	defaultDate        = "Mon Jan 02 15:04:05 MST 2006"
	defaultPrefix      = "::"
	defaultPrefixColor = rgbterm.String("::", 0, 255, 135) // Green
	defaultIndentColor = []uint8{0, 135, 175}              // Grayish blue
)

// Flags are used to control the formatting of the logging output.
const (
	// These flags define which text to prefix to each log entry generated
	// by the Logger. Bits or'ed together to control what's printed.
	Ldate = 1 << iota

	// Full file name and line number: /a/b/c/d.go:23
	LlongFileName

	// Base file name and line number: d.go:23. overrides LshortFileName
	LshortFileName

	// Calling function name
	LfunctionName

	// Calling function line number
	LlineNumber

	// Use color escape sequences
	Lcolor

	// Show indentation with dots and bars
	LshowIndent

	// Disable ansi in file output
	LnoFileAnsi

	// Show prefix output
	Lprefix

	// Show ids for functions generating output. Useful for disabling
	// specific output
	Lid

	// Use indentation. Ltree assumes Lindent.
	Lindent

	// Indent output based on stack position of logging function calls
	Ltree

	// Show the label for output
	Llabel

	// initial values for the standard logger
	LstdFlags = Lprefix | Ldate | Lcolor | LnoFileAnsi | Llabel

	// Special debug output flags
	LdebugFlags = Lcolor | LfunctionName | LlineNumber | Llabel

	// Special debug outpt flags with hierarchy
	LdebugTreeFlags = LdebugFlags | Ltree | Lid | Lindent | LshowIndent
)

// A Logger represents an active logging object that generates lines of output
// to an io.Writer. Each logging operation makes a single call to the Writer's
// Write method. A Logger can be used simultaneously from multiple goroutines;
// it guarantees to serialize access to the Writer.
type Logger struct {
	mu               sync.Mutex         // Ensures atomic writes
	buf              []byte             // For marshaling output to write
	dateFormat       string             // time.RubyDate is the default format
	flags            int                // Properties of the output
	level            level              // The default level is warning
	lastId           int                // The last id level encountered
	ids              map[string]int     // A map of encountered function names with corresponding ID
	template         *template.Template // The format order of the output
	prefix           string             // Inserted into every logging output
	streams          []io.Writer        // Destination for output
	indent           int                // Number of indents to use
	tabStop          int                // Number of spaces considered to be a tab stop
	excludeIDs       []int              // Exclude by whatever things
	excludeFuncNames []string
	excludeStrings   []string
}

var (
	// The default logger
	std = New(LEVEL_CRITICAL, os.Stderr)
)

// New creates a new logger object and returns it.
func New(level level, streams ...io.Writer) (obj *Logger) {
	tmpl := template.Must(template.New("default").Funcs(funcMap).Parse(logFmt))
	obj = &Logger{
		ids:        make(map[string]int),
		streams:    streams,
		dateFormat: defaultDate,
		flags:      LstdFlags,
		level:      level,
		template:   tmpl,
		prefix:     defaultPrefixColor,
		tabStop:    4,
	}
	return
}

// Returns the template of the standard logging object.
func Template() *template.Template { return std.template }

// SetTemplate allocates and parses a new output template for the standard
// logging object. error is returned if the template fails to parse. If the
// template cannot be set, then the default template is used. If data field
// name are misnamed in the template, a panic is produced.
func SetTemplate(temp string) error {
	tmpl, err := template.New("default").Funcs(funcMap).Parse(temp)
	if err != nil {
		return err
	}
	std.template = tmpl
	return nil
}

// Returns the date format used by the standard logging object as a string.
func DateFormat() string { return std.dateFormat }

// Set the date format of the standard logging object. See the date package
// documentation for details on using the date format string.
func SetDateFormat(format string) { std.dateFormat = format }

// Returns the usages flags of the standard logging object.
func Flags() int { return std.flags }

// Set the usage flags for the standard logging object.
func SetFlags(flags int) { std.flags = flags }

// Get the logging level of the standard logging object.
func Level() level { return std.level }

// Set the logging level of the standard logging object.
func SetLevel(level level) { std.level = level }

// Get the logging prefix used by the standard logging object. By default it is
// "::".
func Prefix() string { return std.prefix }

// Set the logging prefix of the standard logging object.
func SetPrefix(prefix string) { std.prefix = prefix }

// Streams get the output streams of the standard logger
func Streams() []io.Writer { return std.streams }

// SetStreams set the output streams of the standard logger
func SetStreams(streams ...io.Writer) { std.streams = streams }

// Indent gets the indent level for all output.
func Indent() int { return std.indent }

// SetIndent allows setting the indent level of all output. level can be
// positive or negative.
func SetIndent(level int) *Logger {
	std.indent = level
	return std
}

// TabStop returns the number of spaces per tab for the standard logging
// object.
func TabStop() int { return std.tabStop }

// SetTabStop sets the number of spaces for each indention. A pointer to the
// standard logging object is returned.
func SetTabStop(stops int) *Logger {
	std.tabStop = stops
	return std
}

// ExcludeByHeirarchyID excludes output if the output is in the heirarchy ID
// identified by ids. ExcludeByHeirarchyID is only available if the Lid flag
// is set.
func ExcludeByHeirarchyID(ids ...int) {
	std.excludeIDs = ids
}

// ExcludeByString excludes output if the output text contains matches for
// strings specified by strs.
func ExcludeByString(strs ...string) {
	std.excludeStrings = strs
}

// ExcludeByFuncName excludes output if it comes from functions matching names.
// ExcludeByFuncName is only available if the LshortFileName or LlongFileName
// flags are used.
func ExcludeByFuncName(names ...string) {
	std.excludeFuncNames = names
}

// WithFlags uses flags to write output using the print function passed as f.
func WithFlags(flags int, f func(...interface{}), args ...interface{}) {
	cFlags := std.flags
	std.SetFlags(flags)
	f(args...)
	std.SetFlags(cFlags)
}

// WithFlagsf uses flags to write output using the print function passed as f
// with the format and arguments specified.
func WithFlagsf(flags int, f func(string, ...interface{}),
	format string, args ...interface{}) {
	cFlags := std.flags
	std.SetFlags(flags)
	f(format, args...)
	std.SetFlags(cFlags)
}

// Printf formats according to a format specifier and writes to standard
// logger output stream(s).
func Printf(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_PRINT, 2, fmt.Sprintf(format, v...), nil)
}

// Print sends output to the standard logger object output stream(s) regardless
// of logging level. The output is formatted using the output template and
// flags. Spaces are added between operands when neither is a string.
func Print(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_PRINT, 2, fmt.Sprint(v...), nil)
}

// Println formats using the default formats for its operands and writes to the
// standard logger output stream(s). Spaces are always added between operands and
// a newline is appended.
func Println(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_PRINT, 2, fmt.Sprintln(v...), nil)
}

// Panicf is equivalent to Printf(), but panic() is called once output is
// complete.
func Panicf(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_CRITICAL, 2, fmt.Sprintf(format, v...), nil)
	panic(v)
}

// Panic is equivalent to Print(), but panic() is called once output is
// complete.
func Panic(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_CRITICAL, 2, fmt.Sprint(v...), nil)
	panic(v)
}

// Panicln is equivalent to Println(), but panic() is called once output is
// complete.
func Panicln(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_CRITICAL, 2, fmt.Sprintln(v...), nil)
	panic(v)
}

// Debugf is similar to Printf(), except the colorized LEVEL_DEBUG label is
// prefixed to the output.
func Debugf(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_DEBUG, 2, fmt.Sprintf(format, v...), nil)
}

// Debug is similar to Print(), except the colorized LEVEL_DEBUG label is
// prefixed to the output.
func Debug(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_DEBUG, 2, fmt.Sprint(v...), nil)
}

// Debugln is similar to Println(), except the colorized LEVEL_DEBUG label is
// prefixed to the output.
func Debugln(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_DEBUG, 2, fmt.Sprintln(v...), nil)
}

// Infof is similar to Printf(), except the colorized LEVEL_INFO label is
// prefixed to the output.
func Infof(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_INFO, 2, fmt.Sprintf(format, v...), nil)
}

// Info is similar to Print(), except the colorized LEVEL_INFO label is prefixed
// to the output.
func Info(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_INFO, 2, fmt.Sprint(v...), nil)
}

// Infoln is similar to Println(), except the colorized LEVEL_INFO label is
// prefixed to the output.
func Infoln(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_INFO, 2, fmt.Sprintln(v...), nil)
}

// Warningf is similar to Printf(), except the colorized LEVEL_WARNING label is
// prefixed to the output.
func Warningf(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_WARNING, 2, fmt.Sprintf(format, v...), nil)
}

// Warning is similar to Print(), except the colorized LEVEL_WARNING label is
// prefixed to the output.
func Warning(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_WARNING, 2, fmt.Sprint(v...), nil)
}

// Warningln is similar to Println(), except the colorized LEVEL_WARNING label
// is prefixed to the output.
func Warningln(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_WARNING, 2, fmt.Sprintln(v...), nil)
}

// Errorf is similar to Printf(), except the colorized LEVEL_ERROR label is
// prefixed to the output.
func Errorf(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_ERROR, 2, fmt.Sprintf(format, v...), nil)
}

// Error is similar to Print(), except the colorized LEVEL_ERROR label is
// prefixed to the output.
func Error(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_ERROR, 2, fmt.Sprint(v...), nil)
}

// Errorln is similar to Println(), except the colorized LEVEL_ERROR label is
// prefixed to the output.
func Errorln(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_ERROR, 2, fmt.Sprintln(v...), nil)
}

// Criticalf is similar to Printf(), except the colorized LEVEL_CRITICAL label is
// prefixed to the output.
func Criticalf(format string, v ...interface{}) {
	std.Fprint(std.flags, LEVEL_CRITICAL, 2, fmt.Sprintf(format, v...), nil)
}

// Critical is similar to Prin()t, except the colorized LEVEL_CRITICAL label is
// prefixed to the output.
func Critical(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_CRITICAL, 2, fmt.Sprint(v...), nil)
}

// Criticalln is similar to Println(), except the colorized LEVEL_CRITICAL label
// is prefixed to the output.
func Criticalln(v ...interface{}) {
	std.Fprint(std.flags, LEVEL_CRITICAL, 2, fmt.Sprintln(v...), nil)
}

// Fprint is used by all of the logging functions to send output to the output
// stream.
//
// flags sets the output flags to use when writing the output.
//
// logLevel is the level of the output.
//
// calldepth is the number of stack frames to skip when getting the file
// name of original calling function for file name output.
//
// text is the string to append to the assembled log format output. If the text
// is prefixed with newlines, they will be stripped out and placed in front of
// the completed output (test with template applied) before writing it to the
// stream.
//
// stream will be used as the output stream the text will be written to. If
// stream is nil, the stream value contained in the logger object is used.
//
// Fprint returns the number of bytes written to the stream or an error.
func (l *Logger) Fprint(flags int, logLevel level, calldepth int,
	text string, stream io.Writer) (n int, err error) {

	if (logLevel != LEVEL_PRINT && l.level != LEVEL_PRINT) &&
		logLevel < l.level {
		return
	}

	// Check for string excludes
	if len(l.excludeStrings) > 0 {
		for _, val := range l.excludeStrings {
			if strings.Contains(text, val) {
				return
			}
		}
	}

	now := time.Now()
	var pgmC uintptr
	var file, fName string
	var line int
	var id string
	var indentCount int

	l.mu.Lock()
	defer l.mu.Unlock()

	if flags&(LlongFileName|LshortFileName|LfunctionName|Lid|Ltree) != 0 ||
		len(l.excludeFuncNames) > 0 {
		// release lock while getting caller info - it's expensive.
		l.mu.Unlock()

		pgmC, file, line, _ = runtime.Caller(calldepth)

		if flags&Ltree != 0 {
			pc := make([]uintptr, 32)
			pcNum := runtime.Callers(4, pc)
			for i := 1; i < pcNum; i++ {
				pcFunc := runtime.FuncForPC(pc[i])
				funcName := pcFunc.Name()
				if funcName == "runtime.goexit" {
					continue
				}
				indentCount += 1
			}
		}

		if flags&Lid != 0 {
			fAtPC := runtime.FuncForPC(pgmC)
			fName = fAtPC.Name()
			var idNum int
			if _, ok := l.ids[fName]; ok {
				idNum = l.ids[fName]
			} else {
				l.ids[fName] = l.lastId
				idNum = l.lastId
				l.lastId += 1
			}
			id = fmt.Sprintf("[%02.f]", float64(idNum))
		}

		if flags&LshortFileName != 0 {
			short := file
			for i := len(file) - 1; i > 0; i-- {
				if file[i] == '/' {
					short = file[i+1:]
					break
				}
			}
			file = short
		}

		if flags&LfunctionName != 0 || len(l.excludeFuncNames) > 0 {
			fAtPC := runtime.FuncForPC(pgmC)
			fName = fAtPC.Name()
			for i := len(fName) - 1; i >= 0; i-- {
				if fName[i] == '.' {
					fName = fName[i+1:]
					break
				}
			}
		}

		l.mu.Lock()
	}

	// Check excludes and skip output if matches are found
	var iId int
	if flags&(Lid) != 0 {
		for _, eId := range l.excludeIDs {
			iId, _ = strconv.Atoi(id[1 : len(id)-1])
			if iId == eId {
				return
			}
		}
	}

	// Check func name excludes and return if matches are found
	if len(fName) > 0 {
		for _, name := range l.excludeFuncNames {
			if strings.Contains(fName, name) {
				return
			}
		}
	}

	// Reset the buffer
	l.buf = l.buf[:0]

	trimText := strings.TrimLeft(text, "\n")
	trimedCount := len(text) - len(trimText)
	if trimedCount > 0 {
		l.buf = append(l.buf, trimText...)
	} else {
		l.buf = append(l.buf, text...)
	}

	var date string
	var prefix string

	if flags&Ldate != 0 {
		date = now.Format(l.dateFormat)
	}

	if flags&Lprefix != 0 {
		prefix = l.prefix
	}

	if flags&LlineNumber == 0 {
		line = 0
	}

	if flags&(LshortFileName|LlongFileName) == 0 {
		file = ""
	}

	if flags&LfunctionName == 0 {
		fName = ""
	}

	var indent string
	if indentCount > 0 || flags&Lindent != 0 {
		for i := 0; i < indentCount+l.indent; i++ {
			for j := 0; j < l.tabStop; j++ {
				if flags&LshowIndent != 0 && j == l.tabStop-1 {
					indent += "|"
				} else if flags&LshowIndent != 0 {
					indent += "."
				} else {
					indent += " "
				}
			}
		}
		if len(indent) > 0 && string(indent[0]) != " " {
			indent = rgbterm.String(indent, defaultIndentColor[0],
				defaultIndentColor[1], defaultIndentColor[2])
		}
	}

	var label string
	if flags&Llabel != 0 {
		if flags&Lcolor != 0 {
			label = logLevel.AnsiLabel()
			if flags&Ltree != 0 {
				label = logLevel.AnsiStdLenLabel()
			}
		} else {
			label = logLevel.Label()
			if flags&Ltree != 0 {
				label = logLevel.StdLenLabel()
			}
		}
	}

	f := &format{
		Prefix:       prefix,
		LogLabel:     label,
		Date:         date,
		FileName:     file,
		FunctionName: fName,
		LineNumber:   line,
		Indent:       indent,
		Id:           id,
		Text:         string(l.buf),
	}

	var out bytes.Buffer
	var strippedText, finalText string

	err = l.template.Execute(&out, f)
	if err != nil {
		panic(err)
	}

	if flags&Lcolor == 0 {
		strippedText = stripAnsi(out.String())
	}

	if trimedCount > 0 && flags&Lcolor == 0 {
		finalText = strings.Repeat("\n", trimedCount) + strippedText
	} else if trimedCount > 0 && flags&Lcolor != 0 {
		finalText = strings.Repeat("\n", trimedCount) + out.String()
	} else if flags&Lcolor == 0 {
		finalText = strippedText
	} else {
		finalText = out.String()
	}

	if stream == nil {
		n, err = l.Write([]byte(finalText))
	} else {
		n, err = stream.Write([]byte(finalText))
	}

	return
}

// Returns the template of the standard logging object.
func (l *Logger) Template() *template.Template { return l.template }

// SetTemplate allocates and parses a new output template for the logging
// object. error is returned if the template fails to parse. If the template
// cannot be set, then the default template is used. If data field name are
// misnamed in the template, a panic is produced.
func (l *Logger) SetTemplate(temp string) error {
	tmpl, err := template.New("default").Funcs(funcMap).Parse(temp)
	if err != nil {
		return err
	}
	l.template = tmpl
	return nil
}

// Returns the date format used by the logging object as a string.
func (l *Logger) DateFormat() string { return l.dateFormat }

// Set the date format of the logging object. See the date package
// documentation for details on using the date format string.
func (l *Logger) SetDateFormat(format string) { l.dateFormat = format }

// Returns the usages flags of the logging object.
func (l *Logger) Flags() int { return l.flags }

// Set the usage flags for the logging object.
func (l *Logger) SetFlags(flags int) { l.flags = flags }

// Get the logging level of the logging object.
func (l *Logger) Level() level { return l.level }

// Set the logging level of the logging object.
func (l *Logger) SetLevel(level level) { l.level = level }

// Get the logging prefix used by the logging object. By default it is "::".
func (l *Logger) Prefix() string { return l.prefix }

// Set the logging prefix of the logging object.
func (l *Logger) SetPrefix(prefix string) { l.prefix = prefix }

// Get the output streams of the logger
func (l *Logger) Streams() []io.Writer { return l.streams }

// Set the output streams of the logger
func (l *Logger) SetStreams(streams ...io.Writer) { l.streams = streams }

// Indent gets the indent level for all output of the logging object.
func (l *Logger) Indent() int { return l.indent }

// SetIndent sets the indent level of all output in the logging object. level
// can be positive or negative. A pointer to the logging object is returned.
func (l *Logger) SetIndent(level int) *Logger {
	l.indent = level
	return l
}

// TabStop returns the number of spaces per tab for the logging object.
func (l *Logger) TabStop() int { return l.tabStop }

// SetTabStop sets the number of spaces for each indention. A pointer to the
// logging object is returned.
func (l *Logger) SetTabStop(stops int) *Logger {
	l.tabStop = stops
	return l
}

// ExcludeByHeirarchyID excludes output if the output is in the heirarchy ID
// identified by ids. ExcludeByHeirarchyID is only available if the Lid flag
// is set.
func (l *Logger) ExcludeByHeirarchyID(ids ...int) {
	l.excludeIDs = ids
}

// ExcludeByString excludes output if the output text contains matches for
// strings specified by strs.
func (l *Logger) ExcludeByString(strs ...string) {
	l.excludeStrings = strs
}

// ExcludeByFuncName excludes output if it comes from functions matching names.
// ExcludeByFuncName is only available if the LshortFileName or LlongFileName
// flags are used.
func (l *Logger) ExcludeByFuncName(names ...string) {
	l.excludeFuncNames = names
}

// WithFlags uses flags to write output using the print function passed as f.
func (l *Logger) WithFlags(flags int, f func(...interface{}), args ...interface{}) {
	cFlags := l.flags
	l.SetFlags(flags)
	f(args...)
	l.SetFlags(cFlags)
}

// WithFlagsf uses flags to write output using the print function passed as f
// with the format and arguments specified.
func (l *Logger) WithFlagsf(flags int, f func(string, ...interface{}),
	format string, args ...interface{}) {
	cFlags := l.flags
	l.SetFlags(flags)
	f(format, args...)
	l.SetFlags(cFlags)
}

// Write writes the array of bytes (p) to all of the logger.Streams. If the
// Lcolor flag is set, ansi escape codes are used to add coloring to the output.
func (l *Logger) Write(p []byte) (wLen int, err error) {
	var write = func(w io.Writer, isStdFile bool) {
		x := p
		if !isStdFile && l.flags&LnoFileAnsi != 0 {
			// TODO: If Lcolor is used, then no coloring should
			// have to be stripped. Inefficient.
			x = stripAnsiByte(x)
		}
		wLen, err = w.Write(x)
		if wLen != len(p) {
			err = io.ErrShortWrite
		}
	}
	for _, w := range l.streams {
		wIface := reflect.ValueOf(w).Interface()
		switch wType := wIface.(type) {
		case *os.File:
			if wType == os.Stdout || wType == os.Stderr {
				write(w, true)
				continue
			}
			write(w, false)
		default:
			write(w, false)
		}
	}
	return
}

// Printf is equivalent to log.Printf().
func (l *Logger) Printf(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_PRINT, 2, fmt.Sprintf(format, v...), nil)
}

// Print is equivalent to log.Print().
func (l *Logger) Print(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_PRINT, 2, fmt.Sprint(v...), nil)
}

// Println is equivalent to log.Println().
func (l *Logger) Println(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_PRINT, 2, fmt.Sprintln(v...), nil)
}

// Panicf is equivalent to log.Panicf().
func (l *Logger) Panicf(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_CRITICAL, 2, fmt.Sprintf(format, v...), nil)
	panic(v)
}

// Panic is equivalent to log.Panic().
func (l *Logger) Panic(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_CRITICAL, 2, fmt.Sprint(v...), nil)
	panic(v)
}

// Panicln is equivalent to log.Panicln().
func (l *Logger) Panicln(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_CRITICAL, 2, fmt.Sprintln(v...), nil)
	panic(v)
}

// Debugf is equivalent to log.Debugf().
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_DEBUG, 2, fmt.Sprintf(format, v...), nil)
}

// Debug is equivalent to log.Debug().
func (l *Logger) Debug(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_DEBUG, 2, fmt.Sprint(v...), nil)
}

// Debugln is equivalent to log.Debugln().
func (l *Logger) Debugln(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_DEBUG, 2, fmt.Sprintln(v...), nil)
}

// Infof is equivalent to log.Infof().
func (l *Logger) Infof(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_INFO, 2, fmt.Sprintf(format, v...), nil)
}

// Info is equivalent to log.Info().
func (l *Logger) Info(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_INFO, 2, fmt.Sprint(v...), nil)
}

// Infoln is equivalent to log.Infoln().
func (l *Logger) Infoln(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_INFO, 2, fmt.Sprintln(v...), nil)
}

// Warningf is equivalent to log.Warningf().
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_WARNING, 2, fmt.Sprintf(format, v...), nil)
}

// Warning is equivalent to log.Warning().
func (l *Logger) Warning(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_WARNING, 2, fmt.Sprint(v...), nil)
}

// Warningln is equivalent to log.Warningln().
func (l *Logger) Warningln(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_WARNING, 2, fmt.Sprintln(v...), nil)
}

// Errorf is equivalent to log.Errorf().
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_ERROR, 2, fmt.Sprintf(format, v...), nil)
}

// Error is equivalent to log.Error().
func (l *Logger) Error(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_ERROR, 2, fmt.Sprint(v...), nil)
}

// Errorln is equivalent to log.Errorln().
func (l *Logger) Errorln(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_ERROR, 2, fmt.Sprintln(v...), nil)
}

// Criticalf is equivalent to log.Criticalf().
func (l *Logger) Criticalf(format string, v ...interface{}) {
	l.Fprint(l.flags, LEVEL_CRITICAL, 2, fmt.Sprintf(format, v...), nil)
}

// Critical is equivalent to log.Critical().
func (l *Logger) Critical(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_CRITICAL, 2, fmt.Sprint(v...), nil)
}

// Criticalln is equivalent to log.Criticalln().
func (l *Logger) Criticalln(v ...interface{}) {
	l.Fprint(l.flags, LEVEL_CRITICAL, 2, fmt.Sprintln(v...), nil)
}
