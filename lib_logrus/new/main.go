package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.New()

	logger.Info("info msg")
}
