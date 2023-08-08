package graph

import (
	"fmt"
	"testing"
)

func TestUnion(t *testing.T) {
	unionFind := NewUnionFind(10)
	fmt.Println(unionFind.Find(1, 2))
	unionFind.Union(1, 2)
	fmt.Println(unionFind.Find(1, 2))
	unionFind.Union(1, 3)
	fmt.Println(unionFind.Find(2, 3))
	unionFind.Union(9, 8)
	unionFind.Union(9, 7)
	unionFind.Union(9, 6)
	fmt.Println(unionFind.Find(1, 6))
	unionFind.Union(6, 3)
	fmt.Println(unionFind.Find(1, 6))
}
