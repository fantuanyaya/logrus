package main

import "github.com/sirupsen/logrus"

// 导入  go get github.com/sirupsen/logrus
func main() {

	// 更改日志级别
	logrus.SetLevel(logrus.DebugLevel)

	logrus.Debugln("Debugln") // debug没有输出是因为logrus的日志输出等级是info
	logrus.Infoln("Infoln")
	logrus.Info("Info")
	logrus.Infof("%v", "Infof")
	logrus.Warnln("Warnln")
	logrus.Warn("Warn")

	logrus.Errorln("Errorln")
	logrus.Println("Println")

	//PanicLevel  // 会抛一个异常
	//FatalLevel  // 打印日志之后就会退出
	//ErrorLevel
	//WarnLevel
	//InfoLevel
	//DebugLevel
	//TraceLevel  // 低级别

}
