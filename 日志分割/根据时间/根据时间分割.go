package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

// 按时间分割日志文件
// 1. 设置输出日志的等级
// 2. 以每天为单位，超过一天从新创建一个日志文件
// 2.1 创建一个目录存放日志文件
// 2.2 拿到当前日期
// 2.3 每次写入日志文件是，判断日期是否和今天一样，如果一样直接写入文件，如果不一样则创建当前获取的日期日志文件并写入日志

type FileDataHook struct {
	file      []*os.File
	logPath   string
	fileDate  string //判断日期切换目录
	errorName string
	warnName  string
}

func (hook FileDataHook) Levels() []logrus.Level {
	// 1. 设置输出日志的等级
	return []logrus.Level{logrus.WarnLevel, logrus.ErrorLevel}
}
func (hook FileDataHook) Fire(entry *logrus.Entry) error {
	// 拿到当前时间
	timer := entry.Time.Format("2006-01-02_15-04")
	// 拿到entry里面的所有数据返回为一个字符串
	line, err := entry.String()
	if err != nil {
		return err
	}

	if hook.fileDate == timer && entry.Level == logrus.ErrorLevel {
		// 日期相等, 直接写入文件
		fileerror := hook.file[0]
		fileerror.Write([]byte(line))
		return nil
	}
	// 时间不相等,先关闭前面打开的文件,然后重新创建一个文件存放日志
	hook.file[0].Close()

	// 创建文件
	err = os.MkdirAll(fmt.Sprintf("%s/%s", hook.logPath, timer), os.ModePerm)
	if err != nil {
		return err
	}
	filename := fmt.Sprintf("%s/%s/%s.log", hook.logPath, timer, hook.errorName)

	file, err := os.OpenFile(filename, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}
	// 将日志写入新创建的文件中
	file.Write([]byte(line))
	return nil
}

// InitFiled 初始化Filed文件
func InitFiled(logPath, errorName string, warnName string) {
	// 拿到当前日期，并以年月日输出
	fileData := time.Now().Format("2006-01-02_15-04")

	// 创建多级目录MkdirAll
	err := os.MkdirAll(fmt.Sprintf("%s/%s", logPath, fileData), os.ModePerm)
	if err != nil {
		panic(err)
		return
	}
	// 先创建一个文件来存储一开始的日志文件
	fileerrosname := fmt.Sprintf("%s/%s/%s.log", logPath, fileData, errorName) // errors文件
	filewarnname := fmt.Sprintf("%s/%s/%s.log", logPath, fileData, warnName)   // warns文件
	fileerror, err := os.OpenFile(fileerrosname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		return
	}
	filewarn, err := os.OpenFile(filewarnname, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
		return
	}

	// 实例化结构体 FileDataHook
	filedatahook := FileDataHook{file: []*os.File{fileerror, filewarn}, logPath: logPath, fileDate: fileData, errorName: errorName, warnName: warnName}
	fmt.Printf("%v\n", filedatahook.file[1])
	fmt.Printf("%v\n", fileerror)

	// 将filedatahook添加到hook钩子中
	logrus.AddHook(filedatahook)
}
func main() {
	InitFiled("hook_log/log", "errors", "warns")

	for {
		logrus.Errorln("Errorln")
		time.Sleep(20 * time.Second)
		logrus.Warnln("Warnln")
	}
}
