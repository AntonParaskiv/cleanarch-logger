package LoggerInteractorStd

import (
	"fmt"
)

func (i *Interactor) Debug(message string) {
	i.log(message)
}

func (i *Interactor) Debugf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *Interactor) Info(message string) {
	i.log(message)
}

func (i *Interactor) Infof(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *Interactor) Warn(message string) {
	i.log(message)
}

func (i *Interactor) Warnf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *Interactor) Error(message string) {
	i.log(message)
}

func (i *Interactor) Errorf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *Interactor) Fatal(message string) {
	i.log(message)
}

func (i *Interactor) Fatalf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *Interactor) logf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	i.log(message)
}

func (i *Interactor) log(message string) {
	i.logger.Println(message)
	return
}
