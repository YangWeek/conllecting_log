package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"renzhi/conf"
	"renzhi/etcd"
	sysinfo "renzhi/gopsutil_demo"
	"renzhi/kafka"
	"renzhi/tailfile"
	"time"
)

// 日志搜集项目
// 国外开源filebeat
// 搜集指定目录的日志   发到kafka中
// 单机版 可以合作多个版本的

// 数据都发到kafka中
func main() {
	err := conf.InitConfig()
	if err != nil {
		logrus.Error("config init failed, err is %v\n", err)
		return
	}

	err = kafka.InitKafka([]string{conf.Cfg.KafkaConfig.Address}, conf.Cfg.KafkaConfig.ChanSize)
	if err != nil {
		logrus.Error("kafka client init failed, err is %v\n", err)
		return
	}

	// etcd 初始化
	err = etcd.Init([]string{conf.Cfg.EtcdConfig.Address})
	if err != nil {
		logrus.Error("etcd connect failed, err is %v\n", err)
		return
	}
	confEtcd, err := etcd.GetConf(conf.Cfg.EtcdConfig.CollectLogKey)
	if err != nil {
		fmt.Errorf("etcd get conf failed, err is %v\n", err)
		return
	}
	fmt.Println(confEtcd)
	go etcd.WatchConf(conf.Cfg.EtcdConfig.CollectLogKey)

	err = tailfile.Init(confEtcd)
	if err != nil {
		logrus.Error("tial init is failed, err is %v\n", err)
		return
	}
	// 循环等待

	go sysinfo.Run(time.Second, "sysinfo")
	select {}
}
