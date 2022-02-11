package reflect

import (
"fmt"
"reflect"
	"testing"
)

var record map[string]interface{}

func init() {
	record = map[string]interface{}{
		"name":   "urx",
		"number": 168,
	}
}

func get(in []reflect.Value, retType reflect.Type) ([]reflect.Value) {
	m := in[0]
	key := in[1]
	result := m.MapIndex(key)
	ok := false

	if result.IsValid() == false {
		return []reflect.Value{reflect.Zero(retType), reflect.ValueOf(ok)}
	}

	resultval := result.Interface()

	if retType != reflect.TypeOf(resultval) {
		return []reflect.Value{reflect.Zero(retType), reflect.ValueOf(ok)}
	}

	ok = true
	return []reflect.Value{reflect.ValueOf(resultval), reflect.ValueOf(ok)}
}

func MakeFuncStub(typ reflect.Type, fn func(args []reflect.Value, retType reflect.Type) (results []reflect.Value)) reflect.Value {
	retType := typ.Out(0)
	return reflect.MakeFunc(typ, func(args []reflect.Value) (results []reflect.Value) {
		return fn(args, retType)
	})
}

func Test_makefunc(t *testing.T) {
	makeGet := func(fptr interface{}) {
		fn := reflect.ValueOf(fptr).Elem()
		fn.Set(MakeFuncStub(fn.Type(), get))
	}

	var getInt func(map[string]interface{}, string) (int, bool)
	makeGet(&getInt)

	var getString func(map[string]interface{}, string) (string, bool)
	makeGet(&getString)

	v1, ok := getInt(record, "number")
	fmt.Printf("%v, %v\n", v1, ok)

	v2, ok := getString(record, "doesnotexist")
	fmt.Printf("%v, %v\n", v2, ok)

	v3, ok := getString(record, "number")
	fmt.Printf("%v, %v\n", v3, ok)

	v4, ok := getString(record, "name")
	fmt.Printf("%v, %v\n", v4, ok)

	v5, ok := getInt(record, "nonum")
	fmt.Printf("%v, %v\n", v5, ok)
}