package main

import "github.com/sirupsen/logrus"

var logger *logrus.Logger

func InitLogrus() {
	logger = logrus.New()
	logger.SetLevel(logrus.DebugLevel)
	logger.SetFormatter(&logrus.TextFormatter{
		TimestampFormat:  "2006-01-02 15:04:05",
		DisableTimestamp: false,
		DisableColors:    true,
	})
}
