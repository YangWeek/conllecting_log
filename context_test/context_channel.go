package context_test

import (
	"context"
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func work(ctx context.Context) {
	defer wg.Done()
LABLE:
	for {
		fmt.Println("worker")
		select {
		case <-ctx.Done():
			break LABLE
		default:

		}
	}
}

func main() {
	ctx, canncel := context.WithCancel(context.Background())
	wg.Add(1)
	go work(ctx)
	// time.Sleep(time.Second * 4)
	canncel() // 类似于向管道写入数据 使协程结束
	wg.Wait()
	fmt.Println("over")
}
