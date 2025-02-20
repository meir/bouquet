package client

import (
	"fmt"
	"io/fs"
	"path/filepath"
	"slices"
)

// matchName returns a function that matches a file with the given name to be used in slices.IndexFunc
func matchName(filename string) func(file fs.DirEntry) bool {
	// matches if the path is either a directory or returns the file that matches with its name without the extension
	return func(file fs.DirEntry) bool {
		if file.IsDir() {
			return file.Name() == filename
		}

		match := fmt.Sprintf("%s.*", filename)
		ok, err := filepath.Match(match, file.Name())
		return ok && (err == nil)
	}
}

// resolvePath will find the correct path for the requested file in the embedded file system
func resolvePath(path string, recurrant bool) string {
	ext := filepath.Ext(path)
	dir := filepath.Dir(path)
	filename := filepath.Base(path)

	// check if the extension is omitted
	// will need to resolve a file for this as js/ts allows importing certain files without an extension
	// such as `import * from ./file`
	if ext == "" {
		directory, err := fs.ReadDir(assets, dir)
		if err != nil {
			return ""
		}

		// find files/directories in directory matching name
		fileIndex := slices.IndexFunc(directory, matchName(filename))
		if fileIndex == -1 {
			return ""
		}
		file := directory[fileIndex]

		// if the file turns out to be a folder we need to check if the folder contains an index
		// index can again be either js/ts/tsx so we need to basically run this method again
		// thus the recurrant variable to check if this is the first or second time going through here
		if file.IsDir() {
			if recurrant {
				return ""
			}

			path = filepath.Join(path, "/index")
			return resolvePath(path, true) // set recurrant to true this time
		}

		// return the file path
		return filepath.Join(dir, file.Name())
	}

	return path // return the path as is
}
