package debug

type DebugScopeType func (scope string) func (text string, args ...interface{})

func DebugScope (scope string) func (text string, args ...interface{}) {
	return Debug
}
