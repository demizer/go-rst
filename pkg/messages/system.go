package messages

// SystemMessageLevel implements four levels for messages and is used in conjunction with the ParserMessage type.
type SystemMessageLevel int

const (
	levelWarning SystemMessageLevel = iota
	levelError
)

var systemMessageLevels = [...]string{
	"WARNING",
	"ERROR",
}

// String implments Stringer and return a string of the SystemMessageLevel.
func (s SystemMessageLevel) String() string { return systemMessageLevels[s] }

// FromString returns the SystemMessageLevel converted from the string name.
func SystemMessageLevelFromString(name string) SystemMessageLevel {
	for num, sLvl := range systemMessageLevels {
		if name == sLvl {
			return SystemMessageLevel(num)
		}
	}
	return -1
}
