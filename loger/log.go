package log

import (
	"fmt"
	"time"
)

var (
	level Level
)

const (
	// DebugLevel defines debug log level.
	DebugLevel Level = iota
	// InfoLevel defines info log level.
	InfoLevel
	// WarnLevel defines warn log level.
	WarnLevel
	// ErrorLevel defines error log level.
	ErrorLevel
	// FatalLevel defines fatal log level.
)

// Level defines log levels.
type Level int8

// SetLevel метод для установки логера
func SetLevel(l Level) {
	level = l
}

func Debug(v ...interface{}) {
	if ok := should(DebugLevel); ok {
		s := fmt.Sprint(v...)
		print(getPrefix("[DEB]"), s)
	}
}

func Debugf(format string, v ...interface{}) {
	if ok := should(DebugLevel); ok {
		s := fmt.Sprintf(format, v...)
		print(getPrefix("[DEB]"), s)
	}
}

func Info(v ...interface{}) {
	if ok := should(InfoLevel); ok {
		s := fmt.Sprint(v...)
		print(getPrefix("[INF]"), s)
	}
}

func Infof(format string, v ...interface{}) {
	if ok := should(InfoLevel); ok {
		s := fmt.Sprintf(format, v...)
		print(getPrefix("[INF]"), s)
	}
}

func Warning(v ...interface{}) {
	if ok := should(WarnLevel); ok {
		s := fmt.Sprint(v...)
		print(getPrefix("[WAR]"), s)
	}
}

func Warningf(format string, v ...interface{}) {
	if ok := should(WarnLevel); ok {
		s := fmt.Sprintf(format, v...)
		print(getPrefix("[WAR]"), s)
	}
}

func Error(v ...interface{}) {
	if ok := should(ErrorLevel); ok {
		s := fmt.Sprint(v...)
		print(getPrefix("[ERR]"), s)
	}
}

func Errorf(format string, v ...interface{}) {
	if ok := should(ErrorLevel); ok {
		s := fmt.Sprintf(format, v...)
		print(getPrefix("[ERR]"), s)
	}
}

func print(pref, format string) {
	fmt.Println(pref, format)
}

func should(lvl Level) bool {
	if lvl > level || lvl == level {
		return true
	}
	return false
}

func getPrefix(l string) string {
	return fmt.Sprint(l, " ", time.Now().Format("2006/01/02 - 15:04:05"), " |")
}
