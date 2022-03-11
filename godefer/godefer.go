package godefer

import (
	"errors"
	"fmt"
)

func testPanic(x int) {
	defer func() {
		println("1")
	}()
	defer func() {
		println("2")
	}()

	defer func() {
		println("3")
		println(100 / x) // x=0 div0 异常未被捕获，逐步往外传递，最终终止进程。
		println("4")
	}()

	defer func() {
		println("5")
	}()

	/*
		5
		3
		2
		1
		panic: runtime error: integer divide by zero
	*/
}

func deferTest1() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y)
	/*
	   x = 20 y = 120
	   defer: 10 120
	*/
}

func deferTest2(a, b int) (i int, err error) {
	defer fmt.Printf("first defer err %v\n", err)
	defer func(err error) { fmt.Printf("second defer err %v\n", err) }(err)
	defer func() { fmt.Printf("third defer err %v\n", err) }()
	if b == 0 {
		err = errors.New("divided by zero!")
		return
	}

	i = a / b
	return
	/*
		third defer err divided by zero!
		second defer err <nil>
		first defer err <nil>
	*/
}

func Maingodefer() {
	// testPanic(0)
	deferTest2(10, 0)
}
