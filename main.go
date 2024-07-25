package main

import (
	"ambassador/src/database"
	"ambassador/src/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// database
	database.Connect()
	database.AutoMigrate()

	// routes
	app := fiber.New()
	routes.Setup(app)
	app.Listen(":8000")
}
