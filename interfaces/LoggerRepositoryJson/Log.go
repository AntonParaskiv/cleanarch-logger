package LoggerRepositoryJson

import (
	"encoding/json"
	"github.com/pkg/errors"
	"strings"
	"time"
)

type LogMessage struct {
	Level     string `json:"level,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message,omitempty"`
}

func (r *Repository) Log(level int, time time.Time, message string) (err error) {
	if r.logLevel&level == level {
		logMessage := LogMessage{}

		message = strings.Trim(message, " \t\r\n")

		if r.prefix != "" {
			message = r.prefix + " " + message
		}
		logMessage.Message = message

		if logLevelTitle[level] != "" {
			logMessage.Level = logLevelTitle[level]
		}

		if r.timeFormat != "" {
			logMessage.Timestamp = time.Format(r.timeFormat)
		}

		var jsn []byte
		jsn, err = json.Marshal(logMessage)
		if err != nil {
			err = errors.Errorf("message marshal failed: %s", err.Error())
			return
		}

		if err = r.storage.Send(string(jsn) + "\n"); err != nil {
			err = errors.Errorf("log storage send failed: %s", err.Error())
			return
		}
	}
	return
}
