package FileMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

type Mock struct {
	Fdescr            uintptr
	Name              string
	Buffer            []byte
	simulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) Write(b []byte) (n int, err error) {
	if m.simulateErrorFlag {
		return 0, m.Error()
	}

	m.Buffer = b
	n = len(b)
	return
}

func (m *Mock) Fd() uintptr {
	return m.Fdescr
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
