package facadeType

import (
	"Go-algorithm/src/design-pattern/structureType/facadeType/shape"
	"fmt"
)

type ShapeMaker struct {
	c *shape.Circle
	s *shape.Square
	r *shape.Rectangle
}

func  NewMaker() *ShapeMaker {
	return &ShapeMaker{
		c: &shape.Circle{},
		s: &shape.Square{},
		r: &shape.Rectangle{},
	}
}

func (s *ShapeMaker) GetCircle() {
	fmt.Println("封装了")
	s.c.Draw()
}
func (s *ShapeMaker) GetSquare() {
	fmt.Println("封装了")
	s.s.Draw()
}
func (s *ShapeMaker) GetRectangle() {
	fmt.Println("封装了")
	s.r.Draw()
}
