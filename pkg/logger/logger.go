package logger

import (
	"os"

	"github.com/rs/zerolog"
)

var Log zerolog.Logger

func Init(level zerolog.Level, format string) {
	zerolog.SetGlobalLevel(level)

	if format == "json" {
		Log = zerolog.New(os.Stdout).With().Timestamp().Logger()
	} else {
		// Красивый консольный вывод
		output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: "2006-01-02 15:04:05"}
		Log = zerolog.New(output).With().Timestamp().Logger()
	}
}

// Helper методы для удобства
func Info() *zerolog.Event {
	return Log.Info()
}

func Error() *zerolog.Event {
	return Log.Error()
}

func Debug() *zerolog.Event {
	return Log.Debug()
}

func Warn() *zerolog.Event {
	return Log.Warn()
}

func Fatal() *zerolog.Event {
	return Log.Fatal()
}
