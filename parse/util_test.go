// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.

package parse

import (
	"github.com/demizer/go-elog"
	"bufio"
	"os"
	"bytes"
	"strings"
	"flag"
	"fmt"
)

func init() { SetDebug() }

// LexTest is the structure that contains parsed test data from the *.dat files in the testdata
// directory.
type LexTest struct {
	name           string
	description    string
	data           string  // The input data to be parsed
	items          string  // The expected lex items output in json
	expectTree         string  // The expected parsed output in json
}

type LexTests []LexTest

// Search l by name for a specific test.
func (l LexTests) SearchByName(name string) *LexTest {
	for _, test := range l {
		if test.name == name {
			return &test
		}
	}
	return nil
}

// ParseTestData parses testdata contained it dat files in the testdata directory. The testdata was
// contained to these files because it became to large to be included legibly inside the *_test.go
// files. ParseTestData is a simple parser for the testdata files and stores the result of the parse
// into the first return variable.
func ParseTestData(filepath string) ([]LexTest, error) {
	testData, err := os.Open(filepath)
	defer testData.Close()
	if err != nil {
		return nil, err
	}

	var LexTests []LexTest
	var curTest = new(LexTest)
	var buffer bytes.Buffer

	scanner := bufio.NewScanner(testData)

	for scanner.Scan() {
		switch scanner.Text() {
		case "#name":
			// buffer = bytes.NewBuffer(buffer.Bytes())
			// name starts a new section
			if buffer.Len() > 0 {
				// Apend the last section to the array and
				// reset
				curTest.expectTree = buffer.String()
				LexTests = append(LexTests, *curTest)
			}
			curTest = new(LexTest)
			buffer.Reset()
		case "#description":
			curTest.name = strings.TrimRight(buffer.String(), "\n")
			buffer.Reset()
		case "#data":
			curTest.description = strings.TrimRight(buffer.String(), "\n")
			buffer.Reset()
		case "#items":
			curTest.data = strings.TrimRight(buffer.String(), "\n")
			buffer.Reset()
		case "#parse-tree":
			curTest.items = buffer.String()
			buffer.Reset()
		default:
			// Collect the text in between sections
			buffer.WriteString(fmt.Sprintln(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	if buffer.Len() > 0 {
		// Apend the last section to the array and
		curTest.expectTree = buffer.String()
		LexTests = append(LexTests, *curTest)
	}

	return LexTests, nil
}

// SetDebug is typically called from the init() function in a test file. SetDebug parses debug flags
// passed to the test binary and also sets the template for logging output.
func SetDebug() {
	var debug bool

	flag.BoolVar(&debug, "debug", false, "Enable debug output.")
	flag.Parse()

	if debug {
		log.SetLevel(log.LEVEL_DEBUG)
	}

	log.SetTemplate("{{if .Date}}{{.Date}} {{end}}" +
		"{{if .Prefix}}{{.Prefix}} {{end}}" +
		"{{if .LogLabel}}{{.LogLabel}} {{end}}" +
		"{{if .FileName}}{{.FileName}}: {{end}}" +
		"{{if .FunctionName}}{{.FunctionName}}{{end}}" +
		"{{if .LineNumber}}#{{.LineNumber}}: {{end}}" +
		"{{if .Text}}{{.Text}}{{end}}")

	log.SetFlags(log.Lansi | log.LnoPrefix | log.LfunctionName |
		log.LlineNumber)
}

