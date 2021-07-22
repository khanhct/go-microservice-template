package logging

import (
	"fmt"
	"os"

	"github.com/sirupsen/logrus"

	"github.com/spf13/viper"
)

var _logger *logrus.Logger

func Initialize() {
	_logger = logrus.New()
	if viper.GetString("environment.mode") == "production" {
		Formatter := new(logrus.JSONFormatter)
		Formatter.TimestampFormat = "02-01-2006 15:04:05"
		Formatter.FieldMap = logrus.FieldMap{
			logrus.FieldKeyTime:  "@timestamp",
			logrus.FieldKeyLevel: "@level",
			logrus.FieldKeyMsg:   "@message",
			logrus.FieldKeyFunc:  "@caller",
		}
		logrus.SetFormatter(Formatter)
	} else {
		Formatter := new(logrus.TextFormatter)
		Formatter.TimestampFormat = "02-01-2006 15:04:05"
		Formatter.FullTimestamp = true
		Formatter.DisableColors = false
		logrus.SetFormatter(Formatter)
	}
	switch viper.GetString("logging.mode") {
	case "panic":
		_logger.Level = logrus.PanicLevel
	case "fatal":
		_logger.Level = logrus.FatalLevel
	case "error":
		_logger.Level = logrus.ErrorLevel
	case "warn":
		_logger.Level = logrus.WarnLevel
	case "info":
		_logger.Level = logrus.InfoLevel
	case "debug":
		_logger.Level = logrus.DebugLevel
	case "trace":
		_logger.Level = logrus.TraceLevel
	default:
		_logger.Level = logrus.InfoLevel
	}

	logPath := fmt.Sprintf("%v%v", viper.GetString("logging.logFolder"), viper.GetString("logging.logFile"))
	file, err := os.OpenFile(logPath, os.O_CREATE|os.O_WRONLY, 0666)
	if err == nil {
		_logger.Out = file
	} else {
		_logger.Warn(err.Error())
	}
}

func GetLogger() *logrus.Logger {
	if _logger == nil {
		Initialize()
	}
	return _logger
}
