package logger

import (
	"encoding/json"
	domain "github.com/AntonParaskiv/cleanarch-logger/domain/logger"
	"github.com/pkg/errors"
	"strings"
	"time"
)

var logLevelJsonTitle = map[int]string{
	domain.LogLevelFatal: "fatal",
	domain.LogLevelError: "error",
	domain.LogLevelWarn:  "warn",
	domain.LogLevelInfo:  "info",
	domain.LogLevelDebug: "debug",
}

type LoggerJsonStorage interface {
	Send(message string) (err error)
}

type LogJsonRepository struct {
	name       string
	storage    LoggerJsonStorage
	logLevel   int
	prefix     string
	timeFormat string
}

func NewLogJsonRepository(name string, storage LoggerJsonStorage, logLevel int, prefix string, timeFormat string) (r *LogJsonRepository) {
	r = new(LogJsonRepository)
	r.name = name
	r.storage = storage
	r.logLevel = logLevel
	r.prefix = prefix
	r.timeFormat = timeFormat
	return
}

type LogMessage struct {
	Level     string `json:"level,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (m *LogMessage) JsonMarshal() string {
	jsoned, _ := json.Marshal(m)
	return string(jsoned)
}

func (r *LogJsonRepository) Log(level int, time time.Time, message string) (err error) {
	if r.logLevel&level == level {
		logMessage := LogMessage{}

		message = strings.Trim(message, " \t\r\n")

		if r.prefix != "" {
			message = r.prefix + " " + message
		}
		logMessage.Message = message

		if logLevelJsonTitle[level] != "" {
			logMessage.Level = logLevelJsonTitle[level]
		}

		if r.timeFormat != "" {
			logMessage.Timestamp = time.Format(r.timeFormat)
		}

		if err = r.storage.Send(logMessage.JsonMarshal() + "\n"); err != nil {
			err = errors.New("log storage send failed:" + err.Error())
			return
		}
	}
	return
}

func (r *LogJsonRepository) GetName() (name string) {
	return r.name
}
