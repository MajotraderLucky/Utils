package logger

import (
	"log"
	"os"
)

func createLogsDirectory() error {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return err
	}
	return nil
}

func openLogFile() (*os.File, error) {
	logFile, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func setLogger(logFile *os.File) {
	log.SetOutput(logFile)
}

func logLine() {
	log.Println("---------------------------------------------------")
}
