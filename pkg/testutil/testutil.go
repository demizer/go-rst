package testutil

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/demizer/go-rst/pkg/logging"
	kit "github.com/go-kit/kit/log"
	jd "github.com/josephburnett/jd/lib"
)

var (
	debug               bool
	excludeNamedContext string // Exclude a log context from being shown in the output
)

func init() { SetDebug() }

// SetDebug is typically called from the init() function in a test file.  SetDebug parses debug flags passed to the test
// binary and also sets the template for logging output.
func SetDebug() {
	flag.StringVar(&excludeNamedContext, "exclude", "test", "Exclude context from output.")
	flag.BoolVar(&debug, "debug", false, "Enable debug output.")
	flag.Parse()
	if debug {
		logging.SetStdLogger(kit.NewLogfmtLogger(os.Stdout))
	}
}

func Log(vals ...interface{}) {
	if debug {
		fmt.Println(vals...)
	}
}

func JsonDiff(expectedItems, parsedItems []interface{}) (string, error) {
	eJson, err := json.Marshal(expectedItems)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal expectedItems: %s", err.Error())
	}

	pJson, err := json.Marshal(parsedItems)
	if err != nil {
		return "", fmt.Errorf("Failed to marshal parsedItems: %s", err.Error())
	}

	a, _ := jd.ReadJsonString(string(eJson))
	b, _ := jd.ReadJsonString(string(pJson))

	return a.Diff(b).Render(), nil
}

// Contains a single test with data loaded from test files in the testdata directory
type Test struct {
	Path     string // The path including directory and basename
	Data     string // The input data to be parsed
	ItemData string // The expected lex items output in json
	NodeData string // The expected parse nodes in json
}

// ExpectNodes returns the expected parse_tree values from the tests as unmarshaled JSON. A panic occurs if there is an error
// unmarshaling the JSON data.
func (l Test) ExpectNodes() (nl []interface{}) {
	if err := json.Unmarshal([]byte(l.NodeData), &nl); err != nil {
		panic(fmt.Sprintln("JSON error: ", err))
	}
	return
}

// ExpectItems unmarshals the expected lex_items into a silce of items. A panic occurs if there is an error decoding the JSON
// data.
func (l Test) ExpectItems() (lexItems []interface{}) {
	if err := json.Unmarshal([]byte(l.ItemData), &lexItems); err != nil {
		panic(fmt.Sprintln("JSON error: ", err))
	}
	return
}

// Contains absolute file paths for the test data
var TestDataFiles []string

// testPathsFromDirectory walks through the file tree in the testdata directory containing all of the tests and returns a
// string slice of all the discovered paths.
func TestPathsFromDirectory(dir string) ([]string, error) {
	var paths []string
	wFunc := func(p string, info os.FileInfo, err error) error {
		path, err := filepath.Abs(p)
		if err != nil {
			return err
		}
		if filepath.Ext(path) == ".rst" {
			paths = append(paths, path[:len(path)-4])
		}
		return nil
	}
	err := filepath.Walk(dir, wFunc)
	if err != nil {
		return nil, err
	}
	return paths, nil
}

// TestPathFromName loops through TestDataFiles until name is matched.
func TestPathFromName(name string) string {
	var err error
	if len(TestDataFiles) < 1 {
		TestDataFiles, err = TestPathsFromDirectory("../../testdata")
		if err != nil {
			panic(err)
		}
	}
	for _, p := range TestDataFiles {
		if len(p)-len(name) > 0 {
			if p[len(p)-len(name):] == name {
				return p
			}
		}
	}
	panic(fmt.Sprintf("Could not find test for %q\n", name))
}
