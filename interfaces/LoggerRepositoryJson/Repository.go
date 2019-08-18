package LoggerRepositoryJson

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"time"
)

var logLevelTitle = map[int]string{
	LoggerDomain.LogLevelFatal: "fatal",
	LoggerDomain.LogLevelError: "error",
	LoggerDomain.LogLevelWarn:  "warn",
	LoggerDomain.LogLevelInfo:  "info",
	LoggerDomain.LogLevelDebug: "debug",
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
	r = new(Repository).
		setStorage(storage).
		SetLogLevelNone().
		SetTimeFormat(time.Stamp)
	return
}

func (r *Repository) SetName(name string) *Repository {
	r.name = name
	return r
}

func (r *Repository) setStorage(storage Storage) *Repository {
	r.storage = storage
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
