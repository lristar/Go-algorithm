package graph

import (
	"fmt"
	"sort"
)

// kruskd 克鲁斯卡尔算法 以边为起点
type Krus struct {
	// 已有节点
	nodes *dsu
	*Graph1
	old Graph1
}
type kEdge struct {
	from  int
	to    int
	width int
}
type ks []kEdge

func (k ks) Len() int {
	return len(k)
}

func (k ks) Less(i, j int) bool {
	if k[i].width < k[j].width {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (k ks) Swap(i, j int) {
	k[i], k[j] = k[j], k[i]
}

func NewKrus(p *Graph1) *Krus {
	dsu := NewDsu(len(p.spot))
	return &Krus{
		nodes:  dsu,
		Graph1: p.Clone(),
		old:    *p,
	}
}

func (k *Krus) GetMinTree() {
	kk := make(ks, 0)
	for i := range k.old.Edges {
		for j := i; j < len(k.old.Edges); j++ {
			if k.old.Edges[i][j] > 0 {
				kk = append(kk, kEdge{
					from:  i,
					to:    j,
					width: k.old.Edges[i][j],
				})
			}
		}
	}
	sort.Sort(kk)
	for i := range kk {
		// 这个查看环路有问题，使用并查集查看环路
		if k.nodes[kk[i].from] == k.nodes[kk[i].to] {
			continue
		}
		if len(k.nodes) != len(k.spot) {
			// 这里如何
			k.nodes[kk[i].from] = struct{}{}
			k.nodes[kk[i].to] = struct{}{}
			k.AddEdges11(kk[i].from, kk[i].to, kk[i].width)
		}
		continue
	}
}

func Kruss(p Graph1) {
	k := NewKrus(&p)
	k.GetMinTree()
	fmt.Println(k.Graph1)
}

// Prim 普利姆算法
type Prim struct {
	nodes    map[int]interface{}
	nowNodes map[int]interface{}
	*Graph1
	old Graph1
}

func NewPrim(p *Graph1) *Prim {
	return &Prim{
		nodes:    make(map[int]interface{}, 0),
		nowNodes: make(map[int]interface{}, 0),
		Graph1:   p.Clone(),
		old:      *p,
	}
}

func (p *Prim) GetMinTree() {
	p.nowNodes[0] = struct{}{}
	p.AddMinEdges()
}

func (p *Prim) AddMinEdges() {
	kk := make(ks, 0)
	for i := range p.nowNodes {
		for i2 := range p.old.Edges[i] {
			if p.old.Edges[i][i2] > 0 {
				kk = append(kk, kEdge{
					from:  i,
					to:    i2,
					width: p.old.Edges[i][i2],
				})
			}
		}
	}
	sort.Sort(kk)
	for i := range kk {
		if _, ok := p.nodes[kk[i].from]; ok {
			if _, ok := p.nodes[kk[i].to]; ok {
				continue
			}
		}
		p.nodes[kk[i].from] = struct{}{}
		p.nodes[kk[i].to] = struct{}{}
		p.AddEdges11(kk[i].from, kk[i].to, kk[i].width)
		p.nowNodes[kk[i].to] = struct{}{}
		if len(p.nodes) != len(p.spot) {
			p.AddMinEdges()
		}
		continue
	}
}

func Prims(p Graph1) {
	k := NewPrim(&p)
	k.GetMinTree()
	fmt.Println(k.Graph1)
}

// 并查集结构
type dsu struct {
	data map[int]int
}

func NewDsu(index int) *dsu {
	d := dsu{data: make(map[int]int, 0)}
	for i := 0; i < index; i++ {
		d.data[i] = i
	}
	return &d
}
func (d *dsu) GetRoot(i int) int {
	if d.data[i] != i {
		d.GetRoot(d.data[i])
	}
}

// todo
func (d *dsu) addNode(from, to int) {
	d.data[from] = d.data[to]

}
