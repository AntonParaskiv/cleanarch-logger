package LoggerRepositoryJson

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
)

func (r *LoggerRepository) SetLogLevelNone() *LoggerRepository {
	r.logLevel = domain.LogLevelNone
	return r
}

func (r *LoggerRepository) AddLogLevelDebug() *LoggerRepository {
	r.logLevel += domain.LogLevelDebug
	return r
}

func (r *LoggerRepository) AddLogLevelInfo() *LoggerRepository {
	r.logLevel += domain.LogLevelInfo
	return r
}

func (r *LoggerRepository) AddLogLevelWarn() *LoggerRepository {
	r.logLevel += domain.LogLevelWarn
	return r
}

func (r *LoggerRepository) AddLogLevelError() *LoggerRepository {
	r.logLevel += domain.LogLevelError
	return r
}

func (r *LoggerRepository) AddLogLevelFatal() *LoggerRepository {
	r.logLevel += domain.LogLevelFatal
	return r
}

func (r *LoggerRepository) RemoveLogLevelDebug() *LoggerRepository {
	r.logLevel -= domain.LogLevelDebug
	return r
}

func (r *LoggerRepository) RemoveLogLevelInfo() *LoggerRepository {
	r.logLevel -= domain.LogLevelInfo
	return r
}

func (r *LoggerRepository) RemoveLogLevelWarn() *LoggerRepository {
	r.logLevel -= domain.LogLevelWarn
	return r
}

func (r *LoggerRepository) RemoveLogLevelError() *LoggerRepository {
	r.logLevel -= domain.LogLevelError
	return r
}

func (r *LoggerRepository) RemoveLogLevelFatal() *LoggerRepository {
	r.logLevel -= domain.LogLevelFatal
	return r
}

func (r *LoggerRepository) SetLogLevelDebugInfoWarn() *LoggerRepository {
	r.SetLogLevelNone()
	r.AddLogLevelDebug()
	r.AddLogLevelInfo()
	r.AddLogLevelWarn()
	return r
}

func (r *LoggerRepository) SetLogLevelInfoWarn() *LoggerRepository {
	r.SetLogLevelNone()
	r.AddLogLevelInfo()
	r.AddLogLevelWarn()
	return r
}

func (r *LoggerRepository) SetLogLevelErrorFatal() *LoggerRepository {
	r.SetLogLevelNone()
	r.AddLogLevelError()
	r.AddLogLevelFatal()
	return r
}
