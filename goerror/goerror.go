package goerror

import "fmt"

func GoErrorTest1() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
	}()

	panic("panic error!") // print panic error!
}
func GoErrorTest2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	var ch chan int = make(chan int, 10)
	close(ch)
	ch <- 1
	// print send on closed channel
}

func GoErrorTest3() {
	defer func() {
		fmt.Println(recover())
	}()

	defer func() {
		panic("defer panic")
	}()

	panic("test panic")
	// print defer panic
}

func GoErrorTest4() {
	defer func() {
		fmt.Println(recover()) //有效
	}()
	defer recover()              //无效！
	defer fmt.Println(recover()) //无效！
	defer func() {
		func() {
			println("defer inner")
			recover() //无效！
		}()
	}()

	panic("test panic")

	// defer inner
	// <nil>
	// test panic
}

// func except() {
//     fmt.Println(recover())
// }

// func test() {
//     defer except()
//     panic("test panic")
// }

// func main() {
//     test() //test panic
// }

func GoErrorTest5(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)

	// func main() {
	// 	test(2, 1)
	// }
	// result => x / y = 0
}

func GoErrorTest6() {
}

func GoErrorTest7() {
}

func GoErrorTest8() {
}

func MainGoerror() {
	GoErrorTest2()
}
