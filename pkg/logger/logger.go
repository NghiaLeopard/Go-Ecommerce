package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

type LoggerZap struct {
	*zap.Logger
}

func NewLog() *LoggerZap {
	LogLevel := "debug"

	var level zapcore.Level

	// debug -> info -> warn -> error -> fatal -> panic
	switch LogLevel {
	case "debug":
		level = zapcore.DebugLevel
	case "info":
		level = zapcore.InfoLevel
	case "warn":
		level = zapcore.WarnLevel
	case "error":
		level = zapcore.ErrorLevel
	default:
		level = zapcore.InfoLevel
	}

	// custom zac
	encoder := getEncoderLog()

	// init and write logs to file
	hook := lumberjack.Logger{
		Filename:   "./storages/logs/dev.001.log",
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     28,   //days
		Compress:   true, // disabled by default
	}

	sync := zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))

	// creates a Core that writes logs to a WriteSyncer.
	core := zapcore.NewCore(encoder, sync, level)

	// new zap
	return &LoggerZap{zap.New(core, zap.AddCaller())}
}

func getEncoderLog() zapcore.Encoder {

	// {"level":"info","ts":1724680001.2874858,"caller":"server/main.go:10","msg":""}
	encoderConfig := zap.NewProductionEncoderConfig()

	// 1724680001.2874858 -> timestamp
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

	// change ts -> time
	encoderConfig.TimeKey = "time"

	// level info -> INFO
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder

	// combine condition zap.New(core, zap.AddCaller()) -> show caller
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	return zapcore.NewJSONEncoder(encoderConfig)
}
