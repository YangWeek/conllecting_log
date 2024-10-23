package goroutinetest

import (
	"fmt"
	"sync"
)

// sync.Map 并发安全的map
// 原生的map 是线程不安全的

var (
	wg3 sync.WaitGroup
)

var m = make(map[int]int)
var m2 = sync.Map{}

func get(key int) int {
	return m[key]
}

func set(key int, value int) {
	m[key] = value
}

//func main() {
//	for i:=0;i<20;i++{
//		wg.Add(1)
//		go func(i int) {
//			set(i, i+100) // map中设置键值对
//			fmt.Printf("key:%v value:%v", i, get(i)) // 打印键值对
//			wg.Done()
//		}(i)
//	}
//	wg.Wait()
//}

func main() {
	for i := 0; i < 20; i++ {
		wg3.Add(1)
		go func(i int) {
			m2.Store(i, i+100) // map中设置键值对
			value, _ := m2.Load(i)
			fmt.Printf("key:%v value:%v\n", i, value) // 打印键值对
			wg3.Done()
		}(i)
	}
	wg3.Wait()
}
