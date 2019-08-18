package LoggerRepository

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (r *Repository) Log(level int, time time.Time, message string) (err error) {
	if !r.isMessageLevelMatchRepositoryLevel(level) {
		return
	}

	message = r.prepareMessage(message, level, time)

	if err = r.sendMessageToStorage(message); err != nil {
		err = errors.Errorf("send message failed: %s", err.Error())
		return
	}

	return
}

func (r *Repository) isMessageLevelMatchRepositoryLevel(level int) (isMatch bool) {
	if r.logLevel&level == level {
		isMatch = true
	}
	return
}

func (r *Repository) prepareMessage(message string, level int, time time.Time) string {
	message = r.clearMessageGarbage(message)
	message = r.addRepositoryPrefix(message)
	message = r.addLogLevelPrefix(message, level)
	message = r.addTimePrefix(message, time)
	return message
}

func (r *Repository) sendMessageToStorage(message string) (err error) {
	if err = r.storage.Send(message); err != nil {
		err = errors.Errorf("storage send failed: %s", err.Error())
		return
	}
	return
}

func (r *Repository) clearMessageGarbage(message string) string {
	message = strings.Trim(message, " \t\r\n") + "\n"
	return message
}

func (r *Repository) addRepositoryPrefix(message string) string {
	if len(r.prefix) > 0 {
		message = r.prefix + " " + message
	}
	return message
}

func (r *Repository) addLogLevelPrefix(message string, level int) string {
	if len(logLevelTitle[level]) > 0 {
		message = logLevelTitle[level] + " " + message
	}
	return message
}

func (r *Repository) addTimePrefix(message string, time time.Time) string {
	if len(r.timeFormat) > 0 {
		message = time.Format(r.timeFormat) + " " + message
	}
	return message
}
