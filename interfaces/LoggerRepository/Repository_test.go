package LoggerRepository

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
	"github.com/pkg/errors"
	"strings"
	"testing"
	"time"
)

type LoggerStorageMock struct {
	message string
}

func (s *LoggerStorageMock) Send(message string) (err error) {
	if message == "error"+"\n" {
		err = errors.New("send failed")
		return
	}

	if message == `{"message":"error"}`+"\n" {
		err = errors.New("send failed")
		return
	}

	s.message = message
	return
}

const (
	logRepoName       = "testLogRepository"
	logRepoPrefix     = "prefix"
	logRepoLevel      = domain.LogLevelMaskInfoWarn
	logRepoTimeFormat = time.Stamp
)

var (
	logStorageMock *LoggerStorageMock
	logRepository  *LoggerRepository
)

func TestNewLogRepository(t *testing.T) {
	// create log storage mock
	logStorageMock = new(LoggerStorageMock)

	// create log logRepository
	logRepository = New(logRepoName, logStorageMock, logRepoLevel, logRepoPrefix, logRepoTimeFormat)
	if logRepository.name != logRepoName {
		t.Error("logRepository name is not match")
		return
	}

	// check log logRepository properties
	if logRepository.storage != logStorageMock {
		t.Error("logRepository storage is not match")
		return
	}

	if logRepository.logLevel != logRepoLevel {
		t.Error("logRepository logLevel is not match")
		return
	}

	if logRepository.prefix != logRepoPrefix {
		t.Error("logRepository prefix is not match")
		return
	}

	if logRepository.timeFormat != logRepoTimeFormat {
		t.Error("logRepository timeFormat is not match")
		return
	}
}

// must send none
func TestLogRepository_Log_1(t *testing.T) {
	timeNow := time.Now()
	message := "abrakadbra"
	if err := logRepository.Log(domain.LogLevelDebug, timeNow, message); err != nil {
		t.Error("log failed: " + err.Error())
		return
	}

	// check empty Message
	if len(logStorageMock.message) > 0 {
		t.Error("log must send none, but it send something")
		return
	}
}

// must send Message
func TestLogRepository_Log_2(t *testing.T) {
	// send Message
	logLevel := domain.LogLevelInfo
	timeNow := time.Now()
	message := " \t\r\nabrakadbra \t\r\n"
	if err := logRepository.Log(logLevel, timeNow, message); err != nil {
		t.Error("log failed: " + err.Error())
		return
	}

	// prepare expect
	messageExpect := timeNow.Format(logRepoTimeFormat) + " " +
		logLevelTitle[logLevel] + " " +
		logRepoPrefix + " " +
		strings.Trim(message, " \t\r\n") + "\n"

	// compare
	if logStorageMock.message != messageExpect {
		t.Error("storage Message is not match")
		return
	}
}

// must return err
func TestLogRepository_Log_3(t *testing.T) {
	// send Message
	logRepository.prefix = ""
	logRepository.timeFormat = ""
	if err := logRepository.Log(0, time.Time{}, "error"); err == nil {
		t.Error("should be error, but not reached")
		return
	}
}

// getname
func TestLogRepository_GetName(t *testing.T) {
	if logRepository.GetName() != logRepoName {
		t.Error("GetName is not match")
		return
	}
}
