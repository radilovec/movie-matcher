package logger

import (
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func SyncLogger() {
	if Logger != nil {
		_ = Logger.Sync()
	}
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func InitFileLogger() error {
	logDir := "../logs"
	logFile := "app.log"

	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return err
	}

	logFilePath := filepath.Join(logDir, logFile)
	file, err := os.OpenFile(logFilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}

	writeSyncer := zapcore.AddSync(file)

	core := zapcore.NewCore(
		zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig()),
		writeSyncer,
		zap.InfoLevel,
	)

	Logger = zap.New(core)
	return nil
}
