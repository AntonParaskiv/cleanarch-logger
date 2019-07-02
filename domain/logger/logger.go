package logger

const (
	_             = 1 << iota // 1
	_             = 1 << iota // 2
	LogLevelFatal = 1 << iota // 4
	LogLevelError = 1 << iota // 8
	LogLevelWarn  = 1 << iota // 16
	_             = 1 << iota // 32
	LogLevelInfo  = 1 << iota // 64
	LogLevelDebug = 1 << iota // 128

	LogLevelMaskAll           = -1 // all bits is 1
	LogLevelMaskDebugInfoWarn = LogLevelDebug | LogLevelInfo | LogLevelWarn
	LogLevelMaskInfoWarn      = LogLevelInfo | LogLevelWarn
	LogLevelMaskErrorFatal    = LogLevelError | LogLevelFatal
)
