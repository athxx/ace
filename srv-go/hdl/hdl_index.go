package hdl

import (
	"github.com/gofiber/fiber/v2"
)

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":     "testetet",
		"Orders":    1,
		"Paginator": 2,
	})
}
