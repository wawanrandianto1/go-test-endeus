package main

import (
	"context"
	"endeus/wawan/internal/config"
	"endeus/wawan/pkg/db"
	"endeus/wawan/pkg/endeus"
	"endeus/wawan/pkg/handler"
	"endeus/wawan/pkg/repository"
	"endeus/wawan/pkg/router"
	"endeus/wawan/pkg/service"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"

	_ "endeus/wawan/docs"

	echoSwagger "github.com/swaggo/echo-swagger"
)

// @title Swagger Endeus API
// @version 1.0

// @host 127.0.0.1:5000
// @BasePath /api

// @schemes http https
// @produce	application/json
// @consumes application/json
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

	dbcon, err := db.New(cfg)
	if err != nil {
		log.Error().Err(err).Msg("Error connection db")
		os.Exit(1)
	}

	repo := repository.NewRepo(dbcon)
	svc := service.NewService(repo)
	ctrl := handler.NewHandler(cfg, svc)

	route := router.New()
	// health check
	route.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Service is healthy!")
	})
	route.GET("/swagger/*", echoSwagger.WrapHandler)

	v1 := route.Group("/api")

	ctrl.Register(v1)
	serveWithGracefulShutdown(cfg, route, dbcon)
}

func serveWithGracefulShutdown(cfg endeus.Config, e *echo.Echo, db *gorm.DB) {
	go func() {
		port := ":" + cfg.ServerPort()
		log.Info().Msgf("HTTP server is listening on: http://localhost%s", port)
		if err := e.Start(port); err != nil && err != http.ErrServerClosed {
			log.Fatal().Err(err).Msg("")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal().Err(err)
	}
	defer sqlDB.Close()

	if err := e.Shutdown(ctx); err != nil {
		log.Fatal().Err(err).Msg("error when shutting down the HTTP server")
	}
	log.Info().Msg("Bye!")
}
