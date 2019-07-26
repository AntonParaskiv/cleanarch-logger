package LoggerStreamStorage

import (
	"github.com/pkg/errors"
	"os"
)

type Storage struct {
	stream *os.File
}

func New(stream *os.File) (s *Storage) {
	s = new(Storage)
	s.stream = stream
	return
}

func NewFromFileName(fileName string, perm os.FileMode) (storage *Storage, err error) {
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm)
	if err != nil {
		err = errors.Errorf("open file failed: %", err.Error())
		return
	}

	storage = New(file)
	return
}
