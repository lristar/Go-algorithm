package leecode_test

import (
	"fmt"
	"sort"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
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

//func show(node *TreeNode,res *TreeNode, root *TreeNode){
//	res.Val=node.Val
//	if node.Left !=nil{
//		res.Right = &TreeNode{}
//		res=res.Right
//		show(node.Left,res,root)
//	}
//
//	if node.Right != nil{
//		res.Right = &TreeNode{}
//		res=res.Right
//		show(node.Right,res,root)
//	}
//}
//leetcode submit region begin(Prohibit modification and deletion)
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)
	fmt.Println(nums1)
	i := len(nums1)
	if i%2 == 0 {
		return (float64(nums1[i/2]) + float64(nums1[(i/2)+1])) / 2
	}
	return float64(nums1[i/2])
}

//leetcode submit region end(Prohibit modification and deletion)

func TestMedianOfTwoSortedArrays(t *testing.T) {
	num1 := []int{1, 2, 3, 5, 8}
	num2 := []int{2, 4, 6, 7, 10}
	fmt.Println(findMedianSortedArrays(num1, num2))
}

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	// 我进行反着加
	tail := 0
	v := x
	for v != 0 {
		tail = tail*10 + v%10
		v = v / 10
	}
	return tail == x
}

func TestPalindromeNumber(t *testing.T) {
	fmt.Printf("是不是回文:%v", isPalindrome(1010101))
}

func longestPalindrome(s string) string {
	l := len(s)
	begin := 0
	maxlen := 0
	for i := 0; i < l-1; i++ {
		for j := i; j < l; j++ {
			if maxlen < j-i+1 && Judge_str(s[i:j+1]) {
				maxlen = j - i + 1
				begin = i
			}
		}
	}
	return s[begin : begin+maxlen]
}

// 判断是不是回文串
func Judge_str(s string) bool {
	l := len(s)
	mid := l/2 + 1
	if l%2 == 0 {
		mid = l / 2
	}
	for i := 0; i < mid; i ++ {
		for s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func TestLongestPalindromicSubstring(t *testing.T) {
	//fmt.Printf("是不是回文:%v", Judge_str("aafaafaa"))
}

//leetcode submit region begin(Prohibit modification and deletion)
func reverse(x int) int {
	// 我进行反着加
	tail := 0
	v := x
	for v != 0 {
		tail = tail*10 + v%10
		v = v / 10
	}
	return tail
}

//func TestFlattenBinaryTreeToLinkedList(t *testing.T) {
//	root := NewTree()
//	flatten(root)
//}

type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(ins []int) *ListNode {
	root := ListNode{}
	head := &root
	for k, v := range ins {
		head.Val = v
		if k != len(ins)-1 {
			head.Next = &ListNode{}
			head = head.Next
		}
	}
	return &root
}

//leetcode submit region end(Prohibit modification and deletion)

func reserve(h1 *ListNode) *ListNode {
	var prev, cur *ListNode = nil, h1
	for cur != nil {
		tmp := cur.Next
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
	for fast != nil && fast.Next != nil {
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

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	mid := middleNode(head)
	l2 := mid.Next
	mid.Next = nil
	rev := reserve(l2)
	mergeList(head, rev)
	showList(head)

}

func showList(l *ListNode) {
	for l != nil {
		fmt.Println("node is ", l.Val)
		l = l.Next
	}
}

func TestReorderList(t *testing.T) {
	ins := []int{1, 3, 4, 5, 6}
	list := NewList(ins)
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
	max := 0
	t := len(height) - 1
	h := 0
	for h < t {
		hg, isLeft := min(height[h], height[t])
		if max < sq(t-h, hg) {
			max = sq(t-h, hg)
		}
		if isLeft {
			h++
		} else {
			t++
		}
	}
	return max
}

func sq(d, h int) int {
	return d * h
}

func min(i, j int) (int, bool) {
	if i > j {
		return j, true
	}
	return i, false
}

//leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {
	height := []int{8, 4, 5, 6, 25, 6}
	fmt.Println(maxArea(height))
}

// 回溯算法
//leetcode submit region begin(Prohibit modification and deletion)
var all [][]int
func permute(nums []int) [][]int {
	all = [][]int{}
	backTrack(len(nums),nums,[]int{})
	return all
}
func backTrack(length int, output []int, path []int)  {
	if length == 0 {
		// 要新开辟一个int数组，新的地址 保证all内的数据不会被修改
		p:=make([]int,len(path))
		copy(p,path)
		all = append(all,p)
	}
	for i:=0;i < length;i ++{
		curNum :=output[i]
		path= append(path, curNum)
		output = append(output[0:i], output[i+1:]...)
		backTrack(len(output),output,path)
		output = append(output[:i], append([]int{curNum},output[i:]...)...)
		path = path[:len(path)-1]
	}
}


//leetcode submit region end(Prohibit modification and deletion)


func TestPermutations(t *testing.T){
	fmt.Println(permute([]int{5,4,6,2}))
}