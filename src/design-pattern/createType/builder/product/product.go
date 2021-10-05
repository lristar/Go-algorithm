package product

import (
	"Go-algorithm/src/design-pattern/createType/builder/item"
)

type Bottle struct {
}

func (b *Bottle) Pack() string {
	return "bottle"
}

type Cake struct {
	Value float64
	Name  string
}

func (c *Cake) Getname() string {
	return c.Name
}
func (c *Cake) Getprice() float64 {
	return c.Value
}
func (c *Cake) Packing() item.Packing {
	return &Bottle{}
}

type Egg struct {
	Value float64
	Name  string
}
func (c *Egg) Getname() string {
	return c.Name
}
func (c *Egg) Getprice() float64 {
	return c.Value
}
func (c *Egg) Packing() item.Packing {
	return &Bottle{}
}

type Apple struct {
	Value float64
	Name  string
}
func (c *Apple) Getname() string {
	return c.Name
}
func (c *Apple) Getprice() float64 {
	return c.Value
}
func (c *Apple) Packing() item.Packing {
	return &Bottle{}
}
