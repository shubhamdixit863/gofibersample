package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"gofibersample/app/middlewares"
)

func main() {
	// Create a new Fiber instance
	app := fiber.New()

	// Create a new route with named params
	app.Get("/user/:name", func(c *fiber.Ctx) error {
		// Print name from params
		fmt.Println(c.Params("name"))

		return nil
	}).Name("username")

	app.Use(middlewares.New(middlewares.Config{}))

	// Create a new route with optional greedy params
	app.Get("/user/*", func(c *fiber.Ctx) error {
		// Print all data, given from '*' param
		fmt.Println(c.Params("*"))

		return nil
	})

	// Start server on port 3000
	app.Listen(":3000")
}
