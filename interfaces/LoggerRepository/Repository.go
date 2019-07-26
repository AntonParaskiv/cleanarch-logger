package LoggerRepository

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
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

type LoggerRepository struct {
	name       string
	storage    LoggerStorage
	logLevel   int
	prefix     string
	timeFormat string
}

func New(storage LoggerStorage) (r *LoggerRepository) {
	r = new(LoggerRepository)
	r.storage = storage
	r.SetLogLevelNone()
	r.timeFormat = time.Stamp
	return
}

func (r *LoggerRepository) SetName(name string) *LoggerRepository {
	r.name = name
	return r
}

func (r *LoggerRepository) SetPrefix(prefix string) *LoggerRepository {
	r.prefix = prefix
	return r
}

func (r *LoggerRepository) SetTimeFormat(timeFormat string) *LoggerRepository {
	r.timeFormat = timeFormat
	return r
}

func (r *LoggerRepository) GetName() (name string) {
	return r.name
}
