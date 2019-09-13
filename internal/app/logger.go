package app

import (
	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const env = "dev"

// InitLogger initializes a logger with timestamp field.
func InitLogger() *logrus.Logger {
	switch env {
	case "dev":
		log.Formatter = &logrus.TextFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
			FullTimestamp:   true,
		}
		log.Level = logrus.InfoLevel

	default:
		log.Formatter = &logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		}
		log.Level = logrus.InfoLevel
	}
	return log
}
