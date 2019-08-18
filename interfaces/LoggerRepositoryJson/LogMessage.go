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
	logMessage = new(LogMessage).
		SetMessage(message).
		SetLevel(level).
		SetTimestamp(time)
	return
}

func (lm *LogMessage) SetMessage(message string) *LogMessage {
	lm.Message = message
	return lm
}

func (lm *LogMessage) SetLevel(level string) *LogMessage {
	lm.Level = level
	return lm
}

func (lm *LogMessage) SetTimestamp(time string) *LogMessage {
	lm.Timestamp = time
	return lm
}

func (lm *LogMessage) MakeJson() (jsnString string, err error) {
	// Marshal
	jsn, err := json.Marshal(lm)
	if err != nil {
		err = errors.Errorf("log message marshal failed: %s", err.Error())
		return
	}

	// Add '\n'
	jsn = append(jsn, byte('\n'))

	jsnString = string(jsn)
	return
}
