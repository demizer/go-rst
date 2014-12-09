// Copyright 2013 The go-elog Authors. All rights reserved.
// This code is MIT licensed. See the LICENSE file for more info.

package log

import "text/template"

// funcMap contains the available functions to the log format template.
var (
	funcMap = template.FuncMap{}
	logFmt  = "{{if .Date}}{{.Date}} {{end}}" +
		"{{if .Prefix}}{{.Prefix}} {{end}}" +
		"{{if .LogLabel}}{{.LogLabel}} {{end}}" +
		"{{if .Id}}{{.Id}} {{end}}" +
		"{{if .Indent}}{{.Indent}}{{end}}" +
		"{{if .FileName}}{{.FileName}}: {{end}}" +
		"{{if .FunctionName}}{{.FunctionName}}: {{end}}" +
		"{{if .LineNumber}}Line {{.LineNumber}}: {{end}}" +
		"{{if .Text}}{{.Text}}{{end}}"
)

// format is the possible values that can be used in a log output format
type format struct {
	Prefix       string
	LogLabel     string
	Date         string
	FileName     string
	FunctionName string
	LineNumber   int
	Indent       string
	Id           string
	Text         string
}
