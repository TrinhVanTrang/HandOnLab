package main

import (
	"bufio"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"os"
	"strconv"
	"syscall"
	"unsafe"
)

func main() {
	// reader := bufio.NewReader(os.Stdin)
	// fmt.Print("Nhập tên file muốn tạo: ")
	// filename, err := reader.ReadString('\n')
	// if err != nil {
	// 	fmt.Println("Lỗi đọc tên file: ", err)
	// 	return
	// }

	// filename = strings.TrimSpace(filename)

	if len(os.Args) != 3 {
		fmt.Println("Usage: lab2 <filename> <size_in_MB>")
		return
	}
	filename := os.Args[1]

	// Mở hoặc tạo file
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Lỗi tạo file: ", err)
		return
	}
	defer file.Close()

	writer := bufio.NewWriter(file)

	// fmt.Print("Nhập kích thước file(Mb, lớn hơn 10 Mb): ")
	// var size int
	// _, err = fmt.Scanf("%d", &size)
	// if err != nil {
	// 	fmt.Println("Lỗi đọc kích thước file: ", err)
	// 	return
	// }

	size, err := strconv.Atoi(os.Args[2])
	if err != nil || size <= 10 {
		fmt.Println("Lỗi: Kích thước file phải là số nguyên dương lớn hơn 10 (MB).")
		return
	}

	fileSize := size * 1024 * 1024
	lineLength := 256
	linesNeeded := fileSize / lineLength

	// Kiểm tra dung lượng đĩa trống
	hasSpace, err := checkDiskSpace("/", int64(fileSize))
	if err != nil {
		fmt.Println("Error checking disk space:", err)
		return
	}
	if !hasSpace {
		fmt.Println("Error: Not enough disk space available")
		return
	}

	for i := 0; i < linesNeeded; i++ {
		line, err := generateRandomLine(lineLength)
		if err != nil {
			fmt.Println("Lỗi tạo dòng ký tự ngẫu nhiên: ", err)
			return
		}
		_, err = writer.WriteString(line + "\n")
		if err != nil {
			fmt.Println("Lỗi ghi file: ", err)
			return
		}
	}

	err = writer.Flush()
	if err != nil {
		fmt.Println("Lỗi flush dữ liệu: ", err)
		return
	}

}

func generateRandomLine(length int) (string, error) {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(bytes)[:length], nil
}

// Hàm kiểm tra dung lượng đĩa trống
func checkDiskSpace(path string, requiredSpace int64) (bool, error) {
	// if runtime.GOOS == "windows" {
	return checkDiskSpaceWindows(path, requiredSpace)
	// }
	// return checkDiskSpaceUnix(path, requiredSpace)
}

// // Kiểm tra dung lượng đĩa trên Unix (Linux, macOS)
// func checkDiskSpaceUnix(path string, requiredSpace int64) (bool, error) {
// 	var stat unix.Statfs_t
// 	err := unix.Statfs(path, &stat)
// 	if err != nil {
// 		return false, err
// 	}
// 	freeSpace := int64(stat.Bavail) * int64(stat.Bsize)
// 	return freeSpace >= requiredSpace, nil
// }

// Kiểm tra dung lượng đĩa trên Windows
func checkDiskSpaceWindows(path string, requiredSpace int64) (bool, error) {
	h := syscall.MustLoadDLL("kernel32.dll")
	c := h.MustFindProc("GetDiskFreeSpaceExW")
	lpFreeBytesAvailable := int64(0)
	_, _, err := c.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(path))),
		uintptr(unsafe.Pointer(&lpFreeBytesAvailable)),
		uintptr(0), uintptr(0))
	if err != nil && err.Error() != "The operation completed successfully." {
		return false, err
	}
	return lpFreeBytesAvailable >= requiredSpace, nil
}
