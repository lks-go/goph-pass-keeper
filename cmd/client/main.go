package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"github.com/lks-go/goph-pass-keeper/internal/app/client"
)

func main() {
	zerolog.SetGlobalLevel(zerolog.DebugLevel)

	log.Info().Msg("Setup config")
	cfg := client.SetupConfig()

	log.Info().Msg("Building app")
	app := client.New(cfg)
	if err := app.Build(); err != nil {
		log.Error().Err(err).Msg("filed to build app")
		return
	}

	log.Info().Msg("Running app")
	if err := app.Run(); err != nil {
		log.Error().Err(err).Msg("filed to run app")
		return
	}

	log.Info().Msg("Service successfully stopped")
}
