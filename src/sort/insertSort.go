package sort

func InsertSort(a []int) []int {
	for i := 0; i < len(a); i++ {
		aaa:=a[i]
		for j := i - 1; j >= 0; j-- {
			if aaa <= a[j] {
				a[j+1] = a[j]
			} else {
				a[j+1] = aaa
				break
			}
		}
	}
	return a
}
