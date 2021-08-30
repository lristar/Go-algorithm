package tree

import "fmt"

type BSTreeNode struct {
	val   int
	left  *BSTreeNode
	right *BSTreeNode
}

//创建根节点
func NewBsRoot(val int) *BSTreeNode {
	return &BSTreeNode{val: val}
}

//插入新数据 递归版本
func (b *BSTreeNode) InsertBst(val int) {
	if b.val < val {
		if b.right == nil {
			b.right = &BSTreeNode{val: val}
			return
		}
		b.right.InsertBst(val)
	}
	if b.val > val {
		if b.left == nil {
			b.left = &BSTreeNode{val: val}
			return
		}
		b.left.InsertBst(val)
	}
}

//插入新数据，非递归版本
func (b *BSTreeNode) Insert(val int) {
	for true {
		if b.val > val {
			if b.left == nil {
				b.left = &BSTreeNode{val: val}
				return
			}
			b = b.left
		} else if b.val < val {
			if b.right == nil {
				b.right = &BSTreeNode{val: val}
				return
			}
			b = b.right
		} else {
			fmt.Print("元素已存在")
			return
		}
	}
}

//查找数字
func (b *BSTreeNode) FindNode(val int) {
	if b == nil {
		fmt.Print("没有这个数值")
		return
	}
	if b.val == val {
       fmt.Print(b)
	}
	if b.val > val {
		b.left.FindNode(val)
	}
	if b.val < val {
		b.right.FindNode(val)
	}
}

//层次遍历
func (t *BSTreeNode) LevelTraversal() {
	if t == nil {
		return
	}
	result := []int{}
	nodes := []*BSTreeNode{t}
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
	for _, v := range result {
		fmt.Print(v, " ")
	}
}
