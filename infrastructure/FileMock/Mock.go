package FileMock

import "github.com/pkg/errors"

const ErrorSimulated = "simulated error"

type Mock struct {
	Fdescr            uintptr
	Name              string
	buffer            []byte
	simulateErrorFlag bool
}

func New() (m *Mock) {
	m = new(Mock)
	return
}

func (m *Mock) Write(b []byte) (n int, err error) {
	if m.IsSetSimulateError() {
		return 0, m.Error()
	}

	m.SetBuffer(b)
	n = m.BufferLen()
	return
}

func (m *Mock) SetBuffer(b []byte) *Mock {
	m.buffer = b
	return m
}

func (m *Mock) Buffer() (b []byte) {
	b = m.buffer
	return
}

func (m *Mock) BufferLen() (bufferLen int) {
	bufferLen = len(m.buffer)
	return
}

func (m *Mock) Fd() uintptr {
	return m.Fdescr
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
