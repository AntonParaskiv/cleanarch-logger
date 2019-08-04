package usecases

import (
	"fmt"
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
	"github.com/pkg/errors"
	"reflect"
	"testing"
	"time"
)

const (
	loggerInteractorLogLevel = domain.LogLevelMaskInfoWarn
)

var (
	loggerRepositoryMock *LoggerRepositoryMock
	loggerInteractor     *LoggerInteractor
)

type LoggerRepositoryMock struct {
	name     string
	logLevel int
	logTime  time.Time
	message  string
}

func (m *LoggerRepositoryMock) Log(level int, time time.Time, message string) (err error) {
	if message == "need error" {
		err = errors.New("Log failed successfully")
		return
	}

	m.logLevel = level
	m.logTime = time
	m.message = message
	return
}

func (m *LoggerRepositoryMock) GetName() (name string) {
	return
}

func TestNewLoggerInteractor(t *testing.T) {
	// create repos
	loggerRepositoryMock = new(LoggerRepositoryMock)
	repositories := []LoggerRepository{
		loggerRepositoryMock,
	}

	// create interactor
	loggerInteractor = NewLoggerInteractor(repositories, loggerInteractorLogLevel)

	// check properties
	if !reflect.DeepEqual(loggerInteractor.Repositories, repositories) {
		t.Error("interactor repositories are not match")
		return
	}
	if loggerInteractor.logLevel != loggerInteractorLogLevel {
		t.Error("interactor logLevel is not match")
		return
	}
}

// must Log none
func TestLoggerInteractor_log_1(t *testing.T) {
	// Log message
	loglevel := domain.LogLevelDebug
	message := "message"
	loggerInteractor.log(loglevel, message)

	// check repo
	if len(loggerRepositoryMock.message) > 0 {
		t.Error("interactor must Log none, but it Log something")
		return
	}
}

// must Log successful
func TestLoggerInteractor_log_2(t *testing.T) {
	// Log message with timer
	loglevel := domain.LogLevelInfo
	message := "message"
	timeStart := time.Now()
	loggerInteractor.log(loglevel, message)
	timeStop := time.Now()

	// check repo
	if loggerRepositoryMock.logLevel != loglevel {
		t.Error("interactor logLevel is not match")
		return
	}

	if loggerRepositoryMock.logTime.Before(timeStart) || loggerRepositoryMock.logTime.After(timeStop) {
		t.Error("interactor timestamp is not match")
		return
	}

	if loggerRepositoryMock.message != message {
		t.Error("interactor message is not match")
		return
	}
}

// must return error
func TestLoggerInteractor_log_3(t *testing.T) {
	// Log message
	loglevel := domain.LogLevelInfo
	message := "need error"
	loggerInteractor.log(loglevel, message)
}

// must logf successful
func TestLoggerInteractor_logf(t *testing.T) {
	// Log message
	loglevel := domain.LogLevelInfo
	messageFormat := "%s %d %t"
	args := []interface{}{"message", 123, true}
	loggerInteractor.logf(loglevel, messageFormat, args...)

	// check repo
	messageExpect := fmt.Sprintf(messageFormat, args...)
	if loggerRepositoryMock.message != messageExpect {
		t.Error("interactor messageExpect is not match")
		return
	}
}

// must debugMode successful
func TestLoggerInteractor_Debug(t *testing.T) {
	loggerInteractor.logLevel = domain.LogLevelMaskAll

	// Log message
	message := "debugMode"
	loggerInteractor.Debug(message)

	// check repo
	if loggerRepositoryMock.message != message {
		t.Error("interactor message is not match")
		return
	}
}

// must debugf successful
func TestLoggerInteractor_Debugf(t *testing.T) {
	// Log message
	messageFormat := "debugf %s %d %t"
	args := []interface{}{"message", 123, true}
	loggerInteractor.Debugf(messageFormat, args...)

	// check repo
	messageExpect := fmt.Sprintf(messageFormat, args...)
	if loggerRepositoryMock.message != messageExpect {
		t.Error("interactor message is not match")
		return
	}
}

// must info successful
func TestLoggerInteractor_Info(t *testing.T) {
	// Log message
	message := "info"
	loggerInteractor.Info(message)

	// check repo
	if loggerRepositoryMock.message != message {
		t.Error("interactor message is not match")
		return
	}
}

// must infof successful
func TestLoggerInteractor_Infof(t *testing.T) {
	// Log message
	messageFormat := "infof %s %d %t"
	args := []interface{}{"message", 123, true}
	loggerInteractor.Infof(messageFormat, args...)

	// check repo
	messageExpect := fmt.Sprintf(messageFormat, args...)
	if loggerRepositoryMock.message != messageExpect {
		t.Error("interactor message is not match")
		return
	}
}

// must warn successful
func TestLoggerInteractor_Warn(t *testing.T) {
	// Log message
	message := "warn"
	loggerInteractor.Warn(message)

	// check repo
	if loggerRepositoryMock.message != message {
		t.Error("interactor message is not match")
		return
	}
}

// must warnf successful
func TestLoggerInteractor_Warnf(t *testing.T) {
	// Log message
	messageFormat := "warnf %s %d %t"
	args := []interface{}{"message", 123, true}
	loggerInteractor.Warnf(messageFormat, args...)

	// check repo
	messageExpect := fmt.Sprintf(messageFormat, args...)
	if loggerRepositoryMock.message != messageExpect {
		t.Error("interactor message is not match")
		return
	}
}

// must error successful
func TestLoggerInteractor_Error(t *testing.T) {
	// Log message
	message := "error"
	loggerInteractor.Error(message)

	// check repo
	if loggerRepositoryMock.message != message {
		t.Error("interactor message is not match")
		return
	}
}

// must errorf successful
func TestLoggerInteractor_Errorf(t *testing.T) {
	// Log message
	messageFormat := "errorf %s %d %t"
	args := []interface{}{"message", 123, true}
	loggerInteractor.Errorf(messageFormat, args...)

	// check repo
	messageExpect := fmt.Sprintf(messageFormat, args...)
	if loggerRepositoryMock.message != messageExpect {
		t.Error("interactor message is not match")
		return
	}
}

// must fatal successful
func TestLoggerInteractor_Fatal(t *testing.T) {
	// Log message
	message := "fatal"
	loggerInteractor.Fatal(message)

	// check repo
	if loggerRepositoryMock.message != message {
		t.Error("interactor message is not match")
		return
	}
}

// must fatalf successful
func TestLoggerInteractor_Fatalf(t *testing.T) {
	// Log message
	messageFormat := "fatalf %s %d %t"
	args := []interface{}{"message", 123, true}
	loggerInteractor.Fatalf(messageFormat, args...)

	// check repo
	messageExpect := fmt.Sprintf(messageFormat, args...)
	if loggerRepositoryMock.message != messageExpect {
		t.Error("interactor message is not match")
		return
	}
}
