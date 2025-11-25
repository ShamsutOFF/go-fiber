package vacancy

import (
	"html/template"

	"github.com/gofiber/fiber/v2"
)

type VacancyHandler struct {
	router fiber.Router
	tmpl   *template.Template
}

func NewVacancyHandler(router fiber.Router) {
	handler := &VacancyHandler{
		router: router,
	}

	vacancyGroup := handler.router.Group("/vacancy")
	vacancyGroup.Post("/", handler.createVacancy)
}

func (h *VacancyHandler) createVacancy(c *fiber.Ctx) error {
	return c.SendString("createVacancy")
}