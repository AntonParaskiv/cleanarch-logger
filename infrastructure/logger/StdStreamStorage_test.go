package logger

import (
	"errors"
	"io/ioutil"
	"os"
	"testing"
)

const (
	testFileStreamName    = "/tmp/loggerStdStreamStorage.test.log"
	testFileStreamPerm    = 0644
	testFileStreamMessage = "abrakadbra"
)

var (
	testStream    *os.File
	streamStorage *StdStreamStorage
)

func TestNewStdStreamStorage(t *testing.T) {
	var err error

	// if exist - remove file
	if _, err = os.Stat(testFileStreamName); err == nil {
		if err = os.Remove(testFileStreamName); err != nil {
			t.Error("file exist and remove failed: " + err.Error())
			return
		}
	}

	// create file
	testStream, err = os.OpenFile(testFileStreamName, os.O_RDWR|os.O_APPEND|os.O_CREATE, testFileStreamPerm)
	if err != nil {
		err = errors.New("open file failed: " + err.Error())
		return
	}

	// create stream
	streamStorage = NewStdStreamStorage(testStream)
	if streamStorage.stream != testStream {
		t.Error("streams is not match")
		return
	}

}

func TestStdStreamStorage_Send(t *testing.T) {
	var err error

	// send message
	if err = streamStorage.Send(testFileStreamMessage); err != nil {
		t.Error("send method failed: " + err.Error())
		return
	}

	// check read file
	if _, err = streamStorage.stream.Seek(0, 0); err != nil {
		t.Error("stream seek to start failed: " + err.Error())
		return
	}
	fileContent, err := ioutil.ReadAll(streamStorage.stream)
	if err != nil {
		t.Error("read stream failed: " + err.Error())
		return
	}

	// compare content
	if string(fileContent) != testFileStreamMessage {
		t.Error("content is not match")
		return
	}
}
