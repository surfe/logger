package zap

import (
	"context"

	"github.com/surfe/logger/key"
	"github.com/surfe/logger/logi"
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func (w *Logger) With(ctx context.Context, keysAndValues ...any) logi.Logger {
	fields := []any{}
	addNotEmpty := func(key string, value any) {
		val, isOk := value.(string)
		if isOk && key != "" && val != "" {
			fields = append(fields, key, value)
		}
	}

	if ctx != nil {
		// We do not want to add empty key-value pairs
		addNotEmpty(key.Email, ctx.Value(key.CtxEmail))
		addNotEmpty(key.CompanyKey, ctx.Value(key.CtxCompany))
		addNotEmpty(key.CorrelationID, ctx.Value(key.CtxCorrelationID))
		addNotEmpty(key.Tool, ctx.Value(key.CtxTool))
		addNotEmpty(key.ProductFeature, ctx.Value(key.CtxProductFeature))
	}

	fields = append(fields, keysAndValues...)

	return &Logger{
		log: w.log.With(fields...),
	}
}

func (w *Logger) Err(err error) logi.Logger {
	return &Logger{
		log: w.log.With(zap.Error(err)),
	}
}

func (w *Logger) Errorf(template string, args ...any) {
	w.log.Errorf(template, args...)
}

func (w *Logger) Error(args ...any) {
	w.log.Error(args...)
}

func (w *Logger) Warnf(template string, args ...any) {
	w.log.Warnf(template, args...)
}

func (w *Logger) Warn(args ...any) {
	w.log.Warn(args...)
}

func (w *Logger) Infof(template string, args ...any) {
	w.log.Infof(template, args...)
}

func (w *Logger) Info(args ...any) {
	w.log.Info(args...)
}

func (w *Logger) Debugf(template string, args ...any) {
	w.log.Debugf(template, args...)
}

func (w *Logger) Debug(args ...any) {
	w.log.Debug(args...)
}

func (w *Logger) Fatalf(template string, args ...any) {
	w.log.Fatalf(template, args...)
}

func (w *Logger) Fatal(args ...any) {
	w.log.Fatal(args...)
}

func (w *Logger) Sync() {
	_ = w.log.Sync()
}
