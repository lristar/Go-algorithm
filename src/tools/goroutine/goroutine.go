package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	cStop chan int
	exitSign  chan int
)

func main() {
	exitSign = make(chan int, 1)
	go GetSign()
	select {
	case <-exitSign:
		fmt.Println("优雅关机")
	}
	fmt.Println("结束")
	// ctx := context.Background()
	// go goroutineMain(ctx)
	// cStop=make(chan int,1)
	// cStop<-0
	// for true{
	// 	fmt.Println("hahaah")
	// 	time.Sleep(time.Second*10)
	// }

}

func GetSign() {
	sign := make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGTERM, syscall.SIGINT, syscall.SIGHUP)
	fmt.Println("hahah",<-sign)
	exitSign<-0
}

func goroutineMain(ctx context.Context) {
	fmt.Println("进入了")
	context.WithCancel(ctx)
	for true {
		select {
		case <-cStop:
			fmt.Println("启动子协程")
			goroutineChildren(cStop)
		}
	}
}

// func goroutineChildren(c chan<- int) {
// 	cStop<-0
// 	return
// }

func goroutineChildren(c chan<- int) {
	filePath := "/home/lzy/test/aaaa"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	//及时关闭file句柄
	defer file.Close()
	//写入文件时，使用带缓存的 *Writer
	write := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		write.WriteString(strconv.Itoa(i) + " \r\n")
		write.Flush()
		time.Sleep(5 * time.Second)
	}
	write.WriteString("end" + " \r\n")
	write.Flush()
	cStop <- 0
	return
}
