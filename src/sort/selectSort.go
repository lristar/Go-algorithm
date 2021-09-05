package sort

//选择排序 --查找最小的插入到前面

func SelectSort(a []int) []int {
	var min int
	var index int
	for i := 0; i < len(a)-1; i++ {
		min = a[i]
		index = i
		for j := i + 1; j < len(a); j++ {
			if min > a[j] {
				min=a[j]
				index = j
			}
		}
		temp := a[i]
		a[i] = a[index]
		a[index] = temp
	}
	return a
}
