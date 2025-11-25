package vacancy

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/views/components"
	"html/template"
	"log"

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
	emai := c.FormValue("email")
	log.Println(emai)
	component := components.Notification("Вакансия успешно создана")
	return tadapter.Render(c, component)
}