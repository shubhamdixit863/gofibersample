package middlewares

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
)

type Config struct {
	Filter func(c *fiber.Ctx) bool
	Check  func(c *fiber.Ctx)
}

func New(config Config) fiber.Handler {

	return func(c *fiber.Ctx) error {
		location, _ := c.GetRouteURL("username", nil)

		fmt.Println("Midddle was run", location)
		return nil
	}
}
