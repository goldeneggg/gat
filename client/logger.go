package client

import (
	"log"
	"os"
)

type LogI interface {
	Err(v ...interface{})
	Debug(v ...interface{})
	DebugF(format string, v ...interface{})
}

type Log struct {
	errorLogger *log.Logger
	debugLogger *log.Logger
	isDebug     bool
}

func (l *Log) Err(v ...interface{}) {
	if l.isDebug {
		l.errorLogger.Println(v)
	}
}

func (l *Log) Debug(v ...interface{}) {
	if l.isDebug {
		l.debugLogger.Println(v)
	}
}

func (l *Log) DebugF(format string, v ...interface{}) {
	if l.isDebug {
		l.debugLogger.Printf(format, v)
	}
}

var L LogI

func PrepareLogger(debug bool) {
	L = &Log{
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
		isDebug:     debug,
	}
}
