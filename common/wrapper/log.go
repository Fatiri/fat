package wrapper

import (
	"time"

	"github.com/sirupsen/logrus"
)

type log struct {
	Entry      *logrus.Entry
	Collection error
}

func newResponseLog(collection error) log {
	logrus.SetFormatter(&logrus.TextFormatter{})
	fields := logrus.Fields{
		"component":  "response",
		"collection": collection,
		"time":       time.Now(),
	}

	return log{
		Entry:      logrus.WithFields(fields),
		Collection: collection,
	}
}

func (l log) Show() {
	if newError, isNewError := l.Collection.(*Error); isNewError {
		l.Entry.Warningln("err: ", newError.Err, ", location_error: ", newError.ErrLocation)
	} else {
		l.Entry.Warningln(l.Collection.Error())
	}
}
