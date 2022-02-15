package leecode_test

import (
	"fmt"
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

type ListNode struct {
	     Val int
	     Next *ListNode
}



func NewList(ins []int)*ListNode{
	root :=ListNode{}
	head :=&root
	for k, v := range ins {
		head.Val=v
		if k != len(ins)-1{
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return &root
}
//leetcode submit region end(Prohibit modification and deletion)

func reserve(h1 *ListNode)*ListNode{
	var prev, cur *ListNode = nil, h1
	for cur !=nil{
		tmp :=cur.Next
		cur.Next = prev
		prev = cur
		cur = tmp
	}
	return prev
}


func middleNode(head *ListNode) *ListNode {
	// 快慢指针
	slow := head
	fast := head
	for fast!=nil &&fast.Next != nil{
		slow = slow.Next
		fast = fast.Next.Next
	}
	return slow
}

func mergeList(l1, l2 *ListNode) {
	var l1Tmp, l2Tmp *ListNode
	for l1 != nil && l2 != nil {
		l1Tmp = l1.Next
		l2Tmp = l2.Next

		l1.Next = l2
		l1 = l1Tmp

		l2.Next = l1
		l2 = l2Tmp
	}
}

func reorderList(head *ListNode)  {
	if head == nil{
		return
	}

	mid := middleNode(head)
	l2 :=mid.Next
	mid.Next =nil
	rev :=reserve(l2)
	mergeList(head,rev)
	showList(head)

}

func showList( l *ListNode){
	for l!=nil{
		fmt.Println("node is ",l.Val)
		l = l.Next
	}
}

func TestReorderList(t *testing.T){
	ins :=[]int{1,3,4,5,6}
	list :=NewList(ins)
	reorderList(list)
}

//leetcode submit region begin(Prohibit modification and deletion)
// 暴力破解法
// func maxArea(height []int) int {
// 	max :=0
// 	l := len(height)
// 	for i:=0; i<l-1;i++ {
// 		for j:=i+1;j<l;j ++{
// 			h :=min(height[i],height[j])
// 			if max < h *(j -i){
// 				max = h *(j -i)
// 			}
// 		}
// 	}
// 	return max
// }
// 双指针方法
func maxArea(height []int) int {
	max :=0
	t := len(height)-1
	h :=0
	for h < t {
		hg,isLeft :=min(height[h],height[t])
		if max <sq(t-h,hg){
			max =sq(t-h,hg)
		}
		if isLeft {
			h++
		}else{
			t++
		}
	}
	return max
}

func sq(d,h int)int{
	return d*h
}

func min(i,j int)(int,bool){
	if i >j {
		return j,true
	}
	return i,false
}
//leetcode submit region end(Prohibit modification and deletion)


func TestContainerWithMostWater(t *testing.T){
	height :=[]int{8,4,5,6,25,6}
	fmt.Println(maxArea(height))
}
// 回溯算法
var all [][]int
func permute(nums []int) [][]int {
	other := []int{}
	for _,v := range nums {
		other = append(other, v)
	}
	length := len(nums)
	backTrack(length,other)
	return nil
}

func backTrack(len int,output []int){

}

//leetcode submit region end(Prohibit modification and deletion)


func TestPermutations(t *testing.T){
	nums :=[]int{11,2,4,6,8,9}
	makeNums(nums)
	fmt.Println(nums)
}

func makeNums(nums []int){
	nums =append(nums, []int{15,6}...)
}
