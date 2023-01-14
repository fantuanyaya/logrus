package main

import (
	"fmt"
	"time"
)

func InitFiled(logPath, appName string) {
	// 拿到当前日期，并以年月日输出
	fileData := time.Now().Format("2006-01-02")
	fmt.Println(fileData)
	// 创建目录

}
func main() {
	fmt.Println("hello git")
	fmt.Println("hello git")
	fmt.Println("hello git")
	fmt.Println("hello master02")
	fmt.Println("hello hot-fix commit02")
	fmt.Println("push commit")
	fmt.Println("push commit2")
	fmt.Println("push commit3")

	fmt.Println("pull commit")
	fmt.Println("pull2 commit")
	fmt.Println("pull2 commit")
	fmt.Println("pull2 commit")
}
