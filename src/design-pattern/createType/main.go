package main

import (
	"Go-algorithm/src/design-pattern/createType/abstractFactory"
	"Go-algorithm/src/design-pattern/createType/abstractFactory/sharpFactory"
	"time"
)

func main() {
	var abf abstractFactory.AbstractFactory
	var dir abstractFactory.IDirectory

	abf = &sharpFactory.ShapeFactory{}
	dir = abf.CreateDirectory()
	dir.Start()
	time.Sleep(time.Second*50)
	dir.Stop()
}
