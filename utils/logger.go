package utils

import (
    "io"
    "log"
    "os"
)



func InitLogger() {
    logFile, err := os.OpenFile("application.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
    if err != nil {
        log.Fatalln("Failed to open log file:", err)
    }


    InfoLogger  := log.New(io.MultiWriter(os.Stdout, logFile),
        "INFO: ",
        log.Ldate|log.Ltime|log.Lshortfile)

    

    ErrorLogger := log.New(io.MultiWriter(os.Stderr, logFile),
        "ERROR: ",
        log.Ldate|log.Ltime|log.Lshortfile)
}


