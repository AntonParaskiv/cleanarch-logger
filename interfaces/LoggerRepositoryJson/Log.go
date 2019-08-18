package LoggerRepositoryJson

import (
	"github.com/pkg/errors"
	"strings"
	"time"
)

func (r *Repository) Log(level int, time time.Time, message string) (err error) {
	if !r.isMessageLevelMatchRepositoryLevel(level) {
		return
	}

	message, err = r.prepareMessage(message, level, time)
	if err != nil {
		err = errors.Errorf("prepare message failed: %s", err.Error())
		return
	}

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

func (r *Repository) prepareMessage(message string, level int, time time.Time) (string, error) {
	message = r.clearMessageGarbage(message)
	message = r.addRepositoryPrefix(message)
	logLevelPrefix := r.getLogLevelPrefix(level)
	timePrefix := r.getTimePrefix(time)

	logMessage := NewLogMessage(message, logLevelPrefix, timePrefix)
	jsn, err := logMessage.MakeJson()
	if err != nil {
		err = errors.Errorf("make json failed: %s", err.Error())
		return "", err
	}

	return jsn, err
}

func (r *Repository) sendMessageToStorage(message string) (err error) {
	if err = r.storage.Send(message); err != nil {
		err = errors.Errorf("storage send failed: %s", err.Error())
		return
	}
	return
}

func (r *Repository) clearMessageGarbage(message string) string {
	message = strings.Trim(message, " \t\r\n")
	return message
}

func (r *Repository) addRepositoryPrefix(message string) string {
	if len(r.prefix) > 0 {
		message = r.prefix + " " + message
	}
	return message
}

func (r *Repository) getLogLevelPrefix(level int) (logLevelPrefix string) {
	if len(logLevelTitle[level]) > 0 {
		logLevelPrefix = logLevelTitle[level]
	}
	return logLevelPrefix
}

func (r *Repository) getTimePrefix(time time.Time) (timePrefix string) {
	if len(r.timeFormat) > 0 {
		timePrefix = time.Format(r.timeFormat)
	}
	return
}
