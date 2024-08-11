package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: lab8 <directory> ")
		return
	}

	root := os.Args[1]
	rootInfo, err := os.Stat(root)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if !rootInfo.IsDir() {
		fmt.Println("Error: specified path is not a directory")
		return
	}

	fmt.Println(rootInfo.Name())
	files, err := os.ReadDir(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	for i, file := range files {
		path := filepath.Join(root, file.Name())
		if i == len(files)-1 {
			printTree(path, "", true)
		} else {
			printTree(path, "", false)
		}
	}

}

func printTree(root string, indent string, last bool) {
	fileInfo, err := os.Stat(root)
	if err != nil {
		fmt.Println(err)
		return
	}

	if fileInfo.IsDir() {
		fmt.Println(indent + "├── " + fileInfo.Name())
		files, err := os.ReadDir(root)
		if err != nil {
			fmt.Println(err)
			return
		}

		for i, file := range files {
			path := filepath.Join(root, file.Name())
			if i == len(files)-1 {
				printTree(path, indent+"│   ", true)
			} else {
				printTree(path, indent+"│   ", false)
			}
		}
	} else {
		if last {
			fmt.Println(indent + "└── " + fileInfo.Name())
		} else {
			fmt.Println(indent + "├── " + fileInfo.Name())
		}
	}
}
