package home

import (
	"go-fiber/internal/vacancy"
	"go-fiber/pkg/tadapter"
	"go-fiber/views"
	"go-fiber/views/components"
	"math"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/rs/zerolog"
)

type HomeHandler struct {
	router       fiber.Router
	customLogger *zerolog.Logger
	repository   *vacancy.VacancyRepository
	store        *session.Store
}

func NewHomeHandler(
	router fiber.Router,
	customLogger *zerolog.Logger,
	repository *vacancy.VacancyRepository,
	store *session.Store,
) {
	handler := &HomeHandler{
		router:       router,
		repository:   repository,
		customLogger: customLogger,
		store:        store,
	}

	handler.router.Get("/", handler.home)

	handler.router.Get("/login", handler.login)
	handler.router.Post("/api/login", handler.apiLogin)
	handler.router.Get("/api/logout", handler.apiLogout)

	handler.router.Get("/error", handler.error)
}

func (h *HomeHandler) home(c *fiber.Ctx) error {
	PAGE_ITEMS := 2
	page := c.QueryInt("page", 1)

	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	userEmail := ""
	if email, ok := sess.Get("email").(string); ok {
		userEmail = email
	}
	c.Locals("email", userEmail)

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
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	userEmail := ""
	if email, ok := sess.Get("email").(string); ok {
		userEmail = email
	}
	c.Locals("email", userEmail)
	return tadapter.Render(c, component, http.StatusOK)
}

func (h *HomeHandler) apiLogin(c *fiber.Ctx) error {
	form := LoginForm{
		Email:    c.FormValue("email"),
		Password: c.FormValue("password"),
	}
	if form.Email == "a@a.ru" && form.Password == "1" { //Тут может быть реальная логика авторизации
		sess, err := h.store.Get(c)
		if err != nil {
			panic(err)
		}
		sess.Set("email", form.Email)
		if err := sess.Save(); err != nil {
			panic(err)
		}
		c.Response().Header.Add("Hx-Redirect", "/")
		c.Redirect("/", http.StatusOK)
	}
	component := components.Notification("Неверный логин или пароль", components.NotificationFail)
	return tadapter.Render(c, component, http.StatusBadRequest)
}

func (h *HomeHandler) apiLogout(c *fiber.Ctx) error {
	sess, err := h.store.Get(c)
	if err != nil {
		panic(err)
	}
	sess.Delete("email")
	if err := sess.Save(); err != nil {
		panic(err)
	}
	c.Response().Header.Add("Hx-Redirect", "/")
	return c.Redirect("/", http.StatusOK)
}
