package logger

import (
	"context"

	"github.com/surfe/logger/logi"
	"github.com/surfe/logger/simple"
	"github.com/surfe/logger/zap"
)

var logger logi.Logger = &simple.Logger{}
var _ logi.Logger = &zap.Logger{}

// Log is the getter for global `logger` variable
func Log(ctx context.Context) logi.Logger {
	return logger.Ctx(ctx)
}
