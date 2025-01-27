package zap_logger

import (
	"log"
	"sync"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	once   sync.Once
	Logger *zap.Logger
	err    error
)

func InitLogger() *zap.Logger {
	once.Do(func() {
		encoder_config := zapcore.EncoderConfig{
			MessageKey:       "message",
			LevelKey:         "level",
			TimeKey:          "time",
			NameKey:          "logger",
			CallerKey:        "caller",
			StacktraceKey:    "stacktrace",
			LineEnding:       zapcore.DefaultLineEnding,
			EncodeLevel:      zapcore.CapitalColorLevelEncoder,
			EncodeTime:       zapcore.TimeEncoderOfLayout("15:04:05"),
			EncodeCaller:     zapcore.ShortCallerEncoder,
			ConsoleSeparator: "\n",
		}

		config := zap.Config{
			Encoding:         "console",
			Level:            zap.NewAtomicLevelAt(zap.InfoLevel),
			OutputPaths:      []string{"stdout"},
			ErrorOutputPaths: []string{"stderr"},
			InitialFields:    zap.NewDevelopmentConfig().InitialFields,
			EncoderConfig:    encoder_config,
		}

		Logger, err = config.Build()
		if err != nil {
			log.Fatal("Error creating logger =>", err)
		}
	})

	return Logger
}
