package LoggerRepositoryMock

import "time"

func (m *Mock) Log(level int, time time.Time, message string) (err error) {
	if m.simulateErrorFlag {
		return m.Error()
	}

	m.logLevel = level
	m.logTime = time
	m.message = message
	return
}
