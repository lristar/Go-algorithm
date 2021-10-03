package main

import (
	"Go-algorithm/src/sort"
	"fmt"
	"time"
)

func main() {
	aa := []int{31, 47, 9, 5, 30, 72, 15, 1, 2, 3, 5, 7, 6, 100}
	fmt.Println(sort.HeapSort(aa, 0))
	time.Sleep(time.Second*1)
}
