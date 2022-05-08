package main

import "fmt"

type person struct {
	name string
	age  int
}

func (p *person) say() {
	fmt.Println(p.name)
}

type student struct {
	person //用于将student继承person
	grade int
}

func main() {
	fmt.Print("hello world\n")

	var s student = student{
		person: person{
			name: "张三",
		},
		grade: 1,
	}
	s.name = "王五"
	s.say()

	fmt.Println(s.name)
}
