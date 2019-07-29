package LoggerStreamStorage

import (
	"github.com/pkg/errors"
	"os"
)

type File interface {
	Write(b []byte) (n int, err error)
	Fd() uintptr
}

type Storage struct {
	stream File
}

func New(stream *os.File) (s *Storage) {
	s = new(Storage)
	s.stream = stream
	return
}

func NewFromFileName(fileName string, perm os.FileMode) (storage *Storage, err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		err = errors.Errorf("open file failed: %s", err.Error())
		return
	}

	storage = New(file)
	return
}
