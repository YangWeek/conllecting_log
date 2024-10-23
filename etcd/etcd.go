package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"renzhi/common"
	"renzhi/tailfile"
	"time"
)

var (
	client *clientv3.Client
)

func Init(addres []string) (err error) {
	client, err = clientv3.New(clientv3.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		logrus.Error("etcd connect failed, err is %v\n", err)
		return err
	}
	return nil
}

// 拉取etcd 日志收集的配置项
func GetConf(key string) (confetcd []*common.CollectEntry, err error) {
	// get
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	resp, err := client.Get(ctx, key)
	if err != nil {
		logrus.Errorf("get from etcd failed, err:%v\n", err)
		return
	}
	if len(resp.Kvs) == 0 {
		logrus.Warnf("can't get any value by key:%s from etcd\n", key)
		return
	}
	keyValues := resp.Kvs[0]
	err = json.Unmarshal(keyValues.Value, &confetcd)
	if err != nil {
		logrus.Errorf("unmarshal value from etcd failed, err:%v\n", err)
		return
	}
	logrus.Debugf("load conf from etcd success, conf:%v\n", confetcd)
	return confetcd, err
}

// 监控etcd
func WatchConf(key string) {
	for {
		watchChan := client.Watch(context.Background(), key)
		var newConf []*common.CollectEntry
		for wresp := range watchChan {
			logrus.Infof("get new conf from etcd")
			for _, evt := range wresp.Events {
				fmt.Printf("type:%s, key;%s, values:%s\n", evt.Type, evt.Kv.Key, evt.Kv.Value)
				if evt.Type == clientv3.EventTypeDelete {
					tailfile.Manager.PutEtcdNewConf(newConf)
					continue
				}
				err := json.Unmarshal(evt.Kv.Value, &newConf)
				if err != nil {
					logrus.Errorf("json unmarshal new conf failed, err is %v\n", err)
					continue
				}
				// 创建新的tailTask
				tailfile.Manager.PutEtcdNewConf(newConf)
			}
		}
	}
}
