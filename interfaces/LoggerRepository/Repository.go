package LoggerRepository

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"time"
)

var logLevelTitle = map[int]string{
	LoggerDomain.LogLevelFatal: "[  FATAL  ]",
	LoggerDomain.LogLevelError: "[  ERROR  ]",
	LoggerDomain.LogLevelWarn:  "[ WARNING ]",
	LoggerDomain.LogLevelInfo:  "[  INFO   ]",
	LoggerDomain.LogLevelDebug: "[  DEBUG  ]",
}

type Storage interface {
	Send(message string) (err error)
}

type Repository struct {
	name       string
	storage    Storage
	logLevel   int
	prefix     string
	timeFormat string
}

func New(storage Storage) (r *Repository) {
	r = new(Repository)
	r.storage = storage
	r.SetLogLevelNone()
	r.timeFormat = time.Stamp
	return
}

func (r *Repository) SetName(name string) *Repository {
	r.name = name
	return r
}

func (r *Repository) SetPrefix(prefix string) *Repository {
	r.prefix = prefix
	return r
}

func (r *Repository) SetTimeFormat(timeFormat string) *Repository {
	r.timeFormat = timeFormat
	return r
}

func (r *Repository) GetName() (name string) {
	return r.name
}
