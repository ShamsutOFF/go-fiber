package main

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/gofiber/storage/postgres/v3"
	"github.com/rs/zerolog"

	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/internal/sitemap"
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/database"
	logGlob "go-fiber/pkg/logger"
	"go-fiber/pkg/middleware"

	"github.com/gofiber/fiber/v2"
)

func main() {
	config.Init()
	logConfig := config.NewLogConfig()
	dbConfig := config.NewDatabaseConfig()

	// Инициализируем глобальный логгер
	logGlob.Init(zerolog.Level(logConfig.Level), logConfig.Format)

	app := fiber.New()

	// HTTP middleware (только для access логов)
	if logConfig.Format == "json" {
		app.Use(logger.New(logger.Config{
			Format:     `{"time":"${time}","method":"${method}","path":"${path}","status":${status},"latency":${latency},"ip":"${ip}"}` + "\n",
			TimeFormat: time.RFC3339,
		}))
	} else {
		app.Use(logger.New())
	}

	app.Use(recover.New())
	app.Static("/public", "./public")

	logGlob.Info().Msg("Application starting...")

	dbpool := database.CreateDbPool(dbConfig, &logGlob.Log)
	defer dbpool.Close()

	storage := postgres.New(postgres.Config{
		DB:         dbpool,
		Table:      "sessions",
		Reset:      false,
		GCInterval: 10 * time.Second,
	})
	store := session.New(session.Config{
		Storage: storage,
	})
	app.Use(middleware.AuthMiddleware(store))

	// Репозитории
	vacancyRepo := vacancy.NewVacancyRepository(dbpool, &logGlob.Log)

	// Регистрируем обработчики
	home.NewHomeHandler(app, &logGlob.Log, vacancyRepo, store)
	vacancy.NewVacancyHandler(app, &logGlob.Log, vacancyRepo)
	sitemap.NewSitemapHandler(app)

	logGlob.Info().Str("port", "3000").Msg("Server started")
	logGlob.Fatal().Err(app.Listen(":3000")).Msg("Server failed to start")
}
