package logging

import (
	"fmt"
	"io"
	"os"
	"runtime"
	"strings"
	"sync"
	"time"
)

// logging types
const (
	Error = iota
	Warning
	Info
	Debug
)

const timeFormat = "2006-01-02 15:04:05"

type logger struct {
	sync.Mutex
	level    uint8
	isLogger bool
	out      io.Writer
}

type Logger interface {
	Error(msg interface{})
	Warn(msg interface{})
	Info(msg interface{})
	Debug(msg interface{})
	Errorf(error string, msg ...interface{})
	Warnf(error string, msg ...interface{})
	Infof(error string, msg ...interface{})
	Debugf(error string, msg ...interface{})
	Custom(msg string)
	SetLevel(level uint8)
	SetOutput(out io.Writer)
	IsLogger(isOk bool)
}

func NewLogger(args ...int) Logger {
	l := new(logger)
	l.level = 1
	if len(args) > 0 {
		l.level = uint8(args[0])
	}
	l.out = os.Stdout
	return l
}

var DefaultLogger = NewLogger()

func (l *logger) SetOutput(out io.Writer) {
	l.Lock()
	defer l.Unlock()
	l.out = out
}

func (l *logger) SetLevel(level uint8) {
	l.Lock()
	defer l.Unlock()
	l.level = level
}

func (l *logger) IsLogger(p bool) {
	l.Lock()
	defer l.Unlock()
	l.isLogger = p
}

var mappingLevel = map[uint8]string{
	Error:   "Error",
	Warning: "Warn",
	Info:    "Info",
	Debug:   "Debug",
}

func (l *logger) logWrite(msg interface{}, level uint8) (string, bool) {
	var msgText string
	switch v := msg.(type) {
	case error:
		msgText = v.Error()
	case string:
		msgText = v
	}

	if level > l.level && !l.isLogger {
		return "", false
	}

	if !l.isLogger {
		_, file, lineno, ok := runtime.Caller(2)

		src := ""

		if ok {
			src = strings.Replace(
				fmt.Sprintf("%s:%d", file, lineno), "%2e", ".", -1)
		}
		msgText = fmt.Sprintf("[%s] %s | (%s) %s", mappingLevel[level], time.Now().Format(timeFormat), src, msgText)
	} else {
		msgText = fmt.Sprintf("[%s] %s | %s", mappingLevel[level], time.Now().Format(timeFormat), msgText)

	}

	return msgText, true
}

func (l *logger) print(msg string) {
	l.Lock()
	defer l.Unlock()
	_, err := l.out.Write(append([]byte(msg), '\n'))
	if err != nil {
		_, _ = os.Stdout.Write(append([]byte(msg), '\n'))
	}
}

func (l *logger) Custom(msg string) {
	l.print(msg)
}

func (l *logger) Error(msg interface{}) {
	if msg, ok := l.logWrite(msg, Error); ok {
		l.print(msg)
	}
}

func (l *logger) Errorf(error string, msg ...interface{}) {
	if msg, ok := l.logWrite(fmt.Sprintf(error, msg...), Error); ok {
		l.print(msg)
	}
}

func (l *logger) Warn(msg interface{}) {
	if msg, ok := l.logWrite(msg, Warning); ok {
		l.print(msg)
	}
}

func (l *logger) Warnf(error string, msg ...interface{}) {
	if msg, ok := l.logWrite(fmt.Sprintf(error, msg...), Warning); ok {
		l.print(msg)
	}
}

func (l *logger) Info(msg interface{}) {
	if msg, ok := l.logWrite(msg, Info); ok {
		l.print(msg)
	}
}

func (l *logger) Infof(error string, msg ...interface{}) {
	if msg, ok := l.logWrite(fmt.Sprintf(error, msg...), Info); ok {
		l.print(msg)
	}
}

func (l *logger) Debugf(error string, msg ...interface{}) {
	if msg, ok := l.logWrite(fmt.Sprintf(error, msg...), Debug); ok {
		l.print(msg)
	}
}

func (l *logger) Debug(msg interface{}) {
	if msg, ok := l.logWrite(msg, Debug); ok {
		l.print(msg)
	}
}
