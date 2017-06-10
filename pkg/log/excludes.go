package log

import "fmt"

// LoggerExcludes satisfies the flag Value interface
type LoggerExcludes []string // Exclude a log context from being shown in the output

// String returns LoggerExcludes as a formatted string
func (e *LoggerExcludes) String() string { return fmt.Sprintf("%#v", e) }

// Set appends value
func (e *LoggerExcludes) Set(value string) error {
	*e = append(*e, value)
	return nil
}
