package leecode_test

import (
	"testing"
)

 type TreeNode struct {
	     Val int
	     Left *TreeNode
	     Right *TreeNode
	 }
func NewTree() *TreeNode {
	root := TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 5}
	left := root.Left
	left.Left = &TreeNode{Val: 3}
	left.Right = &TreeNode{Val: 4}
	return &root
}
func show(node *TreeNode,res *TreeNode, root *TreeNode){

	res.Val=node.Val
	println("val:",node.Val)
	if node.Left !=nil{
		res.Right = &TreeNode{}
		res=res.Right
		show(node.Left,res,root)
	}
	if node.Right != nil{
		res.Right = &TreeNode{}
		res=res.Right
		show(node.Right,res,root)
	}
}
func ShowAll(node *TreeNode){
	println(node.Val)
	if node.Left !=nil{
		ShowAll(node.Left)
	}
	if node.Right != nil{
		ShowAll(node.Right)
	}
}

func flatten(root *TreeNode)  {
	temp := root
	res :=TreeNode{}
	temp2 := &res
	show(temp,temp2,&res)
	ShowAll(&res)

}
//leetcode submit region end(Prohibit modification and deletion)


func TestFlattenBinaryTreeToLinkedList(t *testing.T){
	root :=NewTree()
	flatten(root)
}