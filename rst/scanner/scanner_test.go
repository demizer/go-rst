package scanner

import (
	"bytes"
	"fmt"
	"testing"

	"strings"

	"github.com/demizer/rst/rst/token"
)

type tokenPair struct {
	tok  token.Type
	text string
}

var tokenLists = map[string][]tokenPair{
	"comment": []tokenPair{
		{token.COMMENT, ".."},
	},
	"string": []tokenPair{
		{token.STRING, `" "`},
		{token.STRING, `"a"`},
		{token.STRING, `"本"`},
		{token.STRING, `"${file("foo")}"`},
		{token.STRING, `"${file(\"foo\")}"`},
		{token.STRING, `"\a"`},
		{token.STRING, `"\b"`},
		{token.STRING, `"\f"`},
		{token.STRING, `"\n"`},
		{token.STRING, `"\r"`},
		{token.STRING, `"\t"`},
		{token.STRING, `"\v"`},
		{token.STRING, `"\""`},
		{token.STRING, `"\000"`},
		{token.STRING, `"\777"`},
		{token.STRING, `"\x00"`},
		{token.STRING, `"\xff"`},
		{token.STRING, `"\u0000"`},
		{token.STRING, `"\ufA16"`},
		{token.STRING, `"\U00000000"`},
		{token.STRING, `"\U0000ffAB"`},
	},
}

var orderedTokenLists = []string{
	"comment",
	"operator",
	"bool",
	"ident",
	"heredoc",
	"string",
	"number",
	"float",
}

// func TestPosition(t *testing.T) {
// // create artifical source code
// buf := new(bytes.Buffer)

// for _, listName := range orderedTokenLists {
// for _, ident := range tokenLists[listName] {
// fmt.Fprintf(buf, "\t\t\t\t%s\n", ident.text)
// }
// }

// s := New(buf.Bytes())

// pos := token.Pos{"", 4, 1, 5}
// s.Scan()
// for _, listName := range orderedTokenLists {

// for _, k := range tokenLists[listName] {
// curPos := s.tokPos
// // fmt.Printf("[%q] s = %+v:%+v\n", k.text, curPos.Offset, curPos.Column)

// if curPos.Offset != pos.Offset {
// t.Fatalf("offset = %d, want %d for %q", curPos.Offset, pos.Offset, k.text)
// }
// if curPos.Line != pos.Line {
// t.Fatalf("line = %d, want %d for %q", curPos.Line, pos.Line, k.text)
// }
// if curPos.Column != pos.Column {
// t.Fatalf("column = %d, want %d for %q", curPos.Column, pos.Column, k.text)
// }
// pos.Offset += 4 + len(k.text) + 1     // 4 tabs + token bytes + newline
// pos.Line += countNewlines(k.text) + 1 // each token is on a new line
// s.Scan()
// }
// }
// // make sure there were no token-internal errors reported by scanner
// if s.ErrorCount != 0 {
// t.Errorf("%d errors", s.ErrorCount)
// }
// }

func TestNullChar(t *testing.T) {
	s := New([]byte("\"\\0"))
	s.Scan() // Used to panic
}

func TestComment(t *testing.T) {
	testTokenList(t, tokenLists["comment"])
}

// func TestOperator(t *testing.T) {
// testTokenList(t, tokenLists["operator"])
// }

// func TestBool(t *testing.T) {
// testTokenList(t, tokenLists["bool"])
// }

// func TestIdent(t *testing.T) {
// testTokenList(t, tokenLists["ident"])
// }

// func TestString(t *testing.T) {
// testTokenList(t, tokenLists["string"])
// }

// func TestNumber(t *testing.T) {
// testTokenList(t, tokenLists["number"])
// }

// func TestFloat(t *testing.T) {
// testTokenList(t, tokenLists["float"])
// }

// func TestWindowsLineEndings(t *testing.T) {
// hcl := `// This should have Windows line endings
// resource "aws_instance" "foo" {
// user_data=<<HEREDOC
// test script
// HEREDOC
// }`
// hclWindowsEndings := strings.Replace(hcl, "\n", "\r\n", -1)

// literals := []struct {
// tokenType token.Type
// literal   string
// }{
// {token.COMMENT, "// This should have Windows line endings\r"},
// {token.IDENT, `resource`},
// {token.STRING, `"aws_instance"`},
// {token.STRING, `"foo"`},
// {token.LBRACE, `{`},
// {token.IDENT, `user_data`},
// {token.ASSIGN, `=`},
// {token.HEREDOC, "<<HEREDOC\r\n    test script\r\nHEREDOC\r\n"},
// {token.RBRACE, `}`},
// }

// s := New([]byte(hclWindowsEndings))
// for _, l := range literals {
// tok := s.Scan()

// if l.tokenType != tok.Type {
// t.Errorf("got: %s want %s for %s\n", tok, l.tokenType, tok.String())
// }

// if l.literal != tok.Text {
// t.Errorf("got:\n%v\nwant:\n%v\n", []byte(tok.Text), []byte(l.literal))
// }
// }
// }

// func TestRealExample(t *testing.T) {
// complexHCL := `// This comes from Terraform, as a test
// variable "foo" {
// default = "bar"
// description = "bar"
// }

// provider "aws" {
// access_key = "foo"
// secret_key = "${replace(var.foo, ".", "\\.")}"
// }

// resource "aws_security_group" "firewall" {
// count = 5
// }

// resource aws_instance "web" {
// ami = "${var.foo}"
// security_groups = [
// "foo",
// "${aws_security_group.firewall.foo}"
// ]

// network_interface {
// device_index = 0
// description = <<EOF
// Main interface
// EOF
// }

// network_interface {
// device_index = 1
// description = <<-EOF
// Outer text
// Indented text
// EOF
// }
// }`

// literals := []struct {
// tokenType token.Type
// literal   string
// }{
// {token.COMMENT, `// This comes from Terraform, as a test`},
// {token.IDENT, `variable`},
// {token.STRING, `"foo"`},
// {token.LBRACE, `{`},
// {token.IDENT, `default`},
// {token.ASSIGN, `=`},
// {token.STRING, `"bar"`},
// {token.IDENT, `description`},
// {token.ASSIGN, `=`},
// {token.STRING, `"bar"`},
// {token.RBRACE, `}`},
// {token.IDENT, `provider`},
// {token.STRING, `"aws"`},
// {token.LBRACE, `{`},
// {token.IDENT, `access_key`},
// {token.ASSIGN, `=`},
// {token.STRING, `"foo"`},
// {token.IDENT, `secret_key`},
// {token.ASSIGN, `=`},
// {token.STRING, `"${replace(var.foo, ".", "\\.")}"`},
// {token.RBRACE, `}`},
// {token.IDENT, `resource`},
// {token.STRING, `"aws_security_group"`},
// {token.STRING, `"firewall"`},
// {token.LBRACE, `{`},
// {token.IDENT, `count`},
// {token.ASSIGN, `=`},
// {token.NUMBER, `5`},
// {token.RBRACE, `}`},
// {token.IDENT, `resource`},
// {token.IDENT, `aws_instance`},
// {token.STRING, `"web"`},
// {token.LBRACE, `{`},
// {token.IDENT, `ami`},
// {token.ASSIGN, `=`},
// {token.STRING, `"${var.foo}"`},
// {token.IDENT, `security_groups`},
// {token.ASSIGN, `=`},
// {token.LBRACK, `[`},
// {token.STRING, `"foo"`},
// {token.COMMA, `,`},
// {token.STRING, `"${aws_security_group.firewall.foo}"`},
// {token.RBRACK, `]`},
// {token.IDENT, `network_interface`},
// {token.LBRACE, `{`},
// {token.IDENT, `device_index`},
// {token.ASSIGN, `=`},
// {token.NUMBER, `0`},
// {token.IDENT, `description`},
// {token.ASSIGN, `=`},
// {token.HEREDOC, "<<EOF\nMain interface\nEOF\n"},
// {token.RBRACE, `}`},
// {token.IDENT, `network_interface`},
// {token.LBRACE, `{`},
// {token.IDENT, `device_index`},
// {token.ASSIGN, `=`},
// {token.NUMBER, `1`},
// {token.IDENT, `description`},
// {token.ASSIGN, `=`},
// {token.HEREDOC, "<<-EOF\n\t\t\tOuter text\n\t\t\t\tIndented text\n\t\t\tEOF\n"},
// {token.RBRACE, `}`},
// {token.RBRACE, `}`},
// {token.EOF, ``},
// }

// s := New([]byte(complexHCL))
// for _, l := range literals {
// tok := s.Scan()
// if l.tokenType != tok.Type {
// t.Errorf("got: %s want %s for %s\n", tok, l.tokenType, tok.String())
// }

// if l.literal != tok.Text {
// t.Errorf("got:\n%+v\n%s\n want:\n%+v\n%s\n", []byte(tok.String()), tok, []byte(l.literal), l.literal)
// }
// }

// }

// func TestScan_crlf(t *testing.T) {
// complexHCL := "foo {\r\n  bar = \"baz\"\r\n}\r\n"

// literals := []struct {
// tokenType token.Type
// literal   string
// }{
// {token.IDENT, `foo`},
// {token.LBRACE, `{`},
// {token.IDENT, `bar`},
// {token.ASSIGN, `=`},
// {token.STRING, `"baz"`},
// {token.RBRACE, `}`},
// {token.EOF, ``},
// }

// s := New([]byte(complexHCL))
// for _, l := range literals {
// tok := s.Scan()
// if l.tokenType != tok.Type {
// t.Errorf("got: %s want %s for %s\n", tok, l.tokenType, tok.String())
// }

// if l.literal != tok.Text {
// t.Errorf("got:\n%+v\n%s\n want:\n%+v\n%s\n", []byte(tok.String()), tok, []byte(l.literal), l.literal)
// }
// }

// }

// func TestError(t *testing.T) {
// testError(t, "\x80", "1:1", "illegal UTF-8 encoding", token.ILLEGAL)
// testError(t, "\xff", "1:1", "illegal UTF-8 encoding", token.ILLEGAL)

// testError(t, "ab\x80", "1:3", "illegal UTF-8 encoding", token.IDENT)
// testError(t, "abc\xff", "1:4", "illegal UTF-8 encoding", token.IDENT)

// testError(t, `"ab`+"\x80", "1:4", "illegal UTF-8 encoding", token.STRING)
// testError(t, `"abc`+"\xff", "1:5", "illegal UTF-8 encoding", token.STRING)

// testError(t, `01238`, "1:6", "illegal octal number", token.NUMBER)
// testError(t, `01238123`, "1:9", "illegal octal number", token.NUMBER)
// testError(t, `0x`, "1:3", "illegal hexadecimal number", token.NUMBER)
// testError(t, `0xg`, "1:3", "illegal hexadecimal number", token.NUMBER)
// testError(t, `'aa'`, "1:1", "illegal char", token.ILLEGAL)

// testError(t, `"`, "1:2", "literal not terminated", token.STRING)
// testError(t, `"abc`, "1:5", "literal not terminated", token.STRING)
// testError(t, `"abc`+"\n", "1:5", "literal not terminated", token.STRING)
// testError(t, `"${abc`+"\n", "2:1", "literal not terminated", token.STRING)
// testError(t, `/*/`, "1:4", "comment not terminated", token.COMMENT)
// testError(t, `/foo`, "1:1", "expected '/' for comment", token.COMMENT)
// }

// func testError(t *testing.T, src, pos, msg string, tok token.Type) {
// s := New([]byte(src))

// errorCalled := false
// s.Error = func(p token.Pos, m string) {
// if !errorCalled {
// if pos != p.String() {
// t.Errorf("pos = %q, want %q for %q", p, pos, src)
// }

// if m != msg {
// t.Errorf("msg = %q, want %q for %q", m, msg, src)
// }
// errorCalled = true
// }
// }

// tk := s.Scan()
// if tk.Type != tok {
// t.Errorf("tok = %s, want %s for %q", tk, tok, src)
// }
// if !errorCalled {
// t.Errorf("error handler not called for %q", src)
// }
// if s.ErrorCount == 0 {
// t.Errorf("count = %d, want > 0 for %q", s.ErrorCount, src)
// }
// }

// func testTokenList(t *testing.T, tokenList []tokenPair) {
// // create artifical source code
// buf := new(bytes.Buffer)
// for _, ident := range tokenList {
// fmt.Fprintf(buf, "%s\n", ident.text)
// }

// s := New(buf.Bytes())
// for _, ident := range tokenList {
// tok := s.Scan()
// if tok.Type != ident.tok {
// t.Errorf("tok = %q want %q for %q\n", tok, ident.tok, ident.text)
// }

// if tok.Text != ident.text {
// t.Errorf("text = %q want %q", tok.String(), ident.text)
// }

// }
// }

// func countNewlines(s string) int {
// n := 0
// for _, ch := range s {
// if ch == '\n' {
// n++
// }
// }
// return n
// }