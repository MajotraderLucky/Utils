package logger_test

import (
	"bytes"
	"os"
	"testing"

	"github.com/MajotraderLucky/Utils/logger"
)

func TestCleanLogCountLines_NegativeN(t *testing.T) {
	// Создаем временный файл с некоторым содержимым
	fileName := "log.txt"
	content := "Line 1\nLine 2\nLine 3\nLine 4\nLine 5\n"
	err := os.WriteFile(fileName, []byte(content), 0666)
	if err != nil {
		t.Fatalf("Failed to create file: %s", err)
	}
	defer os.Remove(fileName)

	// Open the file and assign it to Logger
	logFile, err := os.OpenFile(fileName, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil {
		t.Fatalf("Cannot open file: %s", err)
	}
	defer logFile.Close()

	// Создаем новый экземпляр Logger и вызываем функцию CleanLogCountLines с отрицательным значением N
	l := &logger.Logger{logFile: logFile}
	err = l.CleanLogCountLines(-5)
	if err != nil {
		t.Fatalf("Cannot clean log lines: %s", err)
	}

	// Проверяем, что файл не содержит ни одной строки
	fileContent, err := os.ReadFile(fileName)
	if err != nil {
		t.Fatalf("Cannot read file: %s", err)
	}

	// You may want to increase the counter by 1 because the last line contains "\n"
	if len(bytes.Split(fileContent, []byte("\n"))) != 1 {
		t.Errorf("Expected empty file, got content: %s", string(fileContent))
	}
}
