package LoggerStorageMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

type Mock struct {
	message           string
	simulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) Send(message string) (err error) {
	if m.IsSetSimulateError() {
		return m.Error()
	}

	m.SetMessage(message)
	return
}

func (m *Mock) SetMessage(message string) *Mock {
	m.message = message
	return m
}

func (m *Mock) Message() (message string) {
	message = m.message
	return
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
