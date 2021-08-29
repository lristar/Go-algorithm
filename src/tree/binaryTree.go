package tree

import "fmt"

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

//非递归方法
