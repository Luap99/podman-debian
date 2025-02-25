package events

import (
	"fmt"
	"strings"

	"github.com/sirupsen/logrus"
)

// NewEventer creates an eventer based on the eventer type
func NewEventer(options EventerOptions) (Eventer, error) {
	logrus.Debugf("Initializing event backend %s", options.EventerType)
	switch strings.ToUpper(options.EventerType) {
	case strings.ToUpper(LogFile.String()):
		return EventLogFile{options}, nil
	case strings.ToUpper(Null.String()):
		return newNullEventer(), nil
	default:
		return nil, fmt.Errorf("unknown event logger type: %s", strings.ToUpper(options.EventerType))
	}
}
