package main

import "Go-algorithm/src/tree"

func main() {
	root:=tree.NewBsRoot(9)
	root.InsertBst(8)
	root.InsertBst(13)
	root.InsertBst(6)
	root.LevelTraversal()
}
