package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	var catalog string
	fmt.Println("Please enter the name of the catalog: ")
	fmt.Fscan(os.Stdin, &catalog)

	var dirs []string
	dirs = getDirectories(catalog)

	fmt.Println("Directorys found:")
	for _, dir := range dirs {
		fmt.Println(dir)
	}
}

func getDirectories(catalog string) []string {
	var dirs []string

	filepath.WalkDir("../"+catalog, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && !strings.Contains(path, "vendor") && !strings.Contains(path, "cache") && !strings.Contains(path, ".git") && !strings.Contains(path, ".idea") {
			dirs = append(dirs, path)
		}
		return nil
	})

	return dirs
}
