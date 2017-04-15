package logging

import (
	"fmt"
	"os"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/term"
	"github.com/go-stack/stack"
)

var loggers []*LogContext

func init() {
	loggers = make([]*LogContext, 0)
	RegisterNewLogContext("default", log.NewNopLogger())
}

type LogContext struct {
	Name   string
	Logger *log.Logger
}

// func NewLogContext(name string, l log.Logger) *LogContext {
// return &LogContext{Name: name, Context: log.NewContext(l)}
// }

func RegisterNewLogContext(name string, l log.Logger) *LogContext {
	ctx := &LogContext{Name: name, Context: log.NewContext(l)}
	loggers = append(loggers, ctx)
	return ctx
}

// func StdContext() *LogContext { return stdContext }

func StdLogger() *log.Logger { return loggers[0].Context }

func SetStdLogger(l *log.Logger) { loggers[0].Context = l }

// NewColorLogCtx creates a new logger context with ansi coloring.
func NewColorLogCtx(name string, colorFn func(keyvals ...interface{}) term.FgBgColor) *LogContext {
	return &LogContext{Name: name, Context: log.NewContext(term.NewLogger(os.Stdout, log.NewLogfmtLogger, colorFn))}
}

// Msg logs a message to the log context.
func (l *LogContext) Msg(message string) { l.Log("msg", message) }

// Error logs an error to the log context.
func (l *LogContext) Err(err error) { l.Log("msg", err.Error()) }

// Log writes log output to the LogContext of the package with added context
func (l *LogContext) Log(keyvals ...interface{}) error {
	fmt.Printf("%+#v\n", l.Context)
	// if strings.Contains(l.ExcludeNamedContext, l.Name) {
	// return nil
	// }
	cs := stack.Caller(2)
	funcName := fmt.Sprintf("%s", cs)
	file := cs.String()
	return l.Context.WithPrefix("name", l.Name, "line", file, "func", funcName).Log(keyvals...)
}
