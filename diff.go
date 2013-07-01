package debug

import (
	"fmt"
	"time"
)

var last int64 = 0

func diff() string {
	if last == 0 {
		last = now()
		return ""
	}

	d := now() - last

	if d == 0 {
		return ""
	}

	return fmt.Sprintf("+%d", d)
}

func now() int64 {
	return time.Now().UnixNano() / 1000000
}
