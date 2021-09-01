package tree

import (
	"errors"
	"fmt"
)

type AvlTree struct {
	val    int
	height int
	left   *AvlTree
	right  *AvlTree
}

//创建树跟二叉排序树是一样的规则，但创建的过程中有一个检查和修正的过程
func (b *AvlTree) Insert(val int) {
	for true {
		if b.val > val {
			if b.left == nil {
				b.left = &AvlTree{val: val, height: b.height + 1}
				return
			}
			b = b.left
		} else if b.val < val {
			if b.right == nil {
				b.right = &AvlTree{val: val, height: b.height + 1}
				return
			}
			b = b.right
		} else {
			fmt.Print("元素已存在")
			return
		}
	}
}

//检查是否有绝对值大于1的节点

//判断是哪种类型的平衡方式

//LL型
//RR型
//LR型
//RL型

// 获得节点高度
func (b *AvlTree) GetHeight()(int,error) {
	if b == nil {
       return 0,errors.New("不存在")
	}
	return b.height,nil
}

//计算平衡因子
