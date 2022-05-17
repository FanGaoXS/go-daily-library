package my_logger

import (
	"log"
	"os"
)

const (
	LogInfo  = "INFO"
	LogDebug = "DEBUG"
	LogWarn  = "WARN"
	LogError = "ERROR"
)

type MyLogger struct {
	logger *log.Logger
}

func NewLogger(prefix string) *MyLogger {
	return &MyLogger{
		logger: log.New(os.Stdout, prefix, log.LstdFlags),
	}
}

func (m MyLogger) println(level string, args ...interface{}) {
	var msg []interface{}
	msg = append(msg, level)   // level of logger
	msg = append(msg, args...) // content of logger
	m.logger.Println(msg...)
}

func (m *MyLogger) Info(args ...interface{}) {
	m.println(LogInfo, args...)
}

func (m *MyLogger) Debug(args ...interface{}) {
	m.println(LogDebug, args...)
}

func (m *MyLogger) Warn(args ...interface{}) {
	m.println(LogWarn, args...)
}

func (m *MyLogger) Error(args ...interface{}) {
	m.println(LogError, args...)
}
