// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/demizer/go-elog"
	"github.com/demizer/go-spew/spew"
	"os"
	"strings"
	"testing"
)

type lexTest struct {
	name           string
	description    string
	data           string
	items          string
	expect         string
	collectedItems []item
}

var (
	tEOF = item{ElementType: itemEOF, Position: 0, Value: ""}
)

var spd = spew.ConfigState{Indent: "\t"}

func TestAll(t *testing.T) {
	log.SetLevel(log.LEVEL_DEBUG)
	err := log.SetTemplate("{{if .Date}}{{.Date}} {{end}}" +
		"{{if .Prefix}}{{.Prefix}} {{end}}" +
		"{{if .LogLabel}}{{.LogLabel}} {{end}}" +
		"{{if .FileName}}{{.FileName}}: {{end}}" +
		"{{if .FunctionName}}{{.FunctionName}}{{end}}" +
		"{{if .LineNumber}}#{{.LineNumber}}: {{end}}" +
		"{{if .Text}}{{.Text}}{{end}}")
	if err != nil {
		t.Error(err)
	}
	log.SetFlags(log.Lansi | log.LnoPrefix | log.LfunctionName |
		log.LlineNumber)
}

func parseTestData(t *testing.T, filepath string) ([]lexTest, error) {
	testData, err := os.Open(filepath)
	defer testData.Close()
	if err != nil {
		return nil, err
	}

	var lexTests []lexTest
	var curTest = new(lexTest)
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
				lexTests = append(lexTests, *curTest)
			}
			curTest = new(lexTest)
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
		case "#expect":
			curTest.items = buffer.String()
			buffer.Reset()
		default:
			// Collect the text in between sections
			if len(scanner.Text()) == 0 ||
				strings.TrimLeft(scanner.Text(), " ")[0] == '#' {
				continue
			}
			buffer.WriteString(fmt.Sprintln(scanner.Text()))
		}
	}

	if err := scanner.Err(); err != nil {
		t.Error(err)
	}

	if buffer.Len() > 0 {
		// Apend the last section to the array and
		curTest.expect = buffer.String()
		lexTests = append(lexTests, *curTest)
	}

	return lexTests, nil
}

// collect gathers the emitted items into a slice.
func collect(t *lexTest) (items []item) {
	l := lex(t.name, t.data)
	for {
		item := l.nextItem()
		items = append(items, item)
		if item.ElementType == itemEOF || item.ElementType == itemError {
			break
		}
	}
	return
}

func TestSection(t *testing.T) {
	lexTests, err := parseTestData(t, "../testdata/test_lex_sections.dat")
	if err != nil {
		t.FailNow()
	}
	for _, test := range lexTests {
		if test.name == "ST-UNEXP-TITLES" {
			log.Printf("Test Name: \t%s\n", test.name)
			log.Printf("Description: \t%s\n", test.description)
			items := collect(&test)
			b, err := json.MarshalIndent(items, "", "\t")
			if err != nil {
				t.Errorf("JSON Error: %s, IN: %+v\n", err, test)
			}
			log.Println("Collected items:\n\n", spd.Sdump(items))
			log.Println("items JSON object:\n\n", string(b))
			var i interface{}
			err = json.Unmarshal([]byte(test.expect), &i)
			if err != nil {
				t.Errorf("JSON Error: %s, IN: %+v\n", err, test)
			}
			log.Println("JSON object:\n\n", spd.Sdump(i))
		}
	}
}
