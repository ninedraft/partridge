package vec

import (
	"fmt"
	"strings"
)

func PanicF(f string, args ...interface{}) {
	panic(fmt.Sprintf(f, args...))
}

func AssertNil(value interface{}, msgs ...string) {
	var msg = "expected nil, got %v"
	if len(msgs) > 0 {
		msg += "\n" + strings.Join(msgs, "\n")
	}
	switch value {
	case nil:
		//
	default:
		PanicF(msg, value)
	}
}
