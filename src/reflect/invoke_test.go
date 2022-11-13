package reflect

import (
	"fmt"
	"testing"
)

type GOGOGO struct {
}

func (g *GOGOGO) hello() {
	fmt.Println("hello")
}

func Override(typ interface{}) {
	if a, ok := typ.(*GOGOGO); ok {
		a.hello()
		fmt.Println("----", a)
		return
	}
	fmt.Println("failed")
}

func invoke() {
	fmt.Println()
	Override(new(GOGOGO))
}

func Test_invoke(t *testing.T) {
	invoke()
}
