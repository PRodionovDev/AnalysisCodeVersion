package main

import (
	"bufio"
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

func checkFiles(path string) {
	phpFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, phpFile := range phpFiles {
		if !phpFile.IsDir() && strings.HasSuffix(strings.ToLower(phpFile.Name()), ".php") {
			checkFile(path + "/" + phpFile.Name())
		}
	}
}

func checkFile(path string) bool {
	// todo: replace to read file with version
	searchWords := []string{"php", "function"}

	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR couldn't open file:", err)
		return false
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		wordsInLine := strings.Fields(line)

		for _, word := range wordsInLine {
			for _, searchWord := range searchWords {
				if word == searchWord {
					fmt.Printf("Find: '%s' in '%s'\n", word, path)
					break
				}
			}
		}
	}

	return true
}
