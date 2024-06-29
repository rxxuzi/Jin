// Package pkg, lang.go!
package pkg

import "path/filepath"

var langMap = map[string]string{
	".c":   "C",
	".h":   "C",
	".s":   "ASM",
	".asm": "ASM",
	".cpp": "C++",
	".hpp": "C++",
	".f":   "Fortran",
	".f90": "Fortran",
}

func DetectLanguages(files []string) []string {
	var languages []string
	seen := make(map[string]bool)
	for _, file := range files {
		ext := filepath.Ext(file)
		if lang, ok := langMap[ext]; ok {
			if !seen[lang] {
				languages = append(languages, lang)
				seen[lang] = true
			}
		}
	}
	return languages
}

// CategorizeFilesByLanguage sorts files into languages based on their extensions
func CategorizeFilesByLanguage(files []string) map[string][]string {
	categorized := make(map[string][]string)
	for _, file := range files {
		ext := filepath.Ext(file)
		if lang, ok := langMap[ext]; ok {
			categorized[lang] = append(categorized[lang], file)
		}
	}
	return categorized
}
