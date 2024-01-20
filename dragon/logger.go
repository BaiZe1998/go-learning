package main

import (
	"log"
	"os"
)

var logger = initLogger()

func initLogger() *log.Logger {
	if !DEBUG {
		return nil
	}

	f, err := os.Create("log.txt")
	if err != nil {
		panic(err)
	}
	logger := log.New(f, "dragon: ", log.LstdFlags)
	return logger
}

func loggerInfo(format string, v ...interface{}) {
	if logger != nil {
		logger.Printf(format, v...)
	}
}
