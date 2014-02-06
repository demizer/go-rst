// go-rst - A reStructuredText parser for Go
// 2014 (c) The go-rst Authors
// MIT Licensed. See LICENSE for details.
package rst

import (
	"github.com/demizer/go-rst/parse"
)

type Document struct {
	name string
	*parse.Tree
}

func New(name string) *Document {

	return &Document{
		name: name,
	}
}
