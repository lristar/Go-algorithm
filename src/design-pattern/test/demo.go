package main


import "fmt"

type IInfo interface {
	Info()
	Detail()
	Run()
	Stop()
}

type CommInfo struct {
	ctx string
}

func (c *CommInfo)Info(){
	fmt.Println("info")
}
func (c *CommInfo)Detail(){
	fmt.Println("Detail")
}
func (c *CommInfo)Run(){
}
func (c *CommInfo)Stop(){
	fmt.Println("Stop")
}

type Demo struct {

}

type IDemo interface {
	GOGO()
}

func (d *Demo) GOGO(){
	fmt.Println("gogogo")
}

type BatchInfo struct {
	CommInfo
	Demo
}

func (b *BatchInfo)Run(){
	fmt.Println("runrunrun")
	b.GOGO()
}

type BatchModifyInfo struct {
	BatchInfo
}

func (b *BatchModifyInfo)GOGO(){
	fmt.Println("GOGOGOGOGOGO")
}


func main() {
	b := new(BatchModifyInfo)
	b.ctx="halo"
	b.Run()
}



