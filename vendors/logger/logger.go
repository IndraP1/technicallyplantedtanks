package logger

import (
    "log"
    "fmt"
    "io"
    "os"
)

// TODO create handler for log page

func InitServerLogger() *os.File {
    //TODO change to /var/log
	logFile, err := os.OpenFile("./gobbledygook.log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		log.Panicln(err)
	}

	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	return logFile
}

func Error(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	log.Printf("[ERROR] %s\n", msg)
}

func Fatal(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	log.Fatalf("[ERROR] %s\n", msg)
}

func Print(s string, i ...interface{}) {
	msg := fmt.Sprintf(s, i...)
	log.Printf("[STATUS] %s\n", msg)
}
