package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var x int64
var wg sync.WaitGroup
var lock sync.Mutex

func add() {
	for i := 0; i < 1000000000; i++ {
		// lock.Lock() // 加锁
		// x = x + 1
		// lock.Unlock() // 解锁
		atomicAdd()
	}
	wg.Done()
}
func atomicAdd() {
	atomic.AddInt64(&x, 1)
}
func main() {
	// start:=time.Now().Unix()
	// wg.Add(2)
	// go add()
	// go add()
	// wg.Wait()
	// fmt.Println(x)
	// end:=time.Now().Unix()
	// fmt.Println("运行了",end-start,"s")
	aa:=make(map[string]string)
	aa["aaaa"]="1"

	fmt.Println(aa)
}