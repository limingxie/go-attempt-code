package goroutine

import (
	"fmt"
	"sync"
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
	GosyncTest3()
}
