package slice

import (
	"fmt"
	"strings"
)

//slice声明方式
func sliceTest() {
	//1.声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空")
	} else {
		fmt.Println("不是空")
	}
	// 2.:=
	s2 := []int{}
	// 3.make()
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)
	// 4.初始化赋值
	// make([]type, len, cap)
	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)
	s5 := []int{1, 2, 3}
	fmt.Println(s5)
	// 5.从数组切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6)
}

//slice截取
func sliceTest1() {
	var arr = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	var slice0 []int = arr[2:8]
	var slice1 []int = arr[0:6]        //可以简写为 var slice []int = arr[:end]
	var slice2 []int = arr[5:10]       //可以简写为 var slice[]int = arr[start:]
	var slice3 []int = arr[0:len(arr)] //var slice []int = arr[:]
	var slice4 = arr[:len(arr)-1]

	fmt.Printf("arr: %v\n", arr)
	fmt.Printf("slice0: %v\n", slice0)
	fmt.Printf("slice1: %v\n", slice1)
	fmt.Printf("slice2: %v\n", slice2)
	fmt.Printf("slice3: %v\n", slice3)
	fmt.Printf("slice4: %v\n", slice4)
}

//slice截取的值，会指向原有的值。
func sliceTest2() {
	data := []int{0, 1, 2, 3, 4, 5}

	s := data[2:4]
	s[0] += 100
	s[1] += 200

	fmt.Println(s)
	fmt.Println(data)

	//需要注意的现象
	data1 := make([]int, 6, 30)
	for i := 0; i < 6; i++ {
		data1[i] = i
	}
	s1 := data1[0:2]
	s1 = append(s1, 10) // 这个操作会影响 s1,data1 值， 因为都看同一个值

	fmt.Printf("%v\n", s1)
	fmt.Printf("%v\n", data1)

	data2 := make([]int, 6, 30)
	for i := 0; i < 6; i++ {
		data2[i] = i
	}
	s2 := data2[3:] //s2的值是 指向s2[0] 所以s2和data2的指针地址不一样，但是s2[1],和data2[4] 指向的都是同一个内存空间。
	s2[0] = 10
	// s2 = append(s2, 10)

	fmt.Printf("%v, %v, %v, %p\n", s2, len(s2), cap(s2), &s2[1])
	fmt.Printf("%v, %v, %v, %p\n", data2, len(data2), cap(data2), &data2[4])

}

//声明slice的时候 len 和 cap 的关系
func sliceTest3() {
	s1 := []int{0, 1, 2, 3, 8: 100} // 通过初始化表达式构造，可使用索引号。
	fmt.Println(s1, len(s1), cap(s1))

	s2 := make([]int, 6, 8) // 使用 make 创建，指定 len 和 cap 值。
	fmt.Println(s2, len(s2), cap(s2))

	s3 := make([]int, 6) // 省略 cap，相当于 cap = len。
	fmt.Println(s3, len(s3), cap(s3))
}

func sliceTest4() {
	s := []int{0, 1, 2, 3}
	p := &s[2] // *int, 获取底层数组元素指针。
	*p += 100

	fmt.Println(s)
}

func sliceTest5() {
	data := [][]int{
		[]int{1, 2, 3},
		[]int{100, 200},
		[]int{11, 22, 33, 44},
	}
	fmt.Println(data)
}

func sliceTest6() {
	d := [5]struct {
		x int
	}{}

	s := d[:] //指向同一个内存空间

	d[1].x = 10
	s[2].x = 20

	fmt.Printf("%p, %p, %p, %p, %p\n", &d, &d[0], s, &s, &s[0])
	fmt.Printf("d: %v,  s: %v\n", d, s)

	s1 := d //新建一个内存空间
	s1[3].x = 30

	fmt.Printf("%p, %p, %p, %p\n", &d, &d[0], &s1, &s1[0])
	fmt.Printf("d: %v,  s: %v\n", d, s1)
}

//slice追加添加元素，append
func sliceTest7() {
	var a = []int{1, 2, 3}
	fmt.Printf("slice a : %v\n", a)
	var b = []int{4, 5, 6}
	fmt.Printf("slice b : %v\n", b)
	c := append(a, b...)
	fmt.Printf("slice c : %v\n", c)
	d := append(c, 7)
	fmt.Printf("slice d : %v\n", d)
	e := append(d, 8, 9, 10)
	fmt.Printf("slice e : %v\n", e)
}

//虽然有多余的容量cap但是超出长度，append后
func sliceTest8() {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 1)
	fmt.Printf("%p\n", &s2)

	s2[0] = 3

	fmt.Println(s1, s2)
}

func sliceTest9() {
	s1 := make([]int, 0, 5)
	fmt.Printf("%p\n", &s1)

	s1 = append(s1, 1)
	fmt.Printf("%p\n", &s1)

	s2 := append(s1, 2)
	fmt.Printf("%p\n", &s2)

	s2[0] = 10

	fmt.Println(s1, s2)
}

func sliceTest10() {
	data := [...]int{0, 1, 2, 3, 4, 10: 0}
	s := data[:2:3]

	s = append(s, 100, 200) // 一次 append 两个值，超出 s.cap 限制。

	fmt.Println(s, data)         // 重新分配底层数组，与原数组无关。
	fmt.Println(&s[0], &data[0]) // 比对底层数组起始指针。
}

func sliceTest11() {
	s := make([]int, 0, 1)
	c := cap(s)

	for i := 0; i < 50; i++ {
		s = append(s, i)
		if n := cap(s); n > c {
			fmt.Printf("cap: %d -> %d\n", c, n)
			c = n
		}
	}
}

func sliceTest12() {
	s1 := []int{1, 2, 3, 4, 5}
	fmt.Printf("slice s1 : %v\n", s1)
	s2 := make([]int, 10)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	s3 := []int{1, 2, 3}
	fmt.Printf("slice s3 : %v\n", s3)
	s3 = append(s3, s2...)
	fmt.Printf("appended slice s3 : %v\n", s3)
	s3 = append(s3, 4, 5, 6)
	fmt.Printf("last slice s3 : %v\n", s3)
}

func sliceTest13() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println("array data : ", data)
	s1 := data[8:]
	s2 := data[:5]
	fmt.Printf("slice s1 : %v\n", s1)
	fmt.Printf("slice s2 : %v\n", s2)
	copy(s2, s1)
	fmt.Printf("copied slice s1 : %v\n", s1)
	fmt.Printf("copied slice s2 : %v\n", s2)
	fmt.Println("last array data : ", data)
}

func sliceTest14() {
	data := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice := data[:]
	for index, value := range slice {
		fmt.Printf("inde : %v , value : %v\n", index, value)
	}
}

func sliceTest15() {
	var a = []int{1, 3, 4, 5}
	fmt.Printf("slice a : %v , len(a) : %v\n", a, len(a))
	b := a[1:2]
	fmt.Printf("slice b : %v , len(b) : %v\n", b, len(b))
	c := b[0:3]
	fmt.Printf("slice c : %v , len(c) : %v\n", c, len(c))
}

func sliceTest16() {
	str := "hello world"
	s1 := str[0:5]
	fmt.Println(s1)

	s2 := str[6:]
	fmt.Println(s2)
}

func sliceTest17() {
	str := "Hello world"
	s := []byte(str) //中文字符需要用[]rune(str)
	s[6] = 'G'
	s = s[:8]
	s = append(s, '!')
	str = string(s)
	fmt.Println(str)
}

func sliceTest18() {
	str := "你好，世界！hello world！"
	s := []rune(str)
	s[3] = '够'
	s[4] = '浪'
	s[12] = 'g'
	s = s[:14]
	str = string(s)
	fmt.Println(str)
}

func sliceTest19() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	d1 := slice[6:8]
	fmt.Println(d1, len(d1), cap(d1))
	d2 := slice[:6:8]
	fmt.Println(d2, len(d2), cap(d2))
}

func sliceTest20() {
	slice := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	str := strings.Replace(strings.Trim(fmt.Sprint(slice), "[]"), " ", ",", -1)
	fmt.Println(str)
}

func MainListSlice() {
	sliceTest9()

	fmt.Println("-----------------------------")
	// fmt.Printf("%p, %p, %p", s1, s2, s3)
}
