package main

import (
	"github.com/sirupsen/logrus"
	"io"
	"os"
)

// 输出日志到文件
func _SetOutputField() {
	// 先打开一个文件对象,如果没有就创建当前文件
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	// 调用方法SetOutput
	logrus.SetOutput(file)

	logrus.Errorln("nihao")
}

// 输出日志到文件和控制台
func _SetOutputFieldStdout() {
	file, err := os.OpenFile("logrus.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		return
	}
	logrus.SetOutput(io.MultiWriter(file, os.Stdout))

	logrus.Errorln("nihao")
	logrus.Errorln("123")
}
func main() {
	//_SetOutputField()
	_SetOutputFieldStdout()

}
