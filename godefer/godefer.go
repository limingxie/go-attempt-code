package godefer

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
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
	y += 10
	println("x =", x, "y =", y)
	/*
	   x = 20 y = 30
	   defer: 10 30
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

func deferTest3() error {
	res, err := http.Get("http://xxxxxxxxxx")
	if res != nil {
		defer res.Body.Close()
	}

	if err != nil {
		return err
	}

	// ..code...

	return nil
}

func deferTest4() {
	var run func() = nil
	defer run()
	fmt.Println("runs")
}

/*
func main() {
    defer func() {
        if err := recover(); err != nil {
            fmt.Println(err)
        }
    }()
    deferTest4()
}
runs
runtime error: invalid memory address or nil pointer dereference
*/

func deferTest5() (i int) {

	i = 0
	defer func() {
		fmt.Println(i)
	}()

	return 2 // print 2
}

func deferTest6() (err error) {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}

	if f != nil {
		defer func() {
			if ferr := f.Close(); ferr != nil {
				err = ferr
			}
		}()
	}

	// ..code...

	return nil
}

func deferTest7() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}()
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}()
	}

	return nil //输出结果： defer close book.txt err close ./another-book.txt: file already closed
}

func deferTest7_() error {
	f, err := os.Open("book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close book.txt err %v\n", err)
			}
		}(f)
	}

	// ..code...

	f, err = os.Open("another-book.txt")
	if err != nil {
		return err
	}
	if f != nil {
		defer func(f io.Closer) {
			if err := f.Close(); err != nil {
				fmt.Printf("defer close another-book.txt err %v\n", err)
			}
		}(f)
	}

	return nil
}

func Maingodefer() {
	// testPanic(0)
	deferTest2(10, 0)
}
