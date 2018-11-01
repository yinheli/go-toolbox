package sugar

import (
	"github.com/yinheli/go-toolbox/logger/log"
	"go.uber.org/zap"
)

func Named(name string) *zap.SugaredLogger {
	return log.Sugar.Named(name)
}

func With(args ...interface{}) *zap.SugaredLogger {
	return log.Sugar.With(args...)
}

func Debug(args ...interface{}) {
	log.Sugar.Debug(args...)
}

func Info(args ...interface{}) {
	log.Sugar.Info(args...)
}

func Warn(args ...interface{}) {
	log.Sugar.Warn(args...)
}

func Error(args ...interface{}) {
	log.Sugar.Error(args...)
}

func DPanic(args ...interface{}) {
	log.Sugar.DPanic(args...)
}

func Panic(args ...interface{}) {
	log.Sugar.Panic(args...)
}

func Fatal(args ...interface{}) {
	log.Sugar.Fatal(args...)
}

func Debugf(template string, args ...interface{}) {
	log.Sugar.Debugf(template, args...)
}

func Infof(template string, args ...interface{}) {
	log.Sugar.Infof(template, args...)
}

func Warnf(template string, args ...interface{}) {
	log.Sugar.Warnf(template, args...)
}

func Errorf(template string, args ...interface{}) {
	log.Sugar.Errorf(template, args...)
}

func DPanicf(template string, args ...interface{}) {
	log.Sugar.DPanicf(template, args...)
}

func Panicf(template string, args ...interface{}) {
	log.Sugar.Panicf(template, args...)
}

func Fatalf(template string, args ...interface{}) {
	log.Sugar.Fatalf(template, args...)
}

func Debugw(msg string, keysAndValues ...interface{}) {
	log.Sugar.Debugw(msg, keysAndValues...)
}

func Infow(msg string, keysAndValues ...interface{}) {
	log.Sugar.Infow(msg, keysAndValues...)
}

func Warnw(msg string, keysAndValues ...interface{}) {
	log.Sugar.Warnw(msg, keysAndValues...)
}

func Errorw(msg string, keysAndValues ...interface{}) {
	log.Sugar.Errorw(msg, keysAndValues...)
}

func DPanicw(msg string, keysAndValues ...interface{}) {
	log.Sugar.DPanicw(msg, keysAndValues...)
}

func Panicw(msg string, keysAndValues ...interface{}) {
	log.Sugar.Panicw(msg, keysAndValues...)
}

func Fatalw(msg string, keysAndValues ...interface{}) {
	log.Sugar.Fatalw(msg, keysAndValues...)
}
