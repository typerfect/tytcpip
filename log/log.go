package log

import "github.com/sirupsen/logrus"

var (
	Log *logrus.Logger
)

func init() {
	Log = logrus.New()
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.JSONFormatter{})
}
