package home

import "github.com/gofiber/fiber/v3"

type HomeHandler struct {
	router fiber.Router
}

func NewHomeHandler(router fiber.Router) {
	handler := &HomeHandler{
		router: router,
	}
	api := handler.router.Group("/api")
	api.Get("/", handler.home)
	api.Get("/error", handler.error)
}

func (h *HomeHandler) home(c fiber.Ctx) error {
	return c.SendString("Hello, World from Home 👋!")
}

func (h *HomeHandler) error(c fiber.Ctx) error {
	return c.SendString("Hello, World from Error 👋!")
}
