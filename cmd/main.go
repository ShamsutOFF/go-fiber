package main

import (
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

	home.NewHomeHandler(app)

	log.Fatal(app.Listen(":3000"))
}
