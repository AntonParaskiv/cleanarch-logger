package LoggerDomain

const (
	LogLevelAllBits = -1        // all bits set 1
	LogLevelNone    = 0         // 0
	_               = 1 << iota // 1
	_               = 1 << iota // 2
	LogLevelFatal   = 1 << iota // 4
	LogLevelError   = 1 << iota // 8
	LogLevelWarn    = 1 << iota // 16
	_               = 1 << iota // 32
	LogLevelInfo    = 1 << iota // 64
	LogLevelDebug   = 1 << iota // 128
)
