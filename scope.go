package debug

import "os"

type DebugScopeType func (scope string) func (text string, args ...interface{})

func DebugScope (scope string, keys ...string) func (text string, args ...interface{}) {
	if os.Getenv("DEBUG") == "" {
		enabled = append(enabled, keys...)
		return Debug
	}

	return Debug
}
