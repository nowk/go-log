package log

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

type Logger interface {
	Output(calldepth int, s string) error
}

type Level int

const (
	DEBUG Level = iota
	INFO
	NOTICE
	WARN
	ERROR
	FATAL
)

var levels = map[Level]string{
	DEBUG:  "DEBUG",
	INFO:   "INFO",
	NOTICE: "NOTICE",
	WARN:   "WARN",
	ERROR:  "ERROR",
	FATAL:  "FATAL",
}

var (
	LogLevel Level = INFO
)

func SetLogLevel(l Level) {
	LogLevel = l

	log.Printf("log: level: %s\n", levels[LogLevel])
}

func SetLogLevelStr(s string) {
	s = strings.ToUpper(s)
	for k, v := range levels {
		if s == v {
			SetLogLevel(k)
			return
		}
	}

	log.Printf("log: error: unknown level %s", s)
}

// logger wraps go built-in log
type logger struct {
	*log.Logger
}

// New returns a logger for one or more Writers
func New(prefix string, flags int, w ...io.Writer) *logger {
	var wr io.Writer
	if len(w) == 1 {
		wr = w[0]
	} else if len(w) > 1 {
		wr = io.MultiWriter(w...)
	} else {
		wr = os.Stderr
	}

	return &logger{
		Logger: log.New(wr, prefix, flags),
	}
}

func (l logger) Output(calldepth int, s string) error {
	return l.Logger.Output(calldepth, s)
}

// Log provides log level checks before writing. As well as level string prefix
// before the message
func Log(l Logger, lvl Level, p string, v ...interface{}) error {
	s, ok := levels[lvl]
	if !ok || lvl < LogLevel {
		return nil
	}

	return l.Output(2, f(f("[%s] %s", s, p), v...))
}

// f is a short cut to fmt.Sprintf
func f(p string, v ...interface{}) string {
	return fmt.Sprintf(p, v...)
}

func (l logger) Log(lvl Level, p string, v ...interface{}) error {
	return Log(l, lvl, p, v...)
}

func (l logger) Debug(p string, v ...interface{}) error {
	return l.Log(DEBUG, p, v...)
}

func (l logger) Info(p string, v ...interface{}) error {
	return l.Log(INFO, p, v...)
}

func (l logger) Notice(p string, v ...interface{}) error {
	return l.Log(NOTICE, p, v...)
}

func (l logger) Warn(p string, v ...interface{}) error {
	return l.Log(WARN, p, v...)
}

func (l logger) Error(p string, v ...interface{}) error {
	return l.Log(ERROR, p, v...)
}

// Fatal exits process
func (l logger) Fatal(p string, v ...interface{}) error {
	l.Log(FATAL, p, v...)
	os.Exit(1)
	return nil
}
