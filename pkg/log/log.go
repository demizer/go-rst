package log

import (
	"os"

	"github.com/davecgh/go-spew/spew"
	klog "github.com/go-kit/kit/log"
)

// Used for debugging only
var spd = spew.ConfigState{ContinueOnMethod: true, Indent: "\t", MaxDepth: 0} //, DisableMethods: true}

// NewLogger wraps a logger with a name context and caller information. If a named context is specified in the excludes
// slice, then any logging to that context will be ignored.
func NewLogger(name string, caller bool, excludes []string, logr klog.Logger) Logger {
	return Logger{
		ctx:      name,
		caller:   caller,
		log:      logr,
		excludes: excludes,
	}
}

// Logger implements the go-kit logger type.
type Logger struct {
	ctx      string      // Shows up in the log output as field "name".
	caller   bool        // Include the caller field name in the log output.
	log      klog.Logger // The standard logger for which wrapped loggers are based
	excludes []string    // exclude named contexts from output
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
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(4))
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
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(4), "msg", message)
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
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(4))
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
		logr = klog.WithPrefix(l.log, "name", l.ctx, "caller", klog.Caller(4))
	}
	return logr.Log(keyvals...)
}

// Dump pretty prints the var interface to standard output.
func Dump(v interface{}) {
	spd.Dump(v)
}

// DumpExit pretty prints the var interface to standard output and terminates program execution.
func DumpExit(v interface{}) {
	spd.Dump(v)
	os.Exit(1)
}
