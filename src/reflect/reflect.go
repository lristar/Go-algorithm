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

func (s *Student) GetName() string {
	return s.Name
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
	Stu *Student `inject:"Student"`
}

func (f *Container) SetSingleton(str string, fuc interface{}) {
	f.Lock()
	f.singleton[str] = fuc
	f.Unlock()
}

func (f *Container) GetSingleton(name string) (interface{}, error) {
	fuc, ok := f.singleton[name]
	if !ok {
		return nil, errors.New("没有这个key")
	}
	return fuc, nil
}

// 注入实例

func (f *Container) Entry(instance interface{}) error {
	if err := f.entryValue(reflect.ValueOf(instance)); err != nil {
		return err
	}
	return nil
}

func (f *Container) entryValue(value reflect.Value) error {
	fmt.Println("start:",value.Type())
	if value.Kind() != reflect.Ptr {
		return errors.New("不是指针")
	}
	elemType, elemValue := value.Type().Elem(), value.Elem()
	for i := 0; i < elemType.NumField(); i++ {
		if !elemValue.Field(i).CanSet() {
			continue
		}
		fieldType := elemType.Field(i)
		if fieldType.Anonymous {
			fmt.Println(fieldType.Name + "是匿名字段")
			item := reflect.New(elemValue.Field(i).Type())
			f.entryValue(item)
			elemValue.Field(i).Set(item.Elem())
		} else {
			if elemValue.Field(i).IsZero() {
				tag := fieldType.Tag.Get(injectTagName)
				fmt.Println(tag)
				ufc, err := f.GetInstance(tag)
				if err != nil {
					return err
				}
				f.entryValue(reflect.ValueOf(ufc))
				elemValue.Field(i).Set(reflect.ValueOf(ufc))
				fmt.Println("注入了",elemType.Field(i),"-",reflect.ValueOf(ufc).Type())
			} else {
				fmt.Println("非空", fieldType)
			}
		}
	}
	return nil
}

func (f *Container) GetInstance(tag string) (interface{}, error) {
	return f.GetSingleton(tag)
}

func Init() {
	GFactory.SetSingleton("Student", &Student{Name: "张三", Age: 18})
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
	stu, err := GFactory.GetSingleton("Student")
	if err != nil {
		fmt.Println("失败")
		return
	}
	fu := stu.(*Student)
	fmt.Println(fu.GetName())
}
