package LoggerStreamStorage

//import (
//	"io/ioutil"
//	"os"
//	"testing"
//)
//
//const (
//	testFileName    = "/tmp/loggerFileStorage.test.log"
//	testFilePerm    = 0644
//	testFileMessage = "abrakadbra"
//)
//
//var fileStorage *FileStorage
//
//// failed create
//func TestNewFileStorage_1(t *testing.T) {
//	var err error
//
//	// if exist - remove dir
//	if _, err = os.Stat(testFileName); err == nil {
//		if err = os.Remove(testFileName); err != nil {
//			t.Error("dir exist and remove failed: " + err.Error())
//			return
//		}
//	}
//
//	// create dir
//	if err = os.Mkdir(testFileName, testFilePerm); err != nil {
//		t.Error("dir create failed: " + err.Error())
//		return
//	}
//
//	// try create file
//	fileStorage, err = NewFileStorage("/tmp/", testFilePerm)
//	if err == nil {
//		t.Error("should be error, but not reached")
//		return
//	}
//}
//
//// success create
//func TestNewFileStorage_2(t *testing.T) {
//	var (
//		fileInfo os.FileInfo
//		err      error
//	)
//
//	// if exist - remove file
//	if _, err = os.Stat(testFileName); err == nil {
//		if err = os.Remove(testFileName); err != nil {
//			t.Error("file exist and remove failed: " + err.Error())
//			return
//		}
//	}
//
//	// create file
//	fileStorage, err = NewFileStorage(testFileName, testFilePerm)
//	if err != nil {
//		t.Error("file create failed: " + err.Error())
//		return
//	}
//
//	// if not exist - error
//	if fileInfo, err = os.Stat(testFileName); err != nil {
//		t.Error("file was created and not exist: " + err.Error())
//		return
//	}
//
//	// check perm
//	if fileInfo.Mode() != testFilePerm {
//		t.Error("permissions is not match")
//		return
//	}
//}
//
//func TestFileStorage_Send(t *testing.T) {
//	var err error
//
//	// send message
//	if err = fileStorage.Send(testFileMessage); err != nil {
//		t.Error("send method failed: " + err.Error())
//		return
//	}
//
//	// check read file
//	fileContent, err := ioutil.ReadFile(testFileName)
//	if err != nil {
//		t.Error("read file failed: " + err.Error())
//		return
//	}
//
//	// compare content
//	if string(fileContent) != testFileMessage {
//		t.Error("content is not match")
//		return
//	}
//
//}
