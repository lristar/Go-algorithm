package tree

import (
	"fmt"
)

type TreeNode struct {
	val   string
	left  *TreeNode
	right *TreeNode
}

func NewRoot(val string) *TreeNode {
	return &TreeNode{val: val}
}

func (t *TreeNode) AddLeft(val string) {
	t.left = &TreeNode{val: val}
}

func (t *TreeNode) AddRight(val string) {
	t.right = &TreeNode{val: val}
}

func (t *TreeNode) GetLeft() *TreeNode {
	return t.left
}

func (t *TreeNode) GetRight() *TreeNode {
	return t.right
}

//递归方法
//前序遍历，根左右
func (t *TreeNode) PreOrder() {
	if t == nil {
		return
	}
	fmt.Print(t.val)
	t.left.PreOrder()
	t.right.PreOrder()
}

//中序遍历
func (t *TreeNode) MidOrder() {
	if t == nil {
		return
	}
	t.left.PreOrder()
	fmt.Print(t.val)
	t.right.PreOrder()
}

//后序遍历
func (t *TreeNode) PostOrder() {
	if t == nil {
		return
	}
	t.left.PreOrder()
	t.right.PreOrder()
	fmt.Print(t.val)
}

//层次遍历 递归版
func (t *TreeNode) Layer(res []string, level int) {
	if t == nil {
		return
	}
	if t.left != nil {
		t.left.Layer(res, 2*level)
	}
	if t.right != nil {
		t.right.Layer(res, 2*level+1)
	}
	res[level] = t.val
}

// 处理递归版层次遍历
func (t *TreeNode) DealLayer(res []string, level int) {
	t.Layer(res, level+1)
	res = res[1:]
}

// 层次遍历 非递归
func (t *TreeNode) LevelTraversal() {
	if t == nil {
		return
	}
	result := []string{}
	nodes := []*TreeNode{t}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curNode.val)
		if curNode.left != nil {
			nodes = append(nodes, curNode.left)
		}
		if curNode.right != nil {
			nodes = append(nodes, curNode.right)
		}
	}
	for e := range result {
		fmt.Print(e, " ")
	}
}
