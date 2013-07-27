package debug

type DebugScopeType func (scope string) func (text string, args ...interface{})

func DebugScope (scope string, keys ...string) func (text string, args ...interface{}) {
	enabled = append(enabled, keys...)
	return Debug
}
