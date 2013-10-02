package debug

import (
	"path"
	"runtime"
	"strings"
	"sync"
)

var (
	colors   = []string{"cyan", "green", "red", "yellow", "magenta", "blue"}
	colorMap = make(map[string]string)
	mutex    sync.RWMutex
)

func caller() string {
	_, filename, _, _ := runtime.Caller(2)
	return strings.Split(path.Base(filename), ".")[0]
}

func colorOf(name string) string {
	mutex.RLock()
	if color, ok := colorMap[name]; ok {
		mutex.RUnlock()
		return color
	}
	mutex.RUnlock()

	mutex.Lock()
	colorMap[name] = colors[ctr%len(colors)]
	mutex.Unlock()
	ctr++

	return colorMap[name]
}
