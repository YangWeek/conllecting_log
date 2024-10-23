package main

import (
	"context"
	"github.com/sirupsen/logrus"
	"go.etcd.io/etcd/api/v3/v3rpc/rpctypes"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// 客户端 连接etcd 这个是另外的一个程序
func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logrus.Error("etcd connect failed, err is %v\n", err)
		return
	}
	defer cli.Close()

	// put 放值
	str := `[{"path":"C:/Users/SF/Desktop/works_go/renzhi/log/logagent.log","topic":"log1"}]` // json文件序列化 要用正斜杠
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)                   // 设置context时间
	_, err = cli.Put(ctx, "log_key", str)
	if err != nil {
		// handle error!
		switch err {
		case context.Canceled:
			log.Fatalf("ctx is canceled by another routine: %v", err)
		case context.DeadlineExceeded:
			log.Fatalf("ctx is attached with a deadline is exceeded: %v", err)
		case rpctypes.ErrEmptyKey:
			log.Fatalf("client-side error: %v", err)
		default:
			log.Fatalf("bad cluster endpoints, which are not etcd servers: %v", err)
		}
	}
	cancel()

	//// 取值
	//ctx, cancel = context.WithTimeout(context.Background(), time.Second*2)
	//gr, err1 := cli.Get(ctx, "log_key")
	//if err1 != nil {
	//	fmt.Printf("err  is %v\n", err)
	//}
	//for _, ev := range gr.Kvs {
	//	fmt.Printf("key is %s, value is %s\n", ev.Key, ev.Value)
	//}
}
