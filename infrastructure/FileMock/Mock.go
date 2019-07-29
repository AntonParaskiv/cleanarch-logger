package FileMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

type Mock struct {
	Fdescr            uintptr
	Name              string
	Buffer            []byte
	SimulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) Write(b []byte) (n int, err error) {
	if m.SimulateErrorFlag {
		m.SimulateErrorFlag = false
		err = errors.Errorf(ErrorSimulated)
		return
	}
	m.Buffer = b
	n = len(b)
	return
}

func (m *Mock) Fd() uintptr {
	return m.Fdescr
}

func (m *Mock) SimulateError() *Mock {
	m.SimulateErrorFlag = true
	return m
}
