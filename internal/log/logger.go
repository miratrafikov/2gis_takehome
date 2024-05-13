package log

type Logger struct {
	level loggingLevel
}

func NewLogger() Logger {
	logger := Logger{}
	logger.SetLevel(debugLevel)
	return logger
}
