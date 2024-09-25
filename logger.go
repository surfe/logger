package logger

import (
	"context"

	"github.com/surfe/logger/v2/logi"
	"github.com/surfe/logger/v2/simple"
	"github.com/surfe/logger/v2/zap"
)

var logger logi.Logger = &simple.Logger{}
var _ logi.Logger = &zap.Logger{}

// Log is the getter for global `logger` variable
func Log(ctx context.Context) logi.Logger {
	return logger.Ctx(ctx)
}
