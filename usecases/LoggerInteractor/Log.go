package LoggerInteractor

import (
	"fmt"
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"github.com/pkg/errors"
	"time"
)

func (i *Interactor) log(level int, message string) {
	if i.logLevel&level == level {
		timeNow := time.Now()
		for _, repo := range i.repositories {
			if err := repo.Log(level, timeNow, message); err != nil {
				err = errors.Errorf("%s %s", repo.GetName(), err.Error())
				fmt.Println(err)
			}
		}
	}
	return
}

func (i *Interactor) logf(level int, format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	i.log(level, message)
}

func (i *Interactor) Debug(message string) {
	i.log(domain.LogLevelDebug, message)
}

func (i *Interactor) Debugf(format string, a ...interface{}) {
	i.logf(domain.LogLevelDebug, format, a...)
}

func (i *Interactor) Info(message string) {
	i.log(domain.LogLevelInfo, message)
}

func (i *Interactor) Infof(format string, a ...interface{}) {
	i.logf(domain.LogLevelInfo, format, a...)
}

func (i *Interactor) Warn(message string) {
	i.log(domain.LogLevelWarn, message)
}

func (i *Interactor) Warnf(format string, a ...interface{}) {
	i.logf(domain.LogLevelWarn, format, a...)
}

func (i *Interactor) Error(message string) {
	i.log(domain.LogLevelError, message)
}

func (i *Interactor) Errorf(format string, a ...interface{}) {
	i.logf(domain.LogLevelError, format, a...)
}

func (i *Interactor) Fatal(message string) {
	i.log(domain.LogLevelFatal, message)
}

func (i *Interactor) Fatalf(format string, a ...interface{}) {
	i.logf(domain.LogLevelFatal, format, a...)
}
