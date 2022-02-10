package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

var wg sync.WaitGroup
var in = make(chan int, 10)
var quit = make(chan int, 1)
var quito0k = make(chan int, 1)

func ListenQuit() {
	defer wg.Done()
	var sign = make(chan os.Signal, 1)
	signal.Notify(sign, syscall.SIGINT)
	<-sign
	SQuit()

}

func run() {
	defer wg.Done()
Quit:
	for {
		select {
		case out := <-in:
			fmt.Println(out)
		case <-quit:
			time.Sleep(time.Second * 1)
			break Quit
		}
	}

	if len(in) > 0 { //处理完通道中数据
		for i := 0; i < len(in); i++ {
			<-in
			fmt.Println("11111")
		}
	}

}

func producerin() {
Quit:
	for {
		select {
		case in <- 0:
		case <-quit: //停止写数据
			break Quit
		}
	}

}

func SQuit() {
	close(quit)
	go func() {
		wg.Wait()
		quito0k <- 0
	}()
}
func main() {
	go ListenQuit()
	go func() {
		for i := 0; i < 10; i++ {
			wg.Add(1)
			go run()
		}
	}()
	wg.Add(1)
	go producerin()
	<-quito0k
	fmt.Println("qiut ok")

}
