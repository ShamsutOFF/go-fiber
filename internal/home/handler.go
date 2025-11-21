package home

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/views"
	"html/template"

	"github.com/gofiber/fiber/v2"
)

type HomeHandler struct {
	router fiber.Router
	tmpl   *template.Template
}

func NewHomeHandler(router fiber.Router) {
	handler := &HomeHandler{
		router: router,
	}

	api := handler.router.Group("/api")
	api.Get("/", handler.home)
	api.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	component := views.Hello("Adel")
	return tadapter.Render(c, component)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Error 👋!")
}
