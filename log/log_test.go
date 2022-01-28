package logger

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"testing"
)

func TestLogger_Error(t *testing.T) {
	NewDefaultLogger("SVC", true, zapcore.DebugLevel)
	ZapLogger.Error("test", zap.String("info", "test"))
}

func TestLogger_Info(t *testing.T) {
	NewDefaultLogger("SVC", true, zapcore.DebugLevel)
	ZapLogger.Info("test", zap.String("info", "test"))
}
