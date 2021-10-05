package main

import (
	"Go-algorithm/src/design-pattern/createType/builder"
	"fmt"
)

func main() {
	b:=builder.Builder{}
	m:=b.PrepareVegMeal()
	m.ShowItems()
	m.GetCost()
	fmt.Println("----------")
	m1:=b.PrepareSweetMeal()
	m1.ShowItems()
	m1.GetCost()
}
