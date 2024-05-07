package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
)

var (
	version      = ""
	podName      = ""
	podNamespace = ""
	serviceName  = ""
	logLevel     = ""
	stderrLogger zerolog.Logger
	stdoutLogger zerolog.Logger
)

func initZerolog(InputPodName string, InputPodNamespace string, logLevelInput string, serviceNameInput string, versionInput string) {
	podName = InputPodName
	podNamespace = InputPodNamespace
	serviceName = serviceNameInput
	logLevel = logLevelInput
	version = versionInput

	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix

	level := zerolog.InfoLevel // Default to Info level
	if logLevel != "" {
		switch logLevel {
		case "Trace":
			level = zerolog.TraceLevel
		case "Debug":
			level = zerolog.DebugLevel
		case "Info":
			level = zerolog.InfoLevel
		case "Warn":
			level = zerolog.WarnLevel
		case "Error":
			level = zerolog.ErrorLevel
		case "Fatal":
			level = zerolog.FatalLevel
		}
	}
	zerolog.SetGlobalLevel(level)

	stderrLogger = zerolog.New(os.Stderr).With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
	stdoutLogger = zerolog.New(os.Stdout).With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339})
}

func getLoggerForLevel(level zerolog.Level) zerolog.Logger {
	if level == zerolog.ErrorLevel || level == zerolog.FatalLevel {
		return stderrLogger
	}
	return stdoutLogger
}

func errorZeroLog(err error, msg string, extra interface{}) {
	logger := getLoggerForLevel(zerolog.ErrorLevel)
	logger.Error().
		CallerSkipFrame(2).
		Str("serviceContext.service", serviceName).
		Str("serviceContext.version", version).
		AnErr("context.reportLocation", err).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func fatalZeroLog(err error, msg string, extra interface{}) {
	logger := getLoggerForLevel(zerolog.FatalLevel)
	logger.Fatal().
		CallerSkipFrame(2).
		Str("serviceContext.service", serviceName).
		Str("serviceContext.version", version).
		AnErr("context.reportLocation", err).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func warnZeroLog(msg string, extra interface{}) {
	logger := getLoggerForLevel(zerolog.WarnLevel)
	logger.Warn().
		CallerSkipFrame(2).
		Str("serviceContext.service", serviceName).
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func infoZeroLog(msg string, extra interface{}) {
	logger := getLoggerForLevel(zerolog.InfoLevel)
	logger.Info().
		CallerSkipFrame(2).
		Str("serviceContext.service", serviceName).
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func debugZeroLog(msg string, extra interface{}) {
	logger := getLoggerForLevel(zerolog.DebugLevel)
	logger.Debug().
		CallerSkipFrame(2).
		Str("serviceContext.service", serviceName).
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}
