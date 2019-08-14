package LoggerInteractor

import (
	"fmt"
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
	"github.com/pkg/errors"
	"time"
)

func (i *Interactor) Debugf(format string, a ...interface{}) {
	i.Debug(i.formatMessage(format, a...))
}

func (i *Interactor) Infof(format string, a ...interface{}) {
	i.Info(i.formatMessage(format, a...))
}

func (i *Interactor) Warnf(format string, a ...interface{}) {
	i.Warn(i.formatMessage(format, a...))
}

func (i *Interactor) Errorf(format string, a ...interface{}) {
	i.Error(i.formatMessage(format, a...))
}

func (i *Interactor) Fatalf(format string, a ...interface{}) {
	i.Fatal(i.formatMessage(format, a...))
}

func (i *Interactor) Debug(message string) {
	i.log(domain.LogLevelDebug, message)
}

func (i *Interactor) Info(message string) {
	i.log(domain.LogLevelInfo, message)
}

func (i *Interactor) Warn(message string) {
	i.log(domain.LogLevelWarn, message)
}

func (i *Interactor) Error(message string) {
	i.log(domain.LogLevelError, message)
}

func (i *Interactor) Fatal(message string) {
	i.log(domain.LogLevelFatal, message)
}

func (i *Interactor) formatMessage(format string, a ...interface{}) (message string) {
	return fmt.Sprintf(format, a...)
}

func (i *Interactor) logf(level int, format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	i.log(level, message)
}

func (i *Interactor) log(level int, message string) {
	if !i.isMessageLevelMatchInteractorLevel(level) {
		return
	}

	timeNow := i.getTimeNow()
	i.logToRepositories(message, level, timeNow)
	return
}

func (i *Interactor) isMessageLevelMatchInteractorLevel(level int) (isMatch bool) {
	if i.logLevel&level == level {
		isMatch = true
	}
	return
}

func (i *Interactor) getTimeNow() time.Time {
	return time.Now()
}

func (i *Interactor) logToRepositories(message string, level int, time time.Time) {
	for _, repository := range i.repositories {
		if err := repository.Log(level, time, message); err != nil {
			err = errors.Errorf("%s %s", repository.GetName(), err.Error())
			fmt.Println(err)
		}
	}
}
