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
	i.ClearRepositories()
	i.SetLogLevelNone()
	return
}

func (i *Interactor) Fork() (forkedInteractor *Interactor) {
	forkedInteractor = New()
	forkedInteractor.setRepositories(i.getRepositories())
	forkedInteractor.setLogLevel(i.getLogLevel())
	return
}
