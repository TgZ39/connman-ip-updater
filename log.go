package main

import (
	"log"
	"os"
)

const logFile = "updater.log"

// const maxSize = 1024

func getLogFile() (*os.File, error) {
	//err := shortenFile(logFile, maxSize)
	//if err != nil {
	//	log.Println(err)
	//}

	f, err := os.OpenFile(logFile, os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		log.Println("error opening log file: ", err)
		return nil, err
	}

	return f, nil
}

//func shortenFile(file string, maxSize int64) error {
//	f, err := os.OpenFile(file, os.O_CREATE|os.O_RDWR, 0644)
//	if err != nil {
//		log.Println("error opening log file: ", err)
//		return err
//	}
//
//	stats, err := f.Stat()
//	if err != nil {
//		return err
//	}
//
//	startPos := stats.Size() - maxSize
//	if startPos < 1 {
//		return nil
//	}
//
//	_, err = f.Seek(startPos, io.SeekStart)
//	if err != nil {
//		return err
//	}
//
//	remBytes := make([]byte, maxSize)
//	_, err = io.ReadFull(f, remBytes)
//	if err != nil {
//		return err
//	}
//	f.Close()
//
//	err = os.WriteFile(f.Name(), remBytes, 0644)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
