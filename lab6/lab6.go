package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Nhập tên file muốn đọc: ")
	// filename, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Lỗi đọc tên file: ", err)
	// 	return
	// }
	// filename = strings.TrimSpace(filename)

	if len(os.Args) != 3 {
		fmt.Println("Usage: lab6 <filename> <line number>")
		return
	}
	filename := os.Args[1]

	// Mở file
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("Lỗi mở file: ", err)
		return
	}
	defer file.Close()

	// Kiểm tra kích thước file
	maxFileSize := int64(2 * 1024 * 1024 * 1024) // 2GB
	fileSize, err := getFileSize(filename)
	if err != nil {
		fmt.Println("lỗi lấy kích thước file: ", err)
		return
	}
	if fileSize > maxFileSize {
		fmt.Println("Lỗi: Kích thước file vượt quá giới hạn 2G.")
		return
	}

	// fmt.Print("Nhập dòng muốn đọc: ")
	// var lineNumber int
	// _, err = fmt.Scanf("%d", &lineNumber)
	// if err != nil {
	// 	fmt.Println("Lỗi đọc số thứ tự dòng: ", err)
	// 	return
	// }

	lineNumber, err := strconv.Atoi(os.Args[2])
	if err != nil || lineNumber <= 0 {
		fmt.Println("Vui lòng nhập một số nguyên dương hợp lệ cho dòng cần đọc.")
		return
	}

	// Đọc file dòng theo dòng
	scanner := bufio.NewScanner(file)
	currentLine := 1
	for scanner.Scan() {
		if currentLine == lineNumber {
			fmt.Println("Dòng", lineNumber, ": ", scanner.Text())
			return
		}
		currentLine++
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Lỗi đọc file: ", err)
		return
	}

	if currentLine < lineNumber {
		fmt.Println("Lỗi: File không có dòng", lineNumber)
	}

}

// Hàm kiểm tra kích thước file
func getFileSize(filename string) (int64, error) {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		return 0, err
	}
	return fileInfo.Size(), nil
}
