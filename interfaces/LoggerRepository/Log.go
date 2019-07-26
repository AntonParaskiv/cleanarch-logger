package LoggerRepository

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (r *LoggerRepository) Log(level int, time time.Time, message string) (err error) {
	if r.logLevel&level == level {
		message = strings.Trim(message, " \t\r\n") + "\n"

		if r.prefix != "" {
			message = r.prefix + " " + message
		}

		if logLevelTitle[level] != "" {
			message = logLevelTitle[level] + " " + message
		}

		if r.timeFormat != "" {
			message = time.Format(r.timeFormat) + " " + message
		}

		if err = r.storage.Send(message); err != nil {
			err = errors.Errorf("log storage send failed: %s", err.Error())
			return
		}
	}
	return
}
