package LoggerInteractor

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
)

func (i *Interactor) SetLogLevelNone() *Interactor {
	i.logLevel = domain.LogLevelNone
	return i
}

func (i *Interactor) AddLogLevelDebug() *Interactor {
	i.logLevel += domain.LogLevelDebug
	return i
}

func (i *Interactor) AddLogLevelInfo() *Interactor {
	i.logLevel += domain.LogLevelInfo
	return i
}

func (i *Interactor) AddLogLevelWarn() *Interactor {
	i.logLevel += domain.LogLevelWarn
	return i
}

func (i *Interactor) AddLogLevelError() *Interactor {
	i.logLevel += domain.LogLevelError
	return i
}

func (i *Interactor) AddLogLevelFatal() *Interactor {
	i.logLevel += domain.LogLevelFatal
	return i
}

func (i *Interactor) RemoveLogLevelDebug() *Interactor {
	i.logLevel -= domain.LogLevelDebug
	return i
}

func (i *Interactor) RemoveLogLevelInfo() *Interactor {
	i.logLevel -= domain.LogLevelInfo
	return i
}

func (i *Interactor) RemoveLogLevelWarn() *Interactor {
	i.logLevel -= domain.LogLevelWarn
	return i
}

func (i *Interactor) RemoveLogLevelError() *Interactor {
	i.logLevel -= domain.LogLevelError
	return i
}

func (i *Interactor) RemoveLogLevelFatal() *Interactor {
	i.logLevel -= domain.LogLevelFatal
	return i
}

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
