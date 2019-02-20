package client

import (
	"log"
	"os"
)

// A LogI manages logging
type LogI interface {
	// Err writes the output for a error event
	Err(v ...interface{})

	// Debug writes the output for a debug event
	Debug(v ...interface{})

	// DebugF writes the output for a debug event. Arguments are handled in the manner of fmt.Printf.
	DebugF(format string, v ...interface{})
}

// A Log represents an active logging object.
type Log struct {
	errorLogger *log.Logger
	debugLogger *log.Logger
	isDebug     bool
}

// Err writes the output for a error event
func (l *Log) Err(v ...interface{}) {
	if l.isDebug {
		l.errorLogger.Println(v...)
	}
}

// Debug writes the output for a debug event
func (l *Log) Debug(v ...interface{}) {
	if l.isDebug {
		l.debugLogger.Println(v...)
	}
}

// DebugF writes the output for a debug event. Arguments are handled in the manner of fmt.Printf.
func (l *Log) DebugF(format string, v ...interface{}) {
	if l.isDebug {
		l.debugLogger.Printf(format, v...)
	}
}

// L is the default implementation of LogI.
var L LogI

// PrepareLogger prepares L object.
func PrepareLogger(debug bool) {
	L = &Log{
		errorLogger: log.New(os.Stderr, "[ERROR] ", log.LstdFlags),
		debugLogger: log.New(os.Stdout, "[DEBUG] ", log.LstdFlags),
		isDebug:     debug,
	}
}
