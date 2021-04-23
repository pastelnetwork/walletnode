package badger

import "github.com/pastelnetwork/go-commons/log"

const (
	logPrefix = "[badger]"
)

// Logger wraps go-common logger to implement interface `github.com/dgraph-io/badger.Logger`.
type Logger struct {
	prefix string
}

// Errorf logs a error statement.
func (logger *Logger) Errorf(format string, v ...interface{}) {
	log.Errorf(logger.prefix+format, v...)
}

// Warningf logs a warn statement.
func (logger *Logger) Warningf(format string, v ...interface{}) {
	log.Warnf(logger.prefix+format, v...)
}

// Infof logs a notice statement.
func (logger *Logger) Infof(format string, v ...interface{}) {
	log.Infof(logger.prefix+format, v...)
}

// Debugf logs a debug statement.
func (logger *Logger) Debugf(format string, v ...interface{}) {
	log.Debugf(logger.prefix+format, v...)
}

// NewLogger returns a new Logger instance.
func NewLogger() *Logger {
	return &Logger{
		prefix: logPrefix,
	}
}
