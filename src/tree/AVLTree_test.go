package tree

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	root := NewNode(2, 0, nil, nil)
	root.Insert(3)
	root.Insert(1)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	fmt.Println(root)
}

func TestMax(t *testing.T) {
	root := NewNode(2, 0, nil, nil)
	root.Insert(3)
	root.Insert(1)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	node := root.left.Min()
	fmt.Println(node.val)
}

func TestMin(t *testing.T) {
	root := NewNode(2, 0, nil, nil)
	root.Insert(3)
	root.Insert(1)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	node := root.left.Max()
	fmt.Println(node.val)
}

func TestDel(t *testing.T) {
	root := NewNode(2, 0, nil, nil)
	root.Insert(3)
	root.Insert(1)
	root.Insert(4)
	root.Insert(5)
	root.Insert(6)
	if root.FindNum(5) == 1 {
		root.Del(5)
	}
	if root.FindNum(6) == 1 {
		root.Del(6)
	}
	fmt.Println(root)
}
