package workers

import "fmt"

type Rectangle struct {
}
type Square struct {
}
type Circle struct {
}

func (s *Square)CreateSharp(){
	fmt.Println("开始生产Square")
}


func (c *Circle)CreateSharp(){
	fmt.Println("开始生产Circle")
}

func (r *Rectangle)CreateSharp(){
	fmt.Println("开始生产Rectangle")
}