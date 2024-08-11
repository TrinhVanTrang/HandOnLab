package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: search <directory> <keyword>")
		return
	}

	directory := os.Args[1]
	keyword := os.Args[2]

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(info.Name(), ".txt") {
			searchInFile(path, keyword)
		}
		return nil
	})

	if err != nil {
		fmt.Println("Lỗi truy cập đường dẫn thư mục: ", err)
	}

}

func searchInFile(filename string, keyword string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Lỗi mở file: ", err)
		return
	}
	defer file.Close()

	keywordRegex, err := regexp.Compile(`\b` + regexp.QuoteMeta(keyword) + `\b`)
	if err != nil {
		fmt.Println("Error compiling regex:", err)
		return
	}

	scanner := bufio.NewScanner(file)
	lineNumber := 1
	found := false
	for scanner.Scan() {
		line := scanner.Text()
		if keywordRegex.MatchString(line) {
			if !found {
				fmt.Println(filename)
				found = true
			}
			fmt.Printf("%d: %s\n", lineNumber, line)
		}
		lineNumber++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Lỗi đọc file: ", err)
	}
}
