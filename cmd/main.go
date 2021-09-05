package main

import (
	"Go-algorithm/src/sort"
	"fmt"
)

func main() {
	aa:=[]int{31,47,9,5,30,72,15}
	sort.MergeSort(0,6,aa)
	fmt.Println(aa)
}
