// Package config fast.go!
package conf

import (
	"encoding/json"
	"jin/pkg"
	"os"
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
	config.Lang = pkg.DetectLanguages(config.Source)
	return config, nil
}
