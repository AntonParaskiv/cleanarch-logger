package logger

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var logLevelTitle = map[int]string{
	domain.LogLevelFatal: "[  FATAL  ]",
	domain.LogLevelError: "[  ERROR  ]",
	domain.LogLevelWarn:  "[ WARNING ]",
	domain.LogLevelInfo:  "[  INFO   ]",
	domain.LogLevelDebug: "[  DEBUG  ]",
}

type LoggerStorage interface {
	Send(message string) (err error)
}

type LogRepository struct {
	name       string
	storage    LoggerStorage
	logLevel   int
	prefix     string
	timeFormat string
}

func NewLogRepository(name string, storage LoggerStorage, logLevel int, prefix string, timeFormat string) (r *LogRepository) {
	r = new(LogRepository)
	r.name = name
	r.storage = storage
	r.logLevel = logLevel
	r.prefix = prefix
	r.timeFormat = timeFormat
	return
}

func (r *LogRepository) Log(level int, time time.Time, message string) (err error) {
	if r.logLevel&level == level {
		message = strings.Trim(message, " \t\r\n") + "\n"

		if r.prefix != "" {
			message = r.prefix + " " + message
		}

		if logLevelTitle[level] != "" {
			message = logLevelTitle[level] + " " + message
		}

		if r.timeFormat != "" {
			message = time.Format(r.timeFormat) + " " + message
		}

		if err = r.storage.Send(message); err != nil {
			err = errors.New("log storage send failed:" + err.Error())
			return
		}
	}
	return
}

func (r *LogRepository) GetName() (name string) {
	return r.name
}
