package tsort

import (
	"fmt"
)

func QuickSort_b() {
	data := []int{30, 8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}
	quickSort_b10(data, 0, len(data)-1)
	fmt.Println(data)
}

func quickSort_b10(data []int, left, right int) {
	if left >= right {
		return
	}

}

func quickSort_test(data []int, left int, right int) {
	if left >= right {
		return
	}

	index := left + 1
	for i := index; i <= right; i++ {
		if data[left] > data[i] {
			data[i], data[index] = data[index], data[i]
			index++
		}
	}

	data[left], data[index-1] = data[index-1], data[left]
	quickSort_test(data, left, index-2)
	quickSort_test(data, index, right)
}

func quickSort_b9(data []int, left, right int) {
	if left >= right {
		return
	}
	a, t := left, left
	z := right
	temp := data[t]
	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if t < z {
			data[t] = data[z]
			t = z
		}
		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp
	if left < t {
		quickSort_b9(data, left, t)
	}
	if t < right {
		quickSort_b9(data, t+1, right)
	}
}
func quickSort_b8(data []int, left, right int) {
	if left >= right {
		return
	}
	a, t := left, left
	z := right
	temp := data[t]

	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if t < z {
			data[t] = data[z]
			t = z
		}
		for temp > data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp
	if left < t {
		quickSort_b8(data, left, t)
	}
	if t+1 < right {
		quickSort_b8(data, t+1, right)
	}
}
func quickSort_b7(data []int, left, right int) {
	if left >= right {
		return
	}
	a, t := left, left
	z := right
	temp := data[t]
	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if z > t {
			data[t] = data[z]
			t = z
		}
		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp

	if left < t {
		quickSort_b7(data, left, t)
	}
	if t+1 < right {
		quickSort_b7(data, t+1, right)
	}
}
func quickSort_b6(data []int, left, right int) {
	if left >= right {
		return
	}
	a, t := left, left
	z := right
	temp := data[t]
	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if z > t {
			data[t] = data[z]
			t = z
		}
		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp

	if left < t {
		quickSort_b6(data, left, t)
	}
	if t+1 < right {
		quickSort_b6(data, t+1, right)
	}
}
func quickSort_b5(data []int, left, right int) {
	if left >= right {
		return
	}
	a, t := left, left
	z := right
	temp := data[t]

	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if z > t {
			data[t] = data[z]
			t = z

		}

		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			a = t
		}
	}
	data[t] = temp
	if left < t {
		quickSort_b5(data, left, t)
	}
	if right > t {
		quickSort_b5(data, t+1, right)
	}
}
func quickSort_b4(data []int, left, right int) {
	if left >= right {
		return
	}

	a, t := left, left
	z := right
	temp := data[t]

	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if z > t {
			data[t] = data[z]
			t = z
		}

		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp

	if left < t {
		quickSort_b4(data, left, t)
	}
	if t+1 < right {
		quickSort_b4(data, t+1, right)
	}
}
func quickSort_b3(data []int, left, right int) {
	if left >= right {
		return
	}

	a, t := left, left
	z := right
	temp := data[t]

	for a < z {
		for temp <= data[z] && a < z {
			z--
		}
		if z > t {
			data[t] = data[z]
			t = z
		}
		for temp >= data[a] && a < z {
			a++
		}
		if a < t {
			data[t] = data[a]
			t = a
		}
	}
	data[t] = temp
	if left < t {
		quickSort_b3(data, left, t)
	}
	if t+1 < right {
		quickSort_b3(data, t+1, right)
	}
}
func quickSort_b1(data []int, left, right int) {
	if left >= right {
		return
	}

	s := left
	e := right
	t := left
	temp := data[t]

	for e > s {
		for data[e] >= temp && e > s {
			e--
		}
		if t < e {
			data[t] = data[e]
			data[e] = temp
			t = e
		}

		for data[s] <= temp && e > s {
			s++
		}
		if t > e {
			data[t] = data[s]
			data[s] = temp
			t = s
		}

	}

	if t > left {
		quickSort_b1(data, left, t)
	}
	if right > t+1 {
		quickSort_b1(data, t+1, right)
	}
}
func quickSort_b2(data []int, left, right int) {
	if left >= right {
		return
	}

	s := left
	t := s
	e := right

	temp := data[t]

	for s < e {
		for temp <= data[e] && s < e {
			e--
		}
		data[t] = data[e]
		data[e] = temp
		t = e

		for temp >= data[s] && s < e {
			s++
		}
		data[t] = data[s]
		data[s] = temp
		t = s
	}
	quickSort_b2(data, left, t)
	quickSort_b2(data, t+1, right)
}
