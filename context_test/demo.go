package context_test

//import (
//	"fmt"
//	"sync"
//	"time"
//)

//var wg sync.WaitGroup
//
//func worker(ctx context.Context) {
//	go worker2(ctx)
//LOOP:
//	for {
//		fmt.Println("worker")
//		time.Sleep(time.Second)
//		select {
//		case <-ctx.Done(): // 等待上级通知
//			break LOOP
//		default:
//		}
//	}
//	wg.Done()
//}
//
//func worker2(ctx context.Context) {
//LOOP:
//	for {
//		fmt.Println("worker2")
//		time.Sleep(time.Second)
//		select {
//		case <-ctx.Done(): // 等待上级通知
//			break LOOP
//		default:
//		}
//	}
//}
//func main() {
//	ctx, cancel := context.WithCancel(context.Background())
//	wg.Add(1)
//	go worker(ctx)
//	time.Sleep(time.Second * 3)
//	cancel() // 通知子goroutine结束
//	wg.Wait()
//	fmt.Println("over")
//}

//import (
//	"fmt"
//	"sync"
//)
//
//var wg sync.WaitGroup
//
///*
//两个goroutine，两个channel
//	1. 生成0~100的数字发送到ch1
//	2. 从ch1中取出数据计算它的平方，把结果发送到ch2中
//*/
//
//// 生成0~100的数字发送到ch1
//func f1(ch chan<- int) {
//	defer wg.Done()
//	for i := 0; i < 100; i++ {
//		ch <- i
//	}
//	close(ch)
//
//}
//
//// 从ch1中取出数据计算它的平方，把结果发送到ch2中
//func f2(ch1 <-chan int, ch2 chan<- int) {
//	defer wg.Done()
//	// 从通道中取值的方式1
//	for {
//		tmp, ok := <-ch1
//		if !ok {
//			break
//		}
//		ch2 <- tmp * tmp
//	}
//	close(ch2)
//}
//
//func main() {
//	wg.Add(2)
//	ch1 := make(chan int, 100)
//	ch2 := make(chan int, 200)
//
//	go f1(ch1)
//	go f2(ch1, ch2)
//	// 从通道中取值的方式2
//
//	for ret := range ch2 {
//		fmt.Println(ret)
//	}
//	wg.Wait()
//}

//import (
//	"fmt"
//	"sync"
//	"time"
//)
//
//// 读写互斥锁
//
//var (
//	x  int64
//	wg sync.WaitGroup
//	//lock sync.Mutex
//	rwLock sync.RWMutex
//)
//
//func read() {
//	//lock.Lock()
//	rwLock.RLock()
//	time.Sleep(time.Millisecond)
//	//lock.Unlock()
//	rwLock.RUnlock()
//	wg.Done()
//}
//
//func write() {
//	//lock.Lock()
//	rwLock.Lock()
//	x = x + 1
//	time.Sleep(time.Millisecond * 10)
//	//lock.Unlock()
//	rwLock.Unlock()
//	wg.Done()
//}
//
//func main() {
//	start := time.Now()
//
//	for i := 0; i < 1000; i++ {
//		wg.Add(1)
//		go read()
//	}
//
//	for i := 0; i < 10; i++ {
//		wg.Add(1)
//		go write()
//	}
//	wg.Wait()
//	fmt.Println(time.Now().Sub(start))
//}

// 只用互斥锁  时间要用差不多15s  消耗时间较大
// 对于一些共享内存 读的操作比较多的变量， 用读写锁比较节省时间   一共174.6362ms

//var wg sync.WaitGroup
//
//func work(ctx context.Context) {
//	defer wg.Done()
//LABLE:
//	for {
//		fmt.Println("worker")
//		select {
//		case <-ctx.Done():
//			break LABLE
//		default:
//
//		}
//	}
//}
//
//func main() {
//	ctx, canncel := context.WithCancel(context.Background())
//	wg.Add(1)
//	go work(ctx)
//	//time.Sleep(time.Millisecond * 1)
//	canncel() // 类似于向管道写入数据 使协程结束
//	wg.Wait()
//	fmt.Println("over")
//}
