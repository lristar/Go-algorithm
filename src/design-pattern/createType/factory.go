package main

import (
	"errors"
	"fmt"
)

// 简单工厂模式
const (
	CircleType    = "circle"
	SquareType    = "square"
	RectangleType = "rectangle"
)

type Share interface {
	draw()
}

type Rectangle struct {
}
type Square struct {
}
type Circle struct {
}

func (r *Rectangle) draw() {
	fmt.Println(RectangleType)
}

func (s *Square) draw() {
	fmt.Println(SquareType)
}

func (c *Circle) draw() {
	fmt.Println(CircleType)
}

type Factory struct {
}

func (f *Factory) GetShape(shapeType string) (Share,error) {
	switch shapeType {
	case CircleType:
		return &Circle{},nil
	case SquareType:
		return &Square{},nil
	case RectangleType:
		return &Rectangle{},nil
	default:
		return nil,errors.New("错误的入参")
	}
}
