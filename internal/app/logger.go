package app

import (
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()

const env = "local"

// InitLogger initializes a logger for main.go file.
func InitLogger() *logrus.Logger {
	switch env {
	case "local":
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

// InitLoggerEndpoint initilizes a logger for all endpoints.
// The logger includes http_method, request_uri, and user_agent in a message.
func InitLoggerEndpoint(r *http.Request) *logrus.Entry {
	switch env {
	case "local":
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
	logger := log.WithFields(logrus.Fields{
		"http_method": r.Method,
		"request_uri": fmt.Sprintf("%s", r.RequestURI),
		"user_agent":  r.UserAgent(),
	})
	return logger
}
