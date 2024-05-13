package log

import "errors"

type loggingLevel int

const (
	debugLevel loggingLevel = iota
	infoLevel
	errorLevel
	fatalLevel
	temporaryLevel
)

func (logger *Logger) SetLevel(level loggingLevel) {
	logger.level = level
}

func StringToLoggingLevel(s string) (loggingLevel, error) {
	switch s {
	case "debug":
		return debugLevel, nil
	case "info":
		return infoLevel, nil
	case "error":
		return errorLevel, nil
	case "fatal":
		return fatalLevel, nil
	default:
		return 0, errors.New("unknown logging level")
	}
}

func (logger Logger) shouldPrint(level loggingLevel) bool {
	return level >= logger.level
}

func loggingLevelToString(loggingLevel loggingLevel) string {
	switch loggingLevel {
	case debugLevel:
		return "DEBUG"
	case infoLevel:
		return "INFO"
	case errorLevel:
		return "ERROR"
	default:
		return "FATAL"
	}
}
