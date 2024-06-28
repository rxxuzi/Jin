package main

import (
	"fmt"
	"jin/clang"
	"jin/conf"
)

func main() {
	configPath := "build.json"
	config, err := conf.LoadBuildConfig(configPath)
	if err != nil {
		fmt.Println("Error loading config:", err)
		return
	}
	fs := clang.NewFSGO(config.Source, true)

	if fs != nil {
		paths := fs.GetFilePaths()
		for _, path := range paths {
			fmt.Println(path)
		}
		fs.Destroy() // リソースを解放
	} else {
		fmt.Println("Failed to initialize FSGO")
	}
}
