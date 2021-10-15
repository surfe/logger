package zap

import (
	logger "github.com/Leadjet/Logger"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// RegisterLog initiates a new Zap logger instance and sets logger.Log
func RegisterLog() error {
	if logger.Log != nil {
		return nil
	}

	zapLogger, err := initLog()
	if err != nil {
		return errors.Wrap(err, "Init Log")
	}
	defer zapLogger.Sync()
	sugarLog := zapLogger.Sugar()
	logger.SetLogger(&loggerWrapper{sugarLog})
	return nil
}

func initLog() (zap.Logger, error) {
	cfg := zap.NewProductionConfig()
	zLogger, err := cfg.Build()
	if err != nil {
		return *zLogger, errors.Wrap(err, "Build Log Config")
	}

	return *zLogger, nil
}
