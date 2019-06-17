package dircontent

import (
	"fmt"
	"os"
	"path/filepath"
)

// PrintDirContent prints content of directory by dirpath.
// It returns an error.
func PrintDirContent(dirPath string) error {

	return filepath.Walk(
		dirPath,
		printContent,
	)

}

func printContent(path string, info os.FileInfo, err error) error {

	var output string

	if !info.IsDir() {
		output = fmt.Sprintf("%s - %d byte(s)\n", path, info.Size())
	} else {
		output = fmt.Sprintf("%s\n", path)
	}

	fmt.Print(output)

	return nil

}
