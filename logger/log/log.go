package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log"
	"os"
	"os/signal"
	"path/filepath"
	"strings"
	"syscall"
	"time"
)

var (
	currentCfg  *Config
	writer      io.Writer
	logger      *zap.Logger
	Sugar       *zap.SugaredLogger
	atomicLevel zap.AtomicLevel
)

type Config struct {
	Level   string `yaml:"level" toml:"level"`
	File    string `yaml:"file" toml:"file"`
	Format  string `yaml:"format" toml:"format"`
	Caller  bool   `yaml:"caller" toml:"caller"`
	MaxSize int    `yaml:"max-size" toml:"max-size"`
	MaxDays int    `yaml:"max-days" toml:"max-days"`
	Rotate  bool   `yaml:"rotate" toml:"rotate"`
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
		log.Fatal(err)
	}

	atomicLevel = zap.NewAtomicLevelAt(level)

	writeSynced := zapcore.NewMultiWriteSyncer(ws...)
	writer = writeSynced

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
		writeSynced,
		atomicLevel,
	)

	options := make([]zap.Option, 0, 3)
	options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	if cfg.Caller && level.Enabled(zapcore.DebugLevel) {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	lg := zap.New(core, options...)

	return lg
}

func scheduleRotate(log *lumberjack.Logger) {
	// signal
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGHUP, syscall.SIGTERM)
		for {
			<-ch
			_ = logger.Sync()
		}
	}()

	// time
	for {
		n := time.Now().Add(time.Hour * 24)
		next := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
		d := time.Until(next)
		time.Sleep(d)
		_ = logger.Sync()
		_ = log.Rotate()
	}
}
