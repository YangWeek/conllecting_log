package tailfile

import (
	"context"
	"github.com/IBM/sarama"
	"github.com/hpcloud/tail"
	"github.com/sirupsen/logrus"
	"renzhi/kafka"
	"strings"
	"time"
)

//   tail是一个用于监控文件的库，它允许你跟踪文件的末尾几行内容，特别是当文件持续增长时。
//  tail库通常用于日志文件监控，可以实时读取日志文件的新增内容

type tailTask struct {
	path       string
	topic      string
	tailObject *tail.Tail

	ctx    context.Context
	cancel context.CancelFunc
}

//var tailObject *tail.Tail

func gettailconfig() tail.Config {
	cfg := tail.Config{
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, // 从文件末尾开始读取
		Follow:    true,                                 // 跟随文件，监控新增内容
		ReOpen:    true,                                 // 文件被截断后重新打开
		MustExist: false,                                // 文件不需要在开始时就存在
		Poll:      true,                                 // 使用轮询模式，如果文件不支持事件通知
	}
	return cfg
}

func newTailTask(path string, topic string) (*tailTask, error) {
	cfg := gettailconfig()
	ctx, cancel := context.WithCancel(context.Background())
	task := &tailTask{
		path:   path,
		topic:  topic,
		ctx:    ctx,
		cancel: cancel,
	}
	tailObject, err := tail.TailFile(path, cfg)
	task.tailObject = tailObject
	return task, err
}

func (t *tailTask) run() (err error) {
	// logfile ->Tailobj ->client ->kafka -> kafka 消费者
	logrus.Infof("collect file path, path is %v\n", t.path)
	// 死循环监控
	for {
		select {
		case <-t.ctx.Done():
			logrus.Infof("the task for path:%s is stop...", t.path)
			t.tailObject.Cleanup() // 销毁这个 监控
			return nil
		case line, ok := <-t.tailObject.Lines:
			if !ok {
				logrus.Warnf("tail file ropen is failed, filename is %v\n", t.path)
				time.Sleep(time.Second)
				continue
			}
			if len(strings.Trim(line.Text, "\r")) == 0 {
				logrus.Info("出现空行了")
				continue
			}
			msg := &sarama.ProducerMessage{}
			msg.Topic = t.topic
			msg.Value = sarama.StringEncoder(line.Text)
			kafka.SetKafkaChanMes(msg) // 异步进行
		}

	}
}
