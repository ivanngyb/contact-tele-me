package logger

import (
	"os"

	stdlog "log"

	"github.com/rs/zerolog"
)

var L *zerolog.Logger

func NewLogger(environment, logLevel string) {
	log := zerolog.New(os.Stdout).With().
		Timestamp().
		Logger()
	stdlog.SetFlags(0)
	stdlog.SetOutput(log)

	switch logLevel {
	case "PanicLevel":
		zerolog.SetGlobalLevel(zerolog.PanicLevel)
	case "FatalLevel":
		zerolog.SetGlobalLevel(zerolog.FatalLevel)
	case "ErrorLevel":
		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
	case "WarnLevel":
		zerolog.SetGlobalLevel(zerolog.WarnLevel)
	case "InfoLevel":
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	case "DebugLevel":
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	case "TraceLevel":
		zerolog.SetGlobalLevel(zerolog.TraceLevel)
	case "NoLevel":
		zerolog.SetGlobalLevel(zerolog.NoLevel)
	case "Disabled":
		zerolog.SetGlobalLevel(zerolog.Disabled)
	default:
		l := zerolog.New(os.Stdout)
		l.Panic().Msg("Invalid log level, Should be one of PanicLevel, FatalLevel, ErrorLevel, WarnLevel, InfoLevel, DebugLevel, TraceLevel")
	}

	log.Info().Msg("zerolog initialised")
	if L != nil {
		panic("Logger already initialised")
	}
	L = &log
	*L = L.With().Caller().Logger()
}
