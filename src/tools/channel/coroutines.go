package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	// // 取消上下文
	// UseCancel()
	// // 超时上下文
	// UseTimeOut()
	aa:=make(chan int,3)
	aa <-0
	aa<-1
	aa<-2
	for {
		select {
		case a:=<-aa:
			fmt.Println(a)
		default:
			fmt.Println("结束")
			return
		}
	}
}
func UseTimeOut(){

}

func UseSwitch(){

}

func UseCancel(){
	ctx, close := context.WithCancel(context.Background())
	go func(c context.Context) {
		time.Sleep(time.Second*20)
		select  {
		case <-c.Done():
			fmt.Println("协程被关闭")
			return
		default:
		}
		fmt.Println("进行中")

	}(ctx)
	time.Sleep(time.Second*10)
	close()
	fmt.Println("main关闭context")
}
