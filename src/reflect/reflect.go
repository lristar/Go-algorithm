package main

import (
	"errors"
	"fmt"
	"reflect"
	"sync"
)

var injectTagName = "inject" //依赖注入tag名

type Student struct {
	Name string
	Age  int
}

func (s *Student) Hello(str string) {
	fmt.Println(str)
}

var GFactory = &Container{
	singleton: make(map[string]interface{}),
}

type Container struct {
	sync.Mutex
	singleton map[string]interface{}
}
type Factory struct {
	GController *Controller `inject:"Controller"`
}

type Controller struct {
	stu *Student `inject:"Student"`
}

func (f *Container) SetSingleton(str string, fuc interface{}) {
	f.Lock()
	f.singleton[str] = fuc
	f.Unlock()
}

func (f *Container) GetSingleton(name string) (interface{}, error) {
	val, ok := f.singleton[name]
	if !ok {
		return nil, errors.New("没有")
	}
	return val, nil
}

// 注入实例

func (f *Container) Entry(instance interface{}) error {
	fmt.Println("Entry")
	err := f.entryValue(reflect.ValueOf(instance))
	if err != nil {
		return err
	}
	return nil
}

func (f *Container) entryValue(value reflect.Value) error {
	if value.Kind() != reflect.Ptr {
		return errors.New("必须为指针")
	}
	elemType, elemValue := value.Type().Elem(), value.Elem()
	for i := 0; i < elemType.NumField(); i++ {
		if !elemValue.Field(i).CanSet() {
			fmt.Println("continue")
			continue
		}
		fieldType := elemType.Field(i)
		if fieldType.Anonymous {
			fmt.Println(fieldType.Name + "是匿名字段")
		} else {
			if elemValue.Field(i).IsZero() {
				tag := fieldType.Tag.Get(injectTagName)
				injectInstance := f.GetInstance(tag)
				f.entryValue(reflect.ValueOf(injectInstance))
				elemValue.Field(i).Set(reflect.ValueOf(injectInstance))
			} else {
				fmt.Println("非空", fieldType.Name)
			}
		}

	}
	return nil
}

func (f *Container) GetInstance(tag string) interface{} {
	fmt.Println("GetInstance")
	val, err := f.GetSingleton(tag)
	if err != nil {
		return nil
	}
	return val
}

func Init() {
	GFactory.SetSingleton("Student", &Student{})
	GFactory.SetSingleton("Controller", &Controller{})
	ctlFactory := &Factory{}
	GFactory.SetSingleton("CtrlFactory", ctlFactory)
	err := GFactory.Entry(ctlFactory)
	if err != nil {
		return
	}
}

func main() {
	Init()
	injectVal, err := GFactory.GetSingleton("Student")
	if err != nil {
		return
	}
	c := injectVal.(*Student)
	c.Hello("hello")
}
