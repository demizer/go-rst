package log

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	klog "github.com/go-kit/kit/log"
)

// Used for debugging only
var spd = spew.ConfigState{ContinueOnMethod: true, Indent: "\t", MaxDepth: 0} //, DisableMethods: true}

// Logger implements the go-kit logger type.
type Logger struct {
	ctx       string      // Shows up in the log output as field "name".
	caller    bool        // Include the caller field name in the log output.
	callDepth int         // Call stack depth for logging
	log       klog.Logger // The standard logger for which wrapped loggers are based
	excludes  []string    // exclude named contexts from output
}

// NewLogger wraps a logger with a name context and caller information. If a named context is specified in the excludes
// slice, then any logging to that context will be ignored.
func NewLogger(name string, caller bool, callDepth int, excludes []string, logr klog.Logger) Logger {
	return Logger{
		ctx:       name,
		caller:    caller,
		callDepth: callDepth,
		log:       logr,
		excludes:  excludes,
	}
}

func (l Logger) isExcluded() bool {
	if len(l.excludes) > 0 {
		for _, v := range l.excludes {
			if v == l.ctx {
				return true
			}
		}
	}
	return false
}

// StdLogger gets the Logger originally used in creation of the wrapped logger.
func (l Logger) StdLogger() klog.Logger { return l.log }

// SetLogger set the wrapped Logger standard logger
func (l Logger) SetLogger(logr klog.Logger) { l.log = logr }

// Msg logs a message to the log context.
func (l Logger) Msg(message string) error {
	if l.isExcluded() {
		return nil
	}
	logr := klog.WithPrefix(l.log, "name", l.ctx)
	if l.caller {
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(l.callDepth))
	}
	return logr.Log("msg", message)
}

// Msgr logs a message with additional fields.
func (l Logger) Msgr(message string, keyvals ...interface{}) error {
	if l.isExcluded() {
		return nil
	}
	logr := klog.WithPrefix(l.log, "name", l.ctx, "msg", message)
	if l.caller {
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(l.callDepth), "msg", message)
	}
	return logr.Log(keyvals...)
}

// Err logs an error to the log context.
func (l Logger) Err(err error) error {
	if l.isExcluded() {
		return nil
	}
	logr := klog.WithPrefix(l.log, "name", l.ctx)
	if l.caller {
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(l.callDepth))
	}
	return logr.Log("error", err.Error())
}

// Log satisfies the logger interface.
func (l Logger) Log(keyvals ...interface{}) error {
	if l.isExcluded() {
		return nil
	}
	logr := klog.WithPrefix(l.log, "name", l.ctx)
	if l.caller {
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(l.callDepth))
	}
	return logr.Log(keyvals...)
}

// Dump pretty prints the v interface into the msg field
//
// The output is a string containing escaped newlines. It is possoble to show the structured output in one line using:
//
// echo -e $(go test -v ./pkg/parser -test.run=".*<test_id>*_Parse.*" -debug -exclude=lexer | grep "msg=dump" | sed -n "s/.*obj=\"\(.*\)\"/\1/p")
//
func (l Logger) Dump(v interface{}) {
	WithCallDepth(l, l.callDepth+1).Msgr("dump", "obj", spd.Sdump(v))
}

// DumpExit pretty prints the v interface to msg field and terminates program execution.
//
// The output is a string containing escaped newlines. It is possoble to show the structured output in one line using:
//
// echo -e $(go test -v ./pkg/parser -test.run=".*<test_id>*_Parse.*" -debug -exclude=lexer | grep "msg=dump" | sed -n "s/.*obj=\"\(.*\)\"/\1/p")
//
func (l Logger) DumpExit(v interface{}) {
	WithCallDepth(l, l.callDepth+1).Msgr("dump", "obj", spd.Sdump(v))
	os.Exit(1)
}

func WithCallDepth(l Logger, callDepth int) Logger {
	return NewLogger(l.ctx, true, callDepth, l.excludes, l.log)
}
