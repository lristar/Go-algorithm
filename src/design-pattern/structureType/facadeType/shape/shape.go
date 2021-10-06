package shape

import "fmt"

type Shape interface {
	Draw()
}

type Square struct {

}


type Rectangle struct {

}

type Circle struct {

}

func (s *Square)Draw(){
   fmt.Println("square")
}

func (r *Rectangle)Draw(){
	fmt.Println("Rectangle")
}

func (c *Circle)Draw(){
	fmt.Println("Circle")
}
