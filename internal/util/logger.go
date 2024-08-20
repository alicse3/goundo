package util

import (
	"log"
	"os"
)

// LogLevel type to define the log level
type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
)

// parseLogLevel converts a string log level to LogLevel type
func parseLogLevel(logLevel string) LogLevel {
	switch logLevel {
	case "DEBUG":
		return DEBUG
	case "INFO":
		return INFO
	case "WARN":
		return WARN
	case "ERROR":
		return ERROR
	default:
		return INFO
	}
}

// Logger struct to manage logging with levels
type Logger struct {
	level  LogLevel
	logger *log.Logger
}

// Debug logs a message at the DEBUG level
func (l *Logger) Debug(format string, v ...any) {
	if l.level <= DEBUG {
		l.logger.SetPrefix("DEBUG: ")
		l.logger.Printf(format, v...)
	}
}

// Info logs a message at the INFO level
func (l *Logger) Info(format string, v ...any) {
	if l.level <= INFO {
		l.logger.SetPrefix("INFO: ")
		l.logger.Printf(format, v...)
	}
}

// Warn logs a message at the WARN level
func (l *Logger) Warn(format string, v ...any) {
	if l.level <= WARN {
		l.logger.SetPrefix("WARN: ")
		l.logger.Printf(format, v...)
	}
}

// Error logs a message at the ERROR level
func (l *Logger) Error(format string, v ...any) {
	if l.level <= ERROR {
		l.logger.SetPrefix("ERROR: ")
		l.logger.Printf(format, v...)
	}
}

// NewLogger initializes a new Logger with the given log level
func NewLogger(level string, filePath string) (*Logger, error) {
	logLevel := parseLogLevel(level)

	var output *os.File
	var err error

	if filePath != "" {
		output, err = os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			return nil, err
		}
	} else {
		output = os.Stdout
	}

	return &Logger{
		level:  logLevel,
		logger: log.New(output, "", log.LstdFlags),
	}, nil
}
