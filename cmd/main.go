package main

import (
	"endeus/wawan/internal/config"
	"endeus/wawan/pkg/db"
	"fmt"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

func main() {
	godotenv.Load(".env")

	cfg := config.New()
	log.Info().Msg("Initializing...")
	log.Info().Msgf("Current Env: %s", cfg.Environment())

	// Setup logger
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	l, err := zerolog.ParseLevel(cfg.LogLevel())
	if err != nil {
		log.Fatal().Msgf("zerolog.ParseLevel: unable to parse log-level: %v; cfg.Loglevel: %s", err, cfg.LogLevel())
	}

	// Set global log-level
	zerolog.SetGlobalLevel(l)

	_, err = db.New(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Error connection db")
	}

	// repo := repository.NewRepo(db)

	fmt.Println(cfg.LogLevel())
}
