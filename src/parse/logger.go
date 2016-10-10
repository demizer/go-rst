package parse

import (
	"fmt"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-stack/stack"
)

var (
	// LogCtx is the default logger for the parse package
	Log                 = log.NewContext(log.NewNopLogger())
	excludeNamedContext string // Exclude a log context from being shown in the output
	debug               bool
)

type logCtx struct {
	name string
	ctx  *log.Context
}

// Log writes log output to the LogCtx of the package with added context
func (l *logCtx) Log(keyvals ...interface{}) error {
	if strings.Contains(excludeNamedContext, l.name) {
		return nil
	}
	cs := stack.Caller(2)
	funcName := fmt.Sprintf("%n", cs)
	file := cs.String()
	return l.ctx.WithPrefix("name", l.name, "caller", file, "func", funcName).Log(keyvals...)
}

// NewLogCtx creates a new logging context with name and returns o logCtx ready to use.
func NewLogCtx(name string) *logCtx {
	return &logCtx{name: name, ctx: Log}
}
