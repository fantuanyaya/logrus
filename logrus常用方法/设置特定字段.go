package main

import (
	"github.com/sirupsen/logrus"
)

func main() {
	// 显示样式 Text和Json  默认是text
	logrus.SetFormatter(&logrus.JSONFormatter{})

	// 给每个服务添加一个统一的字段名称
	// 设置一个字段，也可以使用链式法，循环添加多个字段
	log := logrus.WithField("user_id", "21")
	// time="2023-01-12T17:02:50+08:00" level=error msg="你好" user_id=21

	log1 := logrus.WithField("user_id", "21").WithField("user_name", "tom")
	// time="2023-01-12T17:04:18+08:00" level=error msg="你好" user_id=21 user_name=tom
	log2 := log.WithField("user_name", "tom") // 输出结果同上

	// 也可以一次性添加多个字段
	log3 := logrus.WithFields(logrus.Fields{
		"user_id":   "21",
		"user_name": "tom",
		"ip":        "127.0.0.1",
	})
	// time="2023-01-12T17:08:03+08:00" level=error msg="你好" ip=127.0.0.1 user_id=21 user_name=tom

	log.Errorln("你好")
	log1.Errorln("你好")
	log2.Errorln("你好")
	log3.Errorln("你好")

}
