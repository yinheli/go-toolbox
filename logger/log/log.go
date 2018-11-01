package log

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Logger *zap.Logger
	Sugar  *zap.SugaredLogger
)

type Log struct {
	Level   string `yaml:"level" toml:"level"`
	File    string `yaml:"file" toml:"file"`
	Format  string `yaml:"format" toml:"format"`
	MaxSize int    `yaml:"max-size" toml:"max-size"`
	MaxDays int    `yaml:"max-days" toml:"max-days"`
	Rotate  bool   `yaml:"rotate" toml:"rotate"`
}

func init() {
	cfg := Log{
		Level: "debug",
	}
	ConfigLogger(&cfg)
}

func ConfigLogger(cfg *Log) {
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

	writeSyncer := zapcore.NewMultiWriteSyncer(ws...)

	encodingCfg := zap.NewProductionEncoderConfig()
	encodingCfg.EncodeTime = zapcore.ISO8601TimeEncoder

	var encoder zapcore.Encoder
	if cfg.Format == "json" {
		encoder = zapcore.NewJSONEncoder(encodingCfg)
	} else {
		encoder = zapcore.NewConsoleEncoder(encodingCfg)
	}
	core := zapcore.NewCore(
		encoder,
		writeSyncer,
		level,
	)

	options := make([]zap.Option, 0, 3)
	options = append(options, zap.AddStacktrace(zapcore.ErrorLevel))
	if level.Enabled(zapcore.DebugLevel) {
		options = append(options, zap.AddCaller(), zap.AddCallerSkip(1))
	}
	Logger = zap.New(core, options...)
	if err != nil {
		log.Fatal(err)
	}
	Sugar = Logger.Sugar()
}

func scheduleRotate(log *lumberjack.Logger) {
	// signal
	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, syscall.SIGHUP)
		for {
			<-ch
			Logger.Sync()
			log.Rotate()
		}
	}()

	// time
	for {
		n := time.Now().Add(time.Hour * 24)
		next := time.Date(n.Year(), n.Month(), n.Day(), 0, 0, 0, 0, time.Local)
		d := next.Sub(time.Now())
		time.Sleep(d)
		Logger.Sync()
		log.Rotate()
	}
}
