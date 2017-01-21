package parse

import "github.com/davecgh/go-spew/spew"

// Used for debugging only
var spd = spew.ConfigState{ContinueOnMethod: true, Indent: "\t", MaxDepth: 0} //, DisableMethods: true}

const (
	// The middle of the Tree.token buffer so that there are three possible "backup" token positions and three forward
	// "peek" positions.
	zed = 4

	// Default indent width
	indentWidth = 4
)
