package tsort

import "fmt"

//桶排序
func BucketSort() {
	aa := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}
	bb := make(map[int]int)

	for _, a := range aa {
		bb[a] = bb[a] + 1
	}
	for i := 0; i < 1001; i++ {
		for j := 0; j < bb[i]; j++ {
			fmt.Println(i)
		}

	}
}
