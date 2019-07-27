package LoggerInteractor

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
)

func (i *LoggerInteractor) SetLogLevelNone() *LoggerInteractor {
	i.logLevel = domain.LogLevelNone
	return i
}

func (i *LoggerInteractor) AddLogLevelDebug() *LoggerInteractor {
	i.logLevel += domain.LogLevelDebug
	return i
}

func (i *LoggerInteractor) AddLogLevelInfo() *LoggerInteractor {
	i.logLevel += domain.LogLevelInfo
	return i
}

func (i *LoggerInteractor) AddLogLevelWarn() *LoggerInteractor {
	i.logLevel += domain.LogLevelWarn
	return i
}

func (i *LoggerInteractor) AddLogLevelError() *LoggerInteractor {
	i.logLevel += domain.LogLevelError
	return i
}

func (i *LoggerInteractor) AddLogLevelFatal() *LoggerInteractor {
	i.logLevel += domain.LogLevelFatal
	return i
}

func (i *LoggerInteractor) RemoveLogLevelDebug() *LoggerInteractor {
	i.logLevel -= domain.LogLevelDebug
	return i
}

func (i *LoggerInteractor) RemoveLogLevelInfo() *LoggerInteractor {
	i.logLevel -= domain.LogLevelInfo
	return i
}

func (i *LoggerInteractor) RemoveLogLevelWarn() *LoggerInteractor {
	i.logLevel -= domain.LogLevelWarn
	return i
}

func (i *LoggerInteractor) RemoveLogLevelError() *LoggerInteractor {
	i.logLevel -= domain.LogLevelError
	return i
}

func (i *LoggerInteractor) RemoveLogLevelFatal() *LoggerInteractor {
	i.logLevel -= domain.LogLevelFatal
	return i
}

func (i *LoggerInteractor) SetLogLevelAll() *LoggerInteractor {
	i.SetLogLevelNone()
	i.AddLogLevelDebug()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	i.AddLogLevelError()
	i.AddLogLevelFatal()
	return i
}

func (i *LoggerInteractor) SetLogLevelDebugInfoWarn() *LoggerInteractor {
	i.SetLogLevelNone()
	i.AddLogLevelDebug()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	return i
}

func (i *LoggerInteractor) SetLogLevelInfoWarn() *LoggerInteractor {
	i.SetLogLevelNone()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	return i
}

func (i *LoggerInteractor) SetLogLevelErrorFatal() *LoggerInteractor {
	i.SetLogLevelNone()
	i.AddLogLevelError()
	i.AddLogLevelFatal()
	return i
}
