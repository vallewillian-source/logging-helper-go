package logging

import "fmt"

func Init(podName string, podNamespace string, logLevel string, serviceName string, version string, backend string) error {
	if backend == "zerolog" {
		initZerolog(podName, podNamespace, logLevel, serviceName, version)
	} else {
		fmt.Println("Invalid backend")
		return fmt.Errorf("Invalid backend")
	}

	return nil
}

// Error logs a message with severity "ERROR"
func Error(err error, msg string, extra interface{}) {
	errorZeroLog(err, msg, extra)
}

// Fatal logs a message with severity "FATAL"
func Fatal(err error, msg string, extra interface{}) {
	fatalZeroLog(err, msg, extra)
}

// Warn logs a message with severity "WARN"
func Warn(msg string, extra interface{}) {
	warnZeroLog(msg, extra)
}

// Info logs a message with severity "INFO"
func Info(msg string, extra interface{}) {
	infoZeroLog(msg, extra)
}

// Debug logs a message with severity "DEBUG"
func Debug(msg string, extra interface{}) {
	debugZeroLog(msg, extra)
}
