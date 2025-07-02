package zap

import (
	"context"

	"github.com/surfe/logger/v2/key"
	"github.com/surfe/logger/v2/logi"
	"go.uber.org/zap"
	"gopkg.in/DataDog/dd-trace-go.v1/ddtrace/tracer"
)

type Logger struct {
	log *zap.SugaredLogger
}

func (w *Logger) Ctx(ctx context.Context) logi.Logger {
	fields := []any{}

	if ctx != nil {
		// We do not want to add empty key-value pairs
		appendFilledFieldsOnly(&fields, key.Email, ctx.Value(key.CtxEmail))
		appendFilledFieldsOnly(&fields, key.CompanyKey, ctx.Value(key.CtxCompany))
		appendFilledFieldsOnly(&fields, key.CorrelationID, ctx.Value(key.CtxCorrelationID))
		appendFilledFieldsOnly(&fields, key.Tool, ctx.Value(key.CtxTool))
		appendFilledFieldsOnly(&fields, key.ProductFeature, ctx.Value(key.CtxProductFeature))
		appendFilledFieldsOnly(&fields, key.APIVersion, ctx.Value(key.CtxAPIVersion))
		appendFilledFieldsOnly(&fields, key.JobDetails, ctx.Value(key.CtxJobDetails))
		if span, ok := tracer.SpanFromContext(ctx); ok {
			appendFilledFieldsOnly(&fields, key.DataDogSpanID, span.Context().SpanID())
			appendFilledFieldsOnly(&fields, key.DataDogTraceID, span.Context().TraceID())
		}
	}

	return &Logger{
		log: w.log.With(fields...),
	}
}

func (w *Logger) With(keysAndValues ...any) logi.Logger {
	return &Logger{
		log: w.log.With(keysAndValues...),
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
