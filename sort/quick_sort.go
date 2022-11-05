package tsort

import (
	"fmt"
)

var qq map[int]int

//快速排序
func QuickSort() {
	qq = make(map[int]int)
	qq[0] = 6
	qq[1] = 1
	qq[2] = 2
	qq[3] = 12
	qq[4] = 9
	qq[5] = 3
	qq[6] = 4
	qq[7] = 5 // fmt.Print(strconv.Itoa(i) + ":")
	qq[8] = 1
	qq[9] = 8
	qq[10] = 11
	qq[11] = 7
	// fmt.Print(strconv.Itoa(i) + ":")
	for i := 0; i < len(qq); i++ {
		// fmt.Print(strconv.Itoa(i) + ":")
		fmt.Print(qq[i])
		fmt.Print("  ")
	}
	fmt.Println("")
	fmt.Println("---------------------------------------")

	QuickSort2(0, 11)

	for i := 0; i < len(qq); i++ {
		fmt.Print(qq[i])
		fmt.Print("  ")
	}
}

func QuickSort2(left int, right int) {
	if left > right { // fmt.Print(strconv.Itoa(i) + ":")
		return
	}

	var i int = left
	var j int = right
	var temp int = qq[left]

	for i != j {
		for qq[j] >= temp && i < j {
			j--
		}
		for qq[i] <= temp && i < j {
			i++
		}
		t := qq[i]
		qq[i] = qq[j] // fmt.Print(strconv.Itoa(i) + ":")
		qq[j] = t
	}

	qq[left] = qq[i]
	qq[i] = temp

	QuickSort2(left, i-1)
	QuickSort2(i+1, right)
}

func QuickSort1(left int, right int) {
	if left >= right {
		return
	}

	var i int = left
	var j int = right
	var temp int = qq[left]

	for i != j {
		for qq[j] >= temp && i < j {
			j--
		}
		for qq[i] <= temp && i < j {
			i++
		}

		if i < j {
			v := qq[i]
			qq[i] = qq[j]
			qq[j] = v
		}
	}

	qq[left] = qq[i]
	qq[i] = temp

	QuickSort1(left, i-1)
	QuickSort1(i+1, right)

}
