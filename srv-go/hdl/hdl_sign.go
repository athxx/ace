package hdl

import (
	"github.com/gofiber/fiber/v2"
)

func Sign(c *fiber.Ctx) error {
	return nil
}

func SignUp(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":     "testetet",
		"Orders":    1,
		"Paginator": 2,
	})
}

func SignIn(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{
		"Title":     "testetet",
		"Orders":    1,
		"Paginator": 2,
	})
}
