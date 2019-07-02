package logger

import (
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
	"strings"
	"testing"
	"time"
)

const (
	logJsonRepoName       = "testLogJsonRepository"
	logJsonRepoPrefix     = "prefix"
	logJsonRepoLevel      = domain.LogLevelMaskInfoWarn
	logJsonRepoTimeFormat = time.Stamp
)

var (
	logJsonStorageMock *LoggerStorageMock
	logJsonRepository  *LogJsonRepository
)

func TestNewLogJsonRepository(t *testing.T) {
	// create log storage mock
	logJsonStorageMock = new(LoggerStorageMock)

	// create log json repository
	logJsonRepository = NewLogJsonRepository(logJsonRepoName, logJsonStorageMock, logJsonRepoLevel, logJsonRepoPrefix, logJsonRepoTimeFormat)
	if logJsonRepository.name != logJsonRepoName {
		t.Error("logJsonRepository name is not match")
		return
	}

	// check log json repository properties
	if logJsonRepository.storage != logJsonStorageMock {
		t.Error("logJsonRepository storage is not match")
		return
	}

	if logJsonRepository.logLevel != logJsonRepoLevel {
		t.Error("logJsonRepository logLevel is not match")
		return
	}

	if logJsonRepository.prefix != logJsonRepoPrefix {
		t.Error("logJsonRepository prefix is not match")
		return
	}

	if logJsonRepository.timeFormat != logJsonRepoTimeFormat {
		t.Error("logJsonRepository timeFormat is not match")
		return
	}
}

func TestLogMessage_JsonMarshal(t *testing.T) {
	message := LogMessage{
		Level:     "Level",
		Timestamp: "Timestamp",
		Message:   "Message",
	}
	jsnExpect := `{"level":"Level","timestamp":"Timestamp","message":"Message"}`

	jsn := message.JsonMarshal()
	if jsn != jsnExpect {
		t.Error("result is not equal to expect")
		return
	}
}

// must send none
func TestLogJsonRepository_Log_1(t *testing.T) {
	timeNow := time.Now()
	message := "abrakadbra"
	if err := logJsonRepository.Log(domain.LogLevelDebug, timeNow, message); err != nil {
		t.Error("log failed: " + err.Error())
		return
	}

	// check empty Message
	if len(logJsonStorageMock.message) > 0 {
		t.Error("log must send none, but it send something")
		return
	}
}

// must send Message
func TestLogJsonRepository_Log_2(t *testing.T) {
	// send Message
	logLevel := domain.LogLevelInfo
	timeNow := time.Now()
	message := " \t\r\nabrakadbra \t\r\n"
	if err := logJsonRepository.Log(logLevel, timeNow, message); err != nil {
		t.Error("log failed: " + err.Error())
		return
	}

	// prepare expect
	messageExpect := `{"level":"` + logLevelJsonTitle[logLevel] +
		`","timestamp":"` + timeNow.Format(logJsonRepoTimeFormat) +
		`","message":"` + logJsonRepoPrefix + " " + strings.Trim(message, " \t\r\n") + `"}` + "\n"

	// compare
	if logJsonStorageMock.message != messageExpect {
		t.Error("storage Message is not match")
		return
	}
}

// must return err
func TestLogJsonRepository_Log_3(t *testing.T) {
	// send Message
	logJsonRepository.prefix = ""
	logJsonRepository.timeFormat = ""
	if err := logJsonRepository.Log(0, time.Time{}, "error"); err == nil {
		t.Error("should be error, but not reached")
		return
	}
}

// getname
func TestLogJsonRepository_GetName(t *testing.T) {
	if logJsonRepository.GetName() != logJsonRepoName {
		t.Error("GetName is not match")
		return
	}
}
