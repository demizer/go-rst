package log

import klog "github.com/go-kit/kit/log"

// Config allows sharing a log config accross loggers.
type Config struct {
	// The name of the logger. Shows up in the output.
	Name string

	// The logger to use. Allows setting output and format.
	Logger klog.Logger

	// Show call information in log output (can affect performance)
	Caller bool

	// The call stack depth. Useful for setting different depths for apps versus testing.
	CallDepth int

	// Excludes allows excluding logger contexts by name.
	Excludes []string
}

// NewConfig returns a new logger config with the arguments set.
func NewConfig(name string, logr klog.Logger, caller bool, callDepth int) *Config {
	return &Config{
		Name:      name,
		Logger:    logr,
		Caller:    caller,
		CallDepth: callDepth,
	}
}
