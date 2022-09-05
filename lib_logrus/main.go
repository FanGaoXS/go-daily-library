package main

import "lib_logrus/my_logger"

func main() {

	type user struct {
		name string
		age  int
	}

	u := &user{
		name: "test",
		age:  18,
	}

	const logLevel = "info"
	const isReportCaller = false
	const logFormatter = "json"
	logger := my_logger.New(logLevel, isReportCaller, logFormatter)
	logger.Info(u)
	logger.Infof("user = %#v", u)
}
