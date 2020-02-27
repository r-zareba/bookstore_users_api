package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	logger *zap.Logger
)

func init() {
	logConfig := zap.Config{
		Level:    zap.NewAtomicLevelAt(zap.InfoLevel),
		Encoding: "json",
		EncoderConfig: zapcore.EncoderConfig{
			LevelKey:   "level",
			TimeKey:    "time",
			MessageKey: "Msg",
			EncodeTime: zapcore.ISO8601TimeEncoder,
			EncodeLevel: zapcore.LowercaseLevelEncoder,
			EncodeCaller: zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: nil,
		InitialFields:    nil,
	}

	var err error
	logger, err = logConfig.Build()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return logger
}

func Info(message string, tags ...zap.Field) {
	logger.Info(message, tags...)
	logger.Sync()
}

func Error(message string, err error, tags ...zap.Field) {
	tags = append(tags, zap.NamedError("error", err))

	logger.Error(message, tags...)
	logger.Sync()
}
