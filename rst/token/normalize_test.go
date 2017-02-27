package token

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var normTests = [...]struct {
	name   string
	test   []byte // byte slice containing unicode literals
	expect []byte // byte slice of "normalized" unicode literals
	err    string // the expected error
}{
	{
		name:   "Swedish place of interest symbol",
		test:   []byte{'\\', 'u', '2', '3', '1', '8'},
		expect: []byte("⌘"),
	},
	{
		name:   "Should be EN QUAD space",
		test:   []byte{'\\', 'u', '2', '0', '0', '0'},
		expect: []byte(" "),
	},
	{
		name:   "Combining: Cyrillic Small leter Short U",
		test:   []byte{'\\', 'u', '0', '4', '5', 'E'},
		expect: []byte("ў"),
	},
	{
		name:   "Combining: Latin small leter a with grave accent",
		test:   []byte{'\\', 'u', '0', '0', 'E', '0'},
		expect: []byte("à"),
	},
	{
		name:   "Combining with EN QUAD space",
		test:   []byte{'\\', 'u', '2', '0', '0', '0', ' ', '\\', 'u', '0', '0', 'E', '0'},
		expect: []byte("  à"),
	},
	{
		name:   "Pointing double angle quotation mark",
		test:   []byte{'\\', 'x', 'a', 'b'},
		expect: []byte("«"),
	},
	{
		name:   "Escaped newline",
		test:   []byte{'\\', 'x', 'a', 'b', ' ', 'e', 's', 'c', 'a', 'p', 'e', ' ', '\\', '\\'},
		expect: []byte("« escape \\"),
	},
	{
		name: "Bad unicode literal #1",
		test: []byte{'\\', 'u', 'a', 'b'},
		err:  "invalid unicode literal",
	},
	{
		name: "Bad unicode literal #2",
		test: []byte{'\\', 'x', 'a'},
		err:  "invalid unicode literal",
	},
	{
		name: "Bad unicode literal #3",
		test: []byte{'\\', 'u', '0', '0', 'a', ' ', ' ', ' '},
		err:  "invalid unicode literal: strconv.ParseUint: parsing \"00a \": invalid syntax",
	},
}

func TestNormalize(t *testing.T) {
	for _, test := range normTests {
		// fmt.Printf("Running test %q...\n", test.name)
		assert := assert.New(t)
		o, err := normalize(test.test)
		if len(test.err) > 0 && !assert.EqualError(err, test.err) {
			assert.Fail("Should result in error")
		}
		if !assert.Equal(o, test.expect) {
			assert.Fail("Should be the same")
		}
	}
}
