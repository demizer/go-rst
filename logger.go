package rst

import (
	"fmt"
	"os"
	"strings"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/term"
	"github.com/go-stack/stack"
)

var logc **log.Context

var (
	// LogCtx is the default logger for the parse package
	logp                = NewLogCtx("parser")
	logt                = NewLogCtx("test")
	excludeNamedContext string // Exclude a log context from being shown in the output
)

type LogCtx struct {
	name string
	ctx  *log.Context
}

// NewLogCtx creates a new logging context with name and returns o LogCtx ready to use.
func NewLogCtx(name string) *LogCtx {
	return &LogCtx{name: name, ctx: log.NewContext(log.NewNopLogger())}
}

// // LogSetContext sets a logger context.
// func LogSetContext(l *log.Context) {
// logp.ctx = l
// // logl.ctx = l
// logt.ctx = l
// }

// NewColorLogCtx creates a new logger context with ansi coloring.
func NewColorLogCtx(name string, colorFn func(keyvals ...interface{}) term.FgBgColor) *LogCtx {
	return &LogCtx{name: name, ctx: log.NewContext(term.NewLogger(os.Stdout, log.NewLogfmtLogger, colorFn))}
}

// Msg logs a message to the log context.
func (l *LogCtx) Msg(message string) { l.Log("msg", message) }

// Error logs an error to the log context.
func (l *LogCtx) Err(err error) { l.Log("msg", err.Error()) }

// Log writes log output to the logCtx of the package with added context
func (l *LogCtx) Log(keyvals ...interface{}) error {
	if strings.Contains(excludeNamedContext, l.name) {
		return nil
	}
	cs := stack.Caller(2)
	funcName := fmt.Sprintf("%s", cs)
	file := cs.String()
	return l.ctx.WithPrefix("name", l.name, "line", file, "func", funcName).Log(keyvals...)
}
