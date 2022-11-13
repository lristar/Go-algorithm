package graph

import (
	"fmt"
	"testing"
)

func TestGraph2_AddEdge(t *testing.T) {
	g := NewGraph2()
	g.AddEdge("a", "b")
	g.AddEdge("a", "c")
	g.AddEdge("b", "c")
	fmt.Println(g.data)
}
