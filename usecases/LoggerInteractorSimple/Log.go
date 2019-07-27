package LoggerInteractorSimple

import (
	"fmt"
	"log"
)

func (i *LoggerInteractor) log(message string) {
	log.Println(message)
	return
}

func (i *LoggerInteractor) logf(format string, a ...interface{}) {
	message := fmt.Sprintf(format, a...)
	i.log(message)
}

func (i *LoggerInteractor) Debug(message string) {
	i.log(message)
}

func (i *LoggerInteractor) Debugf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *LoggerInteractor) Info(message string) {
	i.log(message)
}

func (i *LoggerInteractor) Infof(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *LoggerInteractor) Warn(message string) {
	i.log(message)
}

func (i *LoggerInteractor) Warnf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *LoggerInteractor) Error(message string) {
	i.log(message)
}

func (i *LoggerInteractor) Errorf(format string, a ...interface{}) {
	i.logf(format, a...)
}

func (i *LoggerInteractor) Fatal(message string) {
	i.log(message)
}

func (i *LoggerInteractor) Fatalf(format string, a ...interface{}) {
	i.logf(format, a...)
}
