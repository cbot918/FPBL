package routes

import "github.com/gofiber/fiber/v2"

func Ping(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"msg": "hihi"})
}
