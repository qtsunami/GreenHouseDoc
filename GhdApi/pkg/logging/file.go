package logging

import (
	"log"
	"fmt"
	"time"
	"os"
)

var (
	// LogBasePath is the base path
	LogBasePath = "runtime/logs/"
	// LogName is the name of logging
	LogName = "log"
	// TimeFormat is the format of time logs
	TimeFormat = "20060102"
)


func getLogFilePath() string {
	return fmt.Sprintf("%s", LogBasePath)
}

func getLogFullPath() string {
	prefixPath := getLogFilePath()
	suffixPath := fmt.Sprintf("%s-%s.log", LogName, time.Now().Format(TimeFormat))

	return fmt.Sprintf("%s%s", prefixPath, suffixPath)
}

func openLogFile(filepath string) *os.File {
	_, err := os.Stat(filepath)
	switch {
		case os.IsNotExist(err):
			mkDir()
		case os.IsPermission(err):
			log.Fatal("Permission：%v", err)
	}

	handle, err := os.OpenFile(filepath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
        log.Fatalf("Fail to OpenFile :%v", err)
    }

    return handle
}


func mkDir() {
	dir, _ := os.Getwd()
	err := os.MkdirAll(dir + "/" + getLogFilePath(), os.ModePerm)
	if err != nil {
		panic(err)
	}
}