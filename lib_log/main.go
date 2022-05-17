package main

import "lib_log/my_logger"

type User struct {
	name string
	age  int
}

func main() {
	logger := my_logger.NewLogger("")
	u := &User{
		name: "test",
		age:  18,
	}
	logger.Info(u)
	logger.Warn(u)
}
