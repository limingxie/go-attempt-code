package goroutine

import (
	"fmt"
	"runtime"
	"time"
)

func hello() {
	fmt.Println("Hello Goroutine!")
}
func GoroutineTest1() {
	go hello() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	//因为启用go hello() 协程需要时间，所以大部分会执行前主程序会结束。要么来个 time.Sleep(time.Second)
}

func hello1(i int) {
	defer wg.Done() // goroutine结束就登记-1
	fmt.Println("Hello Goroutine!", i)
}
func GoroutineTest2() {
	for i := 0; i < 10; i++ {
		wg.Add(1) // 启动一个goroutine就登记+1
		go hello1(i)
	}
	wg.Wait() // 等待所有登记的goroutine都结束
}

func GoroutineTest3() {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 3 {
			break
		}
	}
	//主程序结束了，协程自然会结束
}

//runtime.Gosched()的 用法
func GoroutineTest4() {
	go func(s string) {
		for i := 0; i < 3; i++ {
			fmt.Print(s, " ")
		}
	}("world")
	// 主协程
	for i := 0; i < 3; i++ {
		// 切一下，再次分配任务
		runtime.Gosched()
		fmt.Print("hello", " ")
	}
	//world world world hello hello hello
}

//runtime.Goexit() 的用法
func GoroutineTest5() {
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for i := 0; i < 5; i++ {
		fmt.Print(i, " ")
		time.Sleep(time.Second)
	}
	/*
		0 B.defer
		A.defer
		1 2 3 4
	*/
}

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}

//runtime.GOMAXPROCS(count) 设定使用多少个OS线程(CPU数量)，GOMAXPROCS是m:n调度中的n
func GoroutineTest6() {
	runtime.GOMAXPROCS(2)
	go a()
	go b()
	time.Sleep(time.Second)
}

func MainGoroutine() {
	GoroutineTest6()
}
