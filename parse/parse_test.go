// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package parse

// import (
	// "bufio"
	// "encoding/json"
	// "fmt"
	// "os"
	// "strings"
	// "testing"
// )

// func parseTestData(filepath string) ([]lexTest, error) {
	// testData, err := os.Open(filepath)
	// defer testData.Close()
	// if err != nil {
		// return nil, err
	// }
	// var lexTests []lexTest
	// var curTest lexTest
	// var buffer []byte
	// scanner := bufio.NewScanner(testData)
	// for scanner.Scan() {
		// switch scanner.Text() {
		// case "#description":
			// if len(buffer) > 0 {
				// curTest.expectJson = buffer
				// lexTests = append(lexTests, curTest)
				// curTest = lexTest{}
			// } else {
				// curTest = lexTest{}
			// }
			// buffer = nil
		// case "#data":
			// curTest.description = buffer
			// buffer = nil
		// case "#tree":
			// curTest.input = buffer
			// buffer = nil
		// default:
			// if len(scanner.Text()) == 0 ||
				// strings.TrimLeft(scanner.Text(), " ")[0] == '#' {
				// continue
			// }
			// buffer = append(buffer, scanner.Bytes()...)
		// }
	// }
	// return lexTests, nil
// }

// func TestSection(t *testing.T) {
	// lexTests, err := parseTestData("../testdata/test_section_headers.dat")
	// if err != nil {
		// t.FailNow()
	// }
	// var i interface{}
	// err = json.Unmarshal(lexTests[0].expectJson, &i)
	// if err != nil {
		// t.Errorf("JSON Error: %s, IN: %s", err, lexTests[0])
	// }
	// fmt.Printf("%#v\n", i)
// }
