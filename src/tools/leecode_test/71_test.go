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

// func show(node *TreeNode,res *TreeNode, root *TreeNode){
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
// }
// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)

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
	for i := 0; i < mid; i++ {
		for s[i] != s[l-i-1] {
			return false
		}
	}
	return true
}

func TestLongestPalindromicSubstring(t *testing.T) {
	// fmt.Printf("是不是回文:%v", Judge_str("aafaafaa"))
}

// leetcode submit region begin(Prohibit modification and deletion)
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

// func TestFlattenBinaryTreeToLinkedList(t *testing.T) {
//	root := NewTree()
//	flatten(root)
// }

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

// leetcode submit region end(Prohibit modification and deletion)

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

// leetcode submit region begin(Prohibit modification and deletion)
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

// leetcode submit region end(Prohibit modification and deletion)

func TestContainerWithMostWater(t *testing.T) {
	height := []int{8, 4, 5, 6, 25, 6}
	fmt.Println(maxArea(height))
}

// 回溯算法
// leetcode submit region begin(Prohibit modification and deletion)
var all [][]int

func permute(nums []int) [][]int {
	all = [][]int{}
	used := make([]int, 0, len(nums))
	sort.Ints(nums)
	backTrack(len(nums), nums, []int{}, used, 0)
	return all
}

func backTrack(length int, output []int, path []int, used []int, useNum int) {
	if length == 0 {
		// 要新开辟一个int数组，新的地址 保证all内的数据不会被修改
		p := make([]int, len(path))
		copy(p, path)
		all = append(all, p)
	}
	for i := 0; i < length; i++ {
		if len(used) < useNum+1 || used[useNum] != output[i] {
			curNum := output[i]
			if len(used) < useNum+1 {
				used = append(used, curNum)
			} else {
				used[useNum] = curNum
			}
			path = append(path, curNum)
			output = append(output[0:i], output[i+1:]...)
			backTrack(len(output), output, path, used, useNum+1)
			output = append(output[:i], append([]int{curNum}, output[i:]...)...)
			path = path[:len(path)-1]
		}
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func TestPermutations(t *testing.T) {
	fmt.Println(permute([]int{1, 1, 2}))
}

// 112题
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// 展示这个Tree
func ShowTree(root *TreeNode) {
	if root == nil {
		return
	}
	fmt.Printf("node is %d", root.Val)
	ShowTree(root.Left)
	ShowTree(root.Right)
	fmt.Println(".")
}

// [5,4,8,11,null,13,4,7,2,null,null,null,1]
// 使用堆构建二叉树
func NewTreeAccordingHeap(heap []int) *TreeNode {
	root := &TreeNode{}
	CreateTree(root, heap, 0)
	return root
}

func CreateTree(root *TreeNode, heap []int, num int) {
	root.Val = heap[num]
	if len(heap) > 2*num+1 {
		if heap[2*num+1] != -1 {
			root.Left = &TreeNode{}
			CreateTree(root.Left, heap, 2*num+1)
		}
	}
	if len(heap) > 2*num+2 {
		if heap[2*num+2] != -1 {
			root.Right = &TreeNode{}
			CreateTree(root.Right, heap, 2*num+2)
		}
	}
}

var isTrue = false

func hasPathSum(root *TreeNode, targetSum int) bool {
	isTrue = false
	backTrackPath(root, 0, targetSum)
	return isTrue
}

func backTrackPath(root *TreeNode, preSum int, targetSum int) {
	if root == nil {
		return
	}
	if root.Left == nil && root.Right == nil {
		if preSum+root.Val == targetSum {
			isTrue = true
		}
		return
	}
	backTrackPath(root.Left, preSum+root.Val, targetSum)
	backTrackPath(root.Right, preSum+root.Val, targetSum)
}

// leetcode submit region end(Prohibit modification and deletion)

func TestPathSum(t *testing.T) {
	// fmt.Println(hasPathSum(NewTree(),6))
	root := NewTreeAccordingHeap([]int{1, 2, 3})
	fmt.Println(hasPathSum(root, 22))
}

func preorderTraversal(root *TreeNode) (nums []int) {
	preorder(root, &nums)
	return nums
}

func preorder(root *TreeNode, nums *[]int) {
	if root == nil {
		return
	}
	*nums = append(*nums, root.Val)
	preorder(root.Left, nums)
	preorder(root.Right, nums)
}

// 中序遍历
func inorderTraversal(root *TreeNode) (nodes []int) {
	inoTraversal(root, &nodes)
	return nodes
}

func inoTraversal(root *TreeNode, node *[]int) {
	if root == nil {
		return
	}

	inoTraversal(root.Left, node)
	*node = append(*node, root.Val)
	inoTraversal(root.Right, node)
}
func postorderTraversal(root *TreeNode) (val []int) {
	postOrder(root, &val)
	return val
}

func postOrder(root *TreeNode, result *[]int) {
	if root == nil {
		return
	}
	postOrder(root.Left, result)
	postOrder(root.Right, result)
	*result = append(*result, root.Val)
}

// leetcode submit region end(Prohibit modification and deletion)

func TestBinaryTreeInorderTraversal(t *testing.T) {
	root := NewTreeAccordingHeap([]int{1, -1, 2, 3})
	fmt.Println(preorderTraversal(root))
}

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return createTree(1, n)
}

func createTree(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	allTree := []*TreeNode{}
	for i := start; i <= end; i++ {
		lTree := createTree(start, i-1)
		rTree := createTree(i+1, end)
		for _, left := range lTree {
			for _, right := range rTree {
				curTree := &TreeNode{i, nil, nil}
				curTree.Left = left
				curTree.Right = right
				allTree = append(allTree, curTree)
			}
		}
	}
	return allTree
}

func TestUniqueBinarySearchTreesIi(t *testing.T) {
	n := generateTrees(2)
	fmt.Println(n)
	fmt.Println("hahaha")
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	nums := selectNum(digits)
	rs := []string{}
	addNums(nums, 0, &rs, "")
	return rs
}

func addNums(nums [][]string, i int, rs *[]string, str string) {
	if len(nums) == i {
		*rs = append(*rs, str)
		return
	}
	for _, v := range nums[i] {
		addNums(nums, i+1, rs, str+v)
	}
}

func selectNum(digits string) [][]string {
	nums := [][]string{}
	for _, num := range digits {
		switch num {
		case 50:
			nums = append(nums, []string{"a", "b", "c"})
		case 51:
			nums = append(nums, []string{"d", "e", "f"})
		case 52:
			nums = append(nums, []string{"g", "h", "i"})
		case 53:
			nums = append(nums, []string{"j", "k", "l"})
		case 54:
			nums = append(nums, []string{"m", "n", "o"})
		case 55:
			nums = append(nums, []string{"p", "q", "r", "s"})
		case 56:
			nums = append(nums, []string{"t", "u", "v"})
		case 57:
			nums = append(nums, []string{"w", "x", "y", "z"})
		}
	}
	return nums
}

func TestCombinations(t *testing.T) {
	//fmt.Println(letterCombinations())
	result := letterCombinations("23")
	fmt.Println(result)
}

// 32题最长有效字串
func longestValidParentheses(s string) int {
	// 定义两个数组作为栈
	max := 0
	mid := make([]rune, 0)
	mid = append(mid, -1)
	for k, v := range s {
		if v == 40 {
			mid = append(mid, rune(k))
		} else {
			mid = mid[:len(mid)-1]
			if len(mid) < 1 {
				mid = append(mid, rune(k))
			} else {
				before := mid[len(mid)-1]
				length := k - int(before)
				if max < length {
					max = length
				}
			}
		}
	}

	return max
}

func Test_longestValidParentheses(t *testing.T) {
	fmt.Println("len is ", longestValidParentheses(")()((())"))
}

func judgeMin(a []int) int {
	length := len(a)
	if length == 1 {
		return a[0]
	}
	if length == 2 {
		if a[0] > a[1] {
			return a[1]
		}
		return a[0]
	}
	mid := length / 2
	a1 := judgeMin(a[0:mid])
	a2 := judgeMin(a[mid:])
	if a1 > a2 {
		return a2
	}
	return a1
}

func Test_judgeMin(t *testing.T) {
	aa := []int{1, 4325, 6, 4, 7, 6}
	fmt.Println(judgeMin(aa))
}

func trap(height []int) int {
	capacity := 0
	stack := make([]int, 0)
	for _, v := range height {
		if len(stack) == 0 || stack[len(stack)-1] > v {
			stack = append(stack, v)
		} else {
			i :=len(stack) -1
			for i >=0 {
				if stack[len(stack)-1] >= v {
					break
				}else{
					mid:=stack[i]
					i--
					h := getmin(stack[i],v)
					capacity += h-mid
				}
			}
		}
	}
	return capacity
}
func getmin(a,b int)int{
	if a >b {
		return b
	}
	return a

}

func Test_trap(t *testing.T) {
	fmt.Println("trap is", trap([]int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}))

}
