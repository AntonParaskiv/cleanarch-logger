package LoggerStreamStorage

//
//import (
//	"github.com/pkg/errors"
//	"os"
//)
//
//type FileStorage struct {
//	file *os.File
//}
//
//func NewFileStorage(fileName string, perm os.FileMode) (f *FileStorage, err error) {
//	f = new(FileStorage)
//	f.file, err = os.OpenFile(fileName, os.O_WRONLY|os.O_APPEND|os.O_CREATE, perm)
//	if err != nil {
//		err = errors.New("open file failed: " + err.Error())
//		return
//	}
//	return
//}
//
//func (s *FileStorage) Send(message string) (err error) {
//	_, err = s.file.Write([]byte(message))
//	return
//}
