package LoggerRepositoryMock

import (
	"github.com/pkg/errors"
	"time"
)

const ErrorSimulated = "simulated error"

type Mock struct {
	name              string
	logLevel          int
	logTime           time.Time
	message           string
	simulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) SetName(name string) *Mock {
	m.name = name
	return m
}

func (m *Mock) GetName() (name string) {
	return m.name
}

func (m *Mock) SetLogLevel(level int) *Mock {
	m.logLevel = level
	return m
}

func (m *Mock) GetLogLevel() (level int) {
	return m.logLevel
}

func (m *Mock) SetLogTime(time time.Time) *Mock {
	m.logTime = time
	return m
}

func (m *Mock) GetLogTime() (time time.Time) {
	return m.logTime
}

func (m *Mock) SetMessage(message string) *Mock {
	m.message = message
	return m
}

func (m *Mock) GetMessage() (message string) {
	return m.message
}

func (m *Mock) SimulateError() *Mock {
	m.simulateErrorFlag = true
	return m
}

func (m *Mock) IsSetSimulateError() (IsSetSimulateError bool) {
	return m.simulateErrorFlag
}

func (m *Mock) Error() (err error) {
	m.simulateErrorFlag = false
	err = errors.Errorf(ErrorSimulated)
	return
}
