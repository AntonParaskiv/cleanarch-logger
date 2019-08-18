package LoggerInteractorStd

import (
	"log"
	"os"
)

type Interactor struct {
	logger *log.Logger
}

func New() (i *Interactor) {
	i = new(Interactor)
	i.SetLogger(nil)
	return
}

func (i *Interactor) SetLogger(logger *log.Logger) *Interactor {
	if logger == nil {
		logger = i.makeStdLogger()
	}
	i.setLogger(logger)
	return i
}

func (i *Interactor) makeStdLogger() (logger *log.Logger) {
	logger = log.New(os.Stdout, "", log.LstdFlags)
	return
}

func (i *Interactor) setLogger(logger *log.Logger) {
	i.logger = logger
}
