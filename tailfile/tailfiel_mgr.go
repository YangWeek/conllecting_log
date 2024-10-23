package tailfile

import (
	"github.com/sirupsen/logrus"
	"renzhi/common"
)

// tail 日志收集的管理者 管理日志操作的对象
var (
	Manager *tailTaskManager
)

type tailTaskManager struct {
	tailTaskMap      map[string]*tailTask
	collectEntryList []*common.CollectEntry
	newConfChan      chan []*common.CollectEntry
}

func Init(allconf []*common.CollectEntry) error {
	Manager = &tailTaskManager{
		tailTaskMap:      make(map[string]*tailTask, 32),
		collectEntryList: allconf,
		newConfChan:      make(chan []*common.CollectEntry), // 无缓冲区 chan
	}
	for _, conf := range allconf {
		task, err := newTailTask(conf.Path, conf.Topic)
		if err != nil {
			logrus.Errorf("new tailtask failed, err is %v\n", err)
			continue
		}
		Manager.tailTaskMap[task.path] = task
		logrus.Infof("make a task succsue")
		go task.run()
	}

	go Manager.WatchNewConfChan() // 携程监控
	return nil
}

// 日志监听  a
func (t *tailTaskManager) WatchNewConfChan() {
	for {
		newconf := <-t.newConfChan // newconf是切片
		logrus.Infof("get new conf from etcd,conf: %v\n", newconf)
		for _, conf := range newconf {
			// 判断原来有没有conf
			if IsExist(conf) {
				continue
			}
			// 原来没有 创建
			task, err := newTailTask(conf.Path, conf.Topic)
			if err != nil {
				logrus.Errorf("new tailtask failed, err is %v\n", err)
				continue
			}
			Manager.tailTaskMap[task.path] = task
			logrus.Infof("make a task succsue")
			go task.run()
		}

		// 找出应该被删除的task
		for key, task := range Manager.tailTaskMap {
			var found int
			for _, conf := range newconf {
				if key == conf.Path {
					found = 1
					break
				}
			}
			if found == 0 {
				logrus.Infof("the task of path:%s is remove from tailTaskMap", key)
				delete(Manager.tailTaskMap, key)
				task.cancel()
			}
		}
	}
}

func IsExist(conf *common.CollectEntry) bool {
	_, ok := Manager.tailTaskMap[conf.Path]
	return ok
}

func (t *tailTaskManager) PutEtcdNewConf(conf []*common.CollectEntry) {
	t.newConfChan <- conf
}
