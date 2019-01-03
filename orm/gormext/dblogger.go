package gormext

import (
	"fmt"
	"go.uber.org/zap"
	"regexp"
	"strings"
)

var (
	scopeRE = regexp.MustCompile(`(?m)sql\s+.*?\.go\:\d+`)
)

type dbLogger struct {
	l *zap.Logger
}

func NewDBLogger(logger *zap.Logger) *dbLogger {
	return &dbLogger{
		l: logger,
	}
}

func (lg *dbLogger) Print(v ...interface{}) {
	lg.l.Debug(lg.tidySQLLog(strings.TrimSpace(fmt.Sprintln(v...))))
}

func (lg *dbLogger) tidySQLLog(log string) string {
	return scopeRE.ReplaceAllLiteralString(log, "sql")
}
