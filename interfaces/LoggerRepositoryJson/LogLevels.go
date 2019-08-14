package LoggerRepositoryJson

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
)

func (r *Repository) SetLogLevelAll() *Repository {
	r.SetLogLevelNone()
	r.AddLogLevelDebug()
	r.AddLogLevelInfo()
	r.AddLogLevelWarn()
	r.AddLogLevelError()
	r.AddLogLevelFatal()
	return r
}

func (r *Repository) SetLogLevelDebugInfoWarn() *Repository {
	r.SetLogLevelNone()
	r.AddLogLevelDebug()
	r.AddLogLevelInfo()
	r.AddLogLevelWarn()
	return r
}

func (r *Repository) SetLogLevelInfoWarn() *Repository {
	r.SetLogLevelNone()
	r.AddLogLevelInfo()
	r.AddLogLevelWarn()
	return r
}

func (r *Repository) SetLogLevelErrorFatal() *Repository {
	r.SetLogLevelNone()
	r.AddLogLevelError()
	r.AddLogLevelFatal()
	return r
}

func (r *Repository) SetLogLevelNone() *Repository {
	r.setLogLevel(LoggerDomain.LogLevelNone)
	return r
}

func (r *Repository) AddLogLevelDebug() *Repository {
	r.addLogLevel(LoggerDomain.LogLevelDebug)
	return r
}

func (r *Repository) AddLogLevelInfo() *Repository {
	r.addLogLevel(LoggerDomain.LogLevelInfo)
	return r
}

func (r *Repository) AddLogLevelWarn() *Repository {
	r.addLogLevel(LoggerDomain.LogLevelWarn)
	return r
}

func (r *Repository) AddLogLevelError() *Repository {
	r.addLogLevel(LoggerDomain.LogLevelError)
	return r
}

func (r *Repository) AddLogLevelFatal() *Repository {
	r.addLogLevel(LoggerDomain.LogLevelFatal)
	return r
}

func (r *Repository) RemoveLogLevelDebug() *Repository {
	r.removeLogLevel(LoggerDomain.LogLevelDebug)
	return r
}

func (r *Repository) RemoveLogLevelInfo() *Repository {
	r.removeLogLevel(LoggerDomain.LogLevelInfo)
	return r
}

func (r *Repository) RemoveLogLevelWarn() *Repository {
	r.removeLogLevel(LoggerDomain.LogLevelWarn)
	return r
}

func (r *Repository) RemoveLogLevelError() *Repository {
	r.removeLogLevel(LoggerDomain.LogLevelError)
	return r
}

func (r *Repository) RemoveLogLevelFatal() *Repository {
	r.removeLogLevel(LoggerDomain.LogLevelFatal)
	return r
}

func (r *Repository) setLogLevel(level int) *Repository {
	r.logLevel = level
	return r
}

func (r *Repository) addLogLevel(level int) *Repository {
	r.logLevel = r.logLevel | level
	return r
}

func (r *Repository) removeLogLevel(level int) *Repository {
	r.logLevel = r.logLevel &^ level
	return r
}
