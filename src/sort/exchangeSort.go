package main

//冒泡排序
func BubbleSort(a []int) []int {
	for k := 0; k < len(a); k++ {
		for i := 0; i < len(a)-k-1; i++ {
			if a[i] > a[i+1] {
				temp := a[i+1]
				a[i+1] = a[i]
				a[i] = temp
			}
		}
	}
	return a
}

//快速排序
func QuickSort(left, right int, aa []int) {
	if left < right {
		partion := Partition(left, right, aa)
		QuickSort(left,partion-1,aa)
		QuickSort(partion+1,right,aa)
	}
}

//小的在左边，大的在右边
func Partition(left, right int, aa []int) int {
	val := aa[left]
	for left != right {
		for true {
			if aa[right] <= val {
				aa[left] = aa[right]
				left++
				break
			}
			right--
		}
		for true {
			if aa[left] > val {
				aa[right] = aa[left]
				right--
				break
			}
			left++
		}
	}
	aa[left] = val
	return left
}

