package LoggerRepositoryMock

import "time"

func (m *Mock) Log(level int, time time.Time, message string) (err error) {
	if m.simulateErrorFlag {
		return m.Error()
	}

	m.LogLevel = level
	m.LogTime = time
	m.Message = message
	return
}
