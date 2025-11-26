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
		Email:    c.FormValue("email"),
		Location: c.FormValue("location"),
		Type:     c.FormValue("type"),
		Company:  c.FormValue("company"),
		Salary:   c.FormValue("salary"),
		Role:     c.FormValue("role"),
	}

	errors := validate.Validate(
		&validators.EmailIsPresent{
			Name:    "Email",
			Field:   form.Email,
			Message: "Email не задан или не верный",
		},
		&validators.StringIsPresent{
			Name: "Location",
			Field:   form.Location,
			Message: "Расположение не задано",
		},
		&validators.StringIsPresent{
			Name: "Type",
			Field:   form.Type,
			Message: "Сфера компании не задано",
		},
		&validators.StringIsPresent{
			Name: "Company",
			Field:   form.Company,
			Message: "Название компании не задано",
		},
		&validators.StringIsPresent{
			Name: "Role",
			Field:   form.Role,
			Message: "Должность не задана",
		},
		&validators.StringIsPresent{
			Name: "Salary",
			Field:   form.Salary,
			Message: "Зарплата не задана",
		},
	)

	//time.Sleep(time.Second * 2)
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
