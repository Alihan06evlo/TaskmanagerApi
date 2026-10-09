package logger

import (
	"os"

	"github.com/sirupsen/logrus"
)

func Init(){
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true, DisableColors: true})
	logrus.SetLevel(logrus.InfoLevel)
}