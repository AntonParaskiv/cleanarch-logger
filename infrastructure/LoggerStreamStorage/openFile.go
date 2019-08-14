package LoggerStreamStorage

import (
	"github.com/pkg/errors"
	"os"
)

const openFileFlag = os.O_WRONLY | os.O_APPEND | os.O_CREATE

func openFile(fileName string, perm os.FileMode) (file *os.File, err error) {
	file, err = os.OpenFile(fileName, openFileFlag, perm)
	if err != nil {
		err = errors.Errorf("os open file failed: %s", err.Error())
		return
	}
	return
}
