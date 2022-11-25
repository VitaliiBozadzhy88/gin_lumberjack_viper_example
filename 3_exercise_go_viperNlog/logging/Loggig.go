package logging

import (
	"io"
	"time"

	"github.com/Sirupsen/logrus"
	"gopkg.in/natefinch/lumberjack.v2"
)

func LogINFO[T comparable](s T) {
	logger := &lumberjack.Logger{
		Filename:   "3_exercise_go_viperNlog/logging/logfile.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}
	writer := io.Writer(logger)

	logFormatter := new(logrus.TextFormatter)
	logFormatter.TimestampFormat = time.RFC1123 // or RFC3339

	logrus.SetFormatter(logFormatter)
	logrus.SetOutput(writer)
	logrus.Info(s)
}
func LogERROR[T comparable](s T) {
	logger := &lumberjack.Logger{
		Filename:   "3_exercise_go_viperNlog/logging/logfile.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28, //days
	}
	writer := io.Writer(logger)

	logFormatter := new(logrus.TextFormatter)
	logFormatter.TimestampFormat = time.RFC1123 // or RFC3339

	logrus.SetFormatter(logFormatter)
	logrus.SetOutput(writer)
	logrus.Error(s)
}
