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
	s = new(Storage).
		setStream(stream)
	return
}

func NewFromFileName(fileName string, perm os.FileMode) (storage *Storage, err error) {
	fileStream, err := openFile(fileName, perm)
	if err != nil {
		err = errors.Errorf("open file failed: %s", err.Error())
		return
	}
	storage = New(fileStream)
	return
}

func (s *Storage) setStream(stream *os.File) *Storage {
	s.stream = stream
	return s
}
