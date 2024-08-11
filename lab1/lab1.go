package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	// // Nhập tên file
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Nhập tên file muốn đọc: ")
	// filename, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Lỗi đọc tên file: ", err)
	// 	return
	// }

	// filename = strings.TrimSpace(filename)

	if len(os.Args) != 2 {
		fmt.Println("Usage: lab1 <filename>")
		return
	}
	filename := os.Args[1]

	// mở file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Lỗi mở file: ", err)
		return
	}
	defer file.Close()

	// đọc từng dòng
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Lỗi đọc file: ", err)
	}
}
