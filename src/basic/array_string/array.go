package main


import "fmt"

func main() {
	var a [3]int                    // 定义长度为 3 的 int 型数组, 元素全部为 0
	var b = [...]int{1, 2, 3}       // 定义长度为 3 的 int 型数组, 元素为 1, 2, 3
	var c = [...]int{2: 3, 1: 2}    // 定义长度为 3 的 int 型数组, 元素为 0, 2, 3
	var d = [...]int{1, 2, 4: 5, 6} // 定义长度为 6 的 int 型数组, 元素为 1, 2, 0, 0, 5, 6

	fmt.Println("a is :",a)
	fmt.Println("b is :",b)
	fmt.Println("c is :",c)
	fmt.Println("d is :",d)

	var e = [...]int{1, 2, 3} // a 是一个数组
	var f = &e                // b 是指向数组的指针

	fmt.Println(e[0], e[1])   // 打印数组的前 2 个元素
	fmt.Println(f[0], f[1])   // 通过数组指针访问数组元素的方式和数组类似

	for i, v := range f {     // 通过数组指针迭代数组的元素
		fmt.Println(i, v)
	}
}


