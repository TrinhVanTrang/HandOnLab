package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
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
		fmt.Println("Usage: lab2 <filename>")
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

	// Các từ nhạy cảm cần kiểm tra
	sensitiveWords := []string{"sex", "fuck", "drug", "kill"}

	// đọc từng dòng
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		line := scanner.Text()
		censorLine := censorLine(line, sensitiveWords)
		fmt.Println(censorLine)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Lỗi đọc file: ", err)
	}
}

// Hàm thay thế nguyên âm bằng ký tự '*'
func censorWord(word string) string {
	vowels := regexp.MustCompile(`[aeiouAEIOU]`)
	return vowels.ReplaceAllString(word, "*")
}

// Hàm kiểm duyệt từ nhạy cảm
func censorLine(line string, sensitiveWords []string) string {
	for _, word := range sensitiveWords {
		re := regexp.MustCompile(`\b(?i)` + word + `\b`)
		line = re.ReplaceAllStringFunc(line, censorWord)
	}
	return line
}
