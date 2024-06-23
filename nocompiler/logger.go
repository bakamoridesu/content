package main

import "fmt"

type Logger struct {
	disabled bool
}

func (l *Logger) Disable() {
	l.disabled = true
}

func (l *Logger) Log(args ...interface{}) {
	if l.disabled {
		return
	}
	fmt.Println(args...)
}
