package tree

type AvlTree struct {
	val    int
	height int
	left   *AvlTree
	right  *AvlTree
}

func NewNode(val, height int, left, right *AvlTree) *AvlTree {
	return &AvlTree{
		val:    val,
		height: height,
		left:   left,
		right:  right,
	}
}

// 创建树跟二叉排序树是一样的规则，但创建的过程中有一个检查和修正的过程
func (b *AvlTree) Insert(val int) {
	if b.val == val {
		return
	} else if b.val > val {
		if b.left == nil {
			b.left = NewNode(val, 0, nil, nil)
		} else {
			b.left.Insert(val)
		}
	} else if b.val < val {
		if b.right == nil {
			b.right = NewNode(val, 0, nil, nil)
		} else {
			b.right.Insert(val)
		}
	}
	// 计算高度
	leftHeight := Height(b.left, 0)
	rightHeight := Height(b.right, 0)
	balanceF := rightHeight - leftHeight
	if balanceF >= 2 {
		if b.right.val < val {
			b.LeftRotate()
		} else {
			b.right.RightRotate()
			b.LeftRotate()
		}
	} else if balanceF <= -2 {
		if b.left.val > val {
			b.RightRotate()
		} else {
			b.left.LeftRotate()
			b.RightRotate()
		}
	}
}

func (b *AvlTree) Del(val int) (isOne bool) {
	if b.val < val {
		if b.right.Del(val) {
			b.right = nil
		}
		// 平衡
		balanceF := b.Diff()
		if balanceF < -1 {
			lBalance := b.left.Diff()
			if lBalance < 0 {
				b.RightRotate()
			} else {
				b.left.LeftRotate()
				b.RightRotate()
			}
		}
	} else if b.val > val {
		if b.left.Del(val) {
			b.left = nil
		}
		// 平衡
		balanceF := b.Diff()
		if balanceF > 1 {
			rBalance := b.right.Diff()
			if rBalance > 0 {
				b.LeftRotate()
			} else {
				b.right.RightRotate()
				b.LeftRotate()
			}
		}
	} else {
		// 这里要进行补充节点
		// 计算高度
		balanceF := b.Diff()
		if balanceF > 0 {
			if b.right != nil {
				min := b.right.Min()
				if min.val == b.right.val {
					// 说明最小节点是该节点的右节点
					b.right = b.right.right
					b.val = min.val
				} else {
					b.val = min.val
				}
			} else {
				isOne = true
			}
		} else {
			if b.left != nil {
				max := b.left.Max()
				if max.val == b.left.val {
					b.left = b.left.left
					b.val = max.val
				} else {
					b.val = max.val
				}
			} else {
				isOne = true
			}
		}
	}
	return
}

func (b *AvlTree) FindNum(val int) int {
	if b.val > val {
		if b.left == nil {
			return 0
		}
		return b.left.FindNum(val)
	} else if b.val < val {
		if b.right == nil {
			return 0
		}
		return b.right.FindNum(val)
	} else {
		return 1
	}
}

// RightRotate 右旋
func (b *AvlTree) RightRotate() {
	lNode := b.left
	newRootVal := lNode.val
	b.right = NewNode(b.val, 0, lNode.right, b.right)
	b.left = lNode.left
	b.val = newRootVal
}

// LeftRotate 左旋
func (b *AvlTree) LeftRotate() {
	rNode := b.right
	newRootVal := rNode.val
	b.left = NewNode(b.val, 0, b.left, rNode.left)
	b.right = rNode.right
	b.val = newRootVal
}

// 遍历最小值
func (b *AvlTree) Min() *AvlTree {
	if b.left != nil {
		if b.left.left == nil {
			min := b.left
			b.left = nil
			return min
		} else {
			return b.left.Min()
		}
	} else {
		return b
	}
}

func (b *AvlTree) Max() *AvlTree {
	if b.right != nil {
		if b.right.right == nil {
			max := b.right
			b.right = nil
			return max
		} else {
			return b.right.Max()
		}
	} else {
		return b
	}
}

// Height 获取树的高度
func Height(b *AvlTree, max int) int {
	if b == nil {
		return 0
	}
	return maxHeight(Height(b.left, max), Height(b.right, max)) + 1
}

func maxHeight(x, y int) int {
	if x >= y {
		return x
	}
	return y
}

func (b *AvlTree) Diff() int {
	leftHeight := Height(b.left, 0)
	rightHeight := Height(b.right, 0)
	return rightHeight - leftHeight
}
