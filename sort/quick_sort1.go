package tsort

import "fmt"

func QuickSort_a() {
	data := []int{30, 8, 100, 99, 50, 22, 15, 16, 2, 99, 1000, 999, 1}
	quickSort_a3(data, 0, len(data)-1)
	fmt.Println(data)
}

func quickSort_a1(values []int, left, right int) {
	temp := values[left]
	t := left
	s, e := left, right

	for s <= e {
		for e >= t && values[e] >= temp {
			e--
		}
		if e >= t {
			values[t] = values[e]
			t = e
		}

		for s <= t && values[s] <= temp {
			s++
		}
		if s <= t {
			values[t] = values[s]
			t = s
		}
	}
	values[t] = temp
	if t-left > 1 {
		quickSort_a1(values, left, t-1)
	}
	if right-t > 1 {
		quickSort_a1(values, t+1, right)
	}
}

// 第二种写法
func quickSort_a2(values []int) {
	if len(values) <= 1 {
		return
	}
	mid, i := values[0], 1
	head, tail := 0, len(values)-1
	for head < tail {
		if values[i] > mid {
			values[i], values[tail] = values[tail], values[i]
			tail--
		} else {
			values[i], values[head] = values[head], values[i]
			head++
			i++
		}
	}
	values[head] = mid
	quickSort_a2(values[:head])
	quickSort_a2(values[head+1:])
}

// 第三种写法
func quickSort_a3(a []int, left int, right int) {
	if left >= right {
		return
	}
	explodeIndex := left
	for i := left + 1; i <= right; i++ {
		if a[left] >= a[i] {
			//分割位定位++
			explodeIndex++
			a[i], a[explodeIndex] = a[explodeIndex], a[i]
		}
	}
	//起始位和分割位
	a[left], a[explodeIndex] = a[explodeIndex], a[left]

	quickSort_a3(a, left, explodeIndex-1)
	quickSort_a3(a, explodeIndex+1, right)
}
