package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

func main() {
	filePath := "/home/lzy/test/aaaa"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for true {
		write.WriteString("1 \r\n")
		write.WriteString("2 \r\n")
		write.Flush()
		time.Sleep(5*time.Second)
	}
	//Flush将缓存的文件真正写入到文件中
}