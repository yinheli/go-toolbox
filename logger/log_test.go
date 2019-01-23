package logger

import (
	"github.com/yinheli/go-toolbox/logger/log"
	"go.uber.org/zap"
	"testing"
)

func TestFormat(t *testing.T) {
	cfg := &log.Config{
		Level:  "debug",
		Format: "json",
	}
	log.ConfigLogger(cfg)
	log.Debug("test", zap.Any("test", "test"))

	cfg.Format = "console"
	log.ConfigLogger(cfg)
	log.Debug("test", zap.Any("test", "test"))
}
