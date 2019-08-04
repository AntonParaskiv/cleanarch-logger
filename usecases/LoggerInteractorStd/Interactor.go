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
	i.logger = log.New(os.Stdout, "", log.LstdFlags)
	return
}

func (i *Interactor) SetLogger(logger *log.Logger) *Interactor {
	if logger == nil {
		logger = log.New(os.Stdout, "", log.LstdFlags)
	}
	i.logger = logger
	return i
}
