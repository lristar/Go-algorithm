package main

import (
	"Go-algorithm/src/tree"
	"fmt"
)

func main() {
	root := tree.NewRoot("0")
	root.AddLeft("1")
	root.AddRight("2")
	left := root.GetLeft()
	left.AddLeft("3")
	left.AddRight("4")
	//root.LevelTraversal()
	list := make([]string, 20)
	root.DealLayer(list, 0)
	fmt.Println(list)
}
