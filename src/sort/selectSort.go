package main

//选择排序 --查找最小的插入到前面

func SelectSort(a []int) []int {
	var min int
	var index int
	for i := 0; i < len(a)-1; i++ {
		min = a[i]
		index = i
		for j := i + 1; j < len(a); j++ {
			if min > a[j] {
				min = a[j]
				index = j
			}
		}
		temp := a[i]
		a[i] = a[index]
		a[index] = temp
	}
	return a
}

func HeapSort(aa []int, index int) []int {
	var bb []int
	length:=len(aa)
	for i := 0; i < length; i++ {
		BigHeapSort(aa, index)
		bb = append(bb, aa[0])

		aa[0]=aa[len(aa)-1]
		aa = aa[:len(aa)-1]
	}
	return bb
}

//大顶堆构建
func BigHeapSort(aa []int, index int) {
	if index > len(aa)-1 {
		return
	}
	BigHeapSort(aa, index*2+1)
	BigHeapSort(aa, index*2+2)
	if len(aa)-1 >= index*2+2 {
		if aa[index] < aa[index*2+2] {
			swap(aa, index, index*2+2)
		}
	}
	if len(aa)-1 >= index*2+1 {
		if aa[index] < aa[index*2+1] {
			swap(aa, index, index*2+1)
		}
	}
}
func swap(aa []int, i, j int) {
	temp := aa[i]
	aa[i] = aa[j]
	aa[j] = temp
}
