package main

import "github.com/sirupsen/logrus"

func main() {
	logger := logrus.WithFields(logrus.Fields{
		"user": "admin",
		"host": "localhost",
	})
	logger.Info("info msg")

	logEntry := logrus.WithField("user", "admin")
	logEntry = logEntry.WithField("host", "localhost")
	logEntry.Info("info msg")
}
