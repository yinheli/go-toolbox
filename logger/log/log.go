package log

import (
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	currentCfg  *Config
	writer      io.Writer
	logger      *zap.Logger
	Sugar       *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

type Config struct {
	Level   string
	File    string
	Format  string
	Caller  bool
	MaxSize int
	MaxDays int
	Rotate  bool
}

func init() {
	cfg := Config{
		Level: "debug",
	}
	ConfigLogger(&cfg)
}

func ConfigLogger(cfg *Config) {
	currentCfg = cfg
	logger = NewLoggerWithConfig(cfg)
	logger = logger.WithOptions(zap.AddCallerSkip(1))
	Sugar = logger.Sugar()
}

func NewLogger(file string) *zap.Logger {
	cfg := *currentCfg
	if cfg.File != "" {
		cfg.File = filepath.Join(filepath.Dir(cfg.File), file)
	}
	lg := NewLoggerWithConfig(&cfg)
	if cfg.File == "" && file != "" {
		lg = lg.Named(file)
	}
	return lg
}

func NewLoggerWithConfig(cfg *Config) *zap.Logger {
	var err error

	ws := make([]zapcore.WriteSyncer, 0, 2)
	ws = append(ws, zapcore.AddSync(os.Stdout))
	if cfg.File != "" {
		rotateLogger := &lumberjack.Logger{
			Filename:  cfg.File,
			MaxSize:   cfg.MaxSize,
			MaxAge:    cfg.MaxDays,
			LocalTime: true,
			Compress:  true,
		}
		ws = append(ws, zapcore.AddSync(rotateLogger))

		if cfg.Rotate {
			go scheduleRotate(rotateLogger)
		}
	}

	var level zapcore.Level
	err = level.UnmarshalText([]byte(cfg.Level))
	if err != nil {
		level = zap.InfoLevel
	}

	writer := zapcore.NewMultiWriteSyncer(ws...)

	encodingCfg := zap.NewProductionEncoderConfig()
	encodingCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if strings.ToLower(cfg.Format) == "json" {
		encoder = zapcore.NewJSONEncoder(encodingCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encodingCfg)
	}
	core := zapcore.NewCore(
		encoder,
		writer,
		zap.NewAtomicLevelAt(level),
	)

	options := make([]zap.Option, 0)
	options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	if cfg.Caller && level.Enabled(zapcore.DebugLevel) {
		options = append(options, zap.AddCaller())
	}
	lg := zap.New(core, options...)

	return lg
}

func scheduleRotate(log *lumberjack.Logger) {
	for {
		n := time.Now().Add(time.Hour * 24)
		next := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
		d := time.Until(next)
		time.Sleep(d)
		_ = log.Rotate()
	}
}
