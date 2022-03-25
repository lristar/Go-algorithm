package main

import (
	"fmt"
)

//defer 顺序 越下边的越先
func test_defer() {
	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three")
	panic("触发异常")
}

// 显示的map的value都为3，因为val指向的地址是数值中最后赋值的3，所以结果都为3
func test_slice() {
	slice := []int{0, 1, 2, 3}
	m := make(map[int]*int)

	for key, val := range slice {
		m[key] = &val
	}

	for k, v := range m {
		fmt.Println(k, "->", *v)
	}
}

func test_slice_append() {
	s := make([]int, 5)    //已经创建五个零，如果用append的话不会覆盖前面的零
	s = append(s, 1, 2, 3) // append是在尾巴添加数值
	fmt.Println(s)
}
func test_new() {
	list := new([]int) //通过new 常见的对象是指针
	*list = append(*list, 1)
}

var (
//size := 1024     //这里用:=会失败
//max_size = size*2
)

//1.必须使用显示初始化；
//2.不能提供数据类型，编译器会自动推导；
//3.只能在函数内部使用简短模式；

func test_equal() {
	st1 := struct {
		name string
		age  int
	}{name: "张三", age: 18}

	st2 := struct {
		name string
		age  int
	}{name: "张三", age: 18}

	if st1 == st2 { //这里编译是正常的
		fmt.Println("hahah")
	}
	//sm1 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}
	//sm2 := struct {
	//	age int
	//	m   map[string]string
	//}{age: 11, m: map[string]string{"a": "1"}}

	//if sm1 == sm2 {  // 这里编译是失败的，像切片、map、函数等是不能比较的
	//	fmt.Println("sm1 == sm2")
	//}
}

type Student struct {
	name string
}

func test_getChildren() {
	s := &Student{name: "lisi"}
	// 取地址变量取成员变量有两种方式
	fmt.Println(s.name)
	fmt.Println((*s).name)
}

type MyInt1 int   //类型定义  这个必须强转
type MyInt2 = int //类型别名

//func test_type() {
//	var i int = 0
//	var i3 MyInt1 = i  //这里会报错 ，不是同一类型
//	var i1 MyInt1 = MyInt1(i) // 属于类型，不能赋值,只能强转
//	var i2 MyInt2 = i // 这种就可以
//}

func test_string_join() {
	// 字符串定义不能用单引号，只能用双引号
}

const ( // 常量的定义 中间插值不影响他的顺序
	_   = 1 << (10 * iota)
	KiB  // 1024
	MiB  // 1048576
	GiB  // 1073741824
	TiB  // 1099511627776             (exceeds 1 << 32)
	PiB  // 1125899906842624
	EiB  // 1152921504606846976
	// 这里往下已经超过阈值了 const使用的是int 阈值对应的int
	ZiB  // 1180591620717411303424    (exceeds 1 << 64)
	YiB  // 1208925819614629174706176
)

func hello(num []int) {
	num = append(num, 18)
}

// defer 中的i是先保存的之前的数据
//func hello(i int) {
//	fmt.Println(i)
//}
//func main() {
//	i := 5
//	defer hello(i)
//	i = i + 10
//}

// 如果 返回值是匿名的话 顺序则是把return n的值保存在一个临时变量，再执行defer ,最后再return n
// 如果 返回值不是匿名，如下，就没有保存到临时变量，保存再返回值上，再执行defer，最后return n
// defer 关键字后面的函数或者方法想要执行必须先注册，return 之后的 defer 是不能注册的， 也就不能执行后面的函数或方法
func demo01() (n int) {
	n = 10
	defer func() {
		n++
		fmt.Println("defer1的值为：", n) // 输出13
	}()
	defer func() {
		n++
		fmt.Println("defer2的值为：", n) // 输出12
	}()
	n++
	return 5
}

func f1() (r int) {
	defer func() {
		r++
	}()
	return r
}

func f2() (r int) {
	t := 5
	defer func() {
		t = t + 5
	}()
	return t
}

// 只是在局部变量中的执行，所以并没有变化
func f3() (r int) {
	defer func(r int) {
		r = r + 5
	}(r)
	return 0
}

func test_for_range() {
	v := []int{1, 2, 3}
	for i := range v {
		v = append(v, i)
	}
	fmt.Println(v)
}

//func test_array(s ...int){
//	s = append(s, 3)
//
//}
//func main() {
//	a := []int{0,0,0,0,0}
//	a[1]=1
//	a[2]=2
//	test_array(a[1:3]...)
//	fmt.Println("a is ",a)
//}

// 切片和数组的细节区别
func test_slice_array() {
	var b = []int{1, 2, 3, 4, 5}
	var a = [5]int{1, 2, 3, 4, 5}
	var r [5]int

	for i, v := range a {
		if i == 0 {
			a[1] = 12
			a[2] = 13
		}
		r[i] = v
	}
	// 如果数组，for range则是复制了一个数组
	fmt.Println("r1 = ", r)
	for i, v := range b {
		if i == 0 {
			b[1] = 12
			b[2] = 13
		}
		r[i] = v
	}
	//如果是切片，for range有引用则是完成了修改
	fmt.Println("r2 = ", r)
	fmt.Println("b = ", b)
	fmt.Println("a = ", a)
}

type ConfigOne struct {
	Daemon string
}

// 重写输出成string
/*func (c *ConfigOne) String() string {
	return c.Daemon
}
*/
type N int

func (n N) value() {
	n++
	fmt.Printf("v:%p,%v\n", &n, n)
}

func (n *N) pointer() {
	*n++
	fmt.Printf("v:%p,%v\n", n, *n)
}
// ：锁失效。将 Mutex 作为匿名字段时，相关的方法必须使用指针接收者，否则会导致锁机制失效。


//参考答案及解析：4。知识点：多重赋值。
//
//多重赋值分为两个步骤，有先后顺序：
//
//计算等号左边的索引表达式和取址表达式，接着计算等号右边的表达式；
//defer() 后面的函数如果带参数，会优先计算参数，并将结果存储在栈中，到真正执行 defer() 的时候取出

