package sort


//归并排序
func MergeSort(left, right int, aa []int) {
	if left < right {
		MergeSort(left, (right+left)/2, aa)
		MergeSort(((right+left)/2)+1, right, aa)
		MSort(left, right, aa)
	}
}
func MSort(left, right int, aa []int) {
	for k := left; k <= right; k++ {
		for i := left; i < right-k; i++ {
			if aa[i] > aa[i+1] {
				temp := aa[i+1]
				aa[i+1] = aa[i]
				aa[i] = temp
			}
		}
	}
}
