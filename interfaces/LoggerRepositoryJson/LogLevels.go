package LoggerRepositoryJson

import (
	"github.com/AntonParaskiv/cleanarch-logger/domain/LoggerDomain"
)

func (r *Repository) SetLogLevelNone() *Repository {
	r.logLevel = LoggerDomain.LogLevelNone
	return r
}

func (r *Repository) AddLogLevelDebug() *Repository {
	r.logLevel = r.logLevel | LoggerDomain.LogLevelDebug
	return r
}

func (r *Repository) AddLogLevelInfo() *Repository {
	r.logLevel = r.logLevel | LoggerDomain.LogLevelInfo
	return r
}

func (r *Repository) AddLogLevelWarn() *Repository {
	r.logLevel = r.logLevel | LoggerDomain.LogLevelWarn
	return r
}

func (r *Repository) AddLogLevelError() *Repository {
	r.logLevel = r.logLevel | LoggerDomain.LogLevelError
	return r
}

func (r *Repository) AddLogLevelFatal() *Repository {
	r.logLevel = r.logLevel | LoggerDomain.LogLevelFatal
	return r
}

func (r *Repository) RemoveLogLevelDebug() *Repository {
	r.logLevel = r.logLevel &^ LoggerDomain.LogLevelDebug
	return r
}

func (r *Repository) RemoveLogLevelInfo() *Repository {
	r.logLevel = r.logLevel &^ LoggerDomain.LogLevelInfo
	return r
}

func (r *Repository) RemoveLogLevelWarn() *Repository {
	r.logLevel = r.logLevel &^ LoggerDomain.LogLevelWarn
	return r
}

func (r *Repository) RemoveLogLevelError() *Repository {
	r.logLevel = r.logLevel &^ LoggerDomain.LogLevelError
	return r
}

func (r *Repository) RemoveLogLevelFatal() *Repository {
	r.logLevel = r.logLevel &^ LoggerDomain.LogLevelFatal
	return r
}

func (i *Repository) SetLogLevelAll() *Repository {
	i.SetLogLevelNone()
	i.AddLogLevelDebug()
	i.AddLogLevelInfo()
	i.AddLogLevelWarn()
	i.AddLogLevelError()
	i.AddLogLevelFatal()
	return i
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
