package example

import (
	"github.com/theharshon/gocommon/logger"
)

func Logger() {
	logger.Initialize("production")
	logger.Info("This info log")
}
