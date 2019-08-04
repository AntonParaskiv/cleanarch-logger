package logger

import "os"

type StdStreamStorage struct {
	stream *os.File
}

func NewStdStreamStorage(stream *os.File) (s *StdStreamStorage) {
	s = new(StdStreamStorage)
	s.stream = stream
	return
}

func (s *StdStreamStorage) Send(message string) (err error) {
	_, err = s.stream.Write([]byte(message))
	return
}
