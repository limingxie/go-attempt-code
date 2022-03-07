package list

import "fmt"

func SparseArrayTest() {
	var array2 [10][10]int
	array2[1][1] = 1
	array2[2][2] = 2
	array2[3][3] = 3

	fmt.Println("-----------------打印二维数组-----------------")

	for i := range array2 {
		for j := range array2[i] {
			fmt.Printf("%d  ", array2[i][j])
		}
		fmt.Println("")
	}
	sparseArray := ArrayToSparseArray(array2)

	fmt.Println("-----------------打印稀疏数组-----------------")

	for i := range sparseArray {
		for j := range sparseArray[i] {
			fmt.Printf("%d  ", sparseArray[i][j])
		}
		fmt.Println("")
	}

	array := SparseToArray(sparseArray)

	fmt.Println("-----------------打印稀疏数组恢复到二维数组的值-----------------")
	for i := range array {
		for j := range array2[i] {
			fmt.Printf("%d  ", array2[i][j])
		}
		fmt.Println("")
	}

}

func ArrayToSparseArray(array2 [10][10]int) [][3]int {
	var sum int = 0
	for i := range array2 {
		for j := range array2[i] {
			if array2[i][j] != 0 {
				sum++
			}
		}
	}

	var sparseArray [][3]int
	sparseArray = append(sparseArray, [3]int{10, 10, sum})

	for i := range array2 {
		for j := range array2[i] {
			if array2[i][j] != 0 {
				sparseArray = append(sparseArray, [3]int{i, j, array2[i][j]})
			}
		}
	}
	return sparseArray
}

func SparseToArray(sparseArray [][3]int) [10][10]int {
	var array [10][10]int
	for i := range sparseArray {
		for j := range sparseArray[i] {
			array[i][j] = sparseArray[i][j]
		}
	}
	return array
}

func MainSparseArray() {
	SparseArrayTest()
}
