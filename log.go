package main

import (
	"io"
	"log"
	"os"
)

func initLogger() {

	logFileName := os.Getenv("LOG_FILE_NAME")

	if logFileName == "" {
		log.Fatalln("missing enviroment variable LOG_FILE_NAME")
		return
	}

	logFile, error := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	errorWriter := io.MultiWriter(os.Stdout, logFile)

	if error != nil {
		log.Fatalln("Unable to create logger file: ", error.Error())
		return
	}

	log.SetOutput(errorWriter)
}
