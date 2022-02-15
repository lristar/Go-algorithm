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

//leetcode submit region end(Prohibit modification and deletion)

func TestReverseInteger(t *testing.T) {

}
