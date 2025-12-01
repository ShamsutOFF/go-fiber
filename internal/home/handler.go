package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/tadapter"
	"go-fiber/views"
	"math"

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
	handler.router.Get("/login", handler.login)
	handler.router.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)
	count := h.repository.GetCountAll()
	vacancies, err := h.repository.GetAllVacancies(PAGE_ITEMS, (page-1)*PAGE_ITEMS)
	if err != nil {
		h.customLogger.Error().Msg(err.Error())
		return c.SendStatus(http.StatusInternalServerError)
	}
	component := views.Main(vacancies, int(math.Ceil(float64(count/PAGE_ITEMS))), page)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) error(c *fiber.Ctx) error {
	return c.SendString("Hello, World from Error 👋!")
}

func (h *HomeHandler) login(c *fiber.Ctx) error {
	component := views.Login()
	return tadapter.Render(c, component, http.StatusOK)
}
