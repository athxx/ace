package hdl

import "github.com/gofiber/fiber/v2"

// PostList forum post list
func PostList(c *fiber.Ctx) error {
	return nil
}

func PostDetail(c *fiber.Ctx) error {
	return nil
}

func PostReply(c *fiber.Ctx) error {
	args := &struct {
		PostId   int64  `json:"pid"`
		RelateId int64  `json:"relate_id"`
		Content  string `json:"content"`
	}{}
	if err := c.BodyParser(args); err != nil {
		return err
		//return c.Status(fiber.StatusUnprocessableEntity).JSON(fiber.Map{
		//	"errors": err.Error(),
		//})
	}

	return nil
}

func PostComment() {

}
