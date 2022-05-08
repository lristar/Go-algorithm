package main


import "fmt"

type IInfo interface {
	Info()
	Detail()
	Run()
	Stop()
}

type CommInfo struct {

}

func (c *CommInfo)Info(){
	fmt.Println("info")
}
func (c *CommInfo)Detail(){
	fmt.Println("Detail")
}
func (c *CommInfo)Run(){
	fmt.Println("Run")
}
func (c *CommInfo)Stop(){
	fmt.Println("Stop")
}


type BatchInfo struct {
	CommInfo
}

func main() {
	b := BatchInfo{}
	b.Run()
}



