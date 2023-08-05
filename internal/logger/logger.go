package logger

import (
	"log"
)

type Logger struct {
	debug bool
}

func New(debug bool) *Logger {
	return &Logger{debug: debug}
}

func (l *Logger) Debugf(format string, v ...any) {
	if l.debug {
		log.Printf(format, v...)
	}
}

func FatalErr(err error) {
	log.Fatalf("Houston, we have a problem: %v", err)
}
