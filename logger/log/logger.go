package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Named(s string) *zap.Logger {
	return Logger.Named(s)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return Logger.WithOptions(opts...)
}

func With(fields ...zap.Field) *zap.Logger {
	return Logger.WithOptions(zap.AddCallerSkip(-1)).With(fields...)
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return Logger.Check(lvl, msg)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	Logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	Logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}

func Sync() {
	Logger.Sync()
}
