package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"io"
)

func Named(s string) *zap.Logger {
	return logger.Named(s)
}

func WithOptions(opts ...zap.Option) *zap.Logger {
	return logger.WithOptions(opts...)
}

func With(fields ...zap.Field) *zap.Logger {
	return logger.WithOptions(zap.AddCallerSkip(-1)).With(fields...)
}

func Check(lvl zapcore.Level, msg string) *zapcore.CheckedEntry {
	return logger.Check(lvl, msg)
}

func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

func DPanic(msg string, fields ...zap.Field) {
	logger.DPanic(msg, fields...)
}

func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}

func SetLevel(level zapcore.Level) {
	atomicLevel.SetLevel(level)
}

func Logger() *zap.Logger {
	return logger
}

func Writer() io.Writer {
	return writer
}

func Sync() {
	_ = logger.Sync()
}
