package conf

import (
	"github.com/sirupsen/logrus"
	"gopkg.in/ini.v1"
)

var Cfg config

type config struct {
	KafkaConfig   `ini:"kafka"`
	CollectConfig `ini:"collect"`
	EtcdConfig    `ini:"etcd"`
}

type KafkaConfig struct {
	Address  string `ini:"address"`
	ChanSize int64  `ini:"chan_size"` // 地址大小
	Topic    string `ini:"topic"`
}

type CollectConfig struct {
	Logfile string `ini:"logfile"` // 日志地址
}

type EtcdConfig struct {
	Address           string `ini:"address"`
	CollectLogKey     string `ini:"collect_log_key"`
	CollectSysInfoKey string `ini:"collect_sysinfo_key"`
}

func InitConfig() (err error) {
	err = ini.MapTo(&Cfg, "./conf/config.ini")
	if err != nil {
		logrus.Error("ini MapTo failed, err is %v\n", err)
		return err
	}
	return nil
}
