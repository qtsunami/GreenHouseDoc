package logging

import (
	"runtime"
	"log"
	"os"
	"fmt"
	"path/filepath"
)

// Level is the logging level
type Level int 

var (
	// F is the file
	F *os.File
	// DefaultPrefix is the default of log prefix
	DefaultPrefix = ""
	// DefaultCallerDepth is the 2
	DefaultCallerDepth = 2
	// logger
	logger *log.Logger
	logPrefix = ""
	levelFlags = []string{"DEBUG", "INFO", "WARNING", "ERROR", "FATAL"}
)

const (
	// DEBUG is primary level of logging
	DEBUG = iota
	// INFO is secondary level of logging
	INFO
	//WARNING is the third level of logging
	WARNING
	// ERROR is the fourth level of logging
	ERROR
	// FATAL is the fifth level of logging
	FATAL
)

func init() {
	filePath := getLogFullPath()
	F = openLogFile(filePath)

	logger = log.New(F, DefaultPrefix, log.LstdFlags)
}

// Debug logging
func Debug(v ...interface{}) {
	setPrefix(DEBUG)
	logger.Println(v)
}

// Info logging
func Info(v ...interface{}) {
	setPrefix(INFO)
	logger.Println(v)
}

// Warn logging
func Warn(v ...interface{}) {
	setPrefix(WARNING)
	logger.Println(v)
}

// Error logging
func Error(v ...interface{}) {
	setPrefix(ERROR)
	logger.Println(v)
}

// Fatal logging
func Fatal(v ...interface{}) {
	setPrefix(FATAL)
	logger.Println(v)
}

func setPrefix(level Level) {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)

	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", levelFlags[level], filepath.Base(file), line)
    } else {
        logPrefix = fmt.Sprintf("[%s]", levelFlags[level])
	}

	logger.SetPrefix(logPrefix)
}

