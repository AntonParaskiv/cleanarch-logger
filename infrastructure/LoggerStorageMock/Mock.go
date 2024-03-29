package LoggerStorageMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

type Mock struct {
	Message           string
	simulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) Send(message string) (err error) {
	if m.simulateErrorFlag {
		return m.Error()
	}

	m.Message = message
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
