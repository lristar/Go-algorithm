package graph

// 邻接矩阵实现图结构
// 稠密图的时候使用
//
type Graph1 struct {
	// 记录点
	index int
	spot  map[string]int
	// 记录边
	Edges [][]int
}

func NewGraph1() *Graph1 {
	return &Graph1{spot: make(map[string]int, 0)}
}

func (g *Graph1) GetSpot(name string) (res int) {
	if v, ok := g.spot[name]; ok {
		return v
	}
	return -1
}

func (g *Graph1) AddSpot(names ...string) {
	for i := range names {
		if _, ok := g.spot[names[i]]; !ok {
			g.spot[names[i]] = g.index
			g.index++
		}
	}
	g.Edges = make([][]int, len(g.spot))
	for i := 0; i < len(g.spot); i++ {
		g.Edges[i] = make([]int, len(g.spot))
	}
}

func (g *Graph1) AddEdges(from, to string, width int) {
	f := g.GetSpot(from)
	t := g.GetSpot(to)
	if f == -1 || t == -1 {
		return
	}
	// 无向图实现
	g.Edges[f][t] = width
	g.Edges[t][f] = width
}

func (g *Graph1) AddEdges11(from, to, width int) {
	g.Edges[from][to] = width
	g.Edges[to][from] = width
}

// AddEdges2 有向图实现
func (g *Graph1) AddEdges2(from, to string) {
	f := g.GetSpot(from)
	t := g.GetSpot(to)
	if f == -1 || t == -1 {
		return
	}
	// 无向图实现
	g.Edges[f][t] = 1
}

func (g *Graph1) Clone() *Graph1 {
	edges := make([][]int, g.index)
	for i := range edges {
		edges[i] = make([]int, g.index)
	}
	return &Graph1{
		index: g.index,
		spot:  g.spot,
		Edges: edges,
	}
}
