package tsort

import "fmt"

func MergeSort() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	fmt.Println(mergeSort(data))
}

func mergeSort(r []int) []int {
	length := len(r)
	if length <= 1 {
		return r
	}
	num := length / 2
	left := mergeSort(r[:num])
	right := mergeSort(r[num:])
	return merge(left, right)
}
func merge(left, right []int) []int {
	var result []int

	l, r := 0, 0
	for l < len(left) && r < len(right) {
		if left[l] < right[r] {
			result = append(result, left[l])
			l++
		} else {
			result = append(result, right[r])
			r++
		}
	}
	result = append(result, left[l:]...)
	result = append(result, right[r:]...)
	return result
}

func MergeSort1() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	fmt.Println(data)
}
