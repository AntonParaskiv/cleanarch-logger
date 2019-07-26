package LoggerInteractor

import (
	"time"
)

type LoggerRepository interface {
	Log(level int, time time.Time, message string) (err error)
	GetName() (name string)
}

type LoggerInteractor struct {
	repositories map[string]LoggerRepository
	logLevel     int
}

func New() (i *LoggerInteractor) {
	i = new(LoggerInteractor)
	i.repositories = make(map[string]LoggerRepository)
	i.SetLogLevelNone()
	return
}

func (i *LoggerInteractor) Fork() (forkedInteractor *LoggerInteractor) {
	forkedInteractor = New()
	forkedInteractor.repositories = i.repositories
	forkedInteractor.logLevel = i.logLevel
	return
}
