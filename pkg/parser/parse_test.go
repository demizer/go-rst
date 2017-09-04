package parser

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"

	doc "github.com/demizer/go-rst/pkg/document"
)

func nodeListToInterface(v *doc.NodeList) []interface{} {
	v2 := []doc.Node(*v)
	s := make([]interface{}, len(v2))
	for i, j := range v2 {
		s[i] = j
	}
	return s
}

// checkParseNodes compares the expected parser output (*_nodes.json) against the actual parser output using the jd library.
func checkParseNodes(t *testing.T, eParser []interface{}, pNodes *doc.NodeList, testPath string) {
	pJson, err := json.MarshalIndent(pNodes, "", "    ")
	if err != nil {
		t.Errorf("Error Marshalling JSON: %s", err.Error())
		return
	}

	// Json diff output has a syntax:
	// https://github.com/josephburnett/jd#diff-language
	o, err := testutil.JsonDiff(eParser, nodeListToInterface(pNodes))
	if err != nil {
		fmt.Println(o)
		fmt.Printf("Error diffing JSON: %s", err.Error())
		return
	}

	// There should be no output from the diff
	if len(o) != 0 {
		// Give all other output time to print
		// time.Sleep(time.Second / 2)

		testutil.Log("\nFAIL: parsed nodes do not match expected nodes!")

		testutil.Log("\n[Parsed Nodes JSON]\n\n")
		testutil.Log(string(pJson))

		testutil.Log("\n\n[JSON DIFF]\n\n")
		testutil.Log(o)

		t.FailNow()
	}
}

func LoadParserTest(t *testing.T, path string) (test *testutil.Test) {
	iDPath := path + ".rst"
	inputData, err := ioutil.ReadFile(iDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(inputData) == 0 {
		t.Fatalf("\"%s\" is empty!", iDPath)
	}
	nDPath := path + "-nodes.json"
	nodeData, err := ioutil.ReadFile(nDPath)
	if err != nil {
		t.Fatal(err)
	}
	if len(nodeData) == 0 {
		t.Fatalf("\"%s\" is empty!", nDPath)
	}
	return &testutil.Test{
		Path:     path,
		Data:     string(inputData[:len(inputData)-1]),
		NodeData: string(nodeData),
	}
}

// parseTest initiates the parser and parses a test using test.data is input.
func parseTest(t *testing.T, test *testutil.Test) *Parser {
	testutil.Log(fmt.Sprintf("Test path: %s", test.Path))
	testutil.Log(fmt.Sprintf("Test Input:\n\n%s\n", test.Data))
	p, _ := Parse(test.Path, test.Data, testutil.StdLogger)
	return p
}
