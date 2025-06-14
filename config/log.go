package config

import (
	"log"
	"os"
	"strings"
	"time"

	"github.com/rs/zerolog"
)

var (
	logger      = zerolog.Logger{}
	logFile     *os.File
	logFileBase = "logs/"
)

func InitLog() {
	zerolog.DurationFieldInteger = true
	zerolog.DurationFieldUnit = time.Millisecond
	zerolog.DurationFieldInteger = true
	zerolog.DurationFieldUnit = time.Millisecond

	dailyLogFile := logFileBase + time.Now().Format("2006-01-02") + ".log"
	var err error
	logFile, err = os.OpenFile(dailyLogFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal("Failed to open log file")
		return
	}
	if strings.ToLower(GetAppConfig().ENVIRONMENT) == "production" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		logger = zerolog.New(logFile).With().Timestamp().Caller().Logger()
	} else {
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
		logger = zerolog.New(os.Stdout).With().Timestamp().Caller().Logger()
	}
}

// GetLogger returns the logger instance
func GetLogger() zerolog.Logger {
	return logger
}
