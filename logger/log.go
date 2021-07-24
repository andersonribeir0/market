package logger

import (
	"github.com/gemnasium/logrus-graylog-hook/v3"
	logger "github.com/sirupsen/logrus"
)


func ConfigLogger(level logger.Level) {
	hook := graylog.NewAsyncGraylogHook("logstash:12201", map[string]interface{}{"app": "market"})
	defer hook.Flush()
	logger.AddHook(hook)
	logger.SetLevel(level)
	logger.Info("graylog activated.")
}