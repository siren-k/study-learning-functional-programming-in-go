package decorator

import (
	"io"
	"log"
	"os"
)

var (
	Debug       *log.Logger
	Info        *log.Logger
	Error       *log.Logger
	InfoHandler io.Writer
)

func InitLog(traceFileName string, debugHanlder io.Writer, infoHandler io.Writer, errorHandler io.Writer) {
	if len(traceFileName) > 0 {
		_ = os.Remove(traceFileName)
		file, err := os.OpenFile(traceFileName, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
		if err != nil {
			log.Fatalf("Failed to create log file: %s", traceFileName)
		}
		debugHanlder = io.MultiWriter(file, debugHanlder)
		infoHandler = io.MultiWriter(file, infoHandler)
		errorHandler = io.MultiWriter(file, errorHandler)
	}

	InfoHandler = infoHandler

	Debug = log.New(debugHanlder, "DEBUG : ", log.Ldate|log.Ltime|log.Lshortfile)
	Info = log.New(debugHanlder, "INFO : ", log.Ltime)
	Error = log.New(debugHanlder, "DEBUG : ", log.Ldate|log.Ltime|log.Lshortfile)
}
