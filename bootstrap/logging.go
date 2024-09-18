package bootstrap

import (
	"log"
	"log/slog"
	"os"
	"path/filepath"
	"runtime/debug"
)

const (
	ControllerLogsFileName = "controller.json"

	BaseLogsDir    = "logs/"
	BaseUserDir    = BaseLogsDir + "users/"
	BaseWorkoutDir = BaseLogsDir + "users/"

	UserControllersLoggingPath     = BaseUserDir + ControllerLogsFileName
	WorkoutControllersLoggingPath  = BaseWorkoutDir + ControllerLogsFileName
	ExerciseControllersLoggingPath = BaseWorkoutDir + ControllerLogsFileName
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

func (sl *SlogLogger) Debug(msg string, args ...any) {
	sl.logger.Debug(msg, args...)
}

func (sl *SlogLogger) Info(msg string, args ...any) {
	sl.logger.Info(msg, args...)
}

func (sl *SlogLogger) Warn(msg string, args ...any) {
	trace := debug.Stack()
	sl.logger.Warn(msg, append(args, "traceback", string(trace))...)
}

func (sl *SlogLogger) Error(msg string, args ...any) {
	trace := debug.Stack()
	sl.logger.Error(msg, append(args, "traceback", string(trace))...)
}

type LoggerGroup struct {
	User     *Logger
	Workout  *Logger
	Exercise *Logger
}

func NewLoggerGroup(
	userLogger *Logger,
	workoutLogger *Logger,
	exerciseLogger *Logger,
) *LoggerGroup {
	return &LoggerGroup{
		User:     userLogger,
		Workout:  workoutLogger,
		Exercise: exerciseLogger,
	}
}

func initLogger(logFilePath string) Logger {
	logDir := filepath.Dir(logFilePath)

	err := os.MkdirAll(logDir, os.ModePerm)
	if err != nil {
		log.Fatalf("Failed to create log directory: %v", err)
	}

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

func InitUserLogger() Logger     { return initLogger(UserControllersLoggingPath) }
func InitWorkoutLogger() Logger  { return initLogger(WorkoutControllersLoggingPath) }
func InitExerciseLogger() Logger { return initLogger(ExerciseControllersLoggingPath) }
