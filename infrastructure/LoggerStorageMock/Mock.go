package LoggerStorageMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

type Mock struct {
	Message           string
	SimulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) Send(message string) (err error) {
	if m.SimulateErrorFlag {
		return m.Error()
	}

	m.Message = message
	return
}

func (m *Mock) SimulateError() *Mock {
	m.SimulateErrorFlag = true
	return m
}

func (m *Mock) Error() (err error) {
	m.SimulateErrorFlag = false
	err = errors.Errorf(ErrorSimulated)
	return
}
