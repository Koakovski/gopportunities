package config

import (
	"io"
	"log"
	"os"
)

type Logger struct {
	debug   *log.Logger
	info    *log.Logger
	warning *log.Logger
	err     *log.Logger
	writer  io.Writer
}

func NewLogger(p string) *Logger {
	writer := io.Writer(os.Stdout)

	logger := log.New(writer, p, log.Ldate|log.Ltime)

	return &Logger{
		debug:   log.New(writer, "[DEBUG]: ", logger.Flags()),
		info:    log.New(writer, "[INFO]: ", logger.Flags()),
		warning: log.New(writer, "[WARNING]: ", logger.Flags()),
		err:     log.New(writer, "[ERROR]: ", logger.Flags()),
		writer:  writer,
	}
}

// Non-Formated logs
func (l *Logger) Debug(v ...any) {
	l.debug.Println(v...)
}

func (l *Logger) Info(v ...any) {
	l.info.Println(v...)
}

func (l *Logger) Warning(v ...any) {
	l.warning.Println(v...)
}

func (l *Logger) Err(v ...any) {
	l.err.Println(v...)
}

// Formated logs
func (l *Logger) Debugf(text string, v ...any) {
	l.debug.Printf(text, v...)
}

func (l *Logger) Infof(text string, v ...any) {
	l.info.Printf(text, v...)
}

func (l *Logger) WarningF(text string, v ...any) {
	l.warning.Printf(text, v...)
}

func (l *Logger) Errf(text string, v ...any) {
	l.err.Printf(text, v...)
}
