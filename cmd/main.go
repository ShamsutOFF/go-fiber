package main

import (
	"time"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/rs/zerolog"

	"go-fiber/config"
	"go-fiber/internal/home"
	"go-fiber/internal/vacancy"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	config.Init()
	config.NewDatabaseConfig()
	logConfig := config.NewLogConfig()

	app := fiber.New()

	zerolog.SetGlobalLevel(zerolog.Level(logConfig.Level))

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

	home.NewHomeHandler(app)
	vacancy.NewVacancyHandler(app)

	log.Fatal(app.Listen(":3000"))
}
