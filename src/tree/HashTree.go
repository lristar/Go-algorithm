package tree

import "fmt"

const (
	leftNo  = 48
	RightNo = 49
)

type HashTree struct {
	hash  string
	val   string
	left  *HashTree
	right *HashTree
}

func NewHashRoot(val string) *HashTree {
	return &HashTree{val: val}
}

func (h *HashTree) InsertLeft(val string) {
	h.left = &HashTree{val: val}
	h.left.hash = h.hash + "0"
}

func (h *HashTree) InsertRight(val string) {
	h.right = &HashTree{val: val}
	h.right.hash = h.hash + "1"
}

func (h *HashTree) FindOneByHash(hash string) {
	for _, v := range hash {
		if v == leftNo {
			h = h.left
		}
		if v == RightNo {
			h = h.right
		}
	}
	if h == nil {
		fmt.Println("沒有")
		return
	}
	fmt.Println(h)

}

func (h *HashTree) LevelTraversal() {
	if h == nil {
		return
	}
	result := []*HashTree{}
	nodes := []*HashTree{h}
	for len(nodes) > 0 {
		curNode := nodes[0]
		nodes = nodes[1:]
		result = append(result, curNode)
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
