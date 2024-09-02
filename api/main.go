package main

import (
	_ "embed"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

//go:embed assets/index.html
var indexHTML []byte

func main() {
	app := fiber.New()

	// Initialize default config
	app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Redirect("/app")
	})

	app.Get("/app/*", func(c *fiber.Ctx) error {
		// Return `indexHTML` as string with `text/html` content type
		return c.Type("html").Send(indexHTML)
	})

	app.Get("/hello", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/api/processes", func(c *fiber.Ctx) error {
		// Get running processes
		processList, err := getRunningProcesses()
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		}
		return c.JSON(processList)
	})

	// Start the server on port 3000
	app.Listen(":3001")
}
