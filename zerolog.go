package logging

import (
	"os"
	"time"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var version = "1.1"

var podName = os.Getenv("MY_POD_NAME")
var podNamespace = os.Getenv("MY_POD_NAMESPACE")

func initZerolog() {
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logLevel := os.Getenv("LOG_LEVEL")

	if logLevel == "" {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	} else {
		switch logLevel {
		case "Trace":
			zerolog.SetGlobalLevel(zerolog.TraceLevel)
		case "Debug":
			zerolog.SetGlobalLevel(zerolog.DebugLevel)
		case "Info":
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		case "Warn":
			zerolog.SetGlobalLevel(zerolog.WarnLevel)
		case "Error":
			zerolog.SetGlobalLevel(zerolog.ErrorLevel)
		case "Fatal":
			zerolog.SetGlobalLevel(zerolog.FatalLevel)
		default:
			zerolog.SetGlobalLevel(zerolog.InfoLevel)
		}
	}

	log.Logger = log.With().Caller().Logger().Output(zerolog.ConsoleWriter{Out: os.Stderr, TimeFormat: time.RFC3339})
}

func errorZeroLog(err error, msg string, extra interface{}) {
	log.Error().
		CallerSkipFrame(2).
		Str("serviceContext.service", "warehouse-ms-go").
		Str("serviceContext.version", version).
		AnErr("context.reportLocation", err).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func fatalZeroLog(err error, msg string, extra interface{}) {
	log.Fatal().
		CallerSkipFrame(2).
		Str("serviceContext.service", "warehouse-ms-go").
		Str("serviceContext.version", version).
		AnErr("context.reportLocation", err).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func warnZeroLog(msg string, extra interface{}) {
	log.Warn().
		CallerSkipFrame(2).
		Str("serviceContext.service", "warehouse-ms-go").
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func infoZeroLog(msg string, extra interface{}) {
	log.Info().
		CallerSkipFrame(2).
		Str("serviceContext.service", "warehouse-ms-go").
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func debugZeroLog(msg string, extra interface{}) {
	log.Debug().
		CallerSkipFrame(2).
		Str("serviceContext.service", "warehouse-ms-go").
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}

func traceZeroLog(msg string, extra interface{}) {
	log.Trace().
		CallerSkipFrame(2).
		Str("serviceContext.service", "warehouse-ms-go").
		Str("serviceContext.version", version).
		Str("cloudContext.podName", podName).
		Str("cloudContext.podNamespace", podNamespace).
		Interface("extra", extra).
		Msg(msg)
}
