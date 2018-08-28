package logger

import "github.com/Sirupsen/logrus"

var logger = logrus.New()

type Fields map[string]interface{}

func Info(args ...interface{}) {
	logger.Info(args...)
}

func Debug(args ...interface{}) {
	logger.Debug(args...)
}

func Warn(args ...interface{}) {
	logger.Warn(args...)
}

func Error(args ...interface{}) {
	logger.Error(args...)
}

func Fatal(args ...interface{}) {
	logger.Fatal(args...)
}

func WithFields(fields Fields) *logrus.Entry {
	f := make(logrus.Fields)
	for index, property := range fields {
		f[index] = property
	}
	return logrus.WithFields(f)
}
