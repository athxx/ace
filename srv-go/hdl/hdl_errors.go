package hdl

import "github.com/gofiber/fiber/v2"

func ErrNotFound(c *fiber.Ctx) error {
	if c.Method() == "GET" {
		return c.Render("not_found", fiber.Map{
			"Title": "testetet",
		})
	}
	return c.Status(fiber.StatusNotFound).SendString(`{"code":404,"msg":"not found"}`)
}

func ErrInternal(c *fiber.Ctx, err error) error {
	if c.Method() == "GET" {
		return c.Render("not_found", fiber.Map{
			"Title": "testetet",
		})
	}
	return c.Status(fiber.StatusInternalServerError).SendString(`{"code":500,"msg":"internal server error"}`)
}
