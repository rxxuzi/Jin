package build

import (
	"fmt"
	"jin/pkg"
	"os"
	"os/exec"
	"path/filepath"
)

var CompileMap = map[string]string{
	"C":       "gcc",
	"ASM":     "nasm",
	"C++":     "g++",
	"Fortran": "gfortran",
}

// CompileProject compiles the entire project based on language settings and provided options.
func (b *Build) CompileProject(options map[string][]string) error {
	for lang, details := range b.Lang {
		compiler, ok := CompileMap[lang]
		if !ok {
			fmt.Printf("No compiler found for language %s\n", lang)
			continue
		}
		compileOptions, _ := options[compiler] // Default to empty slice if not found

		fmt.Printf("Compiling [%s]\n", lang)
		for _, src := range details.Src {
			objPath := generateObjPath(src)
			if err := compileFile(src, compiler, compileOptions, objPath); err != nil {
				return err
			}
			b.Obj = append(b.Obj, objPath) // Save the path of the compiled object file
		}
	}
	return nil
}

// generateObjPath creates the output path for the object file based on the source file path.
func generateObjPath(sourcePath string) string {
	baseName := filepath.Base(sourcePath)           // e.g., main.c
	newBaseName := baseName + ".obj"                // e.g., main.c.obj
	return filepath.Join(pkg.BuildDir, newBaseName) // Create full path
}

// compileFile compiles a single source file using the specified compiler and saves the output.
func compileFile(filePath, compiler string, options []string, outputPath string) error {
	args := append([]string{filePath}, options...) // Add the file path to the start of options
	args = append(args, "-o", outputPath)          // Specify the output file
	cmd := exec.Command(compiler, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		fmt.Printf("Failed to compile %s: %s\n", filePath, err)
		return err
	}
	fmt.Printf("Compiled %s successfully.\n", filePath)
	return nil
}
