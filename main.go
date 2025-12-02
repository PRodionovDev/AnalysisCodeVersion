package main

import (
	"bufio"
	"encoding/csv"
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

	searchWords := getSearchWord()

	var dirs []string
	dirs = getDirectories(catalog)

	for _, dir := range dirs {
		checkFiles(dir, searchWords)
	}
}

func getSearchWord() []string {
	searchCsv, err := os.Open("search.csv")
	if err != nil {
		log.Fatal("Error couldn't open file:", err)
	}
	defer searchCsv.Close()

	reader := csv.NewReader(searchCsv)

	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal("Error couldn't read CSV:", err)
	}

	var searchWords []string
	for _, record := range records {
		for _, value := range record {
			searchWords = append(searchWords, value)
		}
	}

	return searchWords
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

func checkFiles(path string, searchWords []string) {
	phpFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, phpFile := range phpFiles {
		if !phpFile.IsDir() && strings.HasSuffix(strings.ToLower(phpFile.Name()), ".php") {
			// todo: gorutine
			checkFile(path+"/"+phpFile.Name(), searchWords)
		}
	}
}

func checkFile(path string, searchWords []string) bool {
	file, err := os.Open(path)
	if err != nil {
		fmt.Println("ERROR couldn't open file:", err)
		return false
	}
	defer file.Close()

	var errorFiles []string

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		wordsInLine := strings.Fields(line)

		for _, word := range wordsInLine {
			for _, searchWord := range searchWords {
				if strings.Contains(word, searchWord) {
					errorFiles = append(errorFiles, path)
				}
			}
		}
	}
	if len(errorFiles) > 0 {
		fmt.Printf("Find %d errors in '%s'\n", len(errorFiles), path)
	}

	return true
}
