package goroutinetest

import (
	"fmt"
	"sync"
	"time"
)

// 读写互斥锁

var (
	x1  int64
	wg2 sync.WaitGroup
	//lock2 sync.Mutex
	rwLock sync.RWMutex
)

func read() {
	//lock.Lock()
	rwLock.RLock()
	time.Sleep(time.Millisecond)
	//lock.Unlock()
	rwLock.RUnlock()
	wg.Done()
}

func write() {
	//lock.Lock()
	rwLock.Lock()
	x = x + 1
	time.Sleep(time.Millisecond * 10)
	//lock.Unlock()
	rwLock.Unlock()
	wg2.Done()
}

func main() {
	start := time.Now()

	for i := 0; i < 1000; i++ {
		wg2.Add(1)
		go read()
	}

	for i := 0; i < 10; i++ {
		wg2.Add(1)
		go write()
	}
	wg2.Wait()
	fmt.Println(time.Now().Sub(start))
}
