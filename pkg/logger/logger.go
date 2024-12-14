package logger

import (
	"github.com/sirupsen/logrus"
	"os"
)

type Logger struct {
	internalLogger *logrus.Logger
}

func NewLogger() *Logger {
	logger := logrus.New()

	logger.SetFormatter(&logrus.JSONFormatter{})
	logger.SetOutput(os.Stdout)

	logger.SetLevel(logrus.DebugLevel)

	return &Logger{
		internalLogger: logger,
	}
}

func (l *Logger) Log(message string) {
	l.internalLogger.Info(message)
}

func (l *Logger) Warn(message string) {
	l.internalLogger.Warn(message)
}

func (l *Logger) Error(message string) {
	l.internalLogger.Error(message)
}

func (l *Logger) Trace(message string) {
	l.internalLogger.Trace(message)
}

func (l *Logger) Debug(message string) {
	l.internalLogger.Debug(message)
}

func (l *Logger) Fatal(message string) {
	l.internalLogger.Fatal(message)
}

func (l *Logger) Panic(message string) {
	l.internalLogger.Panic(message)
}
