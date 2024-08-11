package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	// Các câu hỏi
	questions := []string{
		"Tên bạn là gì?",
		"Bạn sinh ngày nào?",
		"Bạn làm nghề gì?",
	}

	file, err := os.Create("person.txt")
	if err != nil {
		fmt.Println("Lỗi tạo file:", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	for _, question := range questions {
		fmt.Println(question)
		answer, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Lỗi đọc input: ", err)
			return
		}
		answer = strings.TrimSpace(answer)
		// answers[i] = answer

		_, err = writer.WriteString(question + "\n" + answer + "\n\n")
		if err != nil {
			fmt.Println("Lỗi ghi vào file: ", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Lỗi flush dữ liệu: ", err)
		return
	}

}
