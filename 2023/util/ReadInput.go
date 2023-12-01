package util

import (
	"os"
	"path"
	"runtime"
	"strings"
)

func ReadInput(filename string) string {

	_, runner, _, ok := runtime.Caller(1)

	if !ok {
		panic("Could not determine path from caller")
	}

	filePath := path.Join(path.Dir(runner), filename)
	data, err := os.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	// trim if input is wrongly copied with newlines
	return strings.TrimRight(string(data), "\n")
}
