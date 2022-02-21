package zap

import (
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

// Init initiates a new Zap log instance and wraps with Logger
func Init() (*Logger, error) {
	cfg := zap.NewProductionConfig()
	zLogger, err := cfg.Build()
	if err != nil {
		return nil, errors.Wrap(err, "Build Zap Config")
	}

	sugarLog := zLogger.Sugar()

	return &Logger{sugarLog}, nil
}
