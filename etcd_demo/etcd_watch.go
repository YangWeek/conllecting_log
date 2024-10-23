package main

//
//import (
//	"context"
//	"fmt"
//	"github.com/sirupsen/logrus"
//	clientv3 "go.etcd.io/etcd/client/v3"
//	"time"
//)
//
//func main() {
//	cli, err := clientv3.New(clientv3.Config{
//		Endpoints:   []string{"127.0.0.1:2379"},
//		DialTimeout: 5 * time.Second,
//	})
//	if err != nil {
//		logrus.Error("etcd connect failed, err is %v\n", err)
//		return
//	}
//	defer cli.Close()
//
//	watchChan := cli.Watch(context.Background(), "log_key")
//	for wresp := range watchChan {
//		for _, evt := range wresp.Events {
//			fmt.Printf("type:%s, key;%s, values:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
//		}
//	}
//}
