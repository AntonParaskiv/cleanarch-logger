package LoggerInteractor

import (
	"fmt"
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
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
	i.log(LoggerDomain.LogLevelDebug, message)
}

func (i *Interactor) Info(message string) {
	i.log(LoggerDomain.LogLevelInfo, message)
}

func (i *Interactor) Warn(message string) {
	i.log(LoggerDomain.LogLevelWarn, message)
}

func (i *Interactor) Error(message string) {
	i.log(LoggerDomain.LogLevelError, message)
}

func (i *Interactor) Fatal(message string) {
	i.log(LoggerDomain.LogLevelFatal, message)
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
	i.sendToRepositories(message, level, timeNow)
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
