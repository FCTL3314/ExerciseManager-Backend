package bootstrap

import (
	"log"
	"log/slog"
	"os"
	"runtime/debug"
)

const (
	ControllersLoggerFilePath = "logs/controller.json"
)

type Logger interface {
	Debug(msg string, args ...any)
	Info(msg string, args ...any)
	Warn(msg string, args ...any)
	Error(msg string, args ...any)
}

type SlogLogger struct {
	logger *slog.Logger
}

func (l *SlogLogger) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

func (l *SlogLogger) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

func (l *SlogLogger) Warn(msg string, args ...any) {
	trace := debug.Stack()
	l.logger.Warn(msg, append(args, "traceback", string(trace))...)
}

func (l *SlogLogger) Error(msg string, args ...any) {
	trace := debug.Stack()
	l.logger.Error(msg, append(args, "traceback", string(trace))...)
}

type LoggerGroup struct {
	User *Logger
}

func NewLoggerGroup(userLogger *Logger) *LoggerGroup {
	return &LoggerGroup{
		User: userLogger,
	}
}

func initLogger(logFilePath string) Logger {
	file, err := os.OpenFile(logFilePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	fileHandler := slog.NewJSONHandler(file, &slog.HandlerOptions{
		Level: slog.LevelDebug,
	})

	l := slog.New(fileHandler)

	return &SlogLogger{
		logger: l,
	}
}

func InitUserLogger() Logger { return initLogger(ControllersLoggerFilePath) }
