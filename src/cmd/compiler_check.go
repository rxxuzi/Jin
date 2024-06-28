// Package cmd compiler_check.go!
package cmd

import (
	"os/exec"
)

func checkCompiler(compiler string) bool {
	cmd := exec.Command(compiler, "--version") // "--version" オプションを使ってバージョン情報を取得
	if err := cmd.Run(); err != nil {
		return false
	}
	return true
}

func CheckCompilers(compilers []string) []bool {
	results := make([]bool, len(compilers))
	for i, compiler := range compilers {
		results[i] = checkCompiler(compiler)
	}
	return results
}
