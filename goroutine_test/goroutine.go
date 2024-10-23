package goroutinetest

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // sync.WaitGroup 是用于等待一组 goroutine 完成的计数器 等待组

func hell() {
	defer wg.Done()
	fmt.Println("jasdsadj")

}

func main() {
	wg.Add(1)
	go hell()
	fmt.Println("asdkaksd")
	// time.Sleep(1) 这个会使cpu 核心等待 不推荐使用
	wg.Wait()
}
