package main

import (
	"github.com/sirupsen/logrus"
	"time"
)

var logger *logrus.Logger

func initLogger() {
	logger = logrus.StandardLogger()

	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.TextFormatter{
		ForceColors:     false,
		FullTimestamp:   true,
		TimestampFormat: time.RFC3339Nano,
		DisableSorting:  true,
	})
}
