package main

import "github.com/sirupsen/logrus"

func main() {
	// 默认的日志显示级别为：Info
	// 级别由高到低分别是：Panic、Fatal、Error、Warn、Info、Debug、Trace
	// 可以手动修改日志显示的级别，即高于该级别的就会显示
	logrus.SetLevel(logrus.TraceLevel)

	// 输出调用者信息：文件名和函数，默认false
	logrus.SetReportCaller(true)

	// 输出的日志包含：
	// time：输出日志的时间
	// level：日志级别
	// msg：日志信息
	logrus.Trace("trace msg")
	logrus.Debug("debug msg")
	logrus.Info("info msg")
	logrus.Warn("warn msg")
	logrus.Error("error msg")
	logrus.Fatal("fatal msg")
	logrus.Panic("panic msg") // Fatal会直接退出程序，所以panic不会触发
}
