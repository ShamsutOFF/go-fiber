package home

import (
	"html/template"
	"time"

	"github.com/gofiber/fiber/v3"
)

type HomeHandler struct {
	router fiber.Router
	tmpl   *template.Template
}

type News struct {
	Title   string
	Content string
	Date    string
	Link    string
}

type PageData struct {
	News       []News
	NewsCount  int
	LastUpdate string
}

func NewHomeHandler(router fiber.Router) {
	// Предварительно парсим шаблон при запуске
	tmpl := template.Must(template.ParseFiles("templates/home.html"))

	handler := &HomeHandler{
		router: router,
		tmpl:   tmpl,
	}

	api := handler.router.Group("/api")
	api.Get("/", handler.home)
	api.Get("/error", handler.error)
}

func (h *HomeHandler) home(c fiber.Ctx) error {
	// Подготовка данных для шаблона
	news := []News{
		{
			Title:   "Go 1.25 выпущен!",
			Content: "Вышла новая версия языка программирования Go с улучшениями производительности...",
			Date:    "21.11.2025",
			Link:    "/news/go-125",
		},
		{
			Title:   "Fiber v3 - что нового?",
			Content: "Анонсирован релиз кандидат Fiber v3 с поддержкой современных стандартов...",
			Date:    "20.11.2025",
			Link:    "/news/fiber-v3",
		},
		{
			Title:   "Искусственный интеллект в разработке",
			Content: "Как ИИ меняет подход к программированию и какие инструменты стоит попробовать...",
			Date:    "19.11.2025",
			Link:    "/news/ai-development",
		},
	}

	data := PageData{
		News:       news,
		NewsCount:  len(news),
		LastUpdate: time.Now().Format("02.01.2006 15:04"),
	}

	// Рендеринг шаблона с данными
	return c.Render("home", data)
}

func (h *HomeHandler) error(c fiber.Ctx) error {
	return c.SendString("Hello, World from Error 👋!")
}
