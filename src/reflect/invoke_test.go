package reflect

import (
	"fmt"
	"testing"
)

type GOGOGO struct {

}

func Override(typ interface{}) {
	if a,ok:=typ.(int); ok {
		fmt.Println("----",a)
	}
	fmt.Println("failed")
}

func invoke(){
	fmt.Println()
	Override(new(*GOGOGO))
}

func Test_invoke(t *testing.T) {
	invoke()
}