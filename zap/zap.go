package zap

import (
	"reflect"

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

func appendFilledFieldsOnly(fields *[]any, key string, value any) {
	if value == nil {
		return
	}

	reflectedValue := reflect.ValueOf(value)
	if reflectedValue.Kind() != reflect.Bool && reflectedValue.IsZero() {
		return
	}

	if (reflectedValue.Kind() == reflect.Slice || reflectedValue.Kind() == reflect.Map) && reflectedValue.Len() == 0 {
		return
	}

	*fields = append(*fields, key, value)
}
