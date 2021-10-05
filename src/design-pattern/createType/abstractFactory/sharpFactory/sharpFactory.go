package sharpFactory

import (
	"Go-algorithm/src/design-pattern/createType/abstractFactory"
	"Go-algorithm/src/design-pattern/createType/abstractFactory/sharpFactory/workers"
	"fmt"
	"time"
)

type ShapeFactory struct {
}

type ShapeDirectory struct {
	IsStart bool
	c       *workers.Circle
	r       *workers.Rectangle
	s       *workers.Square
}

func (d *ShapeFactory) CreateDirectory() abstractFactory.IDirectory {
	// 创建指挥者去指挥生产
	return &ShapeDirectory{IsStart: true,c: &workers.Circle{},r: &workers.Rectangle{},s: &workers.Square{}}
}

func (d *ShapeFactory) Get() {
	// 获取创建的类型
}

func (d *ShapeDirectory) Start() {
	// 开始生产
	fmt.Println("开始生产")
	d.IsStart=true
	go func() {
		for d.IsStart {
			d.s.CreateSharp()
			d.c.CreateSharp()
			d.r.CreateSharp()
			time.Sleep(time.Second * 2)
		}
	}()
}
func (d *ShapeDirectory) Stop() {
	//结束生产
	d.IsStart=false
	fmt.Println("结束生产")
}
