package reflect

import (
	"fmt"
	"reflect"
	"testing"
)

var _internalField = "Internal"

type Person struct {
	Name string
	Age  int
}
type T struct {
	A string
	Internal struct{
		hello func()
		bye func()
	}
}

type Handler interface {
	hello()
	bye()
}

func (s *Student)hello(){
	fmt.Println("hello")
}

func (s *Student)bye(){
	fmt.Println("bye")
}

func Test_AA(t *testing.T) {
	var x T = T{A:"easy"}
	v := reflect.ValueOf(&x).Elem()
	internal := v.FieldByName(_internalField)
	ii := internal.Addr().Interface()
	fmt.Println("ii:",ii)
	count := v.NumField()
	fmt.Println("count:",count)
	for i := 0; i < count; i++ {
		field := v.Field(i)
		fmt.Println(field)
		if field.Kind() == reflect.String {
			field.SetString("Test Value")
		}
	}
}


func Test_BB(t *testing.T) {
	var x T = T{A:"easy"}
	fmt.Println("type:", reflect.TypeOf(x))
	v := reflect.ValueOf(x)
	fmt.Println("value:", v)
	fmt.Println("type:", v.Type())
	fmt.Println("kind:", v.Kind())
	fmt.Println(v.Interface())
	y := v.Interface().(T)
	fmt.Println(y)
}
