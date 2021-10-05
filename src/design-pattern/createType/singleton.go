package main

import (
	"fmt"
	"sync"
)

var (
	Con  *controller
	Once sync.Once
)

type Singleton interface {
	foo()
}

func (c controller) foo() {
	fmt.Println(c.Aa)
}

type controller struct {
	Aa string
}

// 饿汉模式

func NewController() {
	Con = &controller{Aa: "张三"}
}

func GetController() Singleton {
	return Con
}

// 饱汉模式

func GetController1() Singleton {
	var con *controller
	Once.Do(func() {
		con = &controller{Aa: "张三"}
	})
	return con
}


