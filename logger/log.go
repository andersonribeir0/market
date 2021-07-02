package logger

import (
	"fmt"
	customLog "github.com/sirupsen/logrus"
	"io"
	"os"
	"strings"
)

type Log struct {
	Service string
}

func NewLogger(service string) *Log {
	var level = customLog.DebugLevel
	var file, _ = os.OpenFile("./market/logs", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	customLog.SetOutput(io.MultiWriter(os.Stdout, file))
	customLog.SetFormatter(&customLog.JSONFormatter{})
	customLog.SetLevel(level)
	return &Log{
		Service: service,
	}

}

func (logger *Log) Debug(msg string, tags ...string) {
	if customLog.IsLevelEnabled(customLog.DebugLevel) {
		return
	}
	customLog.WithFields(logger.parseFields(tags...)).Debug(msg)
}

func (logger *Log) Info(msg string, tags ...string) {
	if customLog.IsLevelEnabled(customLog.InfoLevel) {
		return
	}
	customLog.WithFields(logger.parseFields(tags...)).Info(msg)
}

func (logger *Log) Error(msg string, err error, tags ...string) {
	if customLog.IsLevelEnabled(customLog.ErrorLevel) {
		return
	}
	msg = fmt.Sprintf("%s - ERROR - %v", msg, err)
	customLog.WithFields(logger.parseFields(tags...)).Error(msg)
}

func (logger *Log) parseFields(tags ...string) customLog.Fields {
	result := make(customLog.Fields, len(tags))
	for _, tag := range tags {
		els := strings.Split(tag, ":")
		result[strings.TrimSpace(els[0])] = strings.TrimSpace(els[1])
	}
	result["service"] = logger.Service
	return result
}
