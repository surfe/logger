package zap

import (
	"go.uber.org/zap"
)

type Logger struct {
	log *zap.SugaredLogger
}

func (w *Logger) Errorf(template string, err interface{}, args ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err.(error))).Errorf(template, args...)

		return
	}

	w.log.Errorf(template, args...)
}

func (w *Logger) Errorw(msg string, err interface{}, keysAndValues ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err.(error))).Errorw(msg, keysAndValues...)

		return
	}

	w.log.Errorw(msg, keysAndValues...)
}

func (w *Logger) Error(err interface{}, args ...interface{}) {
	if err != nil {
		w.log.With(zap.Error(err.(error))).Errorw("Error", args...)

		return
	}

	w.log.Errorw("Error", args...)
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
