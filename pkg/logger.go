package pkg

import (
	"io"
	"log"
	"os"
)

type CustomLogger struct {
	InfoLogger    *log.Logger
	WarningLogger *log.Logger
	ErrorLogger   *log.Logger
	FatalLogger   *log.Logger
	LogLogger     *log.Logger
}

type CustomLoggerI interface {
	Error(v ...any)
	Warning(v ...any)
	Info(v ...any)
	Fatal(v ...any)
	Log(v ...any)
}

func NewLogger() CustomLoggerI {
	file, err := os.OpenFile("scheduler_logger.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal("Cannot init logger")
		os.Exit(1)
	}
	mw := io.MultiWriter(os.Stdout, file)
	return &CustomLogger{
		InfoLogger:    log.New(mw, "[INFO]\t\t", log.Ldate|log.Ltime),
		WarningLogger: log.New(mw, "[WARNING]\t", log.Ldate|log.Ltime),
		ErrorLogger:   log.New(mw, "[ERROR]\t\t", log.Ldate|log.Ltime),
		FatalLogger:   log.New(mw, "[FATAL]\t\t", log.Ldate|log.Ltime),
		LogLogger:     log.New(mw, "[LOG]\t\t", log.Ldate|log.Ltime),
	}
}

func (c *CustomLogger) Log(v ...any) {
	c.LogLogger.Println(v...)
}

func (c *CustomLogger) Error(v ...any) {
	c.ErrorLogger.Println(v...)
}

func (c *CustomLogger) Warning(v ...any) {
	c.WarningLogger.Println(v...)
}

func (c *CustomLogger) Info(v ...any) {
	c.InfoLogger.Println(v...)
}

// Logs a fatal error and exit the program with code 1
func (c *CustomLogger) Fatal(v ...any) {
	c.FatalLogger.Println(v...)
	os.Exit(1)
}
