package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type Log struct {
	logger *logrus.Logger
	metadata map[string] interface{}
}

var logFile *os.File

func NewLogger() *Log{
	var level = logrus.DebugLevel
	extraData := make(map[string]interface{})
	file, _ := os.OpenFile("app_logs.log", os.O_WRONLY|os.O_CREATE|os.O_APPEND, 0666)
	logger := logrus.New()
	logger.Out = io.MultiWriter(os.Stdout, file)
	logger.Formatter = &logrus.JSONFormatter{}
	logger.Level = level
	logger.ExitFunc = func(i int) {
		if file != nil {
			file.Close()
		}
	}

	return &Log{
		logger: logger,
		metadata: extraData,
	}
}


func (l *Log) Debug(msg string, tags ...string) {
	if !l.logger.IsLevelEnabled(logrus.DebugLevel) {
		return
	}
	logrus.WithFields(l.parseFields(tags...)).Debug(msg)
}

func (l *Log) Info(msg string, tags ...string) {
	if !l.logger.IsLevelEnabled(logrus.InfoLevel) {
		return
	}
	l.logger.WithFields(l.parseFields(tags...)).Info(msg)
}

func (l *Log) Error(msg string, err error, tags ...string) {
	if !l.logger.IsLevelEnabled(logrus.ErrorLevel) {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	l.logger.WithFields(l.parseFields(tags...)).Error(msg)
}

func (l *Log) parseFields(tags ...string) logrus.Fields {
	result := make(logrus.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	for k, v := range l.metadata {
		result[k] = v
	}
	return result
}

func (l *Log) WithField(k string, v string) {
	l.metadata[k] = v
}
