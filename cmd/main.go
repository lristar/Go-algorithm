package main

import "Go-algorithm/src/tree"

func main() {
	root:=tree.NewHashRoot("1")
	root.InsertLeft("2")
	root.InsertRight("3")
	root.FindOneByHash("01")
}
