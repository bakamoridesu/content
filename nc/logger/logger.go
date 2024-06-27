package logger

import "fmt"

type ILogger interface {
	Log(args ...any)
	Logf(format string, args ...any)
}

type Logger struct {
}

func (l *Logger) Log(args ...any) {
	fmt.Println(args...)
}

func (l *Logger) Logf(format string, args ...any) {
	fmt.Printf(format, args...)
}
