
package logging

import (
    "fmt"
    "log"
    "os"
    "time"
)

var (
    infoLogger  *log.Logger
    errorLogger *log.Logger
)

func InitLogger() {
    logFileName := time.Now().Format("2006-01-02") + "_upos.log"
    file, err := os.OpenFile("logs/"+logFileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        fmt.Println("Failed to open log file:", err)
        return
    }

    infoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
    errorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)

    infoLogger.Println("Logger initialized")
}

func Info(msg string) {
    if infoLogger != nil {
        infoLogger.Println(msg)
    }
}

func Error(msg string) {
    if errorLogger != nil {
        errorLogger.Println(msg)
    }
}
