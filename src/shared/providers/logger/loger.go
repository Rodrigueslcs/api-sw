package logger

import (
	"os"

	"github.com/rs/zerolog"
)

type logger struct {
	client zerolog.Logger
}

var Instance *logger = nil

func New() logger {
	if Instance == nil {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
		log := zerolog.New(os.Stderr).With().Timestamp().Logger()

		Instance = &logger{
			client: log,
		}
	}
	return *Instance
}

func (l logger) Error(namespace, message string) {
	l.client.Error().Str("namespace", namespace).Msg(message)
}

func (l logger) Info(namespace, message string) {
	l.client.Info().Str("namespace", namespace).Msg(message)
}
