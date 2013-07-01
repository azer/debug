package debug

import (
	"path"
	"runtime"
	"strings"
)

var (
	colors = []string{ "cyan", "green", "red", "yellow", "magenta", "blue" }
	colorMap = make(map[string]string)
)

func caller () string {
	_, filename, _, _ := runtime.Caller(2)
	return strings.Split(path.Base(filename), ".")[0]
}

func colorOf (name string) string {
	if color, ok := colorMap[name]; ok {
		return color
	}

	colorMap[name] = colors[ctr % len(colors)]
	ctr++

	return colorMap[name]
}
