package LoggerInteractor

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
)

func (i *Interactor) SetLogLevelAll() *Interactor {
	i.SetLogLevelNone()
	i.AddLogLevelDebug()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	i.AddLogLevelError()
	i.AddLogLevelFatal()
	return i
}

func (i *Interactor) SetLogLevelDebugInfoWarn() *Interactor {
	i.SetLogLevelNone()
	i.AddLogLevelDebug()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	return i
}

func (i *Interactor) SetLogLevelInfoWarn() *Interactor {
	i.SetLogLevelNone()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	return i
}

func (i *Interactor) SetLogLevelErrorFatal() *Interactor {
	i.SetLogLevelNone()
	i.AddLogLevelError()
	i.AddLogLevelFatal()
	return i
}

func (i *Interactor) SetLogLevelNone() *Interactor {
	i.setLogLevel(LoggerDomain.LogLevelNone)
	return i
}

func (i *Interactor) AddLogLevelDebug() *Interactor {
	i.addLogLevel(LoggerDomain.LogLevelDebug)
	return i
}

func (i *Interactor) AddLogLevelInfo() *Interactor {
	i.addLogLevel(LoggerDomain.LogLevelInfo)
	return i
}

func (i *Interactor) AddLogLevelWarn() *Interactor {
	i.addLogLevel(LoggerDomain.LogLevelWarn)
	return i
}

func (i *Interactor) AddLogLevelError() *Interactor {
	i.addLogLevel(LoggerDomain.LogLevelError)
	return i
}

func (i *Interactor) AddLogLevelFatal() *Interactor {
	i.addLogLevel(LoggerDomain.LogLevelFatal)
	return i
}

func (i *Interactor) RemoveLogLevelDebug() *Interactor {
	i.removeLogLevel(LoggerDomain.LogLevelDebug)
	return i
}

func (i *Interactor) RemoveLogLevelInfo() *Interactor {
	i.removeLogLevel(LoggerDomain.LogLevelInfo)
	return i
}

func (i *Interactor) RemoveLogLevelWarn() *Interactor {
	i.removeLogLevel(LoggerDomain.LogLevelWarn)
	return i
}

func (i *Interactor) RemoveLogLevelError() *Interactor {
	i.removeLogLevel(LoggerDomain.LogLevelError)
	return i
}

func (i *Interactor) RemoveLogLevelFatal() *Interactor {
	i.removeLogLevel(LoggerDomain.LogLevelFatal)
	return i
}

func (i *Interactor) setLogLevel(level int) *Interactor {
	i.logLevel = level
	return i
}

func (i *Interactor) getLogLevel() (level int) {
	return i.logLevel
}

func (i *Interactor) addLogLevel(level int) *Interactor {
	i.logLevel = i.logLevel | level
	return i
}

func (i *Interactor) removeLogLevel(level int) *Interactor {
	i.logLevel = i.logLevel &^ level
	return i
}
