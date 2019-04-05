package echologger

import (
	"github.com/labstack/gommon/log"
	zlog "github.com/yinheli/go-toolbox/logger/log"
	"github.com/yinheli/go-toolbox/logger/sugar"
	"go.uber.org/zap"
	"io"
)

type Logger struct {
	prefix string
	log    *zap.SugaredLogger
}

func NewLogger(prefix string) *Logger {
	return &Logger{
		prefix: prefix,
		log:    sugar.Named(prefix),
	}
}

func (t *Logger) Output() io.Writer {
	return zlog.Writer()
}

func (t *Logger) SetOutput(w io.Writer) {
	//
}

func (t *Logger) Prefix() string {
	return t.prefix
}

func (t *Logger) SetPrefix(p string) {
	t.prefix = p
	t.log = sugar.Named(p)
}

func (t *Logger) Level() log.Lvl {
	return log.INFO
}

func (t *Logger) SetLevel(v log.Lvl) {
	// not work
}

func (t *Logger) SetHeader(h string) {
	// not head
}

func (t *Logger) Print(i ...interface{}) {
	t.log.Info(i...)
}

func (t *Logger) Printf(format string, args ...interface{}) {
	t.log.Infof(format, args...)
}

func (t *Logger) Printj(j log.JSON) {
	t.log.Info(j)
}

func (t *Logger) Debug(i ...interface{}) {
	t.log.Debug(i...)
}

func (t *Logger) Debugf(format string, args ...interface{}) {
	t.log.Debugf(format, args...)
}

func (t *Logger) Debugj(j log.JSON) {
	t.log.Debug(j)
}

func (t *Logger) Info(i ...interface{}) {
	t.log.Info(i...)
}

func (t *Logger) Infof(format string, args ...interface{}) {
	t.log.Infof(format, args...)
}

func (t *Logger) Infoj(j log.JSON) {
	t.log.Info(j)
}

func (t *Logger) Warn(i ...interface{}) {
	t.log.Warn(i...)
}

func (t *Logger) Warnf(format string, args ...interface{}) {
	t.log.Warnf(format, args...)
}

func (t *Logger) Warnj(j log.JSON) {
	t.log.Warn(j)
}

func (t *Logger) Error(i ...interface{}) {
	t.log.Error(i...)
}

func (t *Logger) Errorf(format string, args ...interface{}) {
	t.log.Errorf(format, args...)
}

func (t *Logger) Errorj(j log.JSON) {
	t.log.Error(j)
}

func (t *Logger) Fatal(i ...interface{}) {
	t.log.Fatal(i...)
}

func (t *Logger) Fatalj(j log.JSON) {
	t.log.Fatal(j)
}

func (t *Logger) Fatalf(format string, args ...interface{}) {
	t.log.Fatalf(format, args...)
}

func (t *Logger) Panic(i ...interface{}) {
	t.log.Panic(i...)
}

func (t *Logger) Panicj(j log.JSON) {
	t.log.Panic(i)
}

func (t *Logger) Panicf(format string, args ...interface{}) {
	t.log.Panicf(format, args...)
}
