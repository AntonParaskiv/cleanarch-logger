package LoggerInteractor

import (
	"fmt"
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
	"github.com/pkg/errors"
	"time"
)

func (i *LoggerInteractor) log(level int, message string) {
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

func (i *LoggerInteractor) logf(level int, format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	i.log(level, message)
}

func (i *LoggerInteractor) Debug(message string) {
	i.log(domain.LogLevelDebug, message)
}

func (i *LoggerInteractor) Debugf(format string, a ...interface{}) {
	i.logf(domain.LogLevelDebug, format, a...)
}

func (i *LoggerInteractor) Info(message string) {
	i.log(domain.LogLevelInfo, message)
}

func (i *LoggerInteractor) Infof(format string, a ...interface{}) {
	i.logf(domain.LogLevelInfo, format, a...)
}

func (i *LoggerInteractor) Warn(message string) {
	i.log(domain.LogLevelWarn, message)
}

func (i *LoggerInteractor) Warnf(format string, a ...interface{}) {
	i.logf(domain.LogLevelWarn, format, a...)
}

func (i *LoggerInteractor) Error(message string) {
	i.log(domain.LogLevelError, message)
}

func (i *LoggerInteractor) Errorf(format string, a ...interface{}) {
	i.logf(domain.LogLevelError, format, a...)
}

func (i *LoggerInteractor) Fatal(message string) {
	i.log(domain.LogLevelFatal, message)
}

func (i *LoggerInteractor) Fatalf(format string, a ...interface{}) {
	i.logf(domain.LogLevelFatal, format, a...)
}
