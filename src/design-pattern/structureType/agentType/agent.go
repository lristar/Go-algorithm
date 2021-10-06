package agentType

import (
	"Go-algorithm/src/design-pattern/structureType/agentType/person"
	"fmt"
)

type IProxy interface {
}

type Proxy struct {
	servers map[string]person.IServer
	c       *person.Client
}

func NewProxy() *Proxy {
	p := &Proxy{}
	p.servers = make(map[string]person.IServer)
	return p
}

func (p *Proxy) AddServers(product string, s *person.Server) {
	_, ok := p.servers[product]
	if !ok {
		p.servers[product] = s
		return
	}
	fmt.Println("该货有稳定货源，不需要")
}
func (p *Proxy) AddClient(c *person.Client) {
	p.c = c
}
func (p *Proxy) SellProduct(productName string, num int) {
	// 这个商品什么价格
	v, _ := p.servers[productName]
	fmt.Println(productName, ":", v.GetProductBalance())
	if num > v.GetProductNum() {
		fmt.Println("暂时没这么多货，顾客你看看其他的")
		return
	}
	// 正常情况是要用事务的
	err := p.c.BuyProduct(productName, num, v.GetProductBalance())
	if err != nil {
		return
	}
	v.SellProduct(num)
	fmt.Println(productName, "这个商品交易完成")
}

func (p *Proxy) ShowProduct() {
	fmt.Println("有以下商品可以卖")
	for productName := range p.servers {
		fmt.Println(productName)
		fmt.Println("--------")
	}
}
