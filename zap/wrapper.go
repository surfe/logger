package zap

import "go.uber.org/zap"

type wrapper struct {
	logger *zap.SugaredLogger
}

func (w *wrapper) Errorf(template string, err interface{}, args ...interface{}) {
	if err != nil {
		w.logger.With(zap.Error(err.(error))).Errorf(template, args...)

		return
	}

	w.logger.Errorf(template, args...)
}

func (w *wrapper) Errorw(msg string, err interface{}, keysAndValues ...interface{}) {
	if err != nil {
		w.logger.With(zap.Error(err.(error))).Errorw(msg, keysAndValues...)

		return
	}

	w.logger.Errorw(msg, keysAndValues...)
}

func (w *wrapper) Error(err interface{}, args ...interface{}) {
	if err != nil {
		w.logger.With(zap.Error(err.(error))).Errorw("Error", args...)

		return
	}

	w.logger.Errorw("Error", args...)
}

func (w *wrapper) Infof(template string, args ...interface{}) {
	w.logger.Infof(template, args...)
}

func (w *wrapper) Infow(msg string, keysAndValues ...interface{}) {
	w.logger.Infow(msg, keysAndValues...)
}

func (w *wrapper) Info(args ...interface{}) {
	w.logger.Info(args...)
}

func (w *wrapper) Debugf(template string, args ...interface{}) {
	w.logger.Debugf(template, args...)
}

func (w *wrapper) Debugw(msg string, keysAndValues ...interface{}) {
	w.logger.Debugw(msg, keysAndValues...)
}

func (w *wrapper) Debug(args ...interface{}) {
	w.logger.Debug(args...)
}

func (w *wrapper) Sync() {
	_ = w.logger.Sync()
}
