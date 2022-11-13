package graph

// 邻接表表示法 稀疏图的时候使用
type Graph2 struct {
	data map[string]*Node
}

type Node struct {
	nextarc *Node
	name    string
	width   int
}

func NewGraph2() *Graph2 {
	return &Graph2{data: make(map[string]*Node, 0)}
}

func (g *Graph2) addEdge(from, to string) {
	if _, ok := g.data[from]; !ok {
		g.addSpot(from)
	}
	v := g.data[from]
	for v.nextarc != nil {
		v = v.nextarc
	}
	v.nextarc = &Node{
		nextarc: nil,
		name:    to,
		width:   0,
	}
}

func (g *Graph2) AddEdge(from, to string) {
	// 无向图用法
	//g.addEdge(from, to)
	//g.addEdge(to, from)

	// 有向图用法
	g.addEdge(from, to)
}

func (g *Graph2) addSpot(name string) {
	g.data[name] = &Node{
		nextarc: nil,
		name:    "",
		width:   0,
	}
}
