package builder

import (
	"Go-algorithm/src/design-pattern/createType/builder/item"
	"Go-algorithm/src/design-pattern/createType/builder/product"
	"fmt"
)

type Meal struct {
	Items []item.Item
}

func (m *Meal) AddItem(i item.Item) {
	m.Items = append(m.Items, i)
}

func (m *Meal) GetCost()  {
	var all float64
	for _, v := range m.Items {
		all += v.Getprice()
	}
	fmt.Println("cost:",all)
}

func (m *Meal) ShowItems() {
	for _, v := range m.Items {
		fmt.Println("item:", v.Getname(), "packing:", v.Packing(), "price:", v.Getprice())
	}
}

type Builder struct {
}

func (b *Builder) PrepareVegMeal() *Meal {
	m := Meal{}
	app := &product.Apple{Name: "苹果", Value: 1.99}
	m.AddItem(app)
	egg := &product.Egg{Name: "鸡蛋", Value: 2.99}
	m.AddItem(egg)
	return &m
}

func (b *Builder) PrepareSweetMeal() *Meal {
	m := Meal{}
	cake := &product.Cake{Name: "蛋糕", Value: 5.99}
	m.AddItem(cake)
	egg := &product.Egg{Name: "鸡蛋", Value: 2.99}
	m.AddItem(egg)
	return &m
}
