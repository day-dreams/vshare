package utils

import (
	"os"

	"github.com/sirupsen/logrus"
)

var (
	logger = logrus.New()
)

func Logger() *logrus.Logger {
	return logger
}

func InitLogger() {
	logger.Out = os.Stdout
	logger.SetReportCaller(true)
	logger.SetLevel(logrus.DebugLevel)
}
