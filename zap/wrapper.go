package zap

import (
	"context"

	"github.com/Leadjet/logger/key"
	"github.com/Leadjet/logger/logi"
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func (w *Logger) WithContext(ctx context.Context) logi.Logger {
	email, _ := ctx.Value(key.CtxEmail).(string)
	company, _ := ctx.Value(key.CtxCompany).(string)
	correlationID, _ := ctx.Value(key.CtxCorrelationID).(string)

	if email == "" {
		return w
	}

	return &Logger{
		log: w.log.With(
			key.Email, email,
			key.CompanyKey, company,
			key.CorrelationID, correlationID,
		),
	}
}

func (w *Logger) Errorf(template string, err error, args ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err)).Errorf(template, args...)

		return
	}

	w.log.Errorf(template, args...)
}

func (w *Logger) Errorw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err)).Errorw(msg, keysAndValues...)

		return
	}

	w.log.Errorw(msg, keysAndValues...)
}

func (w *Logger) Error(err error, args ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err)).Errorw("Error", args...)

		return
	}

	w.log.Errorw("Error", args...)
}

func (w *Logger) Warnf(template string, err error, args ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err)).Warnf(template, args...)

		return
	}

	w.log.Warnf(template, args...)
}

func (w *Logger) Warnw(msg string, err error, keysAndValues ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err)).Warnw(msg, keysAndValues...)

		return
	}

	w.log.Warnw(msg, keysAndValues...)
}

func (w *Logger) Warn(err error, args ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err)).Warnw("Error", args...)

		return
	}

	w.log.Warnw("Error", args...)
}

func (w *Logger) Infof(template string, args ...interface{}) {
	w.log.Infof(template, args...)
}

func (w *Logger) Infow(msg string, keysAndValues ...interface{}) {
	w.log.Infow(msg, keysAndValues...)
}

func (w *Logger) Info(args ...interface{}) {
	w.log.Info(args...)
}

func (w *Logger) Debugf(template string, args ...interface{}) {
	w.log.Debugf(template, args...)
}

func (w *Logger) Debugw(msg string, keysAndValues ...interface{}) {
	w.log.Debugw(msg, keysAndValues...)
}

func (w *Logger) Debug(args ...interface{}) {
	w.log.Debug(args...)
}

func (w *Logger) Sync() {
	_ = w.log.Sync()
}
