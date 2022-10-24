package logger

import (
	"github.com/Leadjet/logger/logi"
	"github.com/Leadjet/logger/simple"
)

var logger logi.Logger = &simple.Logger{}

// Log is the getter for global `logger` variable
func Log() logi.Logger {
	return logger
}
