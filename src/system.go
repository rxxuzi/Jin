package main

import (
	"jin/pkg"
)

var baseDir string = "jin-build"
var subDir = []string{"build", "cache", "debug"}

func Init() {
	pkg.CreateDir(baseDir, subDir)
}
