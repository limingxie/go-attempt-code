package tsort

import (
	"fmt"
)

func ShellSort() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for step := len(data) / 2; step > 0; step /= 2 {
		//开始插入排序
		for i := step; i < len(data); i++ {
			//满足条件则插入
			for j := i - step; j >= 0 && data[j+step] < data[j]; j -= step {
				data[j], data[j+step] = data[j+step], data[j]
			}
		}
	}

	fmt.Println(data)
}

func ShellSort1() {
	data := []int{8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}

	for step := len(data) / 2; step > 0; step = step / 2 {
		for i := 0; i < step; i++ {
			for j := i; j < len(data); j = j + step {
				for k := i; k < j; k = k + step {
					if data[k] > data[j] {
						v := data[j]
						for l := j; l > k; l = l - step {
							data[l] = data[l-step]
						}
						data[k] = v
					}
				}

			}
		}
	}

	fmt.Println(data)
}
