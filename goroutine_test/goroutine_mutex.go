package goroutinetest

import (
	"fmt"
	"sync"
)

// 多个goroutine并发操作全局变量x

var (
	x    int64
	wg1  sync.WaitGroup
	lock sync.Mutex // 互斥锁
)

func add() {
	for i := 0; i < 5000; i++ {
		lock.Lock() // 加锁
		x = x + 1
		lock.Unlock() // 释放锁
	}
	wg1.Done()
}

func main() {
	wg1.Add(2)
	go add()
	go add()
	wg1.Wait()
	fmt.Println(x)
}
