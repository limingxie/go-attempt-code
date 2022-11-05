package tsort

import "fmt"

//插入排序：是在有序的数据中寻找数据
func InsertionSort() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for i := 0; i < len(data); i++ {
		for j := 0; j < i; j++ {
			if data[j] > data[i] {
				v := data[i]
				for k := i; k > j; k-- {
					data[k] = data[k-1]
				}
				data[j] = v
			}
		}
	}

	fmt.Println(data)
}

func InsertionSort2() {
	arr := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for i := range arr {
		preIndex := i - 1
		current := arr[i]
		for preIndex >= 0 && arr[preIndex] > current {
			arr[preIndex+1] = arr[preIndex]
			preIndex -= 1
		}
		arr[preIndex+1] = current
	}

	fmt.Println(arr)
}

func InsertionSort1() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	var v int
	for i := 0; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				v = data[j]
				data[j] = data[j-1]
				data[j-1] = v
			} else {
				break
			}
		}
	}

	fmt.Println(data)
}

func InsertionSort3() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for i := 0; i < len(data); i++ {
		for j := i; j > 0; j-- {
			if data[j] < data[j-1] {
				data[j], data[j-1] = data[j-1], data[j]
			}
		}
	}

	fmt.Println(data)
}
