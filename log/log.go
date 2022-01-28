package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var ZapLogger *Logger

type Logger struct {
	logger *zap.Logger
}

type Field = zap.Field

func NewDefaultLogger(serviceName string, developmentMod bool, debugLevel zapcore.Level) (err error) {
	logger, err := zap.Config{
		Level:       zap.NewAtomicLevelAt(debugLevel),
		Development: developmentMod,
		Encoding:    "json",
		InitialFields:    map[string]interface{}{"serviceName": serviceName},
		EncoderConfig: zapcore.EncoderConfig{
			TimeKey:        "time",
			LevelKey:       "level",
			NameKey:        "logger",
			CallerKey:      "caller",
			FunctionKey:    zapcore.OmitKey,
			MessageKey:     "msg",
			StacktraceKey:  "stacktrace",
			LineEnding:     zapcore.DefaultLineEnding,
			EncodeLevel:    zapcore.CapitalLevelEncoder,
			EncodeTime:     zapcore.EpochNanosTimeEncoder,
			EncodeDuration: zapcore.SecondsDurationEncoder,
			EncodeCaller:   zapcore.ShortCallerEncoder,
		},
		OutputPaths:      []string{"stdout"},
		ErrorOutputPaths: []string{"stderr"},
	}.Build(zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel))
	if err != nil {
		return err
	}
	ZapLogger = &Logger{
		logger: logger,
	}

	return nil
}

func (l *Logger) Info(msg string, fields ...Field) {
	l.logger.Info(msg, fields...)
}

func (l *Logger) Error(msg string, fields ...Field) {
	l.logger.Error(msg, fields...)
}

func (l *Logger) Warn(msg string, fields ...Field) {
	l.logger.Warn(msg, fields...)
}

func (l *Logger) Panic(msg string, fields ...Field) {
	l.logger.Panic(msg, fields...)
}
