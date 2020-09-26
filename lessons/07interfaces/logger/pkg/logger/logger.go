package logger

import "io"

type Logger struct {
	level int
	wr    io.Writer
}

func NewLogger(level int, wr io.Writer) *Logger {
	return &Logger{
		level: level,
		wr:    wr,
	}
}

func (l *Logger) Write(msg string) error {
	_, err := l.wr.Write([]byte(msg))

	return err
}
