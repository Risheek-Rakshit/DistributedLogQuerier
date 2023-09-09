package logger

import (log "github.com/sirupsen/logrus")

type CustomLogger struct{
	cLogger *log.Logger
}

func NewLogger(level string) *CustomLogger{
	
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	if level == "debug" {
		logger.SetLevel(log.DebugLevel)
	} else if level == "error" {
		logger.SetLevel(log.ErrorLevel)
	} else if level == "warn" {
		logger.SetLevel(log.WarnLevel)
	} else {
		logger.SetLevel(log.InfoLevel)
	}

	return &CustomLogger{logger}
}

func (l *CustomLogger) Info(text string){
	l.cLogger.Info(text)
}

func (l *CustomLogger) Warn(text string){
	l.cLogger.Warn(text)
}

func (l *CustomLogger) Debug(text string){
	l.cLogger.Debug(text)
}

func (l *CustomLogger) Error(text string, err error){
	l.cLogger.Error(text, err)
}