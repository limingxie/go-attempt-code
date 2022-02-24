package list

import (
	"fmt"
)

func printArrary(arr *[3]int) {
	arr[0] = 7
	fmt.Printf("%+v", arr)
	fmt.Println("")
}

func MainList() {
	var arr [3]int
	printArrary(&arr)
	fmt.Printf("%+v", arr)
	fmt.Println("")

	var arr1 [3]int = [3]int{1, 2, 3}
	printArrary(&arr1)
	fmt.Printf("%+v", arr1)
	fmt.Println("")
}

func arrayTest(source []int, target int) {
	for i, s := range source {
		for j := i + 1; j < len(source); j++ {
			if s+source[j] == target {
				fmt.Println(i, j)
			}
		}
	}
}

func MainList1() {
	s := []int{1, 3, 5, 8, 7, 9, 5}
	arrayTest(s, 10)
}
