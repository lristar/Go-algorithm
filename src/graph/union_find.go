package graph

type UnionFind struct {
	Data []int
}

func (u *UnionFind) Find(a, b int) bool {
	v1 := u.getRoot(a)
	v2 := u.getRoot(b)
	return v1 == v2
}

func (u *UnionFind) Union(a, b int) {
	v1 := u.getRoot(a)
	v2 := u.getRoot(b)
	num1 := 0
	num2 := 0
	// 不是同一个集合才有合并的意义
	if v1 != -1 && v2 != -1 && v1 != v2 {
		// 找较大的集合
		for i := range u.Data {
			if u.Data[i] == v1 {
				num1++
			}
			if u.Data[i] == v2 {
				num2++
			}
		}
		if num2 > num1 {
			for i := range u.Data {
				if u.Data[i] == v1 {
					u.Data[i] = v2
				}
			}
		} else {
			for i := range u.Data {
				if u.Data[i] == v2 {
					u.Data[i] = v1
				}
			}
		}
	}
}

func (u *UnionFind) getRoot(index int) int {
	if index >= len(u.Data) {
		return -1
	}
	v := u.Data[index]
	if v == index {
		return v
	} else {
		return u.getRoot(v)
	}
}

func NewUnionFind(index int) *UnionFind {
	data := make([]int, index)
	for i := 0; i < index; i++ {
		data[i] = i
	}
	return &UnionFind{data}
}
