package database

import (
	"context"
	"go-fiber/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDbPool(config *config.DatabaseConfig, logger *zerolog.Logger) *pgxpool.Pool {
	logger.Info().Str("url", config.Url).Msg("Connecting to database...")

	dbpool, err := pgxpool.New(context.Background(), config.Url)
	if err != nil {
		logger.Error().Err(err).Msg("Failed to connect to database")
		panic(err)
	}

	// Проверим подключение
	if err := dbpool.Ping(context.Background()); err != nil {
		logger.Error().Err(err).Msg("Database ping failed")
		panic(err)
	}

	logger.Info().Msg("Database connection established successfully")
	return dbpool
}
