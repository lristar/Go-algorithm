package main

import (
	"Go-algorithm/src/tree"
	"fmt"
)

func main() {
	root:=tree.NewRoot("0")
	root.AddLeft("1")
	root.AddRight("2")
	root.PreOrder()
	fmt.Println()
	root.MidOrder()
	fmt.Println()
	root.PostOrder()
}
