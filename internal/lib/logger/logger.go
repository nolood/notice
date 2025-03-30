package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var (
	envDev  = "dev"
	envProd = "prod"
)

func New(env string, logPath string) *zap.Logger {
	var log *zap.Logger
	var core zapcore.Core

	consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
	consoleWriter := zapcore.Lock(os.Stdout)
	consoleCore := zapcore.NewCore(consoleEncoder, consoleWriter, zap.DebugLevel)

	fileEncoder := zapcore.NewJSONEncoder(zap.NewProductionEncoderConfig())
	file, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err)
	}
	fileWriter := zapcore.AddSync(file)
	fileCore := zapcore.NewCore(fileEncoder, fileWriter, zap.InfoLevel)

	switch env {
	case envDev:
		core = zapcore.NewTee(consoleCore)
	case envProd:
		core = zapcore.NewTee(consoleCore, fileCore)
	default:
		core = zapcore.NewTee(consoleCore, fileCore)
	}

	log = zap.New(core)

	return log
}
