package messages

type ParserMessage struct {
	Type          MessageType
	Line          int    // The line in the input that caused the message
	LiteralText   string // Additional text
	StartPosition int    // The start position of the problem resulting in a message
}

// NewParserMessage returns a parser message built from t.
func NewParserMessage(t MessageType) *ParserMessage {
	return &ParserMessage{
		Type: t,
	}
}

// Level returns the MessageType level.
func (p ParserMessage) Level() string { return p.Type.level() }

// Message returns the message of the MessageType as a string.
func (p ParserMessage) Message() string { return p.Type.message() }
