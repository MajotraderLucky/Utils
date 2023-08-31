package logger

import (
	"log"
	"os"
	"strings"
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

func (l *Logger) CleanLog() {
	// Read the file contents log.txt
	data, err := os.ReadFile("logs/log.txt")
	if err != nil {
		log.Println(err)
	}
	// Split the content into lines
	lines := strings.Split(string(data), "\n")

	// Check the number of lines
	if len(lines) > 100 {
		// Open a file log.txt in overwrite mode
		logFile, err := os.OpenFile("logs/log.txt", os.O_TRUNC|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			log.Fatal(err)
		}
		// Close the file
		defer logFile.Close()

		// Write the last 100 lines to log.txt
		for _, line := range lines[len(lines)-100:] {
			logFile.WriteString(line + "\n")
		}
	}
}
