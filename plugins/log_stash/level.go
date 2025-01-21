package log_stash

import "encoding/json"

type Level int

const (
	DebugLevel = 1 + iota
	InfoLevel
	WarnLevel
	ErrorLevel
)

func (r Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}
func (r Level) String() string {
	var str string
	switch r {
	case DebugLevel:
		str = "debug"
	case InfoLevel:
		str = "info"
	case WarnLevel:
		str = "warn"
	case ErrorLevel:
		str = "error"
	default:
		str = "其他"
	}
	return str
}
