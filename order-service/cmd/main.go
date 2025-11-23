package main

import (
	"log/slog"
	"order-service/internal/config"
	"order-service/internal/http-server/handlers/place"
	mwLogger "order-service/internal/http-server/middlware/logger"
	"order-service/internal/storage/postgres"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

const (
	envLocal = "local"
	envDev 	 = "dev"
	envProd  = "prod"
)

func main() {
	cfg := config.MustLoad()
	log := setupLogger(cfg.Env)
	log = log.With(slog.String("env", cfg.Env))

	// storage, err := postgres.NewStorage(cfg.DB.URL)
	// if err != nil {
	// 	log.Error("failed to initialize storage", slWrap.Err(err))
	// }

	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(mwLogger.New(log))
	router.Use(middleware.Recoverer)
	router.Use(middleware.URLFormat)

	router.Route("/order", func(r chi.Router) {
		r.Post("/", place.New(log, postgres.NewStorageWrapper()))	
	})
}

func setupLogger(env string) *slog.Logger {
	var log *slog.Logger

	switch env {
	case envLocal:
		log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	case envDev: 
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug, AddSource: true}))
	case envProd: 
		log = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelInfo}))
	}

	return log
}