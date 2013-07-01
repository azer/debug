package debug

import (
	"os"
	"strings"
)

var enabled []string = strings.Split(os.Getenv("DEBUG"), ",")

func isEnabled (name string) bool {
	if enabled[0] == "*" {
		return true
	}

	for _, v := range enabled {
		if (v == name) {
			return true
		}
	}

	return false
}
