package main

import (
	"test/internal/config"
	"test/internal/tracer"

	"net/http"
	_ "net/http/pprof"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
)

func main() {
	setLogLevel()

	app := fiber.New()

	app.Get("/seperation", func(c *fiber.Ctx) error {
		from := c.Query("from")
		to := c.Query("to")
		if from == "" || to == "" {
			return c.Status(fiber.StatusBadRequest).JSON(map[string]string{"error": "from and to query params are required"})
		}

		seperation, err := tracer.FindSeperation(from, to)
		if err != nil {
			return c.Status(fiber.StatusInternalServerError).JSON(map[string]string{"error": err.Error()})
		}

		return c.Status(fiber.StatusOK).JSON(map[string]interface{}{
			"seperation": seperation,
		})
	})
	go func() {
		log.Fatal(http.ListenAndServe(":"+config.PprofPort, nil)) //for pprof, as fiber doesn't use net/http.
	}()
	log.Fatal(app.Listen(":" + config.Port))
}

func setLogLevel() {
	switch config.LogLevel {
	case "debug", "DEBUG":
		log.SetLevel(log.LevelDebug)
	case "info", "INFO":
		log.SetLevel(log.LevelInfo)
	default:
		log.SetLevel(log.LevelInfo)
	}
}
