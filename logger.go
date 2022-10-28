package logger

import (
	"github.com/Leadjet/logger/logi"
	"github.com/Leadjet/logger/simple"
	"github.com/Leadjet/logger/zap"
)

var logger logi.Logger = &simple.Logger{}
var _ logi.Logger = &zap.Logger{}

// Log is the getter for global `logger` variable
func Log() logi.Logger {
	return logger
}
