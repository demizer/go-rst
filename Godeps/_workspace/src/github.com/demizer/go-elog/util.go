// Copyright 2013 The go-elog Authors. All rights reserved.
// This code is MIT licensed. See the LICENSE file for more info.

package log

import (
	"regexp"
)

// stripAnsi removes all ansi escapes from a string.
func stripAnsi(text string) string {
	reg := regexp.MustCompile("\x1b\\[[\\d;]+m")
	return reg.ReplaceAllString(text, "")
}

// stripAnsiByte removes all ansi escapes from a string and returns the clean
// string.
func stripAnsiByte(text []byte) []byte {
	reg := regexp.MustCompile("\x1b\\[[\\d;]+m")
	return reg.ReplaceAll(text, []byte(""))
}
