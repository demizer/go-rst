package scanner

import (
	"testing"

	"github.com/demizer/go-rst/rst/token"
)

type scanExpect struct {
	srcText string
	,s
}

var tokenLists = map[string][]tokenPair{
	"comment": []tokenPair{
		{token.COMMENT, ".."},
	},
}

func TestNullChar(t *testing.T) {
	s := New([]byte("\"\\0"))
	s.Scan() // Used to panic
}

func TestComment(t *testing.T) {
	testTokenList(t, tokenLists["comment"])
}
