package logger

import (log "github.com/sirupsen/logrus")

type CustomLogger struct{
	cLogger *log.Logger
}

func NewLogger() *CustomLogger{
	logger := log.New()
	logger.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
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