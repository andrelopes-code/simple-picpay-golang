package cfg

import (
	"fmt"
	"io"
	"log"
	"os"
)

type colorCodes struct {
	reset    string
	red      string
	yellow   string
	blue     string
	green    string
	bgred    string
	bggreen  string
	bgyellow string
	bgblue   string
}

var colors = colorCodes{
	reset:    "\033[m",
	red:      "\033[1;31m",
	yellow:   "\033[1;33m",
	blue:     "\033[1;34m",
	green:    "\033[1;32m",
	bgred:    "\033[1;41m",
	bggreen:  "\033[1;42m",
	bgyellow: "\033[1;43m",
	bgblue:   "\033[1;44m",
}

type Logger struct {
	debug   log.Logger
	info    log.Logger
	warning log.Logger
	err     log.Logger
	fatal   log.Logger
	writer  io.Writer
}

func newLogger(prefix string) *Logger {
	writer := io.Writer(os.Stdout)
	logger := log.New(writer, prefix, log.Ldate|log.Ltime)
	return &Logger{
		debug: *log.New(
			writer,
			fmt.Sprintf("%s %s ~ DEBUG %s %s", colors.bggreen, prefix, colors.reset, colors.green),
			logger.Flags(),
		),
		info: *log.New(
			writer,
			fmt.Sprintf("%s %s ~ INFO %s %s", colors.bgblue, prefix, colors.reset, colors.blue),
			logger.Flags(),
		),
		warning: *log.New(
			writer,
			fmt.Sprintf("%s %s ~ WARNING %s %s", colors.bgyellow, prefix, colors.reset, colors.yellow),
			logger.Flags(),
		),
		err: *log.New(
			writer,
			fmt.Sprintf("%s %s ~ ERROR %s %s", colors.bgred, prefix, colors.reset, colors.red),
			logger.Flags(),
		),
		fatal: *log.New(
			writer,
			fmt.Sprintf("%s %s ~ FATAL  ", colors.bgred, prefix),
			logger.Flags(),
		),
		writer: writer,
	}
}

// Create Non-Formatted logs
func (l *Logger) Debug(v ...interface{}) {
	l.debug.Println(v...)
	print(colors.reset)
}
func (l *Logger) Info(v ...interface{}) {
	l.info.Println(v...)
	print(colors.reset)
}
func (l *Logger) Warning(v ...interface{}) {
	l.warning.Println(v...)
	print(colors.reset)
}
func (l *Logger) Error(v ...interface{}) {
	l.err.Println(v...)
	print(colors.reset)
}
func (l *Logger) Fatal(v ...interface{}) {
	l.fatal.Println(v...)
	print(colors.reset)
}

// Create Formatted logs
func (l *Logger) Debugf(format string, v ...interface{}) {
	l.debug.Printf(format, v...)
	print(colors.reset)
}
func (l *Logger) Infof(format string, v ...interface{}) {
	l.info.Printf(format, v...)
	print(colors.reset)
}
func (l *Logger) Warningf(format string, v ...interface{}) {
	l.warning.Printf(format, v...)
	print(colors.reset)
}
func (l *Logger) Errorf(format string, v ...interface{}) {
	l.err.Printf(format, v...)
	print(colors.reset)
}
func (l *Logger) Fatalf(format string, v ...interface{}) {
	l.fatal.Printf(format, v...)
	print(colors.reset)
}
