package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var zapLog *zap.Logger

func init() {
	// var err error
	// config := zap.NewProductionConfig()
	// enccoderConfig := zap.NewProductionEncoderConfig()
	// zapcore.TimeEncoderOfLayout("Jan _2 15:04:05.000000000")
	// enccoderConfig.StacktraceKey = "" // to hide stacktrace info
	// config.EncoderConfig = enccoderConfig

	// zapLog, err = config.Build(zap.AddCallerSkip(1))
	// if err != nil {
	// 	panic(err)
	// }

	const LOG_FILE = "wordlebot.log"
	config := zap.NewProductionEncoderConfig()
	config.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(config)
	consoleEncoder := zapcore.NewConsoleEncoder(config)
	logFile, _ := os.OpenFile(LOG_FILE, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	writer := zapcore.AddSync(logFile)
	defaultLogLevel := zapcore.DebugLevel
	core := zapcore.NewTee(
		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
	)

	zl := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))
	zapLog = zl
	// 	return logger
}

// func fileLogger(filename string) *zap.Logger {
// 	config := zap.NewProductionEncoderConfig()
// 	config.EncodeTime = zapcore.ISO8601TimeEncoder
// 	fileEncoder := zapcore.NewJSONEncoder(config)
// 	consoleEncoder := zapcore.NewConsoleEncoder(config)
// 	logFile, _ := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	writer := zapcore.AddSync(logFile)
// 	defaultLogLevel := zapcore.DebugLevel
// 	core := zapcore.NewTee(
// 		zapcore.NewCore(fileEncoder, writer, defaultLogLevel),
// 		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), defaultLogLevel),
// 	)

// 	logger := zap.New(core, zap.AddCaller(), zap.AddStacktrace(zapcore.ErrorLevel))

// 	return logger
// }

func Info(message string, fields ...zap.Field) {
	zapLog.Info(message, fields...)
}

func Debug(message string, fields ...zap.Field) {
	zapLog.Debug(message, fields...)
}

func Error(message string, fields ...zap.Field) {
	zapLog.Error(message, fields...)
}

func Fatal(message string, fields ...zap.Field) {
	zapLog.Fatal(message, fields...)
}

func Fatalf(message string, errs ...error) {
	zapLog.Sugar().Fatalf(message, errs)
}
