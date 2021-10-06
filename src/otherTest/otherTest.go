package main

import "fmt"

type Student struct {
	ClassNum int
	People
}

type People struct {
	Name string
	Age int
}

func (p *People)Hello(){
	fmt.Println("hello")
}
func main() {
	stu:=Student{
		ClassNum: 0,
		People:   People{},
	}
	stu.Hello()
}
