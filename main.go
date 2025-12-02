package main

import (
	"fmt"
	"io/ioutil"
	"log"
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

	for _, dir := range dirs {
		checkFiles(dir)
	}
}

func getDirectories(catalog string) []string {
	var dirs []string

	filepath.WalkDir("../"+catalog, func(path string, d os.DirEntry, err error) error {
		if err != nil {
			return err
		}
		if d.IsDir() && !strings.Contains(path, "vendor") && !strings.Contains(path, "cache") && !strings.Contains(path, ".git") && !strings.Contains(path, ".idea") && !strings.Contains(path, "var") {
			dirs = append(dirs, path)
		}
		return nil
	})

	return dirs
}

func checkFiles(path string) bool {
	phpFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, phpFile := range phpFiles {
		if !phpFile.IsDir() && strings.HasSuffix(strings.ToLower(phpFile.Name()), ".php") {
			fmt.Println(phpFile.Name())
		}
	}

	return true
}
