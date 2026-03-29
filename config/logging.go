package config

import (
	"github.com/kataras/golog"
)

var logger *golog.Logger

func GetLogger() *golog.Logger {
	if logger == nil {
		logger = golog.New()
	}
	return logger
}
