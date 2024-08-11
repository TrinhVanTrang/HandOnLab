
# Hand On Lab

## Hướng dẫn chạy 
## Lab 1

```bash
cd ./handOnLab
go build lab1/lab1.go  
./lab1 <filename>
```

## Lab 2

```bash
cd ./handOnLab
go build lab2/lab2.go  
./lab2 <filename>
```

## Lab 3

```bash
cd ./handOnLab
go build lab3/lab3.go  
./lab3 <filename> <content>
```
#### Lưu ý
Content phải nằm trong dấu ngoặc kép ("").

## Lab 4

```bash
cd ./handOnLab
go build lab4/lab4.go  
./lab4 
```

## Lab 5

```bash
cd ./handOnLab
go build lab5/lab5.go  
./lab1 <filename> <size_in_MB>
```
Khi tạo file có kích thước vượt quá dung lượng trống của ổ đĩa, chương trình sẽ báo lỗi "There is not enough space on the disk."

Để khắc phục, ta kiểm tra dung lượng trống của ổ đĩa trước khi tạo file.

``` python
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
```

Nếu dung lượng còn lại nhỏ hơn kích thước file cần tạo, chương trình báo lỗi và kết thúc. 

## Lab 6

```bash
cd ./handOnLab
go build lab6/lab6.go  
./lab6 <filename> <line number>
```

## Lab 7

```bash
cd ./handOnLab
go build lab7/search.go  
./search <directory> <keyword>
```
#### Lưu ý
Keyword phải nằm trong dấu ngoặc kép ("").

## Lab 8

```bash
cd ./handOnLab
go build lab8/lab8.go  
./lab8 <directory>
```
