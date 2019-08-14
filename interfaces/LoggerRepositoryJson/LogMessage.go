package LoggerRepositoryJson

import (
	"encoding/json"
	"github.com/pkg/errors"
)

type LogMessage struct {
	Level     string `json:"level,omitempty"`
	Timestamp string `json:"timestamp,omitempty"`
	Message   string `json:"message,omitempty"`
}

func NewLogMessage(message string, level string, time string) (logMessage *LogMessage) {
	logMessage = new(LogMessage)
	logMessage.Message = message
	logMessage.Level = level
	logMessage.Timestamp = time
	return
}

func (lm *LogMessage) MakeJson() (jsnString string, err error) {
	jsn, err := json.Marshal(lm)
	if err != nil {
		err = errors.Errorf("log message marshal failed: %s", err.Error())
		return
	}
	jsn = append(jsn, byte('\n'))
	jsnString = string(jsn)
	return
}
