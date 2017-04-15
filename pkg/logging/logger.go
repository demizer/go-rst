package logging

import (
	"github.com/davecgh/go-spew/spew"
	"github.com/go-kit/kit/log"
)

// Used for debugging only
var spd = spew.ConfigState{ContinueOnMethod: true, Indent: "\t", MaxDepth: 0} //, DisableMethods: true}

var stdLogger log.Logger

func init() {
	stdLogger = Logger{l: log.NewNopLogger()}
}

func NewLogger(name string, l log.Logger) Logger {
	return Logger{log.With(l, "name", name, "caller", log.Caller(4))}
}

func StdLogger() log.Logger { return stdLogger }

func SetStdLogger(l log.Logger) { stdLogger = l }

// Logger implements the go-kit logger type.
type Logger struct {
	l log.Logger
}

// Msg logs a message to the log context.
func (r Logger) Msg(message string) { r.l.Log("msg", message) }

// Err logs an error to the log context.
func (r Logger) Err(err error) { r.l.Log("msg", err.Error()) }

// Log satisfies the logger interface.
func (r Logger) Log(keyvals ...interface{}) error { return r.l.Log(keyvals...) }
