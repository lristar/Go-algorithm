package workers

import "fmt"


type Red struct {
}

func (r *Red)CreateColor(){
	fmt.Println("开始生产red")
}

type Green struct {
}

func (g *Green)CreateColor(){
	fmt.Println("开始生产green")
}
