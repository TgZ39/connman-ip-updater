package main

import (
	"log"
	"os"
)

const logFile = "updater.log"

func getLogFile() (*os.File, error) {
	if err := os.Truncate(logFile, 0); err != nil {
		log.Printf("error truncating log file: %v", err)
	}

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		log.Printf("error opening log file: %v", err)
		return nil, err
	}

	return f, nil
}
