package log

import (
	"fmt"
	"os"
	"time"
)

func (logger Logger) Debug(content string) {
	if logger.shouldPrint(debugLevel) {
		logger.printLog(content, debugLevel)
	}
}

func (logger Logger) Debugf(format string, v ...any) {
	logger.Debug(fmt.Sprintf(format, v...))
}

func (logger Logger) Info(content string) {
	if logger.shouldPrint(infoLevel) {
		logger.printLog(content, infoLevel)
	}
}

func (logger Logger) Infof(format string, v ...any) {
	logger.Info(fmt.Sprintf(format, v...))
}

func (logger Logger) Error(content string) {
	if logger.shouldPrint(errorLevel) {
		logger.printLog(content, errorLevel)
	}
}

func (logger Logger) Errorf(format string, v ...any) {
	logger.Error(fmt.Sprintf(format, v...))
}

// Fatal prints message to stdout and terminates the process.
func (logger Logger) Fatal(content string) {
	logger.printLog(content, fatalLevel)
	os.Exit(1)
}

// Fatalf formats message, prints it to stdout,
// and terminates the process.
func (logger Logger) Fatalf(format string, v ...any) {
	logger.Fatal(fmt.Sprintf(format, v...))
}

func (logger Logger) printLog(content string, level loggingLevel) {
	currentTime := time.Now()
	fmt.Printf(
		"\n[%s][%s]: %s",
		currentTime.Format("2006-01-02 15:04:05"),
		loggingLevelToString(level),
		content,
	)
}
