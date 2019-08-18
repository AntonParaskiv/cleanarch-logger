package LoggerRepositoryMock

import "time"

func (m *Mock) Log(level int, time time.Time, message string) (err error) {
	if m.IsSetSimulateError() {
		return m.Error()
	}

	m.SetLogLevel(level)
	m.SetLogTime(time)
	m.SetMessage(message)
	return
}
