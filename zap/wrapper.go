package zap

import "go.uber.org/zap"

type loggerWrapper struct {
	lw *zap.SugaredLogger
}

func (logger *loggerWrapper) Errorf(template string, err interface{}, args ...interface{}) {
	logger.lw.With(zap.Error(err.(error))).Errorf(template, args...)
}

func (logger *loggerWrapper) Errorw(msg string, err interface{}, keysAndValues ...interface{}) {
	logger.lw.With(zap.Error(err.(error))).Errorw(msg, keysAndValues...)
}

func (logger *loggerWrapper) Error(err interface{}, args ...interface{}) {
	logger.lw.With(zap.Error(err.(error))).Errorw("Error", args...)
}

func (logger *loggerWrapper) Infof(template string, args ...interface{}) {
	logger.lw.Infof(template, args...)
}

func (logger *loggerWrapper) Infow(msg string, keysAndValues ...interface{}) {
	logger.lw.Infow(msg, keysAndValues...)
}

func (logger *loggerWrapper) Info(args ...interface{}) {
	logger.lw.Info(args...)
}

func (logger *loggerWrapper) Debugf(template string, args ...interface{}) {
	logger.lw.Debugf(template, args...)
}

func (logger *loggerWrapper) Debugw(msg string, keysAndValues ...interface{}) {
	logger.lw.Debugw(msg, keysAndValues...)
}

func (logger *loggerWrapper) Debug(args ...interface{}) {
	logger.lw.Debug(args...)
}
