// Package config fast.go!
package conf

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type BuildConfig struct {
	Project   string `json:"project"`
	Lang      []string
	Source    []string            `json:"source"`
	Ignore    []string            `json:"ignore"`
	Libraries []Library           `json:"libraries"`
	Option    map[string][]string `json:"option"`
}

type Library struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

// detectLanguage はファイル名リストから言語を推測する
func detectLanguage(files []string) []string {
	langMap := map[string]string{
		".c":   "C",
		".h":   "C",
		".s":   "ASM",
		".asm": "ASM",
		".go":  "Go",
		".cpp": "C++",
		".hpp": "C++",
		".f":   "Fortran",
		".f90": "Fortran",
	}
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

func LoadBuildConfig(path string) (BuildConfig, error) {
	var config BuildConfig
	data, err := os.ReadFile(path)
	if err != nil {
		return config, err
	}
	err = json.Unmarshal(data, &config)
	if err != nil {
		return config, err
	}
	// 言語を推測する
	config.Lang = detectLanguage(config.Source)
	return config, nil
}
