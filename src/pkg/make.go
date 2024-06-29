// Package pkg , make.go!
package pkg

import (
	"os"
	"path/filepath"
)

const BuildDir = "jin/build"
const CacheDir = "jin/cache"

func createDirIfNotExist(path string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.MkdirAll(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
}

func CreateDir(base string, sub []string) {
	createDirIfNotExist(base)
	for _, dir := range sub {
		fullPath := filepath.Join(base, dir)
		createDirIfNotExist(fullPath)
	}
}
