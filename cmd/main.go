package main

import "Go-algorithm/src/tree"

func main() {
	root:=tree.NewBsRoot(9)
	root.Insert(8)
	root.Insert(13)
	root.Insert(6)
	root.FindNode(6)
}
