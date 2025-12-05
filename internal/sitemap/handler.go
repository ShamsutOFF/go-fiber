package sitemap

import (
	"bytes"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/sabloger/sitemap-generator/smg"
)

type SitemapHandler struct {
	router fiber.Router
}

func NewSitemapHandler(router fiber.Router) {
	handler := &SitemapHandler{
		router: router,
	}
	handler.router.Get("/sitemap.xml", handler.getSitemap)
}

func (h *SitemapHandler) getSitemap(c *fiber.Ctx) error {
	now := time.Now().UTC()
	sm := smg.NewSitemap(false)
	sm.SetHostname("https://rabota.dev") // Должно подтягиваться из конфига
	sm.SetLastMod(&now)
	sm.SetCompress(false)
	sm.Add(&smg.SitemapLoc{
		Loc:        "/",
		LastMod:    &now,
		ChangeFreq: smg.Daily,
		Priority:   0.8,
	})
	sm.Add(&smg.SitemapLoc{
		Loc:        "/login",
		LastMod:    &now,
		ChangeFreq: smg.Weekly,
		Priority:   0.6,
	})
	sm.Finalize()
	var buf bytes.Buffer
	if _, err := sm.WriteTo(&buf); err != nil {
		return err
	}
	c.Set(fiber.HeaderContentType, fiber.MIMETextXMLCharsetUTF8)
	return c.Send(buf.Bytes())
}
