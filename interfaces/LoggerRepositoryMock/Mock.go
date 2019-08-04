package LoggerRepositoryMock

import (
	"github.com/pkg/errors"
	"time"
)

const ErrorSimulated = "simulated error"

type Mock struct {
	Name              string
	LogLevel          int
	LogTime           time.Time
	Message           string
	simulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) SimulateError() *Mock {
	m.simulateErrorFlag = true
	return m
}

func (m *Mock) Error() (err error) {
	m.simulateErrorFlag = false
	err = errors.Errorf(ErrorSimulated)
	return
}

func (m *Mock) GetName() (name string) {
	return m.Name
}
