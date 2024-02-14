package main

import (
	"fpbl/internal"
	"fpbl/internal/config"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {

	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	app := fiber.New()

	app = internal.SetupAPIRouter(app)

	err = app.Listen(cfg.PORT)
	if err != nil {
		log.Fatal(err)
	}

}
