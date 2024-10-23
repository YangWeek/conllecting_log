package goroutinetest

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x5  int64
	wg5 sync.WaitGroup
	//lock sync.Mutex
	rwLock5 sync.RWMutex
)

func read5() {
	//lock.Lock()
	rwLock5.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock5.RUnlock()
	wg5.Done()
}

func write5() {
	//lock.Lock()
	rwLock5.Lock()
	x5 = x5 + 1
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	rwLock5.Unlock()
	wg.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg5.Add(1)
		go read5()
	}

	for i := 0; i < 10; i++ {
		wg5.Add(1)
		go write5()
	}
	wg5.Wait()
	fmt.Println(time.Now().Sub(start))
}

// 只用互斥锁  时间要用差不多15s  消耗时间较大
// 对于一些共享内存 读的操作比较多的变量， 用读写锁比较节省时间   一共174.6362ms
