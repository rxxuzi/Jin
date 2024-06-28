// Package pkg , make.go!
package pkg

import (
	"os"
	"path/filepath"
)

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
