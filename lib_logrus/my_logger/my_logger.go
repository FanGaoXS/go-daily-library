package main

import (
	"log"

	"github.com/sirupsen/logrus"
)

func logLevel(level string) logrus.Level {
	switch level {
	case "trace", "Trace", "TRACE":
		return logrus.TraceLevel
	case "debug", "Debug", "DEBUG":
		return logrus.DebugLevel
	case "info", "Info", "INFO":
		return logrus.InfoLevel
	case "warn", "Warn", "WARN":
		return logrus.WarnLevel
	case "fatal", "Fatal", "FATAL":
		return logrus.FatalLevel
	case "panic", "Panic", "PANIC":
		return logrus.PanicLevel
	}
	log.Println("please enter a correct log level for example trace, debug, info, warn, fatal, panic.")
	return logrus.InfoLevel
}

func logFormatter(formatter string) logrus.Formatter {
	switch formatter {
	case "json", "JSON":
		return &logrus.JSONFormatter{}
	case "text", "Text":
		return &logrus.TextFormatter{}
	}
	return &logrus.TextFormatter{}
}

type MyLogger struct {
	logger *logrus.Entry
}

func New(level string, isReportCaller bool, formatter string) MyLogger {
	logger := logrus.New()
	logger.SetLevel(logLevel(level))
	logger.SetReportCaller(isReportCaller)
	logger.SetFormatter(logFormatter(formatter))

	return MyLogger{
		logger: logger.WithFields(logrus.Fields{"app": "test_demo"}),
	}
}

func (m MyLogger) Trace(args ...interface{}) {
	m.logger.Trace(args...)
}

func (m MyLogger) Debug(args ...interface{}) {
	m.logger.Debug(args...)
}

func (m MyLogger) Info(args ...interface{}) {
	m.logger.Info(args...)
}

func (m MyLogger) Warn(args ...interface{}) {
	m.logger.Warn(args...)
}

func (m MyLogger) Fatal(args ...interface{}) {
	m.logger.Fatal(args...)
}

func (m MyLogger) Panic(args ...interface{}) {
	m.logger.Panic(args...)
}

func (m MyLogger) Tracef(format string, args ...interface{}) {
	m.logger.Tracef(format, args...)
}

func (m MyLogger) Debugf(format string, args ...interface{}) {
	m.logger.Debugf(format, args...)
}

func (m MyLogger) Infof(format string, args ...interface{}) {
	m.logger.Infof(format, args...)
}

func (m MyLogger) Warnf(format string, args ...interface{}) {
	m.logger.Warnf(format, args...)
}

func (m MyLogger) Fatalf(format string, args ...interface{}) {
	m.logger.Fatalf(format, args...)
}

func (m MyLogger) Panicf(format string, args ...interface{}) {
	m.logger.Panicf(format, args...)
}
