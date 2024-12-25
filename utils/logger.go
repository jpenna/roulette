package utils

import (
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var Console zerolog.Logger

func init() {
	Console = log.Logger.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	zerolog.SetGlobalLevel(zerolog.InfoLevel)
}

func SetLevel(level zerolog.Level) {
	zerolog.SetGlobalLevel(level)
}
