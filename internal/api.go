package internal

import (
	"fpbl/internal/routes"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func SetupAPIRouter(app *fiber.App) *fiber.App {

	app.Use(logger.New())

	app.Get("/ping", routes.Ping)
	return app
}
