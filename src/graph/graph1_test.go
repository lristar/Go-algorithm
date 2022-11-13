package graph

import (
	"testing"
)

func TestGraph1_AddEdges(t *testing.T) {
	g := NewGraph1()
	g.AddSpot("a", "b", "c", "d", "e")
	g.AddEdges("a", "b", 6)
	g.AddEdges("a", "c", 8)
	g.AddEdges("a", "d", 5)
	g.AddEdges("a", "e", 1)
	g.AddEdges("b", "c", 3)
	g.AddEdges("b", "d", 4)
	g.AddEdges("c", "d", 7)
	// 最小生成树生成
}

func TestKruss(t *testing.T) {
	g := NewGraph1()
	g.AddSpot("a", "b", "c", "d", "e")
	g.AddEdges("a", "b", 6)
	g.AddEdges("a", "c", 8)
	g.AddEdges("a", "d", 5)
	g.AddEdges("a", "e", 1)
	g.AddEdges("b", "c", 3)
	g.AddEdges("b", "d", 4)
	g.AddEdges("c", "d", 7)
	Kruss(*g)
}

func TestPrim(t *testing.T) {
	g := NewGraph1()
	g.AddSpot("a", "b", "c", "d", "e")
	g.AddEdges("a", "b", 6)
	g.AddEdges("a", "c", 8)
	g.AddEdges("a", "d", 5)
	g.AddEdges("a", "e", 1)
	g.AddEdges("b", "c", 3)
	g.AddEdges("b", "d", 4)
	g.AddEdges("c", "d", 7)
	Kruss(*g)
	Prims(*g)
}
