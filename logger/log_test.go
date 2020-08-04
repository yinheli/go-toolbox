package logger

import (
	"testing"

	"github.com/yinheli/go-toolbox/logger/log"
	"go.uber.org/zap"
)

func TestFormat(t *testing.T) {
	cfg := &log.Config{
		Level:  "debug",
		Format: "json",
		Caller: true,
	}
	log.ConfigLogger(cfg)
	log.Debug("test", zap.Any("test", "test"))

	cfg.Format = "console"
	log.ConfigLogger(cfg)
	log.Debug("test", zap.Any("test", "test"))
}
