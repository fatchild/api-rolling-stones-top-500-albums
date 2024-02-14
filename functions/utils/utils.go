package utils

import (
	"path"
	"path/filepath"
	"runtime"
)

func RootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	// This seems a bit hacky, should I recursively move from current directory backwards until I find the project name? No sure what the best approach is. It seems totally dependant on environment.
	return filepath.Dir(d + "/../../")
}

func StringIn(str string, list []string) bool {
	for _, b := range list {
		if b == str {
			return true
		}
	}
	return false
}
