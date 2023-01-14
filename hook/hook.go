package main

import (
	"github.com/sirupsen/logrus"
	"os"
)

type MyHook struct {
}

// Levels 设置日志生效等级
func (hook *MyHook) Levels() []logrus.Level {
	//return logrus.AllLevels   所有等级都生效
	return []logrus.Level{logrus.ErrorLevel} // 只有error等级生效
}

// Fire Entry 设置文件filed参数
func (hook *MyHook) Fire(entry *logrus.Entry) error {
	// 比如我想在每个日志里添加上前端传过来的user_id
	//entry.Data["user_id"] = 22

	// 将error级别的日志写入到文件
	file, err := os.OpenFile("hook_errorlevel.log", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	line, _ := entry.String()
	if err != nil {
		return err
	}
	file.Write([]byte(line))
	return nil
}

func main() {
	// 使用
	logrus.AddHook(&MyHook{})
	logrus.Errorln("errors")

}
