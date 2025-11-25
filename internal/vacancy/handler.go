package vacancy

import (
	"go-fiber/pkg/tadapter"
	"go-fiber/pkg/validator"
	"go-fiber/views/components"
	"html/template"
	"time"

	"github.com/a-h/templ"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
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
	form := VacancyCreateForm{
		Email: c.FormValue("email"),
	}
	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный",
		},
	)

	time.Sleep(time.Second * 2)
	var component templ.Component

	if len(errors.Errors) > 0 {
		component = components.Notification(
			validator.FormatErrors(errors),
			components.NotificationFail,
		)
		return tadapter.Render(c, component)
	}

	component = components.Notification(
		"Вакансия успешно создана",
		components.NotificationSuccess,
	)
	return tadapter.Render(c, component)
}
