package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// // Nhập tên file
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Nhập tên file muốn ghi: ")
	// filename, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Lỗi đọc tên file: ", err)
	// 	return
	// }

	// filename = strings.TrimSpace(filename)

	if len(os.Args) != 3 {
		fmt.Println("Usage: lab3 <filename> <content>")
		return
	}
	filename := os.Args[1]

	// Tạo hoặc mở file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Lỗi tạo file: ", err)
		return
	}
	defer file.Close()

	// fmt.Print("Nhập nội dung muốn ghi vào file: ")
	// content, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Lỗi đọc nội dung: ", err)
	// 	return
	// }

	// content = strings.TrimSpace(content)
	content := os.Args[2]
	writer := bufio.NewWriter(file)

	// Ghi nội dung vào file
	_, err = writer.WriteString(content)
	if err != nil {
		fmt.Println("Lỗi ghi file:", err)
		return
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Lỗi flush dữ liệu:", err)
		return
	}

}
