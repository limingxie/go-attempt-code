package tsort

import "fmt"

//泡沫排序
func BubbleSort1() {
	aa := make(map[int]int)
	aa[0] = 8
	aa[1] = 100
	aa[2] = 99
	aa[3] = 50
	aa[4] = 22
	aa[5] = 15
	aa[6] = 16
	aa[7] = 2
	aa[8] = 99
	aa[9] = 1000
	aa[10] = 999
	aa[11] = 1

	for i := 0; i < 11; i++ {
		for j := 0; j < 11-i; j++ {
			if aa[j] > aa[j+1] {
				v := aa[j]
				aa[j] = aa[j+1]
				aa[j+1] = v
			}
		}
	}

	fmt.Println(aa)
}

//泡沫排序
func BubbleSort() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-i; j++ {
			if data[j] > data[j+1] {
				v := data[j]
				data[j] = data[j+1]
				data[j+1] = v
			}
		}
	}

	fmt.Println(data)
}
