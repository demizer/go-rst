package main

import (
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

var testFiles []testFile

type testFile struct {
	name        string
	id          string
	path        string
	baseName    string
	bad         bool
	lexerTest   bool
	parserTest  bool
	implemented bool
}

type testPackage struct {
	Name       string
	OutputPath string
	Tests      []test
}

type test struct {
	Package      string
	ID           string
	Type         string
	Name         string
	FileBaseName string
	Outcome      string
	Implemented  bool
}

var testTemplate = `package {{.Name}}

//
// AUTO-GENERATED using tools/gentests.go
//

import (
	"os"
	"testing"

	"github.com/demizer/go-rst/pkg/testutil"
)

{{ range .Tests -}}
func Test_{{.ID}}_{{.Type}}{{.Name}}{{.Outcome}}(t *testing.T) {
	{{- if not .Implemented}}
	if os.Getenv("GO_RST_SKIP_NOT_IMPLEMENTED") == "1" {
		t.SkipNow()
	}
	{{- end }}
	testPath := testutil.TestPathFromName("{{.FileBaseName}}")
	{{- if eq .Type "Lexer"}}
	test := LoadLexTest(t, testPath)
	items := lexTest(t, test)
	equal(t, test.ExpectItems(), items)
	{{- else}}
	test := LoadParserTest(t, testPath)
	pTree := parseTest(t, test)
	checkParseNodes(t, test.ExpectData, pTree, testPath)
	{{- end}}
}

{{ end -}}
`

func capitalize(s string) (o string) {
	n := strings.Split(s, "-")
	for _, v := range n {
		o = o + strings.Title(v)
	}
	return
}

func walkFunc(path string, info os.FileInfo, err error) error {
	if info.IsDir() {
		return err
	}
	fp := strings.Split(path, string(os.PathSeparator))
	re := regexp.MustCompile(`(([\d\.]{4,})-(bad-)?(.*))(-items|-nodes)(-xx)?\.json`)
	js := re.FindStringSubmatch(path)
	if js != nil {
		// Example object:
		// []string{"11.00.00.00-three-short-options-nodes.json", "11.00.00.00-three-short-options", "11.00.00.00", "", "three-short-options", "-nodes", "", "json"}
		re2 := regexp.MustCompile(`[\d]*-test-(.*)`)
		m := testFile{name: re2.FindStringSubmatch(fp[1])[1], baseName: js[1], id: js[2], path: path}
		if js[3] != "" {
			m.bad = true
		}
		if js[5] == "-items" {
			m.lexerTest = true
		} else if js[5] == "-nodes" {
			m.parserTest = true
		}
		if js[6] == "" {
			m.implemented = true
		}
		// fmt.Printf("%#v\n", m)
		testFiles = append(testFiles, m)
	}
	return err
}

func main() {
	testFiles = make([]testFile, 0, 100)
	err := filepath.Walk("testdata", walkFunc)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	tokenTestSuite := testPackage{
		Name:       "token",
		Tests:      make([]test, 0, 100),
		OutputPath: "pkg/token/rst_test.go",
	}
	parserTestSuite := testPackage{
		Name:       "parser",
		Tests:      make([]test, 0, 100),
		OutputPath: "pkg/parser/rst_test.go",
	}

	for _, v := range testFiles {
		t := "Lexer"
		p := "parser"
		if v.parserTest {
			t = "Parser"
			p = "token"
		}

		o := "Good"
		if v.bad {
			o = "Bad"
		}

		id := strings.Replace(v.id, ".", "_", -1)
		m := test{
			Package:      p,
			ID:           id,
			Type:         t,
			Name:         capitalize(v.name),
			FileBaseName: v.baseName,
			Outcome:      o,
			Implemented:  v.implemented,
		}

		if v.lexerTest {
			tokenTestSuite.Tests = append(tokenTestSuite.Tests, m)
		} else {
			parserTestSuite.Tests = append(parserTestSuite.Tests, m)
		}
	}

	tmp := template.Must(template.New("tests").Parse(testTemplate))
	var testObj []testPackage
	testObj = append(testObj, tokenTestSuite)
	testObj = append(testObj, parserTestSuite)

	for _, v := range testObj {
		f, err := os.Create(v.OutputPath)
		if err != nil {
			fmt.Println("\n\nERROR Create file: ", err)
			return
		}
		err = tmp.Execute(f, v)
		if err != nil {
			fmt.Printf("\n\nERROR executing template: %s", err)
			os.Exit(1)
		}
		f.Close()
	}

}
