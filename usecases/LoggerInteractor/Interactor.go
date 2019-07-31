package LoggerInteractor

import (
	"time"
)

type Repository interface {
	Log(level int, time time.Time, message string) (err error)
	GetName() (name string)
}

type Interactor struct {
	repositories map[string]Repository
	logLevel     int
}

func New() (i *Interactor) {
	i = new(Interactor)
	i.repositories = make(map[string]Repository)
	i.SetLogLevelNone()
	return
}

func (i *Interactor) Fork() (forkedInteractor *Interactor) {
	forkedInteractor = New()
	forkedInteractor.repositories = i.repositories
	forkedInteractor.logLevel = i.logLevel
	return
}
