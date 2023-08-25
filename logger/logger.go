package logger

import (
	"log"
	"os"
)

type Logger struct {
	logFile *os.File
}

func (l *Logger) CreateLogsDir() error {
	err := os.MkdirAll("logs", 0755)
	if err != nil {
		return err
	}
	return nil
}

func (l *Logger) OpenLogFile() error {
	logFile, err := os.OpenFile("logs/log.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	l.logFile = logFile
	return nil
}

func (l *Logger) SetLogger() {
	log.SetOutput(l.logFile)
}

func (l *Logger) LogLine() {
	log.Println("-------------------------------------------------------")
}
