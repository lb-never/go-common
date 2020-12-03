package cusUtils

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

// Debug function
var baseFileName = "log-"
var logInfo *log.Logger
var logDebug *log.Logger
var logErrInfo *log.Logger
var logDebugFlag = false
var logInfoFlag = false
var logErrorFlag = false

func init() {
	buildLogFile()
	if strings.Contains(GetLogLevel(), "logInfo") {
		logInfoFlag = true
	}
	if strings.Contains(GetLogLevel(), "logDebug") {
		logDebugFlag = true
	}
	if strings.Contains(GetLogLevel(), "logErrInfo") {
		logErrorFlag = true
	}
}

func Debug(v ...interface{}) {
	if logDebugFlag {
		logDebug.Output(3, fmt.Sprintln(v...))
	}
}

func Info(v ...interface{}) {
	if logInfoFlag {
		logInfo.Output(3, fmt.Sprintln(v...))
	}
}

func ErrInfo(v ...interface{}) {
	if logErrorFlag {
		logErrInfo.Output(3, fmt.Sprintln(v...))
	}
}

func buildLogFile() {
	logFile, err := os.OpenFile(buildLogFileName(), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666) //打开文件，若果文件不存在就创建一个同名文件并打开
	if err != nil {
		log.Fatalln("Failed to open error log file:", err)
	}
	logInfo = log.New(logFile, "[logInfo]", log.Lshortfile)
	logDebug = log.New(logFile, "[logDebug]", log.Lshortfile)
	logErrInfo = log.New(logFile, "[logErrInfo]", log.Lshortfile)
	/*logInfo = log.New(io.MultiWriter(logFile, os.Stderr), "[logInfo]", log.Lshortfile|log.Ldate|log.Ltime) //|log.Ldate|log.Ltime
	logErrInfo = log.New(io.MultiWriter(logFile, os.Stderr), "[logErrInfo]", log.Lshortfile|log.Ldate|log.Ltime) //|log.Ldate|log.Ltime
	logDebug = log.New(io.MultiWriter(logFile, os.Stderr), "[logDebug]", log.Lshortfile|log.Ldate|log.Ltime) //|log.Ldate|log.Ltime*/
}

func buildLogFileName() string {
	dateStr := time.Now().Format("2006-01")
	fileName := baseFileName + dateStr + ".log"
	return fileName
}
