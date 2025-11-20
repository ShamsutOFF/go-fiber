package main

import (
	"github.com/gofiber/fiber/v3/middleware/recover"
	"go-fiber/config"
	"go-fiber/internal/home"
	"log"

	"github.com/gofiber/fiber/v3"
)

func main() {
	config.Init()
	dbConf := config.NewDatabaseConfig()
	log.Println(dbConf)

	app := fiber.New()
	app.Use(recover.New())

	home.NewHomeHandler(app)

	log.Fatal(app.Listen(":3000"))
}
