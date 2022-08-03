package fintech_logger

import (
	"github.com/sirupsen/logrus"
)

type LoggerType string

const (
	JsonFormatter = LoggerType("JSON")
	TextFormatter = LoggerType("Text")
)

type FintechLogger struct {
	logger *logrus.Logger
}

func NewFintechLogger(logType LoggerType) *FintechLogger {
	l := new(FintechLogger)
	l.logger = logrus.New()
	switch logType {
	case JsonFormatter:
		l.logger.SetFormatter(&logrus.JSONFormatter{})
	case TextFormatter:
		l.logger.SetFormatter(&logrus.TextFormatter{})
	}
	l.logger.SetLevel(logrus.InfoLevel)
	return l
}

func (l *FintechLogger) Info(args ...interface{}) {
	l.logger.Info(args)
}

func (l *FintechLogger) Error(args ...interface{}) {
	l.logger.Error(args)
}

func (l *FintechLogger) Fatal(args ...interface{}) {
	l.logger.Fatal(args)
}

func (l *FintechLogger) Debug(args ...interface{}) {
	l.logger.Debug(args)
}

func (l *FintechLogger) Fatalf(format string, args ...interface{}) {
	l.logger.Fatalf(format, args)
}
