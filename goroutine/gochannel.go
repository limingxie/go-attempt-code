package goroutine

import (
	"fmt"
	"math/rand"
	"sync"
)

func Receive(c chan int) {
	defer wg.Done()
	ret := <-c
	fmt.Println("接收成功", ret)
}

func GoChannelTest() {
	ch := make(chan int) //这是声明 无缓存通道
	wg.Add(1)
	go Receive(ch) //没有这一段话会报错  all goroutines are asleep - deadlock!
	ch <- 10       // 发送：把10发送到ch中
	fmt.Println(ch)
	wg.Wait() // 等待所有登记的goroutine都结束

	x := <-ch // 接受：从ch中接收值并赋值给变量x
	<-ch      // 接受：从ch中接收值，忽略结果
	close(ch) // 关闭
	//无缓存通道式必须要有接受者才能用。

	fmt.Println(x)

	/*
		接收成功 10
		0xc00008c060
		fatal error: all goroutines are asleep - deadlock!
	*/
}

var wg sync.WaitGroup

//缓存通道
func GoChannelTest1() {
	ch := make(chan int, 10) //这是声明 有缓存通道

	ch <- 10
	ch <- 20

	fmt.Printf("%p, %+v\n", ch, <-ch) //0xc0000b4000, 10
	fmt.Printf("%p, %+v\n", ch, <-ch) //0xc0000b4000, 20

	ch <- 30
	ch <- 40
	ch <- 50
	wg.Add(3)
	fmt.Printf("%v, %+v\n", len(ch), cap(ch)) //3, 10

	go Receive(ch) //接收成功 30
	go Receive(ch) //接收成功 40
	go Receive(ch) //接收成功 50
	wg.Wait()      // 等待所有登记的goroutine都结束
}

//Close
func GoChannelTest2() {
	c := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
	}()
	for {
		if data, ok := <-c; ok {
			fmt.Print(data, " ")
		} else {
			break
		}
	}
	fmt.Println("main结束")
	//0 1 2 3 4 main结束
}

//Close 判断通道是否被关闭，通常使用的是for range的方式。
func GoChannelTest3() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值ok=false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中从ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range循环
		fmt.Println(i)
	}
}

/*-----------------------------------------------------------------------------*/
//单项通道
func counter(out chan<- int) {
	for i := 0; i < 100; i++ {
		out <- i
	}
	close(out)
}

func squarer(out chan<- int, in <-chan int) {
	for i := range in {
		out <- i * i
	}
	close(out)
}
func printer(in <-chan int) {
	for i := range in {
		fmt.Println(i)
	}
}

func GoChannelTest4() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go counter(ch1)
	go squarer(ch2, ch1)
	printer(ch2)
}

/*-----------------------------------------------------------------------------*/

func MainChannel() {
	gotickerTest()
}

/*-----------------------------------------------------------------------------*/
//goroutine pool
type Job struct {
	// id
	Id int
	// 需要计算的随机数
	RandNum int
}

type Result struct {
	// 这里必须传对象实例
	job *Job
	// 求和
	sum int
}

func GoChannelTest5() {
	// 需要2个管道
	// 1.job管道
	jobChan := make(chan *Job, 128)
	// 2.结果管道
	resultChan := make(chan *Result, 128)
	// 3.创建工作池
	createPool(64, jobChan, resultChan)
	// 4.开个打印的协程
	go func(resultChan chan *Result) {
		// 遍历结果管道打印
		for result := range resultChan {
			fmt.Printf("job id:%v randnum:%v result:%d\n", result.job.Id,
				result.job.RandNum, result.sum)
		}
	}(resultChan)
	var id int
	// 循环创建job，输入到管道
	for {
		id++
		// 生成随机数
		r_num := rand.Int()
		job := &Job{
			Id:      id,
			RandNum: r_num,
		}
		jobChan <- job
	}
}

// 创建工作池
// 参数1：开几个协程
func createPool(num int, jobChan chan *Job, resultChan chan *Result) {
	// 根据开协程个数，去跑运行
	for i := 0; i < num; i++ {
		go func(jobChan chan *Job, resultChan chan *Result) {
			// 执行运算
			// 遍历job管道所有数据，进行相加
			for job := range jobChan {
				// 随机数接过来
				r_num := job.RandNum
				// 随机数每一位相加
				// 定义返回值
				var sum int
				for r_num != 0 {
					tmp := r_num % 10
					sum += tmp
					r_num /= 10
				}
				// 想要的结果是Result
				r := &Result{
					job: job,
					sum: sum,
				}
				//运算结果扔到管道
				resultChan <- r
			}
		}(jobChan, resultChan)
	}
}

/*-----------------------------------------------------------------------------*/
