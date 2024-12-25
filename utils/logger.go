package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Console zerolog.Logger

func init() {
	Console = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	SetLevel(zerolog.InfoLevel)
}

func SetLevel(level zerolog.Level) {
	Console.Level(level)
}
