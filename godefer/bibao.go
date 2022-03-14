package godefer

import "fmt"

func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

/*
func main() {
    c := a()
    c()
    c()
    c()

    a() //不会输出i
}
	1
    2
    3
*/

//闭包复制的是原对象指针，这就很容易解释延迟引用现象。
func test() func() {
	x := 100
	fmt.Printf("x (%p) = %d\n", &x, x)

	return func() {
		fmt.Printf("x (%p) = %d\n", &x, x)
	}
}

/*
func main() {
    f := test()
    f()
}
    x (0xc42007c008) = 100
    x (0xc42007c008) = 100
*/

//外部引用函数参数局部变量
func add(base int) func(int) int {
	return func(i int) int {
		base += i
		return base
	}
}

/*
func MainBibao() {
	tmp1 := add(10)
	fmt.Println(tmp1(1), tmp1(2))
	// 此时tmp1和tmp2不是一个实体了
	tmp2 := add(100)
	fmt.Println(tmp2(1), tmp2(2))
}
	11 13
	101 103
*/

// 返回2个函数类型的返回值
func test01(base int) (func(int) int, func(int) int) {
	// 定义2个函数，并返回
	// 相加
	add := func(i int) int {
		base += i
		return base
	}
	// 相减
	sub := func(i int) int {
		base -= i
		return base
	}
	// 返回
	return add, sub
}

/*
func MainBibao() {
	f1, f2 := test01(10)
	// base一直是没有消
	fmt.Println(f1(1), f2(2))
	// 此时base是9
	fmt.Println(f1(3), f2(4))
}
	11 9
	12 8
*/

func fibonaci(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonaci(i-1) + fibonaci(i-2)
}

/*
func main() {
    var i int
    for i = 0; i < 10; i++ {
        fmt.Printf("%d\n", fibonaci(i))
    }
}
    0
    1
    1
    2
    3
    5
    8
    13
    21
    34
*/
