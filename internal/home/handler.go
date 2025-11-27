package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/tadapter"
	"go-fiber/views"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
}

func NewHomeHandler(
	router fiber.Router,
	customLogger *zerolog.Logger,
	repository *vacancy.VacancyRepository,
) {
	handler := &HomeHandler{
		router:       router,
		repository:   repository,
		customLogger: customLogger,
	}

	handler.router.Get("/", handler.home)
	handler.router.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	vacancies, err := h.repository.GetAllVacancies()
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	component := views.Main(vacancies)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Error 👋!")
}
