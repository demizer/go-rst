package testutil

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"testing"

	"github.com/demizer/go-rst/pkg/log"
	klog "github.com/go-kit/kit/log"
	jd "github.com/josephburnett/jd/lib"
)

func init() { SetDebug() }

var (
	StdLogger    = klog.NewNopLogger()
	debug        bool
	LogExcludes  log.LoggerExcludes
	LoggerConfig log.Config
)

// SetDebug is typically called from the init() function in a test file.  SetDebug parses debug flags passed to the test
// binary and also sets the template for logging output.
func SetDebug() {
	flag.Var(&LogExcludes, "exclude", "Exclude context from output.")
	flag.BoolVar(&debug, "debug", false, "Enable debug output.")
	flag.Parse()
	if debug {
		StdLogger = klog.NewLogfmtLogger(os.Stdout)
	}
	LoggerConfig = log.Config{
		Name:      "test",
		Logger:    StdLogger,
		Caller:    true,
		CallDepth: 4,
		Excludes:  LogExcludes,
	}
}

func LogRun(name string) {
	if testing.Verbose() {
		fmt.Printf("+++ RUN   %s\n", name)
	}
}

func LogPass(name string) {
	if testing.Verbose() {
		fmt.Printf("+++ PASS  %s\n", name)
	}
}

func Log(vals ...interface{}) {
	if testing.Verbose() {
		fmt.Println(vals...)
	}
}

func LogDebug(vals ...interface{}) {
	if debug {
		fmt.Println(vals...)
	}
}

func Logf(s string, vals ...interface{}) {
	if testing.Verbose() {
		fmt.Printf(s, vals...)
	}
}

func JsonDiff(expectedData string, parsedData string) (string, error) {
	a, _ := jd.ReadJsonString(expectedData)
	b, _ := jd.ReadJsonString(parsedData)
	return a.Diff(b).Render(), nil
}

// Contains a single test with data loaded from test files in the testdata directory
type Test struct {
	Path            string // The path including directory and basename
	Data            string // The input data to be parsed
	ExpectItemData  string // The expected lex items output in json
	ExpectParseData string // The expected parse nodes in json
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
