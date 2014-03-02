// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

import (
	"bufio"
	"bytes"
	"fmt"
	"encoding/json"
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

type lexTests []lexTest

func (l lexTests) SearchByName(name string) *lexTest {
	for _, test := range l {
		if test.name == name {
			return &test
		}
	}
	return nil
}

var tests lexTests

var (
	tEOF = item{ElementType: itemEOF, Position: 0, Value: ""}
)

var spd = spew.ConfigState{Indent: "\t"}

func init() {
	log.SetLevel(log.LEVEL_DEBUG)
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

func lexSectionTest(t *testing.T, testName string) []item {
	var err error
	if tests == nil {
		tests, err = parseTestData(t, "../testdata/test_lex_sections.dat")
		if err != nil {
			t.FailNow()
		}
	}
	test := tests.SearchByName(testName)
	if test != nil {
		log.Debugf("Test Name: \t%s\n", test.name)
		log.Debugf("Description: \t%s\n", test.description)
		log.Debugf("Test Input:\n-----------\n%s\n----------\n", test.data)
		items := collect(test)
		return items
	}
	return nil
}

func TestSectionTitlePara(t *testing.T) {
	items := runTest(t, "SectionTitlePara")
	log.Debugln("Collected items:\n\n", spd.Sdump(items))
}

func TestSectionTitleParaNoBlankline(t *testing.T) {
	items := runTest(t, "SectionTitleParaNoBlankline")
	log.Debugln("Collected items:\n\n", spd.Sdump(items))
}

func TestSectionParaHeadPara(t *testing.T) {
	items := runTest(t, "SectionParaHeadPara")
	log.Debugln("Collected items:\n\n", spd.Sdump(items))
}

func TestSectionUnexpectedTitles(t *testing.T) {
	items := runTest(t, "SectionUnexpectedTitles")
	log.Debugln("Collected items:\n\n", spd.Sdump(items))
}
// Unmarshals input into []items, the json input from test data does not include ElementType, so
// this is filled in manually. Returns error if there is a json parsing error.
func jsonToItems(input []byte) ([]item, error) {
	var exp []item
	err := json.Unmarshal(input, &exp)
	if err != nil {
		return nil, err
	}
	// Set the correct ElementType (int), this is not included in the json from the test data.
	for i, item := range exp {
		for j, elm := range elements {
			if item.ElementName == elm {
				exp[i].ElementType = itemElement(j)
			}
		}
	}
	return exp, nil
}
