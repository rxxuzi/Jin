// Package build , build.go!
package build

import (
	"jin/pkg"
	"jin/set"
)

// Build structure to manage build process
type Build struct {
	Src  []string         // List of file paths
	Obj  []string         // List of object files
	Lang map[string]*Lang // Dynamic generation based on {"C":[], "ASM":[]}
}

// Lang represents language-specific data
type Lang struct {
	Path string   // Language name
	Src  []string // Source files for this language
}

// NewBuild initializes a new build context with source and ignore paths
func NewBuild(sourcePaths, ignorePaths []string) *Build {
	files := set.Diff(sourcePaths, ignorePaths)

	// Prepare the language map
	langFiles := pkg.CategorizeFilesByLanguage(files)
	langMap := make(map[string]*Lang)
	for lang, paths := range langFiles {
		langMap[lang] = &Lang{Path: lang, Src: paths}
	}

	b := &Build{
		Src:  files,
		Obj:  []string{},
		Lang: langMap,
	}
	return b
}
