package main

import "github.com/sirupsen/logrus"

func main() {
	// 设置logger在terminal打印的等级，
	// 高于该等级才打印，默认是打印info以上的日志
	logrus.SetLevel(logrus.TraceLevel)

	// 输出文件名
	logrus.SetReportCaller(true)

	//logrus.Trace("trace msg")
	//logrus.Debug("debug msg")
	//logrus.Info("info msg")
	//logrus.Warn("warn msg")
	//logrus.Error("error msg")
	//// 致命错误，出现错误时程序无法正常运转。输出日志后，程序退出
	//logrus.Fatal("fatal msg")
	//// 记录日志，然后panic
	//logrus.Panic("panic msg")

	// LogWithFields
	logger := logrus.WithFields(logrus.Fields{"id": 11, "ip": "192.0.0.1"})
	logger.Info("info msg")
}
