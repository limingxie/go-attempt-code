package goroutine

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

var x int64
var wg1 sync.WaitGroup
var lock sync.Mutex
var rwlock sync.RWMutex

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg1.Done()
}
func GosyncTest1() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x)
}

func add1() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 解锁
	}
	wg1.Done()
}

//使用互斥锁控制流程
func GosyncTest2() {
	wg1.Add(2)
	go add1()
	go add1()
	wg1.Wait()
	fmt.Println(x)
}

func write1() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x = x + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg1.Done()
}

func read1() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg1.Done()
}

//读写锁的简单应用
func GosyncTest3() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg1.Add(1)
		go write1()
	}

	for i := 0; i < 1000; i++ {
		wg1.Add(1)
		go read1()
	}

	wg1.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}

func MainGoSync() {
	GosyncTest4()
}

var x2 int64
var l2 sync.Mutex
var wg2 sync.WaitGroup

// 普通版加函数
func add2() {
	// x = x + 1
	x2++ // 等价于上面的操作
	wg2.Done()
}

// 互斥锁版加函数
func mutexAdd() {
	l2.Lock()
	x2++
	l2.Unlock()
	wg2.Done()
}

// 原子操作版加函数
func atomicAdd() {
	atomic.AddInt64(&x, 1)
	wg2.Done()
}

func GosyncTest4() {
	start := time.Now()
	for i := 0; i < 10000; i++ {
		wg2.Add(1)
		// go add()       // 普通版add函数 不是并发安全的
		// go mutexAdd()  // 加锁版add函数 是并发安全的，但是加锁性能开销大
		go atomicAdd() // 原子操作版add函数 是并发安全，性能优于加锁版
	}
	wg2.Wait()
	end := time.Now()
	fmt.Println(x)
	fmt.Println(end.Sub(start))
}

//atomic包提供了底层的原子级内存操作，对于同步算法的实现很有用。这些函数必须谨慎地保证正确使用。
//除了某些特殊的底层应用，使用通道或者sync包的函数/类型实现同步更好。
