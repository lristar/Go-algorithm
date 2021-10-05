package colorFactory

import (
	"Go-algorithm/src/design-pattern/createType/abstractFactory"
	"Go-algorithm/src/design-pattern/createType/abstractFactory/colorFactory/workers"
	"fmt"
	"time"
)

type ColorFactory struct {
}

type ColorDirectory struct {
	IsStart bool
	r *workers.Red
	g *workers.Green
}

func (c *ColorFactory) CreateDirectory() abstractFactory.IDirectory {
	// 创建指挥者去指挥生产
	return &ColorDirectory{r: &workers.Red{},g: &workers.Green{}}
}


func (c *ColorFactory)Get(){
	//获取创建的类型
}

func (c *ColorDirectory)Start()  {
	// 开始生产
	fmt.Println("开始生产")
	c.IsStart=true
	go func() {
		for c.IsStart{
			c.r.CreateColor()
			c.g.CreateColor()
			time.Sleep(2*time.Second)
		}
	}()

}
func (c *ColorDirectory)Stop()  {
    //结束生产
	c.IsStart=false
	fmt.Println("结束生产")
}




