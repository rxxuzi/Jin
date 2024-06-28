// fis.go

package clang

/*
#include "fis.h"
*/
import "C"
import (
	"unsafe"
)

// FSGO is a wrapper around the C FIS structure to make it more idiomatic in Go.
type FSGO struct {
	fis []*C.FIS
	set map[string]struct{}
}

// NewFSGO creates new file scanners for each provided path and returns a Go wrapper around them.
func NewFSGO(paths []string, recursive bool) *FSGO {
	fs := &FSGO{
		fis: make([]*C.FIS, len(paths)),
		set: make(map[string]struct{}),
	}

	for i, path := range paths {
		cPath := C.CString(path)
		defer C.free(unsafe.Pointer(cPath))
		fis := C.newFIS(cPath, C.bool(recursive))
		if fis == nil {
			continue // Handle the error appropriately
		}
		fs.fis[i] = fis

		// Populate the set with file paths from the FIS structure
		for j := 0; j < int(fis.size); j++ {
			cFilePath := C.getFilePath(fis, C.int(j))
			if cFilePath == nil {
				continue
			}
			filePath := C.GoString(cFilePath)
			fs.set[filePath] = struct{}{}
		}
	}

	return fs
}

// Destroy releases all resources associated with each FSGO.
func (fs *FSGO) Destroy() {
	for _, fis := range fs.fis {
		if fis != nil {
			C.FISDestroy(fis)
		}
	}
	fs.fis = nil
}

// GetFilePaths returns a slice of unique file paths.
func (fs *FSGO) GetFilePaths() []string {
	paths := make([]string, 0, len(fs.set))
	for path := range fs.set {
		paths = append(paths, path)
	}
	return paths
}
