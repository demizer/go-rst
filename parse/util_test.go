package parse

import (
	"github.com/demizer/go-elog"
	"testing"
	"bufio"
	"os"
	"bytes"
	"strings"
	"flag"
	"fmt"
)

type LexTest struct {
	name           string
	description    string
	data           string
	items          string
	expect         string
	collectedItems []item
}

type LexTests []LexTest

func (l LexTests) SearchByName(name string) *LexTest {
	for _, test := range l {
		if test.name == name {
			return &test
		}
	}
	return nil
}

var Tests LexTests

func ParseTestData(t *testing.T, filepath string) ([]LexTest, error) {
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
				curTest.expect = buffer.String()
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
		case "#parse-expect":
			curTest.items = buffer.String()
			buffer.Reset()
		default:
			// Collect the text in between sections
			buffer.WriteString(fmt.Sprintln(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	if buffer.Len() > 0 {
		// Apend the last section to the array and
		curTest.expect = buffer.String()
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

