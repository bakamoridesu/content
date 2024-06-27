package logger

import "fmt"

type ILogger interface {
	Log(args ...interface{})
	Logf(format string, args ...interface{})
}

type Logger struct {
}

func NewLogger() ILogger {
	return &Logger{}
}

func (l *Logger) Log(args ...interface{}) {
	fmt.Println(args...)
}

func (l *Logger) Logf(format string, args ...interface{}) {
	fmt.Printf(format, args...)
}
